// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/distributed/api.proto

package go_micro_api_distributed

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Note struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNoteRequest) Reset()         { *m = CreateNoteRequest{} }
func (m *CreateNoteRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNoteRequest) ProtoMessage()    {}
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{1}
}

func (m *CreateNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNoteRequest.Unmarshal(m, b)
}
func (m *CreateNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNoteRequest.Marshal(b, m, deterministic)
}
func (m *CreateNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNoteRequest.Merge(m, src)
}
func (m *CreateNoteRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNoteRequest.Size(m)
}
func (m *CreateNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNoteRequest proto.InternalMessageInfo

func (m *CreateNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type CreateNoteResponse struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNoteResponse) Reset()         { *m = CreateNoteResponse{} }
func (m *CreateNoteResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNoteResponse) ProtoMessage()    {}
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{2}
}

func (m *CreateNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNoteResponse.Unmarshal(m, b)
}
func (m *CreateNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNoteResponse.Marshal(b, m, deterministic)
}
func (m *CreateNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNoteResponse.Merge(m, src)
}
func (m *CreateNoteResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNoteResponse.Size(m)
}
func (m *CreateNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNoteResponse proto.InternalMessageInfo

func (m *CreateNoteResponse) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type UpdateNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNoteRequest) Reset()         { *m = UpdateNoteRequest{} }
func (m *UpdateNoteRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNoteRequest) ProtoMessage()    {}
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{3}
}

func (m *UpdateNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNoteRequest.Unmarshal(m, b)
}
func (m *UpdateNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNoteRequest.Marshal(b, m, deterministic)
}
func (m *UpdateNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNoteRequest.Merge(m, src)
}
func (m *UpdateNoteRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNoteRequest.Size(m)
}
func (m *UpdateNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNoteRequest proto.InternalMessageInfo

func (m *UpdateNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type UpdateNoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNoteResponse) Reset()         { *m = UpdateNoteResponse{} }
func (m *UpdateNoteResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateNoteResponse) ProtoMessage()    {}
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{4}
}

func (m *UpdateNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNoteResponse.Unmarshal(m, b)
}
func (m *UpdateNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNoteResponse.Marshal(b, m, deterministic)
}
func (m *UpdateNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNoteResponse.Merge(m, src)
}
func (m *UpdateNoteResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateNoteResponse.Size(m)
}
func (m *UpdateNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNoteResponse proto.InternalMessageInfo

type DeleteNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNoteRequest) Reset()         { *m = DeleteNoteRequest{} }
func (m *DeleteNoteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNoteRequest) ProtoMessage()    {}
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{5}
}

func (m *DeleteNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNoteRequest.Unmarshal(m, b)
}
func (m *DeleteNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNoteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNoteRequest.Merge(m, src)
}
func (m *DeleteNoteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNoteRequest.Size(m)
}
func (m *DeleteNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNoteRequest proto.InternalMessageInfo

func (m *DeleteNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type DeleteNoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNoteResponse) Reset()         { *m = DeleteNoteResponse{} }
func (m *DeleteNoteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteNoteResponse) ProtoMessage()    {}
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{6}
}

func (m *DeleteNoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNoteResponse.Unmarshal(m, b)
}
func (m *DeleteNoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNoteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteNoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNoteResponse.Merge(m, src)
}
func (m *DeleteNoteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteNoteResponse.Size(m)
}
func (m *DeleteNoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNoteResponse proto.InternalMessageInfo

type ListNotesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNotesRequest) Reset()         { *m = ListNotesRequest{} }
func (m *ListNotesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNotesRequest) ProtoMessage()    {}
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{7}
}

func (m *ListNotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotesRequest.Unmarshal(m, b)
}
func (m *ListNotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotesRequest.Marshal(b, m, deterministic)
}
func (m *ListNotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotesRequest.Merge(m, src)
}
func (m *ListNotesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNotesRequest.Size(m)
}
func (m *ListNotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotesRequest proto.InternalMessageInfo

type ListNotesResponse struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNotesResponse) Reset()         { *m = ListNotesResponse{} }
func (m *ListNotesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNotesResponse) ProtoMessage()    {}
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b6004bb6792d52, []int{8}
}

func (m *ListNotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotesResponse.Unmarshal(m, b)
}
func (m *ListNotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotesResponse.Marshal(b, m, deterministic)
}
func (m *ListNotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotesResponse.Merge(m, src)
}
func (m *ListNotesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNotesResponse.Size(m)
}
func (m *ListNotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotesResponse proto.InternalMessageInfo

func (m *ListNotesResponse) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Note)(nil), "go.micro.api.distributed.Note")
	proto.RegisterType((*CreateNoteRequest)(nil), "go.micro.api.distributed.CreateNoteRequest")
	proto.RegisterType((*CreateNoteResponse)(nil), "go.micro.api.distributed.CreateNoteResponse")
	proto.RegisterType((*UpdateNoteRequest)(nil), "go.micro.api.distributed.UpdateNoteRequest")
	proto.RegisterType((*UpdateNoteResponse)(nil), "go.micro.api.distributed.UpdateNoteResponse")
	proto.RegisterType((*DeleteNoteRequest)(nil), "go.micro.api.distributed.DeleteNoteRequest")
	proto.RegisterType((*DeleteNoteResponse)(nil), "go.micro.api.distributed.DeleteNoteResponse")
	proto.RegisterType((*ListNotesRequest)(nil), "go.micro.api.distributed.ListNotesRequest")
	proto.RegisterType((*ListNotesResponse)(nil), "go.micro.api.distributed.ListNotesResponse")
}

