// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

/*
Package raftpb is a generated protocol buffer package.

It is generated from these files:
	raft.proto

It has these top-level messages:
	Entry
	AppendEntriesArgs
	AppendEntriesReply
	RequestVotesArgs
	RequestVotesReply
*/
package raftpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntryType int32

const (
	EntryType_Normal EntryType = 0
	EntryType_Config EntryType = 1
)

var EntryType_name = map[int32]string{
	0: "Normal",
	1: "Config",
}
var EntryType_value = map[string]int32{
	"Normal": 0,
	"Config": 1,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}
func (EntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Entry struct {
	Type  EntryType `protobuf:"varint,1,opt,name=Type,enum=raftpb.EntryType" json:"Type,omitempty"`
	Term  uint64    `protobuf:"varint,2,opt,name=Term" json:"Term,omitempty"`
	Index uint64    `protobuf:"varint,3,opt,name=Index" json:"Index,omitempty"`
	Data  []byte    `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetType() EntryType {
	if m != nil {
		return m.Type
	}
	return EntryType_Normal
}

func (m *Entry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Entry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AppendEntriesArgs struct {
	Term         uint64   `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LeaderID     uint64   `protobuf:"varint,2,opt,name=LeaderID" json:"LeaderID,omitempty"`
	PrevLogIndex uint64   `protobuf:"varint,3,opt,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  uint64   `protobuf:"varint,4,opt,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	Entries      []*Entry `protobuf:"bytes,5,rep,name=Entries" json:"Entries,omitempty"`
	LeaderCommit uint64   `protobuf:"varint,6,opt,name=LeaderCommit" json:"LeaderCommit,omitempty"`
}

func (m *AppendEntriesArgs) Reset()                    { *m = AppendEntriesArgs{} }
func (m *AppendEntriesArgs) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesArgs) ProtoMessage()               {}
func (*AppendEntriesArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppendEntriesArgs) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesArgs) GetLeaderID() uint64 {
	if m != nil {
		return m.LeaderID
	}
	return 0
}

func (m *AppendEntriesArgs) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesArgs) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesArgs) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntriesArgs) GetLeaderCommit() uint64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type AppendEntriesReply struct {
	Term    uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
}

func (m *AppendEntriesReply) Reset()                    { *m = AppendEntriesReply{} }
func (m *AppendEntriesReply) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesReply) ProtoMessage()               {}
func (*AppendEntriesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestVotesArgs struct {
	Term         uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	CandidateID  uint64 `protobuf:"varint,2,opt,name=CandidateID" json:"CandidateID,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=LastLogTerm" json:"LastLogTerm,omitempty"`
}

func (m *RequestVotesArgs) Reset()                    { *m = RequestVotesArgs{} }
func (m *RequestVotesArgs) String() string            { return proto.CompactTextString(m) }
func (*RequestVotesArgs) ProtoMessage()               {}
func (*RequestVotesArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVotesArgs) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVotesArgs) GetCandidateID() uint64 {
	if m != nil {
		return m.CandidateID
	}
	return 0
}

func (m *RequestVotesArgs) GetLastLogIndex() uint64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVotesArgs) GetLastLogTerm() uint64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type RequestVotesReply struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (m *RequestVotesReply) Reset()                    { *m = RequestVotesReply{} }
func (m *RequestVotesReply) String() string            { return proto.CompactTextString(m) }
func (*RequestVotesReply) ProtoMessage()               {}
func (*RequestVotesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVotesReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVotesReply) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

