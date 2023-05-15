// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/scheduler/v1/job.proto

package schedulerpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// State of the job.
type Job_State int32

const (
	// Unspecified state.
	Job_STATE_UNSPECIFIED Job_State = 0
	// The job is executing normally.
	Job_ENABLED Job_State = 1
	// The job is paused by the user. It will not execute. A user can
	// intentionally pause the job using
	// [PauseJobRequest][google.cloud.scheduler.v1.PauseJobRequest].
	Job_PAUSED Job_State = 2
	// The job is disabled by the system due to error. The user
	// cannot directly set a job to be disabled.
	Job_DISABLED Job_State = 3
	// The job state resulting from a failed
	// [CloudScheduler.UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob]
	// operation. To recover a job from this state, retry
	// [CloudScheduler.UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob]
	// until a successful response is received.
	Job_UPDATE_FAILED Job_State = 4
)

// Enum value maps for Job_State.
var (
	Job_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "PAUSED",
		3: "DISABLED",
		4: "UPDATE_FAILED",
	}
	Job_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ENABLED":           1,
		"PAUSED":            2,
		"DISABLED":          3,
		"UPDATE_FAILED":     4,
	}
)

func (x Job_State) Enum() *Job_State {
	p := new(Job_State)
	*p = x
	return p
}

func (x Job_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Job_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_scheduler_v1_job_proto_enumTypes[0].Descriptor()
}

func (Job_State) Type() protoreflect.EnumType {
	return &file_google_cloud_scheduler_v1_job_proto_enumTypes[0]
}

