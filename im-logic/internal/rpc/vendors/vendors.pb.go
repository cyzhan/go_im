// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: vendors.proto

package vendors

import (
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

type SensitiveWord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Word        string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	Replacement string `protobuf:"bytes,3,opt,name=replacement,proto3" json:"replacement,omitempty"`
}

func (x *SensitiveWord) Reset() {
	*x = SensitiveWord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensitiveWord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensitiveWord) ProtoMessage() {}

func (x *SensitiveWord) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensitiveWord.ProtoReflect.Descriptor instead.
func (*SensitiveWord) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{0}
}

func (x *SensitiveWord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SensitiveWord) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *SensitiveWord) GetReplacement() string {
	if x != nil {
		return x.Replacement
	}
	return ""
}

type UpdateSensitiveWordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensitiveWord *SensitiveWord `protobuf:"bytes,1,opt,name=sensitiveWord,proto3" json:"sensitiveWord,omitempty"`
	VdId          int64          `protobuf:"varint,2,opt,name=vdId,proto3" json:"vdId,omitempty"`
}

func (x *UpdateSensitiveWordReq) Reset() {
	*x = UpdateSensitiveWordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSensitiveWordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSensitiveWordReq) ProtoMessage() {}

func (x *UpdateSensitiveWordReq) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSensitiveWordReq.ProtoReflect.Descriptor instead.
func (*UpdateSensitiveWordReq) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSensitiveWordReq) GetSensitiveWord() *SensitiveWord {
	if x != nil {
		return x.SensitiveWord
	}
	return nil
}

func (x *UpdateSensitiveWordReq) GetVdId() int64 {
	if x != nil {
		return x.VdId
	}
	return 0
}

type UpdateSensitiveWordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSensitiveWordResp) Reset() {
	*x = UpdateSensitiveWordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSensitiveWordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSensitiveWordResp) ProtoMessage() {}

func (x *UpdateSensitiveWordResp) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSensitiveWordResp.ProtoReflect.Descriptor instead.
func (*UpdateSensitiveWordResp) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{2}
}

type GetSensitiveWordPageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language  string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	PageIndex int32  `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize  int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	VdId      int64  `protobuf:"varint,4,opt,name=vdId,proto3" json:"vdId,omitempty"`
}

func (x *GetSensitiveWordPageReq) Reset() {
	*x = GetSensitiveWordPageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensitiveWordPageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensitiveWordPageReq) ProtoMessage() {}

func (x *GetSensitiveWordPageReq) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensitiveWordPageReq.ProtoReflect.Descriptor instead.
func (*GetSensitiveWordPageReq) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{3}
}

func (x *GetSensitiveWordPageReq) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetSensitiveWordPageReq) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetSensitiveWordPageReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetSensitiveWordPageReq) GetVdId() int64 {
	if x != nil {
		return x.VdId
	}
	return 0
}

type GetSensitiveWordPageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensitiveWords []*SensitiveWord `protobuf:"bytes,2,rep,name=sensitiveWords,proto3" json:"sensitiveWords,omitempty"`
	TotalCount     int32            `protobuf:"varint,3,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *GetSensitiveWordPageResp) Reset() {
	*x = GetSensitiveWordPageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensitiveWordPageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensitiveWordPageResp) ProtoMessage() {}

func (x *GetSensitiveWordPageResp) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensitiveWordPageResp.ProtoReflect.Descriptor instead.
func (*GetSensitiveWordPageResp) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{4}
}

func (x *GetSensitiveWordPageResp) GetSensitiveWords() []*SensitiveWord {
	if x != nil {
		return x.SensitiveWords
	}
	return nil
}

func (x *GetSensitiveWordPageResp) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetSensitiveWordListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageId int64 `protobuf:"varint,1,opt,name=languageId,proto3" json:"languageId,omitempty"`
	VdId       int64 `protobuf:"varint,2,opt,name=vdId,proto3" json:"vdId,omitempty"`
}

