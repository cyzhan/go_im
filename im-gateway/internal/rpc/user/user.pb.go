// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "imgateway/internal/rpc/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VdID     int64  `protobuf:"varint,1,opt,name=vdID,proto3" json:"vdID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Vip      int32  `protobuf:"varint,4,opt,name=vip,proto3" json:"vip,omitempty"`
	Avatar   int32  `protobuf:"varint,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Language string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginData) GetVdID() int64 {
	if x != nil {
		return x.VdID
	}
	return 0
}

func (x *LoginData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginData) GetVip() int32 {
	if x != nil {
		return x.Vip
	}
	return 0
}

func (x *LoginData) GetAvatar() int32 {
	if x != nil {
		return x.Avatar
	}
	return 0
}

func (x *LoginData) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type LoginResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserID int64         `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Link   string        `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *LoginResult) Reset() {
	*x = LoginResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResult) ProtoMessage() {}

func (x *LoginResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResult.ProtoReflect.Descriptor instead.
func (*LoginResult) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResult) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LoginResult) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LoginResult) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetUserListArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size     int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	VendorID int64 `protobuf:"varint,4,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
}

func (x *GetUserListArgs) Reset() {
	*x = GetUserListArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListArgs) ProtoMessage() {}

func (x *GetUserListArgs) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListArgs.ProtoReflect.Descriptor instead.
func (*GetUserListArgs) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserListArgs) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetUserListArgs) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserListArgs) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetUserListArgs) GetVendorID() int64 {
	if x != nil {
		return x.VendorID
	}
	return 0
}

type UserDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AccountType int32  `protobuf:"varint,3,opt,name=accountType,proto3" json:"accountType,omitempty"`
	Date        string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Report      int32  `protobuf:"varint,5,opt,name=report,proto3" json:"report,omitempty"`
	Status      int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Remark      string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UserDTO) Reset() {
	*x = UserDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDTO) ProtoMessage() {}

func (x *UserDTO) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDTO.ProtoReflect.Descriptor instead.
func (*UserDTO) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDTO) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *UserDTO) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *UserDTO) GetReport() int32 {
	if x != nil {
		return x.Report
	}
	return 0
}

func (x *UserDTO) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserDTO) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type GetUserListResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Items      []*UserDTO    `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserListResult) Reset() {
	*x = GetUserListResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResult) ProtoMessage() {}

func (x *GetUserListResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResult.ProtoReflect.Descriptor instead.
func (*GetUserListResult) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListResult) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetUserListResult) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetUserListResult) GetItems() []*UserDTO {
	if x != nil {
		return x.Items
	}
	return nil
}

type PatchUserArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorID int64  `protobuf:"varint,1,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	Id       int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Remark   string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *PatchUserArgs) Reset() {
	*x = PatchUserArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUserArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUserArgs) ProtoMessage() {}

func (x *PatchUserArgs) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUserArgs.ProtoReflect.Descriptor instead.
func (*PatchUserArgs) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *PatchUserArgs) GetVendorID() int64 {
	if x != nil {
		return x.VendorID
	}
	return 0
}

func (x *PatchUserArgs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchUserArgs) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PatchUserArgs) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UserStatusChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserStatusChange) Reset() {
	*x = UserStatusChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatusChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatusChange) ProtoMessage() {}

func (x *UserStatusChange) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatusChange.ProtoReflect.Descriptor instead.
func (*UserStatusChange) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserStatusChange) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserStatusChange) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ReportAbuseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ReportId int64  `protobuf:"varint,2,opt,name=reportId,proto3" json:"reportId,omitempty"`
	Reason   string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ReportAbuseReq) Reset() {
	*x = ReportAbuseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportAbuseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAbuseReq) ProtoMessage() {}

func (x *ReportAbuseReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAbuseReq.ProtoReflect.Descriptor instead.
func (*ReportAbuseReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *ReportAbuseReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ReportAbuseReq) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

func (x *ReportAbuseReq) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x76, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x64,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x76, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x6d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x44, 0x22, 0xab, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x6b, 0x0a, 0x0d, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x42, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x41,
	0x62, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x32, 0xe6, 0x01, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x62, 0x75, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x41, 0x62, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(*LoginData)(nil),         // 0: user.LoginData
	(*LoginResult)(nil),       // 1: user.LoginResult
	(*GetUserListArgs)(nil),   // 2: user.GetUserListArgs
	(*UserDTO)(nil),           // 3: user.UserDTO
	(*GetUserListResult)(nil), // 4: user.GetUserListResult
	(*PatchUserArgs)(nil),     // 5: user.PatchUserArgs
	(*UserStatusChange)(nil),  // 6: user.UserStatusChange
	(*ReportAbuseReq)(nil),    // 7: user.ReportAbuseReq
	(*common.Error)(nil),      // 8: common.Error
}
var file_user_proto_depIdxs = []int32{
	8, // 0: user.LoginResult.error:type_name -> common.Error
	8, // 1: user.GetUserListResult.error:type_name -> common.Error
	3, // 2: user.GetUserListResult.items:type_name -> user.UserDTO
	0, // 3: user.userService.Login:input_type -> user.LoginData
	2, // 4: user.userService.GetUserList:input_type -> user.GetUserListArgs
	5, // 5: user.userService.PatchUser:input_type -> user.PatchUserArgs
	7, // 6: user.userService.ReportAbuse:input_type -> user.ReportAbuseReq
	1, // 7: user.userService.Login:output_type -> user.LoginResult
	4, // 8: user.userService.GetUserList:output_type -> user.GetUserListResult
	8, // 9: user.userService.PatchUser:output_type -> common.Error
	8, // 10: user.userService.ReportAbuse:output_type -> common.Error
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResult); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListArgs); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDTO); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResult); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchUserArgs); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatusChange); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportAbuseReq); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
