// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: aas/v1/aas.proto

package aasv1

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

type Aa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category        int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Classes         int32  `protobuf:"varint,4,opt,name=classes,proto3" json:"classes,omitempty"`
	Races           int32  `protobuf:"varint,5,opt,name=races,proto3" json:"races,omitempty"`
	DrakkinHeritage int32  `protobuf:"varint,6,opt,name=drakkin_heritage,json=drakkinHeritage,proto3" json:"drakkin_heritage,omitempty"`
	Deities         int32  `protobuf:"varint,7,opt,name=deities,proto3" json:"deities,omitempty"`
	Status          int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Type            int32  `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Charges         int32  `protobuf:"varint,10,opt,name=charges,proto3" json:"charges,omitempty"`
	GrantOnly       bool   `protobuf:"varint,11,opt,name=grant_only,json=grantOnly,proto3" json:"grant_only,omitempty"`
	FirstRankId     int32  `protobuf:"varint,12,opt,name=first_rank_id,json=firstRankId,proto3" json:"first_rank_id,omitempty"`
	Enabled         bool   `protobuf:"varint,13,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ResetOnDeath    bool   `protobuf:"varint,14,opt,name=reset_on_death,json=resetOnDeath,proto3" json:"reset_on_death,omitempty"`
}

func (x *Aa) Reset() {
	*x = Aa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aa) ProtoMessage() {}

func (x *Aa) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aa.ProtoReflect.Descriptor instead.
func (*Aa) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{0}
}

func (x *Aa) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Aa) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Aa) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Aa) GetClasses() int32 {
	if x != nil {
		return x.Classes
	}
	return 0
}

func (x *Aa) GetRaces() int32 {
	if x != nil {
		return x.Races
	}
	return 0
}

func (x *Aa) GetDrakkinHeritage() int32 {
	if x != nil {
		return x.DrakkinHeritage
	}
	return 0
}

func (x *Aa) GetDeities() int32 {
	if x != nil {
		return x.Deities
	}
	return 0
}

func (x *Aa) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Aa) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Aa) GetCharges() int32 {
	if x != nil {
		return x.Charges
	}
	return 0
}

func (x *Aa) GetGrantOnly() bool {
	if x != nil {
		return x.GrantOnly
	}
	return false
}

func (x *Aa) GetFirstRankId() int32 {
	if x != nil {
		return x.FirstRankId
	}
	return 0
}

func (x *Aa) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Aa) GetResetOnDeath() bool {
	if x != nil {
		return x.ResetOnDeath
	}
	return false
}

type CreateAaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *CreateAaRequest) Reset() {
	*x = CreateAaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAaRequest) ProtoMessage() {}

func (x *CreateAaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAaRequest.ProtoReflect.Descriptor instead.
func (*CreateAaRequest) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAaRequest) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type CreateAaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *CreateAaResponse) Reset() {
	*x = CreateAaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAaResponse) ProtoMessage() {}

func (x *CreateAaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAaResponse.ProtoReflect.Descriptor instead.
func (*CreateAaResponse) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAaResponse) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type ListAasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAasRequest) Reset() {
	*x = ListAasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAasRequest) ProtoMessage() {}

func (x *ListAasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAasRequest.ProtoReflect.Descriptor instead.
func (*ListAasRequest) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{3}
}

type ListAasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *ListAasResponse) Reset() {
	*x = ListAasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAasResponse) ProtoMessage() {}

func (x *ListAasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAasResponse.ProtoReflect.Descriptor instead.
func (*ListAasResponse) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{4}
}

func (x *ListAasResponse) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type GetAaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAaRequest) Reset() {
	*x = GetAaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAaRequest) ProtoMessage() {}

func (x *GetAaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAaRequest.ProtoReflect.Descriptor instead.
func (*GetAaRequest) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{5}
}

func (x *GetAaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *GetAaResponse) Reset() {
	*x = GetAaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAaResponse) ProtoMessage() {}

func (x *GetAaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAaResponse.ProtoReflect.Descriptor instead.
func (*GetAaResponse) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{6}
}