func (x *GetSensitiveWordListReq) Reset() {
	*x = GetSensitiveWordListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensitiveWordListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensitiveWordListReq) ProtoMessage() {}

func (x *GetSensitiveWordListReq) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensitiveWordListReq.ProtoReflect.Descriptor instead.
func (*GetSensitiveWordListReq) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{5}
}

func (x *GetSensitiveWordListReq) GetLanguageId() int64 {
	if x != nil {
		return x.LanguageId
	}
	return 0
}

func (x *GetSensitiveWordListReq) GetVdId() int64 {
	if x != nil {
		return x.VdId
	}
	return 0
}

type GetSensitiveWordListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensitiveWords []*SensitiveWord `protobuf:"bytes,2,rep,name=sensitiveWords,proto3" json:"sensitiveWords,omitempty"`
}

func (x *GetSensitiveWordListResp) Reset() {
	*x = GetSensitiveWordListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensitiveWordListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensitiveWordListResp) ProtoMessage() {}

func (x *GetSensitiveWordListResp) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensitiveWordListResp.ProtoReflect.Descriptor instead.
func (*GetSensitiveWordListResp) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{6}
}

func (x *GetSensitiveWordListResp) GetSensitiveWords() []*SensitiveWord {
	if x != nil {
		return x.SensitiveWords
	}
	return nil
}

type GetConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VdId int64 `protobuf:"varint,1,opt,name=VdId,proto3" json:"VdId,omitempty"`
}

func (x *GetConfigReq) Reset() {
	*x = GetConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigReq) ProtoMessage() {}

func (x *GetConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigReq.ProtoReflect.Descriptor instead.
func (*GetConfigReq) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{7}
}

func (x *GetConfigReq) GetVdId() int64 {
	if x != nil {
		return x.VdId
	}
	return 0
}

type GetConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageInterval   int32   `protobuf:"varint,1,opt,name=messageInterval,proto3" json:"messageInterval,omitempty"`
	MessagePermission int32   `protobuf:"varint,2,opt,name=messagePermission,proto3" json:"messagePermission,omitempty"`
	ReportLock        int32   `protobuf:"varint,3,opt,name=reportLock,proto3" json:"reportLock,omitempty"`
	EnableLang        []*Lang `protobuf:"bytes,4,rep,name=enableLang,proto3" json:"enableLang,omitempty"`
}

func (x *GetConfigResp) Reset() {
	*x = GetConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResp) ProtoMessage() {}

func (x *GetConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResp.ProtoReflect.Descriptor instead.
func (*GetConfigResp) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{8}
}

func (x *GetConfigResp) GetMessageInterval() int32 {
	if x != nil {
		return x.MessageInterval
	}
	return 0
}

func (x *GetConfigResp) GetMessagePermission() int32 {
	if x != nil {
		return x.MessagePermission
	}
	return 0
}

func (x *GetConfigResp) GetReportLock() int32 {
	if x != nil {
		return x.ReportLock
	}
	return 0
}

func (x *GetConfigResp) GetEnableLang() []*Lang {
	if x != nil {
		return x.EnableLang
	}
	return nil
}

type Lang struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Lang) Reset() {
	*x = Lang{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendors_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lang) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lang) ProtoMessage() {}

func (x *Lang) ProtoReflect() protoreflect.Message {
	mi := &file_vendors_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lang.ProtoReflect.Descriptor instead.
func (*Lang) Descriptor() ([]byte, []int) {
	return file_vendors_proto_rawDescGZIP(), []int{9}
}

func (x *Lang) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lang) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lang) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_vendors_proto protoreflect.FileDescriptor

var file_vendors_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x22, 0x55, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x6a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x64, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x64, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x64, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x76, 0x64, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x56, 0x64, 0x49, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63,
	0x6b, 0x12, 0x2d, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x6e, 0x67,
	0x22, 0x3e, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0xe8, 0x02, 0x0a, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x57, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5d,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x15, 0x2e, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vendors_proto_rawDescOnce sync.Once
	file_vendors_proto_rawDescData = file_vendors_proto_rawDesc
)

