// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comms.proto

package comms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Message struct {
	To                   uint64   `protobuf:"varint,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 uint64   `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Spawn                string   `protobuf:"bytes,4,opt,name=spawn,proto3" json:"spawn,omitempty"`
	SpawnAddress         uint64   `protobuf:"varint,5,opt,name=spawn_address,json=spawnAddress,proto3" json:"spawn_address,omitempty"`
	Kill                 bool     `protobuf:"varint,6,opt,name=kill,proto3" json:"kill,omitempty"`
	Exiting              bool     `protobuf:"varint,7,opt,name=exiting,proto3" json:"exiting,omitempty"`
	Exit                 int32    `protobuf:"varint,8,opt,name=exit,proto3" json:"exit,omitempty"`
	YourAddress          uint64   `protobuf:"varint,9,opt,name=your_address,json=yourAddress,proto3" json:"your_address,omitempty"`
	ParentAddress        uint64   `protobuf:"varint,10,opt,name=parent_address,json=parentAddress,proto3" json:"parent_address,omitempty"`
	Error                int32    `protobuf:"varint,11,opt,name=error,proto3" json:"error,omitempty"`
	Startup              *Startup `protobuf:"bytes,12,opt,name=startup,proto3" json:"startup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Message) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetSpawn() string {
	if m != nil {
		return m.Spawn
	}
	return ""
}

func (m *Message) GetSpawnAddress() uint64 {
	if m != nil {
		return m.SpawnAddress
	}
	return 0
}

func (m *Message) GetKill() bool {
	if m != nil {
		return m.Kill
	}
	return false
}

func (m *Message) GetExiting() bool {
	if m != nil {
		return m.Exiting
	}
	return false
}

func (m *Message) GetExit() int32 {
	if m != nil {
		return m.Exit
	}
	return 0
}

func (m *Message) GetYourAddress() uint64 {
	if m != nil {
		return m.YourAddress
	}
	return 0
}

func (m *Message) GetParentAddress() uint64 {
	if m != nil {
		return m.ParentAddress
	}
	return 0
}

func (m *Message) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *Message) GetStartup() *Startup {
	if m != nil {
		return m.Startup
	}
	return nil
}

type Startup struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Addr                 uint64   `protobuf:"varint,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Parent               uint64   `protobuf:"varint,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Dbs                  []*DB    `protobuf:"bytes,4,rep,name=dbs,proto3" json:"dbs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Startup) Reset()         { *m = Startup{} }
func (m *Startup) String() string { return proto.CompactTextString(m) }
func (*Startup) ProtoMessage()    {}
func (*Startup) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{1}
}

func (m *Startup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Startup.Unmarshal(m, b)
}
func (m *Startup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Startup.Marshal(b, m, deterministic)
}
func (m *Startup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Startup.Merge(m, src)
}
func (m *Startup) XXX_Size() int {
	return xxx_messageInfo_Startup.Size(m)
}
func (m *Startup) XXX_DiscardUnknown() {
	xxx_messageInfo_Startup.DiscardUnknown(m)
}

var xxx_messageInfo_Startup proto.InternalMessageInfo

func (m *Startup) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *Startup) GetAddr() uint64 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *Startup) GetParent() uint64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *Startup) GetDbs() []*DB {
	if m != nil {
		return m.Dbs
	}
	return nil
}

type DB struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Connection           string   `protobuf:"bytes,3,opt,name=connection,proto3" json:"connection,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DB) Reset()         { *m = DB{} }
func (m *DB) String() string { return proto.CompactTextString(m) }
func (*DB) ProtoMessage()    {}
func (*DB) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{2}
}

func (m *DB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DB.Unmarshal(m, b)
}
func (m *DB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DB.Marshal(b, m, deterministic)
}
func (m *DB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DB.Merge(m, src)
}
func (m *DB) XXX_Size() int {
	return xxx_messageInfo_DB.Size(m)
}
func (m *DB) XXX_DiscardUnknown() {
	xxx_messageInfo_DB.DiscardUnknown(m)
}

var xxx_messageInfo_DB proto.InternalMessageInfo

func (m *DB) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DB) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DB) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *DB) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "comms.Message")
	proto.RegisterType((*Startup)(nil), "comms.Startup")
	proto.RegisterType((*DB)(nil), "comms.DB")
}

func init() { proto.RegisterFile("comms.proto", fileDescriptor_db39efb7717b7d47) }