func init() { proto.RegisterFile("proto/distributed/api.proto", fileDescriptor_f2b6004bb6792d52) }

var fileDescriptor_f2b6004bb6792d52 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x93, 0x2a, 0x9d, 0x82, 0xd8, 0xa1, 0x87, 0xa5, 0x82, 0x94, 0x9c, 0x8a, 0x95,
	0x14, 0xa2, 0xff, 0xc0, 0x42, 0x15, 0xc4, 0x43, 0xc0, 0x8b, 0xb7, 0xb6, 0x3b, 0x94, 0x85, 0xda,
	0x8d, 0xd9, 0x29, 0xf8, 0xb3, 0xfc, 0x89, 0xb2, 0x9b, 0x68, 0x82, 0x21, 0x21, 0x48, 0x6f, 0xd9,
	0x97, 0xb7, 0xdf, 0x9b, 0xcc, 0x23, 0x70, 0x95, 0x66, 0x9a, 0xf5, 0x42, 0x2a, 0xc3, 0x99, 0xda,
	0x1c, 0x99, 0xe4, 0x62, 0x9d, 0xaa, 0xc8, 0xa9, 0x28, 0x76, 0x3a, 0x7a, 0x57, 0xdb, 0x4c, 0x47,
	0x56, 0xab, 0x78, 0xc2, 0x37, 0x08, 0x5e, 0x34, 0x13, 0x5e, 0x40, 0x4f, 0x49, 0xe1, 0x4d, 0xbd,
	0xd9, 0x20, 0xe9, 0x29, 0x89, 0x02, 0xce, 0xb7, 0x19, 0xad, 0x99, 0xa4, 0xe8, 0x4d, 0xbd, 0x99,
	0x9f, 0xfc, 0x1c, 0x71, 0x0c, 0x7d, 0x56, 0xbc, 0x27, 0xe1, 0x3b, 0x73, 0x7e, 0x40, 0x84, 0x80,
	0xe9, 0x93, 0x45, 0xe0, 0x44, 0xf7, 0x1c, 0xae, 0x60, 0xf4, 0xe0, 0x2e, 0xd9, 0x84, 0x84, 0x3e,
	0x8e, 0x64, 0x18, 0x63, 0x08, 0x0e, 0x9a, 0xc9, 0x45, 0x0d, 0xe3, 0xeb, 0xa8, 0x69, 0xb2, 0xc8,
	0x5d, 0x72, 0xde, 0xf0, 0x11, 0xb0, 0x0a, 0x32, 0xa9, 0x3e, 0x18, 0xfa, 0x17, 0x69, 0x05, 0xa3,
	0xd7, 0x54, 0x9e, 0x60, 0xa4, 0x31, 0x60, 0x15, 0x94, 0x8f, 0x64, 0xf1, 0x4b, 0xda, 0xd3, 0x49,
	0xf0, 0x55, 0x50, 0x81, 0x47, 0xb8, 0x7c, 0x56, 0x86, 0xad, 0x66, 0x0a, 0x7a, 0xf8, 0x04, 0xa3,
	0x8a, 0x56, 0xac, 0xe6, 0x1e, 0xfa, 0x16, 0x63, 0x84, 0x37, 0xf5, 0x3b, 0x64, 0xe6, 0xe6, 0xf8,
	0xcb, 0x87, 0xe1, 0xb2, 0x7c, 0x87, 0x3b, 0x80, 0x72, 0xed, 0x38, 0x6f, 0x86, 0xd4, 0x5a, 0x9e,
	0xdc, 0x76, 0x33, 0x17, 0xe3, 0x2a, 0x80, 0x72, 0x99, 0x6d, 0x41, 0xb5, 0xee, 0xda, 0x82, 0xea,
	0xfd, 0xcc, 0x3c, 0xfb, 0x4d, 0xe5, 0x62, 0xdb, 0xa2, 0x6a, 0x3d, 0xb6, 0x45, 0xd5, 0xbb, 0x42,
	0x09, 0x83, 0xdf, 0x5e, 0xf0, 0xa6, 0xf9, 0xea, 0xdf, 0x42, 0x27, 0xf3, 0x4e, 0xde, 0x3c, 0x65,
	0x73, 0xe6, 0xfe, 0xef, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb3, 0xdd, 0x99, 0xfe,
	0x03, 0x00, 0x00,
}