func (x *GetAaResponse) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type UpdateAaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Aa *Aa   `protobuf:"bytes,2,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *UpdateAaRequest) Reset() {
	*x = UpdateAaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAaRequest) ProtoMessage() {}

func (x *UpdateAaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAaRequest.ProtoReflect.Descriptor instead.
func (*UpdateAaRequest) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAaRequest) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type UpdateAaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *UpdateAaResponse) Reset() {
	*x = UpdateAaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAaResponse) ProtoMessage() {}

func (x *UpdateAaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAaResponse.ProtoReflect.Descriptor instead.
func (*UpdateAaResponse) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAaResponse) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

type DeleteAaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAaRequest) Reset() {
	*x = DeleteAaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAaRequest) ProtoMessage() {}

func (x *DeleteAaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAaRequest.ProtoReflect.Descriptor instead.
func (*DeleteAaRequest) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aa *Aa `protobuf:"bytes,1,opt,name=aa,proto3" json:"aa,omitempty"`
}

func (x *DeleteAaResponse) Reset() {
	*x = DeleteAaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aas_v1_aas_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAaResponse) ProtoMessage() {}

func (x *DeleteAaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aas_v1_aas_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAaResponse.ProtoReflect.Descriptor instead.
func (*DeleteAaResponse) Descriptor() ([]byte, []int) {
	return file_aas_v1_aas_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAaResponse) GetAa() *Aa {
	if x != nil {
		return x.Aa
	}
	return nil
}

var File_aas_v1_aas_proto protoreflect.FileDescriptor

var file_aas_v1_aas_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x02, 0x41, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x64, 0x72, 0x61, 0x6b, 0x6b, 0x69, 0x6e, 0x5f, 0x68, 0x65, 0x72, 0x69, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x72, 0x61, 0x6b, 0x6b,
	0x69, 0x6e, 0x48, 0x65, 0x72, 0x69, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x74, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x6e, 0x44, 0x65, 0x61, 0x74, 0x68, 0x22, 0x2d, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x2e, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x10, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x1e, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x3d, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x02, 0x61, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x2e, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x02, 0x61, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x52, 0x02, 0x61, 0x61, 0x32, 0xc3, 0x05, 0x0a,
	0x09, 0x41, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x07, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x61, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x92, 0x41, 0x2c, 0x0a, 0x03, 0x41, 0x61,
	0x73, 0x12, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x61, 0x73, 0x1a, 0x1b, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x61, 0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x62, 0x02,
	0x61, 0x61, 0x12, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x30,
	0x01, 0x12, 0x8a, 0x01, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x61, 0x12, 0x17,
	0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4b, 0x92, 0x41, 0x2e, 0x0a, 0x03, 0x41, 0x61, 0x73, 0x12, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x61, 0x1a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x61, 0x61, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x62, 0x02, 0x61,
	0x61, 0x22, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x12, 0x7f,
	0x0a, 0x05, 0x47, 0x65, 0x74, 0x41, 0x61, 0x12, 0x14, 0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x92, 0x41, 0x2a, 0x0a, 0x03, 0x41, 0x61, 0x73, 0x12, 0x08,
	0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x61, 0x61, 0x1a, 0x19, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x61, 0x61, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x62, 0x02, 0x61, 0x61, 0x12, 0x10, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x8f, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x61, 0x12, 0x17, 0x2e, 0x61,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x50, 0x92, 0x41, 0x2e, 0x0a, 0x03, 0x41, 0x61, 0x73, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x61, 0x61, 0x1a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x61, 0x61, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x62, 0x02, 0x61, 0x61, 0x1a,
	0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x8e, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x61, 0x12, 0x17,
	0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4f, 0x92, 0x41, 0x30, 0x0a, 0x03, 0x41, 0x61, 0x73, 0x12, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x61, 0x61, 0x1a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x61, 0x61, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x62, 0x02, 0x61, 0x61,
	0x2a, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x61, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0xea, 0x01, 0x92, 0x41, 0x6e, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x02, 0x72, 0x62, 0x0a, 0x23, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x20, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3b, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f,
	0x68, 0x61, 0x6e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x6f, 0x72, 0x73, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65,
	0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x79, 0x2f, 0x6e, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x61, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x61, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02,
	0x06, 0x41, 0x61, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x61, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x12, 0x41, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x61, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aas_v1_aas_proto_rawDescOnce sync.Once
	file_aas_v1_aas_proto_rawDescData = file_aas_v1_aas_proto_rawDesc
)

func file_aas_v1_aas_proto_rawDescGZIP() []byte {
	file_aas_v1_aas_proto_rawDescOnce.Do(func() {
		file_aas_v1_aas_proto_rawDescData = protoimpl.X.CompressGZIP(file_aas_v1_aas_proto_rawDescData)
	})
	return file_aas_v1_aas_proto_rawDescData
}

var file_aas_v1_aas_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_aas_v1_aas_proto_goTypes = []interface{}{
	(*Aa)(nil),               // 0: aas.v1.Aa
	(*CreateAaRequest)(nil),  // 1: aas.v1.CreateAaRequest
	(*CreateAaResponse)(nil), // 2: aas.v1.CreateAaResponse
	(*ListAasRequest)(nil),   // 3: aas.v1.ListAasRequest
	(*ListAasResponse)(nil),  // 4: aas.v1.ListAasResponse
	(*GetAaRequest)(nil),     // 5: aas.v1.GetAaRequest
	(*GetAaResponse)(nil),    // 6: aas.v1.GetAaResponse
	(*UpdateAaRequest)(nil),  // 7: aas.v1.UpdateAaRequest
	(*UpdateAaResponse)(nil), // 8: aas.v1.UpdateAaResponse
	(*DeleteAaRequest)(nil),  // 9: aas.v1.DeleteAaRequest
	(*DeleteAaResponse)(nil), // 10: aas.v1.DeleteAaResponse
}
var file_aas_v1_aas_proto_depIdxs = []int32{
	0,  // 0: aas.v1.CreateAaRequest.aa:type_name -> aas.v1.Aa
	0,  // 1: aas.v1.CreateAaResponse.aa:type_name -> aas.v1.Aa
	0,  // 2: aas.v1.ListAasResponse.aa:type_name -> aas.v1.Aa
	0,  // 3: aas.v1.GetAaResponse.aa:type_name -> aas.v1.Aa
	0,  // 4: aas.v1.UpdateAaRequest.aa:type_name -> aas.v1.Aa
	0,  // 5: aas.v1.UpdateAaResponse.aa:type_name -> aas.v1.Aa
	0,  // 6: aas.v1.DeleteAaResponse.aa:type_name -> aas.v1.Aa
	3,  // 7: aas.v1.AaService.ListAas:input_type -> aas.v1.ListAasRequest
	1,  // 8: aas.v1.AaService.CreateAa:input_type -> aas.v1.CreateAaRequest
	5,  // 9: aas.v1.AaService.GetAa:input_type -> aas.v1.GetAaRequest
	7,  // 10: aas.v1.AaService.UpdateAa:input_type -> aas.v1.UpdateAaRequest
	9,  // 11: aas.v1.AaService.DeleteAa:input_type -> aas.v1.DeleteAaRequest
	4,  // 12: aas.v1.AaService.ListAas:output_type -> aas.v1.ListAasResponse
	2,  // 13: aas.v1.AaService.CreateAa:output_type -> aas.v1.CreateAaResponse
	6,  // 14: aas.v1.AaService.GetAa:output_type -> aas.v1.GetAaResponse
	8,  // 15: aas.v1.AaService.UpdateAa:output_type -> aas.v1.UpdateAaResponse
	10, // 16: aas.v1.AaService.DeleteAa:output_type -> aas.v1.DeleteAaResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_aas_v1_aas_proto_init() }
func file_aas_v1_aas_proto_init() {
	if File_aas_v1_aas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aas_v1_aas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aa); i {
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
		file_aas_v1_aas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAaRequest); i {
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
		file_aas_v1_aas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAaResponse); i {
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
		file_aas_v1_aas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAasRequest); i {
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
		file_aas_v1_aas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAasResponse); i {
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
		file_aas_v1_aas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAaRequest); i {
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
		file_aas_v1_aas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAaResponse); i {
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
		file_aas_v1_aas_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAaRequest); i {
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
		file_aas_v1_aas_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAaResponse); i {
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
		file_aas_v1_aas_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAaRequest); i {
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
		file_aas_v1_aas_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAaResponse); i {
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
			RawDescriptor: file_aas_v1_aas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aas_v1_aas_proto_goTypes,
		DependencyIndexes: file_aas_v1_aas_proto_depIdxs,
		MessageInfos:      file_aas_v1_aas_proto_msgTypes,
	}.Build()
	File_aas_v1_aas_proto = out.File
	file_aas_v1_aas_proto_rawDesc = nil
	file_aas_v1_aas_proto_goTypes = nil
	file_aas_v1_aas_proto_depIdxs = nil
}
