// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: filesystem.proto

package proto

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FileAction int32

const (
	FileAction_FILE_ACTION_UNKNOWN FileAction = 0
	FileAction_FILE_ACTION_CREATE  FileAction = 1
	FileAction_FILE_ACTION_WRITE   FileAction = 2
	FileAction_FILE_ACTION_REMOVE  FileAction = 3
	FileAction_FILE_ACTION_RENAME  FileAction = 4
	FileAction_FILE_ACTION_CHMOD   FileAction = 5
)

// Enum value maps for FileAction.
var (
	FileAction_name = map[int32]string{
		0: "FILE_ACTION_UNKNOWN",
		1: "FILE_ACTION_CREATE",
		2: "FILE_ACTION_WRITE",
		3: "FILE_ACTION_REMOVE",
		4: "FILE_ACTION_RENAME",
		5: "FILE_ACTION_CHMOD",
	}
	FileAction_value = map[string]int32{
		"FILE_ACTION_UNKNOWN": 0,
		"FILE_ACTION_CREATE":  1,
		"FILE_ACTION_WRITE":   2,
		"FILE_ACTION_REMOVE":  3,
		"FILE_ACTION_RENAME":  4,
		"FILE_ACTION_CHMOD":   5,
	}
)

func (x FileAction) Enum() *FileAction {
	p := new(FileAction)
	*p = x
	return p
}

func (x FileAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileAction) Descriptor() protoreflect.EnumDescriptor {
	return file_filesystem_proto_enumTypes[0].Descriptor()
}

func (FileAction) Type() protoreflect.EnumType {
	return &file_filesystem_proto_enumTypes[0]
}

func (x FileAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileAction.Descriptor instead.
func (FileAction) EnumDescriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{0}
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size    int64                `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mode    uint32               `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Updated *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	IsDir   bool                 `protobuf:"varint,5,opt,name=isDir,proto3" json:"isDir,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfo) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *FileInfo) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *FileInfo) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

type FilesystemDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *FilesystemDirRequest) Reset() {
	*x = FilesystemDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemDirRequest) ProtoMessage() {}

func (x *FilesystemDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemDirRequest.ProtoReflect.Descriptor instead.
func (*FilesystemDirRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{1}
}

func (x *FilesystemDirRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type FilesystemDirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FilesystemDirResponse) Reset() {
	*x = FilesystemDirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemDirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemDirResponse) ProtoMessage() {}

func (x *FilesystemDirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemDirResponse.ProtoReflect.Descriptor instead.
func (*FilesystemDirResponse) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{2}
}

func (x *FilesystemDirResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type FilesystemDirStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *FilesystemDirStreamRequest) Reset() {
	*x = FilesystemDirStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemDirStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemDirStreamRequest) ProtoMessage() {}

func (x *FilesystemDirStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemDirStreamRequest.ProtoReflect.Descriptor instead.
func (*FilesystemDirStreamRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{3}
}

func (x *FilesystemDirStreamRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type FilesystemDirStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File   *FileInfo  `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Action FileAction `protobuf:"varint,2,opt,name=action,proto3,enum=proto.FileAction" json:"action,omitempty"`
	Error  string     `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FilesystemDirStreamResponse) Reset() {
	*x = FilesystemDirStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemDirStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemDirStreamResponse) ProtoMessage() {}

func (x *FilesystemDirStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemDirStreamResponse.ProtoReflect.Descriptor instead.
func (*FilesystemDirStreamResponse) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{4}
}

func (x *FilesystemDirStreamResponse) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FilesystemDirStreamResponse) GetAction() FileAction {
	if x != nil {
		return x.Action
	}
	return FileAction_FILE_ACTION_UNKNOWN
}

func (x *FilesystemDirStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FilesystemFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FilesystemFileRequest) Reset() {
	*x = FilesystemFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemFileRequest) ProtoMessage() {}

func (x *FilesystemFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemFileRequest.ProtoReflect.Descriptor instead.
func (*FilesystemFileRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{5}
}

func (x *FilesystemFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FilesystemFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *FileInfo `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FilesystemFileResponse) Reset() {
	*x = FilesystemFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemFileResponse) ProtoMessage() {}