func init() {
	proto.RegisterType((*Entry)(nil), "raftpb.Entry")
	proto.RegisterType((*AppendEntriesArgs)(nil), "raftpb.AppendEntriesArgs")
	proto.RegisterType((*AppendEntriesReply)(nil), "raftpb.AppendEntriesReply")
	proto.RegisterType((*RequestVotesArgs)(nil), "raftpb.RequestVotesArgs")
	proto.RegisterType((*RequestVotesReply)(nil), "raftpb.RequestVotesReply")
	proto.RegisterEnum("raftpb.EntryType", EntryType_name, EntryType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Raft service

type RaftClient interface {
	AppendEntries(ctx context.Context, in *AppendEntriesArgs, opts ...grpc.CallOption) (*AppendEntriesReply, error)
	RequestVotes(ctx context.Context, in *RequestVotesArgs, opts ...grpc.CallOption) (*RequestVotesReply, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) AppendEntries(ctx context.Context, in *AppendEntriesArgs, opts ...grpc.CallOption) (*AppendEntriesReply, error) {
	out := new(AppendEntriesReply)
	err := grpc.Invoke(ctx, "/raftpb.Raft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) RequestVotes(ctx context.Context, in *RequestVotesArgs, opts ...grpc.CallOption) (*RequestVotesReply, error) {
	out := new(RequestVotesReply)
	err := grpc.Invoke(ctx, "/raftpb.Raft/RequestVotes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	AppendEntries(context.Context, *AppendEntriesArgs) (*AppendEntriesReply, error)
	RequestVotes(context.Context, *RequestVotesArgs) (*RequestVotesReply, error)
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftpb.Raft/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).AppendEntries(ctx, req.(*AppendEntriesArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_RequestVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVotesArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).RequestVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftpb.Raft/RequestVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).RequestVotes(ctx, req.(*RequestVotesArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raftpb.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _Raft_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVotes",
			Handler:    _Raft_RequestVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x65, 0x8b, 0x31, 0x74, 0x0c, 0x15, 0xac, 0x7a, 0x58, 0x7c, 0xb2, 0x5c, 0x55, 0xb5, 0x7a,
	0xe0, 0x40, 0x7f, 0x01, 0x05, 0xd4, 0x22, 0xa1, 0x28, 0xda, 0xa0, 0xdc, 0x17, 0xbc, 0x20, 0x4b,
	0xf8, 0x23, 0xeb, 0x25, 0x0a, 0x7f, 0x22, 0xd7, 0xfc, 0xb3, 0xfc, 0x9e, 0x68, 0xd7, 0x1f, 0x5a,
	0x13, 0x92, 0xdb, 0xcc, 0x7b, 0xa3, 0x37, 0xf3, 0x9e, 0xd7, 0x00, 0x82, 0xed, 0xe5, 0x24, 0x13,
	0xa9, 0x4c, 0xb1, 0xad, 0xea, 0x6c, 0xeb, 0x1f, 0xa1, 0xb3, 0x4c, 0xa4, 0x38, 0xe3, 0x9f, 0x60,
	0x6d, 0xce, 0x19, 0x27, 0xc8, 0x43, 0xc1, 0xb7, 0xe9, 0x68, 0x52, 0xf0, 0x13, 0x4d, 0x2a, 0x82,
	0x6a, 0x1a, 0x63, 0xb0, 0x36, 0x5c, 0xc4, 0xe4, 0x8b, 0x87, 0x02, 0x8b, 0xea, 0x1a, 0x7f, 0x87,
	0xce, 0x2a, 0x09, 0xf9, 0x13, 0x69, 0x6b, 0xb0, 0x68, 0xd4, 0xe4, 0x82, 0x49, 0x46, 0x2c, 0x0f,
	0x05, 0x7d, 0xaa, 0x6b, 0xff, 0x15, 0xc1, 0x68, 0x96, 0x65, 0x3c, 0x09, 0x95, 0x6e, 0xc4, 0xf3,
	0x99, 0x38, 0xe4, 0xb5, 0x26, 0x32, 0x34, 0x5d, 0xe8, 0xad, 0x39, 0x0b, 0xb9, 0x58, 0x2d, 0xca,
	0x5d, 0x75, 0x8f, 0x7d, 0xe8, 0xdf, 0x0a, 0xfe, 0xb8, 0x4e, 0x0f, 0xe6, 0xda, 0x06, 0x86, 0x3d,
	0x70, 0xca, 0x5e, 0x4b, 0x5b, 0x7a, 0xc4, 0x84, 0xf0, 0x2f, 0xe8, 0x96, 0x47, 0x90, 0x8e, 0xd7,
	0x0e, 0x9c, 0xe9, 0xa0, 0xe1, 0x99, 0x56, 0xac, 0x5a, 0x57, 0xac, 0x9e, 0xa7, 0x71, 0x1c, 0x49,
	0x62, 0x17, 0xeb, 0x4c, 0xcc, 0xff, 0x0b, 0xb8, 0xe1, 0x8b, 0xf2, 0xec, 0x78, 0xbe, 0x6a, 0x8c,
	0x40, 0xf7, 0xee, 0xb4, 0xdb, 0xf1, 0x3c, 0xd7, 0xbe, 0x7a, 0xb4, 0x6a, 0xfd, 0x67, 0x04, 0x43,
	0xca, 0x1f, 0x4e, 0x3c, 0x97, 0xf7, 0xa9, 0xfc, 0x24, 0x1b, 0x0f, 0x9c, 0x39, 0x4b, 0xc2, 0x28,
	0x64, 0x92, 0xd7, 0xf1, 0x98, 0x90, 0x3e, 0x99, 0xe5, 0xf2, 0x32, 0x21, 0x13, 0x53, 0x2a, 0x65,
	0x6f, 0x26, 0x64, 0x40, 0xfe, 0x0a, 0x46, 0xe6, 0x3d, 0x1f, 0x7b, 0xf2, 0xc0, 0x51, 0x13, 0xff,
	0x04, 0x4b, 0x24, 0x0f, 0x4b, 0x5f, 0x26, 0xf4, 0xfb, 0x07, 0x7c, 0xad, 0x5f, 0x12, 0x06, 0xb0,
	0x6f, 0x52, 0x11, 0xb3, 0xe3, 0xb0, 0xa5, 0xea, 0x79, 0x9a, 0xec, 0xa3, 0xc3, 0x10, 0x4d, 0x5f,
	0x10, 0x58, 0x94, 0xed, 0x25, 0xfe, 0x0f, 0x83, 0x46, 0x9a, 0x78, 0x5c, 0x7d, 0x9a, 0x77, 0x8f,
	0xc7, 0x75, 0xaf, 0x52, 0xfa, 0x56, 0xbf, 0x85, 0x97, 0xd0, 0x37, 0x2d, 0x60, 0x52, 0x4d, 0x5f,
	0x06, 0xed, 0x8e, 0xaf, 0x31, 0xa5, 0xcc, 0xd6, 0xd6, 0x3f, 0xcd, 0x9f, 0xb7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xce, 0xce, 0xa7, 0x42, 0x03, 0x00, 0x00,
}
