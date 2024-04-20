// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tasks/v1/tasks.proto

package tasksv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTasksRequest) Reset() {
	*x = GetAllTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksRequest) ProtoMessage() {}

func (x *GetAllTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksRequest.ProtoReflect.Descriptor instead.
func (*GetAllTasksRequest) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{1}
}

type GetAllTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetAllTasksResponse) Reset() {
	*x = GetAllTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksResponse) ProtoMessage() {}

func (x *GetAllTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksResponse.ProtoReflect.Descriptor instead.
func (*GetAllTasksResponse) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{4}
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{6}
}

type GetOneTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOneTaskRequest) Reset() {
	*x = GetOneTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneTaskRequest) ProtoMessage() {}

func (x *GetOneTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneTaskRequest.ProtoReflect.Descriptor instead.
func (*GetOneTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *GetOneTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOneTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *GetOneTaskResponse) Reset() {
	*x = GetOneTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneTaskResponse) ProtoMessage() {}

func (x *GetOneTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneTaskResponse.ProtoReflect.Descriptor instead.
func (*GetOneTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{8}
}

func (x *GetOneTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_v1_tasks_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_v1_tasks_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_v1_tasks_proto_rawDescGZIP(), []int{10}
}

var File_tasks_v1_tasks_proto protoreflect.FileDescriptor

var file_tasks_v1_tasks_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x5b,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xe2, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x5a, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01,
	0x2a, 0x22, 0x06, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x5c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x1a, 0x0b, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x97, 0x02, 0x92, 0x41, 0x6c, 0x12, 0x64, 0x0a, 0x08,
	0x54, 0x61, 0x73, 0x6b, 0x20, 0x41, 0x50, 0x49, 0x12, 0x1c, 0x41, 0x20, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x33, 0x0a, 0x0e, 0x44, 0x75, 0x63, 0x6b, 0x79, 0x4d,
	0x6f, 0x6d, 0x6f, 0x32, 0x30, 0x30, 0x31, 0x32, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75, 0x63,
	0x6b, 0x79, 0x4d, 0x6f, 0x6d, 0x6f, 0x32, 0x30, 0x30, 0x31, 0x32, 0x32, 0x05, 0x31, 0x2e, 0x30,
	0x2e, 0x30, 0x22, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x75, 0x63, 0x6b, 0x79, 0x4d, 0x6f, 0x6d, 0x6f, 0x32, 0x30, 0x30, 0x31, 0x32, 0x2f,
	0x67, 0x6f, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tasks_v1_tasks_proto_rawDescOnce sync.Once
	file_tasks_v1_tasks_proto_rawDescData = file_tasks_v1_tasks_proto_rawDesc
)

func file_tasks_v1_tasks_proto_rawDescGZIP() []byte {
	file_tasks_v1_tasks_proto_rawDescOnce.Do(func() {
		file_tasks_v1_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasks_v1_tasks_proto_rawDescData)
	})
	return file_tasks_v1_tasks_proto_rawDescData
}

var file_tasks_v1_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tasks_v1_tasks_proto_goTypes = []interface{}{
	(*Task)(nil),                // 0: tasks.v1.Task
	(*GetAllTasksRequest)(nil),  // 1: tasks.v1.GetAllTasksRequest
	(*GetAllTasksResponse)(nil), // 2: tasks.v1.GetAllTasksResponse
	(*CreateTaskRequest)(nil),   // 3: tasks.v1.CreateTaskRequest
	(*CreateTaskResponse)(nil),  // 4: tasks.v1.CreateTaskResponse
	(*DeleteTaskRequest)(nil),   // 5: tasks.v1.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),  // 6: tasks.v1.DeleteTaskResponse
	(*GetOneTaskRequest)(nil),   // 7: tasks.v1.GetOneTaskRequest
	(*GetOneTaskResponse)(nil),  // 8: tasks.v1.GetOneTaskResponse
	(*UpdateTaskRequest)(nil),   // 9: tasks.v1.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),  // 10: tasks.v1.UpdateTaskResponse
}
var file_tasks_v1_tasks_proto_depIdxs = []int32{
	0,  // 0: tasks.v1.GetAllTasksResponse.tasks:type_name -> tasks.v1.Task
	0,  // 1: tasks.v1.GetOneTaskResponse.task:type_name -> tasks.v1.Task
	1,  // 2: tasks.v1.TaskService.GetAllTasks:input_type -> tasks.v1.GetAllTasksRequest
	3,  // 3: tasks.v1.TaskService.CreateTask:input_type -> tasks.v1.CreateTaskRequest
	5,  // 4: tasks.v1.TaskService.DeleteTask:input_type -> tasks.v1.DeleteTaskRequest
	7,  // 5: tasks.v1.TaskService.GetOneTask:input_type -> tasks.v1.GetOneTaskRequest
	9,  // 6: tasks.v1.TaskService.UpdateTask:input_type -> tasks.v1.UpdateTaskRequest
	2,  // 7: tasks.v1.TaskService.GetAllTasks:output_type -> tasks.v1.GetAllTasksResponse
	4,  // 8: tasks.v1.TaskService.CreateTask:output_type -> tasks.v1.CreateTaskResponse
	6,  // 9: tasks.v1.TaskService.DeleteTask:output_type -> tasks.v1.DeleteTaskResponse
	8,  // 10: tasks.v1.TaskService.GetOneTask:output_type -> tasks.v1.GetOneTaskResponse
	10, // 11: tasks.v1.TaskService.UpdateTask:output_type -> tasks.v1.UpdateTaskResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_tasks_v1_tasks_proto_init() }
func file_tasks_v1_tasks_proto_init() {
	if File_tasks_v1_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tasks_v1_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_tasks_v1_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTasksRequest); i {
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
		file_tasks_v1_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTasksResponse); i {
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
		file_tasks_v1_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_tasks_v1_tasks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_tasks_v1_tasks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_tasks_v1_tasks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskResponse); i {
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
		file_tasks_v1_tasks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneTaskRequest); i {
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
		file_tasks_v1_tasks_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneTaskResponse); i {
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
		file_tasks_v1_tasks_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_tasks_v1_tasks_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tasks_v1_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasks_v1_tasks_proto_goTypes,
		DependencyIndexes: file_tasks_v1_tasks_proto_depIdxs,
		MessageInfos:      file_tasks_v1_tasks_proto_msgTypes,
	}.Build()
	File_tasks_v1_tasks_proto = out.File
	file_tasks_v1_tasks_proto_rawDesc = nil
	file_tasks_v1_tasks_proto_goTypes = nil
	file_tasks_v1_tasks_proto_depIdxs = nil
}