func (x Job_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Job_State.Descriptor instead.
func (Job_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_scheduler_v1_job_proto_rawDescGZIP(), []int{0, 0}
}

// Configuration for a job.
// The maximum allowed size for a job is 1MB.
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optionally caller-specified in
	// [CreateJob][google.cloud.scheduler.v1.CloudScheduler.CreateJob], after
	// which it becomes output only.
	//
	// The job name. For example:
	// `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the job's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optionally caller-specified in
	// [CreateJob][google.cloud.scheduler.v1.CloudScheduler.CreateJob] or
	// [UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob].
	//
	// A human-readable description for the job. This string must not contain
	// more than 500 characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Required.
	//
	// Delivery settings containing destination and parameters.
	//
	// Types that are assignable to Target:
	//	*Job_PubsubTarget
	//	*Job_AppEngineHttpTarget
	//	*Job_HttpTarget
	Target isJob_Target `protobuf_oneof:"target"`
	// Required, except when used with
	// [UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob].
	//
	// Describes the schedule on which the job will be executed.
	//
	// The schedule can be either of the following types:
	//
	// * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview)
	// * English-like
	// [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules)
	//
	// As a general rule, execution `n + 1` of a job will not begin
	// until execution `n` has finished. Cloud Scheduler will never
	// allow two simultaneously outstanding executions. For example,
	// this implies that if the `n+1`th execution is scheduled to run at
	// 16:00 but the `n`th execution takes until 16:15, the `n+1`th
	// execution will not start until `16:15`.
	// A scheduled start time will be delayed if the previous
	// execution has not ended when its scheduled time occurs.
	//
	// If [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] > 0 and
	// a job attempt fails, the job will be tried a total of
	// [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] times,
	// with exponential backoff, until the next scheduled start time.
	Schedule string `protobuf:"bytes,20,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// Specifies the time zone to be used in interpreting
	// [schedule][google.cloud.scheduler.v1.Job.schedule]. The value of this field
	// must be a time zone name from the [tz
	// database](http://en.wikipedia.org/wiki/Tz_database).
	//
	// Note that some time zones include a provision for
	// daylight savings time. The rules for daylight saving time are
	// determined by the chosen tz. For UTC use the string "utc". If a
	// time zone is not specified, the default will be in UTC (also known
	// as GMT).
	TimeZone string `protobuf:"bytes,21,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Output only. The creation time of the job.
	UserUpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=user_update_time,json=userUpdateTime,proto3" json:"user_update_time,omitempty"`
	// Output only. State of the job.
	State Job_State `protobuf:"varint,10,opt,name=state,proto3,enum=google.cloud.scheduler.v1.Job_State" json:"state,omitempty"`
	// Output only. The response from the target for the last attempted execution.
	Status *status.Status `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The next time the job is scheduled. Note that this may be a
	// retry of a previously failed attempt or the next execution time
	// according to the schedule.
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time the last job attempt started.
	LastAttemptTime *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=last_attempt_time,json=lastAttemptTime,proto3" json:"last_attempt_time,omitempty"`
	// Settings that determine the retry behavior.
	RetryConfig *RetryConfig `protobuf:"bytes,19,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// The deadline for job attempts. If the request handler does not respond by
	// this deadline then the request is cancelled and the attempt is marked as a
	// `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according
	// to the [RetryConfig][google.cloud.scheduler.v1.RetryConfig].
	//
	// The default and the allowed values depend on the type of target:
	//
	// * For [HTTP targets][google.cloud.scheduler.v1.Job.http_target], the
	// default is 3 minutes. The deadline must be in the interval [15 seconds, 30
	// minutes].
	//
	// * For [App Engine HTTP
	// targets][google.cloud.scheduler.v1.Job.app_engine_http_target], 0 indicates
	// that the request has the default deadline. The default deadline depends on
	// the scaling type of the service: 10 minutes for standard apps with
	// automatic scaling, 24 hours for standard apps with manual and basic
	// scaling, and 60 minutes for flex apps. If the request deadline is set, it
	// must be in the interval [15 seconds, 24 hours 15 seconds].
	//
	// * For [Pub/Sub targets][google.cloud.scheduler.v1.Job.pubsub_target], this
	// field is ignored.
	AttemptDeadline *durationpb.Duration `protobuf:"bytes,22,opt,name=attempt_deadline,json=attemptDeadline,proto3" json:"attempt_deadline,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_scheduler_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_scheduler_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_google_cloud_scheduler_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Job) GetTarget() isJob_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *Job) GetPubsubTarget() *PubsubTarget {
	if x, ok := x.GetTarget().(*Job_PubsubTarget); ok {
		return x.PubsubTarget
	}
	return nil
}

func (x *Job) GetAppEngineHttpTarget() *AppEngineHttpTarget {
	if x, ok := x.GetTarget().(*Job_AppEngineHttpTarget); ok {
		return x.AppEngineHttpTarget
	}
	return nil
}

func (x *Job) GetHttpTarget() *HttpTarget {
	if x, ok := x.GetTarget().(*Job_HttpTarget); ok {
		return x.HttpTarget
	}
	return nil
}

func (x *Job) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *Job) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Job) GetUserUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UserUpdateTime
	}
	return nil
}

func (x *Job) GetState() Job_State {
	if x != nil {
		return x.State
	}
	return Job_STATE_UNSPECIFIED
}

func (x *Job) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Job) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *Job) GetLastAttemptTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAttemptTime
	}
	return nil
}

func (x *Job) GetRetryConfig() *RetryConfig {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

func (x *Job) GetAttemptDeadline() *durationpb.Duration {
	if x != nil {
		return x.AttemptDeadline
	}
	return nil
}

type isJob_Target interface {
	isJob_Target()
}

type Job_PubsubTarget struct {
	// Pub/Sub target.
	PubsubTarget *PubsubTarget `protobuf:"bytes,4,opt,name=pubsub_target,json=pubsubTarget,proto3,oneof"`
}

type Job_AppEngineHttpTarget struct {
	// App Engine HTTP target.
	AppEngineHttpTarget *AppEngineHttpTarget `protobuf:"bytes,5,opt,name=app_engine_http_target,json=appEngineHttpTarget,proto3,oneof"`
}

type Job_HttpTarget struct {
	// HTTP target.
	HttpTarget *HttpTarget `protobuf:"bytes,6,opt,name=http_target,json=httpTarget,proto3,oneof"`
}

func (*Job_PubsubTarget) isJob_Target() {}

func (*Job_AppEngineHttpTarget) isJob_Target() {}

func (*Job_HttpTarget) isJob_Target() {}

// Settings that determine the retry behavior.
//
// By default, if a job does not complete successfully (meaning that
// an acknowledgement is not received from the handler, then it will be retried
// with exponential backoff according to the settings in
// [RetryConfig][google.cloud.scheduler.v1.RetryConfig].
type RetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of attempts that the system will make to run a job using the
	// exponential backoff procedure described by
	// [max_doublings][google.cloud.scheduler.v1.RetryConfig.max_doublings].
	//
	// The default value of retry_count is zero.
	//
	// If retry_count is zero, a job attempt will *not* be retried if
	// it fails. Instead the Cloud Scheduler system will wait for the
	// next scheduled execution time.
	//
	// If retry_count is set to a non-zero number then Cloud Scheduler
	// will retry failed attempts, using exponential backoff,
	// retry_count times, or until the next scheduled execution time,
	// whichever comes first.
	//
	// Values greater than 5 and negative values are not allowed.
	RetryCount int32 `protobuf:"varint,1,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	// The time limit for retrying a failed job, measured from time when an
	// execution was first attempted. If specified with
	// [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count], the job
	// will be retried until both limits are reached.
	//
	// The default value for max_retry_duration is zero, which means retry
	// duration is unlimited.
	MaxRetryDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// The minimum amount of time to wait before retrying a job after
	// it fails.
	//
	// The default value of this field is 5 seconds.
	MinBackoffDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=min_backoff_duration,json=minBackoffDuration,proto3" json:"min_backoff_duration,omitempty"`
	// The maximum amount of time to wait before retrying a job after
	// it fails.
	//
	// The default value of this field is 1 hour.
	MaxBackoffDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=max_backoff_duration,json=maxBackoffDuration,proto3" json:"max_backoff_duration,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A job's retry interval starts at
	// [min_backoff_duration][google.cloud.scheduler.v1.RetryConfig.min_backoff_duration],
	// then doubles `max_doublings` times, then increases linearly, and finally
	// retries at intervals of
	// [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration]
	// up to [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count]
	// times.
	//
	// For example, if
	// [min_backoff_duration][google.cloud.scheduler.v1.RetryConfig.min_backoff_duration]
	// is 10s,
	// [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration]
	// is 300s, and `max_doublings` is 3, then the a job will first be retried in
	// 10s. The retry interval will double three times, and then increase linearly
	// by 2^3 * 10s.  Finally, the job will retry at intervals of
	// [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration]
	// until the job has been attempted
	// [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] times.
	// Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s,
	// 300s, ....
	//
	// The default value of this field is 5.
	MaxDoublings int32 `protobuf:"varint,5,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
}

func (x *RetryConfig) Reset() {
	*x = RetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_scheduler_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryConfig) ProtoMessage() {}

func (x *RetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_scheduler_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryConfig.ProtoReflect.Descriptor instead.
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_scheduler_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *RetryConfig) GetRetryCount() int32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *RetryConfig) GetMaxRetryDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxRetryDuration
	}
	return nil
}

func (x *RetryConfig) GetMinBackoffDuration() *durationpb.Duration {
	if x != nil {
		return x.MinBackoffDuration
	}
	return nil
}

func (x *RetryConfig) GetMaxBackoffDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxBackoffDuration
	}
	return nil
}

func (x *RetryConfig) GetMaxDoublings() int32 {
	if x != nil {
		return x.MaxDoublings
	}
	return 0
}

var File_google_cloud_scheduler_v1_job_proto protoreflect.FileDescriptor

var file_google_cloud_scheduler_v1_job_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x07,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0d, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x16, 0x61,
	0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x48, 0x74, 0x74, 0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x13, 0x61,
	0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0c,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x58, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41,
	0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x3a, 0x5a, 0xea, 0x41, 0x57, 0x0a, 0x21, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x62, 0x12,
	0x32, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x6a,
	0x6f, 0x62, 0x7d, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0xb6, 0x02,
	0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47,
	0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x6f, 0x66, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6d,
	0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x68, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_scheduler_v1_job_proto_rawDescOnce sync.Once
	file_google_cloud_scheduler_v1_job_proto_rawDescData = file_google_cloud_scheduler_v1_job_proto_rawDesc
)

func file_google_cloud_scheduler_v1_job_proto_rawDescGZIP() []byte {
	file_google_cloud_scheduler_v1_job_proto_rawDescOnce.Do(func() {
		file_google_cloud_scheduler_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_scheduler_v1_job_proto_rawDescData)
	})
	return file_google_cloud_scheduler_v1_job_proto_rawDescData
}

var file_google_cloud_scheduler_v1_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_scheduler_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_scheduler_v1_job_proto_goTypes = []interface{}{
	(Job_State)(0),                // 0: google.cloud.scheduler.v1.Job.State
	(*Job)(nil),                   // 1: google.cloud.scheduler.v1.Job
	(*RetryConfig)(nil),           // 2: google.cloud.scheduler.v1.RetryConfig
	(*PubsubTarget)(nil),          // 3: google.cloud.scheduler.v1.PubsubTarget
	(*AppEngineHttpTarget)(nil),   // 4: google.cloud.scheduler.v1.AppEngineHttpTarget
	(*HttpTarget)(nil),            // 5: google.cloud.scheduler.v1.HttpTarget
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*status.Status)(nil),         // 7: google.rpc.Status
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_google_cloud_scheduler_v1_job_proto_depIdxs = []int32{
	3,  // 0: google.cloud.scheduler.v1.Job.pubsub_target:type_name -> google.cloud.scheduler.v1.PubsubTarget
	4,  // 1: google.cloud.scheduler.v1.Job.app_engine_http_target:type_name -> google.cloud.scheduler.v1.AppEngineHttpTarget
	5,  // 2: google.cloud.scheduler.v1.Job.http_target:type_name -> google.cloud.scheduler.v1.HttpTarget
	6,  // 3: google.cloud.scheduler.v1.Job.user_update_time:type_name -> google.protobuf.Timestamp
	0,  // 4: google.cloud.scheduler.v1.Job.state:type_name -> google.cloud.scheduler.v1.Job.State
	7,  // 5: google.cloud.scheduler.v1.Job.status:type_name -> google.rpc.Status
	6,  // 6: google.cloud.scheduler.v1.Job.schedule_time:type_name -> google.protobuf.Timestamp
	6,  // 7: google.cloud.scheduler.v1.Job.last_attempt_time:type_name -> google.protobuf.Timestamp
	2,  // 8: google.cloud.scheduler.v1.Job.retry_config:type_name -> google.cloud.scheduler.v1.RetryConfig
	8,  // 9: google.cloud.scheduler.v1.Job.attempt_deadline:type_name -> google.protobuf.Duration
	8,  // 10: google.cloud.scheduler.v1.RetryConfig.max_retry_duration:type_name -> google.protobuf.Duration
	8,  // 11: google.cloud.scheduler.v1.RetryConfig.min_backoff_duration:type_name -> google.protobuf.Duration
	8,  // 12: google.cloud.scheduler.v1.RetryConfig.max_backoff_duration:type_name -> google.protobuf.Duration
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_cloud_scheduler_v1_job_proto_init() }
func file_google_cloud_scheduler_v1_job_proto_init() {
	if File_google_cloud_scheduler_v1_job_proto != nil {
		return
	}
	file_google_cloud_scheduler_v1_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_scheduler_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_scheduler_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_scheduler_v1_job_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Job_PubsubTarget)(nil),
		(*Job_AppEngineHttpTarget)(nil),
		(*Job_HttpTarget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_scheduler_v1_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_scheduler_v1_job_proto_goTypes,
		DependencyIndexes: file_google_cloud_scheduler_v1_job_proto_depIdxs,
		EnumInfos:         file_google_cloud_scheduler_v1_job_proto_enumTypes,
		MessageInfos:      file_google_cloud_scheduler_v1_job_proto_msgTypes,
	}.Build()
	File_google_cloud_scheduler_v1_job_proto = out.File
	file_google_cloud_scheduler_v1_job_proto_rawDesc = nil
	file_google_cloud_scheduler_v1_job_proto_goTypes = nil
	file_google_cloud_scheduler_v1_job_proto_depIdxs = nil
}
