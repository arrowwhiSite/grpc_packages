// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: database.proto

package db

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Определение сообщения для клиента
type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login        string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	FirstName    string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Password     string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreateDate   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	ProfilePhoto string                 `protobuf:"bytes,6,opt,name=profile_photo,json=profilePhoto,proto3" json:"profile_photo,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Client) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Client) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Client) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Client) GetCreateDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *Client) GetProfilePhoto() string {
	if x != nil {
		return x.ProfilePhoto
	}
	return ""
}

// Запросы для методов сервиса
type SelectClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *SelectClientRequest) Reset() {
	*x = SelectClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectClientRequest) ProtoMessage() {}

func (x *SelectClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectClientRequest.ProtoReflect.Descriptor instead.
func (*SelectClientRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{1}
}

func (x *SelectClientRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type ChangeProfilePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ChangeProfilePhotoRequest) Reset() {
	*x = ChangeProfilePhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeProfilePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeProfilePhotoRequest) ProtoMessage() {}

func (x *ChangeProfilePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeProfilePhotoRequest.ProtoReflect.Descriptor instead.
func (*ChangeProfilePhotoRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeProfilePhotoRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ChangeProfilePhotoRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type UpdateClientPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login       string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *UpdateClientPasswordRequest) Reset() {
	*x = UpdateClientPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientPasswordRequest) ProtoMessage() {}

func (x *UpdateClientPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientPasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientPasswordRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateClientPasswordRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateClientPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type AddClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Login     string `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AddClientRequest) Reset() {
	*x = AddClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientRequest) ProtoMessage() {}

func (x *AddClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientRequest.ProtoReflect.Descriptor instead.
func (*AddClientRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{4}
}

func (x *AddClientRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AddClientRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AddClientRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AddClientRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetLoginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoginsRequest) Reset() {
	*x = GetLoginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginsRequest) ProtoMessage() {}

func (x *GetLoginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginsRequest.ProtoReflect.Descriptor instead.
func (*GetLoginsRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{5}
}

type GetLoginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logins []string `protobuf:"bytes,1,rep,name=logins,proto3" json:"logins,omitempty"`
}

func (x *GetLoginsResponse) Reset() {
	*x = GetLoginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginsResponse) ProtoMessage() {}

func (x *GetLoginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginsResponse.ProtoReflect.Descriptor instead.
func (*GetLoginsResponse) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{6}
}

func (x *GetLoginsResponse) GetLogins() []string {
	if x != nil {
		return x.Logins
	}
	return nil
}

var File_database_proto protoreflect.FileDescriptor

var file_database_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x45, 0x0a, 0x19,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x56, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x32,
	0xe2, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x4d, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x51, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_proto_rawDescOnce sync.Once
	file_database_proto_rawDescData = file_database_proto_rawDesc
)

func file_database_proto_rawDescGZIP() []byte {
	file_database_proto_rawDescOnce.Do(func() {
		file_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_proto_rawDescData)
	})
	return file_database_proto_rawDescData
}

var file_database_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_database_proto_goTypes = []interface{}{
	(*Client)(nil),                      // 0: chat.Client
	(*SelectClientRequest)(nil),         // 1: chat.SelectClientRequest
	(*ChangeProfilePhotoRequest)(nil),   // 2: chat.ChangeProfilePhotoRequest
	(*UpdateClientPasswordRequest)(nil), // 3: chat.UpdateClientPasswordRequest
	(*AddClientRequest)(nil),            // 4: chat.AddClientRequest
	(*GetLoginsRequest)(nil),            // 5: chat.GetLoginsRequest
	(*GetLoginsResponse)(nil),           // 6: chat.GetLoginsResponse
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_database_proto_depIdxs = []int32{
	7, // 0: chat.Client.create_date:type_name -> google.protobuf.Timestamp
	1, // 1: chat.ClientService.SelectClientByLogin:input_type -> chat.SelectClientRequest
	2, // 2: chat.ClientService.ChangeProfilePhoto:input_type -> chat.ChangeProfilePhotoRequest
	3, // 3: chat.ClientService.UpdateClientPassword:input_type -> chat.UpdateClientPasswordRequest
	4, // 4: chat.ClientService.AddClient:input_type -> chat.AddClientRequest
	5, // 5: chat.ClientService.GetLogins:input_type -> chat.GetLoginsRequest
	0, // 6: chat.ClientService.SelectClientByLogin:output_type -> chat.Client
	8, // 7: chat.ClientService.ChangeProfilePhoto:output_type -> google.protobuf.Empty
	8, // 8: chat.ClientService.UpdateClientPassword:output_type -> google.protobuf.Empty
	0, // 9: chat.ClientService.AddClient:output_type -> chat.Client
	6, // 10: chat.ClientService.GetLogins:output_type -> chat.GetLoginsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_database_proto_init() }
func file_database_proto_init() {
	if File_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectClientRequest); i {
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
		file_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeProfilePhotoRequest); i {
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
		file_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientPasswordRequest); i {
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
		file_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientRequest); i {
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
		file_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginsRequest); i {
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
		file_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginsResponse); i {
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
			RawDescriptor: file_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_database_proto_goTypes,
		DependencyIndexes: file_database_proto_depIdxs,
		MessageInfos:      file_database_proto_msgTypes,
	}.Build()
	File_database_proto = out.File
	file_database_proto_rawDesc = nil
	file_database_proto_goTypes = nil
	file_database_proto_depIdxs = nil
}