func (x *FilesystemFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemFileResponse.ProtoReflect.Descriptor instead.
func (*FilesystemFileResponse) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{6}
}

func (x *FilesystemFileResponse) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

type FilesystemFileStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FilesystemFileStreamRequest) Reset() {
	*x = FilesystemFileStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemFileStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemFileStreamRequest) ProtoMessage() {}

func (x *FilesystemFileStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemFileStreamRequest.ProtoReflect.Descriptor instead.
func (*FilesystemFileStreamRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{7}
}

func (x *FilesystemFileStreamRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FilesystemFileStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File   *FileInfo  `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Action FileAction `protobuf:"varint,2,opt,name=action,proto3,enum=proto.FileAction" json:"action,omitempty"`
	Error  string     `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FilesystemFileStreamResponse) Reset() {
	*x = FilesystemFileStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemFileStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemFileStreamResponse) ProtoMessage() {}

func (x *FilesystemFileStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemFileStreamResponse.ProtoReflect.Descriptor instead.
func (*FilesystemFileStreamResponse) Descriptor() ([]byte, []int) {
	return file_filesystem_proto_rawDescGZIP(), []int{8}
}

func (x *FilesystemFileStreamResponse) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FilesystemFileStreamResponse) GetAction() FileAction {
	if x != nil {
		return x.Action
	}
	return FileAction_FILE_ACTION_UNKNOWN
}

func (x *FilesystemFileStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_filesystem_proto protoreflect.FileDescriptor

var file_filesystem_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44,
	0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x22,
	0x34, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3e, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x1a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x22, 0x83, 0x01, 0x0a, 0x1b, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x3d, 0x0a, 0x16, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x22, 0x31, 0x0a, 0x1b, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x84, 0x01, 0x0a, 0x1c, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x9b, 0x01,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x4d, 0x4f, 0x44, 0x10, 0x05, 0x32, 0xea, 0x02, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x4a, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filesystem_proto_rawDescOnce sync.Once
	file_filesystem_proto_rawDescData = file_filesystem_proto_rawDesc
)

func file_filesystem_proto_rawDescGZIP() []byte {
	file_filesystem_proto_rawDescOnce.Do(func() {
		file_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_filesystem_proto_rawDescData)
	})
	return file_filesystem_proto_rawDescData
}

