package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	modulePaths, err := findModules()
	if err != nil {
		return fmt.Errorf("findModules failed: %w", err)
	}

	modules := []*Module{}
	for _, modulePath := range modulePaths {
		modules = append(modules, &Module{
			Path: modulePath,
			Dir:  filepath.Dir(modulePath),
			Done: make(chan struct{}),
		})
	}

	var downloads errgroup.Group
	downloads.SetLimit(8)

	for _, module := range modules {
		module := module
		downloads.Go(func() error {
			return DownloadDependencies(module.Dir, os.Stdout)
		})
	}
	if err := downloads.Wait(); err != nil {
		return err
	}

	var group errgroup.Group
	group.SetLimit(4 + 1)

	failures := []error{}

	group.Go(func() error {
		for _, module := range modules {
			<-module.Done
			if module.Failed() {
				module.WriteTo(os.Stdout)
				failures = append(failures, fmt.Errorf("Linting %v failed", module.Dir))
			} else {
				fmt.Fprintf(os.Stdout, "# %v OK\n", module.Dir)
			}
		}
		return nil
	})

	for _, module := range modules {
		module := module
		group.Go(func() error {
			module.Lint()
			return nil
		})
	}

	_ = group.Wait()

	return errors.Join(failures...)
}

type Module struct {
	Path string
	Dir  string

	Failures []Failure

	Done chan struct{}
}

func (m *Module) Lint() {
	defer close(m.Done)

	m.Fmt()
	m.Vet()
	m.Staticcheck()
}

var ignoreLines = []string{
	"a blank import should be only in a main or test package",
	"(SA1019)",
}

func (m *Module) Failed() bool { return len(m.Failures) > 0 }

func (m *Module) WriteTo(w io.Writer) {
	for _, fail := range m.Failures {
		fail.WriteTo(w)
	}
}

type Failure struct {
	Tool           string
	Path           string
	Lineno, Column int
	Posn           string
	Message        string
}

func (fail *Failure) WriteTo(w io.Writer) {
	if fail.Lineno >= 0 {
		fmt.Fprintf(w, "::error file=%v,line=%v,title=%v::%v%v\n", fail.Path, fail.Lineno, fail.Tool, fail.Posn, fail.Message)
	} else {
		fmt.Fprintf(w, "::error file=%v,title=%v::%v%v\n", fail.Path, fail.Tool, fail.Posn, fail.Message)
	}
}

func (m *Module) Fmt() {
	cmd := exec.Command("gofmt", "-s", "-d", "-e", "-l", m.Dir+"/.")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		m.Failures = append(m.Failures, Failure{
			Tool:    "gofmt",
			Path:    m.Dir,
			Message: "Code not formatted: " + out.String(),
		})
	}
}

var rxComment = regexp.MustCompile("# .*\n")

type VetResult map[string]VetPackageResult
type VetPackageResult map[string][]VetMessage

type VetMessage struct {
	Posn    string
	Message string
}

func (m *Module) Vet() {
	var stderr bytes.Buffer

	cmd := exec.Command("go", "vet", "-json", "./"+m.Dir+"/...")
	cmd.Stdout = io.Discard
	cmd.Stderr = &stderr

	_ = cmd.Run()

	workdir, _ := os.Getwd()

	result := rxComment.ReplaceAll(stderr.Bytes(), nil)
	dec := json.NewDecoder(bytes.NewReader(result))
	for {
		result := VetResult{}
		err := dec.Decode(&result)
		if errors.Is(err, io.EOF) {
			break
		}

		for pkgname, pkgresult := range result {
			_ = pkgname
			for toolname, messages := range pkgresult {
				for _, message := range messages {
					path, line, col := ParsePosn(message.Posn)

					if relpath, err := filepath.Rel(workdir, path); err == nil {
						path = relpath
					}

					m.Failures = append(m.Failures, Failure{
						Tool:    "go vet " + toolname,
						Path:    path,
						Lineno:  line,
						Column:  col,
						Posn:    fmt.Sprintf("%v:%v:%v: ", path, line, col),
						Message: message.Message,
					})
				}
			}
		}
	}
}

type StaticcheckMessage struct {
	Code     string
	Severity string
	Location struct {
		File string
		Line int
	}
	Message string
}

var ignoreStaticcheckCode = map[string]bool{
	"SA1019": true,
}

func (m *Module) Staticcheck() {
	cmd := exec.Command("staticcheck", "-f", "json", "-go=1.19", "./"+m.Dir+"/...")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()

	if stderr.Len() > 0 {
		m.Failures = append(m.Failures, Failure{
			Tool:    "staticcheck stderr",
			Path:    m.Dir,
			Message: stderr.String(),
		})
	}

	workdir, _ := os.Getwd()

	dec := json.NewDecoder(bytes.NewReader(stdout.Bytes()))
	for {
		msg := StaticcheckMessage{}
		err := dec.Decode(&msg)
		if errors.Is(err, io.EOF) {
			break
		}

		if ignoreStaticcheckCode[msg.Code] {
			continue
		}

		path := msg.Location.File
		if relpath, err := filepath.Rel(workdir, path); err == nil {
			path = relpath
		}

		m.Failures = append(m.Failures, Failure{
			Tool:    "staticcheck " + msg.Code,
			Path:    path,
			Lineno:  msg.Location.Line,
			Posn:    fmt.Sprintf("%v:%v: ", path, msg.Location.Line),
			Message: msg.Message,
		})
	}
}

func findModules() (modules []string, _ error) {
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == "go.mod" {
			modules = append(modules, path)
		}

		return nil
	})

	return modules, err
}

func DownloadDependencies(dir string, out io.Writer) error {
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = dir

	fmt.Fprintln(out, "# go mod download", "./"+dir)

	cmd.Stdout = out
	cmd.Stderr = out

	return cmd.Run()
}

func ParsePosn(v string) (path string, lineno, column int) {
	p := strings.LastIndexByte(v, ':')
	if p < 0 {
		// there's no position at the end
		return v, -1, -1
	}

	last, ok := parseInt(v[p+1:])
	if !ok {
		// there's no position at the end
		return v, -1, -1
	}

	p2 := strings.LastIndexByte(v, ':')
	if p2 < 0 {
		// there's only one,
		// x/y:123
		return v, last, -1
	}

	last2, ok := parseInt(v[p2+1:])
	if !ok {
		// there's only one, and on windows
		// C:/x/y:123
		return v, last, -1
	}

	return v, last2, last
}

func parseInt(data string) (int, bool) {
	x, err := strconv.Atoi(data)
	if err != nil {
		return -1, false
	}
	return x, true
}