func file_vendors_proto_rawDescGZIP() []byte {
	file_vendors_proto_rawDescOnce.Do(func() {
		file_vendors_proto_rawDescData = protoimpl.X.CompressGZIP(file_vendors_proto_rawDescData)
	})
	return file_vendors_proto_rawDescData
}

var file_vendors_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_vendors_proto_goTypes = []interface{}{
	(*SensitiveWord)(nil),            // 0: vendors.SensitiveWord
	(*UpdateSensitiveWordReq)(nil),   // 1: vendors.UpdateSensitiveWordReq
	(*UpdateSensitiveWordResp)(nil),  // 2: vendors.UpdateSensitiveWordResp
	(*GetSensitiveWordPageReq)(nil),  // 3: vendors.GetSensitiveWordPageReq
	(*GetSensitiveWordPageResp)(nil), // 4: vendors.GetSensitiveWordPageResp
	(*GetSensitiveWordListReq)(nil),  // 5: vendors.GetSensitiveWordListReq
	(*GetSensitiveWordListResp)(nil), // 6: vendors.GetSensitiveWordListResp
	(*GetConfigReq)(nil),             // 7: vendors.GetConfigReq
	(*GetConfigResp)(nil),            // 8: vendors.GetConfigResp
	(*Lang)(nil),                     // 9: vendors.Lang
}
var file_vendors_proto_depIdxs = []int32{
	0, // 0: vendors.UpdateSensitiveWordReq.sensitiveWord:type_name -> vendors.SensitiveWord
	0, // 1: vendors.GetSensitiveWordPageResp.sensitiveWords:type_name -> vendors.SensitiveWord
	0, // 2: vendors.GetSensitiveWordListResp.sensitiveWords:type_name -> vendors.SensitiveWord
	9, // 3: vendors.GetConfigResp.enableLang:type_name -> vendors.Lang
	1, // 4: vendors.vendorsService.UpdateSensitiveWord:input_type -> vendors.UpdateSensitiveWordReq
	3, // 5: vendors.vendorsService.GetSensitiveWordPage:input_type -> vendors.GetSensitiveWordPageReq
	5, // 6: vendors.vendorsService.GetSensitiveWordList:input_type -> vendors.GetSensitiveWordListReq
	7, // 7: vendors.vendorsService.GetConfig:input_type -> vendors.GetConfigReq
	2, // 8: vendors.vendorsService.UpdateSensitiveWord:output_type -> vendors.UpdateSensitiveWordResp
	4, // 9: vendors.vendorsService.GetSensitiveWordPage:output_type -> vendors.GetSensitiveWordPageResp
	6, // 10: vendors.vendorsService.GetSensitiveWordList:output_type -> vendors.GetSensitiveWordListResp
	8, // 11: vendors.vendorsService.GetConfig:output_type -> vendors.GetConfigResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vendors_proto_init() }
func file_vendors_proto_init() {
	if File_vendors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vendors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensitiveWord); i {
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
		file_vendors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSensitiveWordReq); i {
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
		file_vendors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSensitiveWordResp); i {
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
		file_vendors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensitiveWordPageReq); i {
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
		file_vendors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensitiveWordPageResp); i {
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
		file_vendors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensitiveWordListReq); i {
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
		file_vendors_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensitiveWordListResp); i {
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
		file_vendors_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigReq); i {
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
		file_vendors_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResp); i {
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
		file_vendors_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lang); i {
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
			RawDescriptor: file_vendors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vendors_proto_goTypes,
		DependencyIndexes: file_vendors_proto_depIdxs,
		MessageInfos:      file_vendors_proto_msgTypes,
	}.Build()
	File_vendors_proto = out.File
	file_vendors_proto_rawDesc = nil
	file_vendors_proto_goTypes = nil
	file_vendors_proto_depIdxs = nil
}