var file_filesystem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_filesystem_proto_goTypes = []interface{}{
	(FileAction)(0),                      // 0: proto.FileAction
	(*FileInfo)(nil),                     // 1: proto.FileInfo
	(*FilesystemDirRequest)(nil),         // 2: proto.FilesystemDirRequest
	(*FilesystemDirResponse)(nil),        // 3: proto.FilesystemDirResponse
	(*FilesystemDirStreamRequest)(nil),   // 4: proto.FilesystemDirStreamRequest
	(*FilesystemDirStreamResponse)(nil),  // 5: proto.FilesystemDirStreamResponse
	(*FilesystemFileRequest)(nil),        // 6: proto.FilesystemFileRequest
	(*FilesystemFileResponse)(nil),       // 7: proto.FilesystemFileResponse
	(*FilesystemFileStreamRequest)(nil),  // 8: proto.FilesystemFileStreamRequest
	(*FilesystemFileStreamResponse)(nil), // 9: proto.FilesystemFileStreamResponse
	(*timestamp.Timestamp)(nil),          // 10: google.protobuf.Timestamp
}
var file_filesystem_proto_depIdxs = []int32{
	10, // 0: proto.FileInfo.updated:type_name -> google.protobuf.Timestamp
	1,  // 1: proto.FilesystemDirResponse.files:type_name -> proto.FileInfo
	1,  // 2: proto.FilesystemDirStreamResponse.file:type_name -> proto.FileInfo
	0,  // 3: proto.FilesystemDirStreamResponse.action:type_name -> proto.FileAction
	1,  // 4: proto.FilesystemFileResponse.file:type_name -> proto.FileInfo
	1,  // 5: proto.FilesystemFileStreamResponse.file:type_name -> proto.FileInfo
	0,  // 6: proto.FilesystemFileStreamResponse.action:type_name -> proto.FileAction
	2,  // 7: proto.Filesystem.FilesystemDir:input_type -> proto.FilesystemDirRequest
	4,  // 8: proto.Filesystem.FilesystemDirStream:input_type -> proto.FilesystemDirStreamRequest
	6,  // 9: proto.Filesystem.FilesystemFile:input_type -> proto.FilesystemFileRequest
	8,  // 10: proto.Filesystem.FilesystemFileStream:input_type -> proto.FilesystemFileStreamRequest
	3,  // 11: proto.Filesystem.FilesystemDir:output_type -> proto.FilesystemDirResponse
	5,  // 12: proto.Filesystem.FilesystemDirStream:output_type -> proto.FilesystemDirStreamResponse
	7,  // 13: proto.Filesystem.FilesystemFile:output_type -> proto.FilesystemFileResponse
	9,  // 14: proto.Filesystem.FilesystemFileStream:output_type -> proto.FilesystemFileStreamResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_filesystem_proto_init() }
func file_filesystem_proto_init() {
	if File_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_filesystem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemDirRequest); i {
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
		file_filesystem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemDirResponse); i {
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
		file_filesystem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemDirStreamRequest); i {
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
		file_filesystem_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemDirStreamResponse); i {
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
		file_filesystem_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemFileRequest); i {
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
		file_filesystem_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemFileResponse); i {
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
		file_filesystem_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemFileStreamRequest); i {
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
		file_filesystem_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemFileStreamResponse); i {
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
			RawDescriptor: file_filesystem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filesystem_proto_goTypes,
		DependencyIndexes: file_filesystem_proto_depIdxs,
		EnumInfos:         file_filesystem_proto_enumTypes,
		MessageInfos:      file_filesystem_proto_msgTypes,
	}.Build()
	File_filesystem_proto = out.File
	file_filesystem_proto_rawDesc = nil
	file_filesystem_proto_goTypes = nil
	file_filesystem_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FilesystemClient is the client API for Filesystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilesystemClient interface {
	// list the contents of a directory
	FilesystemDir(ctx context.Context, in *FilesystemDirRequest, opts ...grpc.CallOption) (*FilesystemDirResponse, error)
	// stream any updates to the contents of a directory
	FilesystemDirStream(ctx context.Context, in *FilesystemDirStreamRequest, opts ...grpc.CallOption) (Filesystem_FilesystemDirStreamClient, error)
	// get information about a file
	FilesystemFile(ctx context.Context, in *FilesystemFileRequest, opts ...grpc.CallOption) (*FilesystemFileResponse, error)
	// stream any updates to a file
	FilesystemFileStream(ctx context.Context, in *FilesystemFileStreamRequest, opts ...grpc.CallOption) (Filesystem_FilesystemFileStreamClient, error)
}

type filesystemClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesystemClient(cc grpc.ClientConnInterface) FilesystemClient {
	return &filesystemClient{cc}
}

func (c *filesystemClient) FilesystemDir(ctx context.Context, in *FilesystemDirRequest, opts ...grpc.CallOption) (*FilesystemDirResponse, error) {
	out := new(FilesystemDirResponse)
	err := c.cc.Invoke(ctx, "/proto.Filesystem/FilesystemDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemClient) FilesystemDirStream(ctx context.Context, in *FilesystemDirStreamRequest, opts ...grpc.CallOption) (Filesystem_FilesystemDirStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Filesystem_serviceDesc.Streams[0], "/proto.Filesystem/FilesystemDirStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesystemFilesystemDirStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Filesystem_FilesystemDirStreamClient interface {
	Recv() (*FilesystemDirStreamResponse, error)
	grpc.ClientStream
}

type filesystemFilesystemDirStreamClient struct {
	grpc.ClientStream
}

func (x *filesystemFilesystemDirStreamClient) Recv() (*FilesystemDirStreamResponse, error) {
	m := new(FilesystemDirStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filesystemClient) FilesystemFile(ctx context.Context, in *FilesystemFileRequest, opts ...grpc.CallOption) (*FilesystemFileResponse, error) {
	out := new(FilesystemFileResponse)
	err := c.cc.Invoke(ctx, "/proto.Filesystem/FilesystemFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemClient) FilesystemFileStream(ctx context.Context, in *FilesystemFileStreamRequest, opts ...grpc.CallOption) (Filesystem_FilesystemFileStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Filesystem_serviceDesc.Streams[1], "/proto.Filesystem/FilesystemFileStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesystemFilesystemFileStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Filesystem_FilesystemFileStreamClient interface {
	Recv() (*FilesystemFileStreamResponse, error)
	grpc.ClientStream
}

type filesystemFilesystemFileStreamClient struct {
	grpc.ClientStream
}

func (x *filesystemFilesystemFileStreamClient) Recv() (*FilesystemFileStreamResponse, error) {
	m := new(FilesystemFileStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilesystemServer is the server API for Filesystem service.
type FilesystemServer interface {
	// list the contents of a directory
	FilesystemDir(context.Context, *FilesystemDirRequest) (*FilesystemDirResponse, error)
	// stream any updates to the contents of a directory
	FilesystemDirStream(*FilesystemDirStreamRequest, Filesystem_FilesystemDirStreamServer) error
	// get information about a file
	FilesystemFile(context.Context, *FilesystemFileRequest) (*FilesystemFileResponse, error)
	// stream any updates to a file
	FilesystemFileStream(*FilesystemFileStreamRequest, Filesystem_FilesystemFileStreamServer) error
}

// UnimplementedFilesystemServer can be embedded to have forward compatible implementations.
type UnimplementedFilesystemServer struct {
}

func (*UnimplementedFilesystemServer) FilesystemDir(context.Context, *FilesystemDirRequest) (*FilesystemDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilesystemDir not implemented")
}
func (*UnimplementedFilesystemServer) FilesystemDirStream(*FilesystemDirStreamRequest, Filesystem_FilesystemDirStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FilesystemDirStream not implemented")
}
func (*UnimplementedFilesystemServer) FilesystemFile(context.Context, *FilesystemFileRequest) (*FilesystemFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilesystemFile not implemented")
}
func (*UnimplementedFilesystemServer) FilesystemFileStream(*FilesystemFileStreamRequest, Filesystem_FilesystemFileStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FilesystemFileStream not implemented")
}

func RegisterFilesystemServer(s *grpc.Server, srv FilesystemServer) {
	s.RegisterService(&_Filesystem_serviceDesc, srv)
}

func _Filesystem_FilesystemDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesystemDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).FilesystemDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filesystem/FilesystemDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).FilesystemDir(ctx, req.(*FilesystemDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filesystem_FilesystemDirStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilesystemDirStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilesystemServer).FilesystemDirStream(m, &filesystemFilesystemDirStreamServer{stream})
}

type Filesystem_FilesystemDirStreamServer interface {
	Send(*FilesystemDirStreamResponse) error
	grpc.ServerStream
}

type filesystemFilesystemDirStreamServer struct {
	grpc.ServerStream
}

func (x *filesystemFilesystemDirStreamServer) Send(m *FilesystemDirStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Filesystem_FilesystemFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesystemFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).FilesystemFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filesystem/FilesystemFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).FilesystemFile(ctx, req.(*FilesystemFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filesystem_FilesystemFileStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilesystemFileStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilesystemServer).FilesystemFileStream(m, &filesystemFilesystemFileStreamServer{stream})
}

type Filesystem_FilesystemFileStreamServer interface {
	Send(*FilesystemFileStreamResponse) error
	grpc.ServerStream
}

type filesystemFilesystemFileStreamServer struct {
	grpc.ServerStream
}

func (x *filesystemFilesystemFileStreamServer) Send(m *FilesystemFileStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Filesystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Filesystem",
	HandlerType: (*FilesystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilesystemDir",
			Handler:    _Filesystem_FilesystemDir_Handler,
		},
		{
			MethodName: "FilesystemFile",
			Handler:    _Filesystem_FilesystemFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FilesystemDirStream",
			Handler:       _Filesystem_FilesystemDirStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FilesystemFileStream",
			Handler:       _Filesystem_FilesystemFileStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "filesystem.proto",
}