var fileDescriptor_db39efb7717b7d47 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x86, 0x69, 0x6d, 0xad, 0x3d, 0xad, 0x5e, 0x04, 0x19, 0x81, 0xc1, 0xc8, 0x1c, 0x83, 0x5c,
	0x79, 0xe1, 0x9e, 0x60, 0xe2, 0xed, 0x6e, 0xb2, 0x07, 0x18, 0xd1, 0x66, 0xd2, 0x69, 0x93, 0x92,
	0x44, 0x36, 0xdf, 0x70, 0x8f, 0x35, 0x72, 0xd2, 0xca, 0xee, 0xfe, 0xff, 0xeb, 0xcf, 0xf9, 0x93,
	0xd3, 0x40, 0x75, 0x30, 0x5d, 0xe7, 0xd6, 0xbd, 0x35, 0xde, 0x90, 0x1c, 0xcd, 0xea, 0x37, 0x85,
	0xe2, 0x4d, 0x39, 0x27, 0x8f, 0x8a, 0x2c, 0x20, 0xf5, 0x86, 0x26, 0x2c, 0xe1, 0x99, 0x48, 0xbd,
	0x21, 0x04, 0xb2, 0x4f, 0x6b, 0x3a, 0x9a, 0x22, 0x41, 0x1d, 0x58, 0x23, 0xbd, 0xa4, 0x13, 0x96,
	0xf0, 0x5a, 0xa0, 0x26, 0x4b, 0xc8, 0x5d, 0x2f, 0xbf, 0x35, 0xcd, 0x58, 0xc2, 0x4b, 0x11, 0x0d,
	0x79, 0x82, 0x39, 0x8a, 0x0f, 0xd9, 0x34, 0x56, 0x39, 0x47, 0x73, 0x1c, 0x53, 0x23, 0x7c, 0x8d,
	0x2c, 0x8c, 0x3b, 0xb5, 0xe7, 0x33, 0x9d, 0xb2, 0x84, 0xcf, 0x04, 0x6a, 0x42, 0xa1, 0x50, 0x3f,
	0xad, 0x6f, 0xf5, 0x91, 0x16, 0x88, 0x47, 0x1b, 0xd2, 0x41, 0xd2, 0x19, 0x4b, 0x78, 0x2e, 0x50,
	0x93, 0x47, 0xa8, 0xaf, 0xe6, 0x62, 0x6f, 0x2d, 0x25, 0xb6, 0x54, 0x81, 0x8d, 0x25, 0xcf, 0xb0,
	0xe8, 0xa5, 0x55, 0xda, 0xdf, 0x42, 0x80, 0xa1, 0x79, 0xa4, 0x63, 0x6c, 0x09, 0xb9, 0xb2, 0xd6,
	0x58, 0x5a, 0xe1, 0xf8, 0x68, 0x08, 0x87, 0xc2, 0x79, 0x69, 0xfd, 0xa5, 0xa7, 0x35, 0x4b, 0x78,
	0xb5, 0x59, 0xac, 0xe3, 0x1a, 0xdf, 0x23, 0x15, 0xe3, 0xe7, 0xd5, 0x17, 0x14, 0x03, 0x23, 0x77,
	0x30, 0xed, 0x4c, 0x73, 0x39, 0x2b, 0xdc, 0x66, 0x29, 0x06, 0x17, 0x2e, 0x10, 0x8e, 0x30, 0x6e,
	0x34, 0xe8, 0x90, 0x8d, 0xe7, 0xc0, 0x9d, 0x66, 0x62, 0x70, 0xe4, 0x1e, 0x26, 0xcd, 0xde, 0xd1,
	0x8c, 0x4d, 0x78, 0xb5, 0x29, 0x87, 0xd2, 0xdd, 0x56, 0x04, 0xba, 0xda, 0x43, 0xba, 0xdb, 0x86,
	0x71, 0xfe, 0xda, 0x8f, 0x25, 0xa8, 0x03, 0xd3, 0xb2, 0x53, 0x58, 0x51, 0x0a, 0xd4, 0xe4, 0x01,
	0xe0, 0x60, 0xb4, 0x56, 0x07, 0xdf, 0x1a, 0x8d, 0x35, 0xa5, 0xf8, 0x47, 0xc2, 0xcd, 0xbd, 0x39,
	0xa9, 0xdb, 0x0f, 0x44, 0xb3, 0x9f, 0xe2, 0x43, 0x79, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0xde, 0xea, 0xec, 0x37, 0x02, 0x00, 0x00,
}