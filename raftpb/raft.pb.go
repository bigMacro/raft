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
	InstallSnapshotArgs
	InstallSnapshotReply
	ProposeArgs
	ProposeReply
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
	EntryType_None   EntryType = 2
)

var EntryType_name = map[int32]string{
	0: "Normal",
	1: "Config",
	2: "None",
}
var EntryType_value = map[string]int32{
	"Normal": 0,
	"Config": 1,
	"None":   2,
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

type InstallSnapshotArgs struct {
	Term             uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LeaderID         uint64 `protobuf:"varint,2,opt,name=LeaderID" json:"LeaderID,omitempty"`
	LastConfig       []byte `protobuf:"bytes,3,opt,name=LastConfig,proto3" json:"LastConfig,omitempty"`
	Data             []byte `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	LastIncludeIndex uint64 `protobuf:"varint,5,opt,name=LastIncludeIndex" json:"LastIncludeIndex,omitempty"`
	LastIncludeTerm  uint64 `protobuf:"varint,6,opt,name=LastIncludeTerm" json:"LastIncludeTerm,omitempty"`
}

func (m *InstallSnapshotArgs) Reset()                    { *m = InstallSnapshotArgs{} }
func (m *InstallSnapshotArgs) String() string            { return proto.CompactTextString(m) }
func (*InstallSnapshotArgs) ProtoMessage()               {}
func (*InstallSnapshotArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InstallSnapshotArgs) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *InstallSnapshotArgs) GetLeaderID() uint64 {
	if m != nil {
		return m.LeaderID
	}
	return 0
}

func (m *InstallSnapshotArgs) GetLastConfig() []byte {
	if m != nil {
		return m.LastConfig
	}
	return nil
}

func (m *InstallSnapshotArgs) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InstallSnapshotArgs) GetLastIncludeIndex() uint64 {
	if m != nil {
		return m.LastIncludeIndex
	}
	return 0
}

func (m *InstallSnapshotArgs) GetLastIncludeTerm() uint64 {
	if m != nil {
		return m.LastIncludeTerm
	}
	return 0
}

type InstallSnapshotReply struct {
	Term uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
}

func (m *InstallSnapshotReply) Reset()                    { *m = InstallSnapshotReply{} }
func (m *InstallSnapshotReply) String() string            { return proto.CompactTextString(m) }
func (*InstallSnapshotReply) ProtoMessage()               {}
func (*InstallSnapshotReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InstallSnapshotReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

// ProposeArgs is the log entry stored in Data which will be replicated to raft servers.
type ProposeArgs struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ProposeArgs) Reset()                    { *m = ProposeArgs{} }
func (m *ProposeArgs) String() string            { return proto.CompactTextString(m) }
func (*ProposeArgs) ProtoMessage()               {}
func (*ProposeArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ProposeArgs) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ProposeReply returns the proposes's result. If Success is True, propose succeeds. Otherwise, ErrCode indicates the error reason. ErrNotReady means you can retry later; ErrNotLeader means the current server is not leader, you should ask for Addr server.
type ProposeReply struct {
	Success bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	ErrCode uint64 `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,3,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	Addr    string `protobuf:"bytes,4,opt,name=Addr" json:"Addr,omitempty"`
}

func (m *ProposeReply) Reset()                    { *m = ProposeReply{} }
func (m *ProposeReply) String() string            { return proto.CompactTextString(m) }
func (*ProposeReply) ProtoMessage()               {}
func (*ProposeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ProposeReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ProposeReply) GetErrCode() uint64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ProposeReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ProposeReply) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*Entry)(nil), "raftpb.Entry")
	proto.RegisterType((*AppendEntriesArgs)(nil), "raftpb.AppendEntriesArgs")
	proto.RegisterType((*AppendEntriesReply)(nil), "raftpb.AppendEntriesReply")
	proto.RegisterType((*RequestVotesArgs)(nil), "raftpb.RequestVotesArgs")
	proto.RegisterType((*RequestVotesReply)(nil), "raftpb.RequestVotesReply")
	proto.RegisterType((*InstallSnapshotArgs)(nil), "raftpb.InstallSnapshotArgs")
	proto.RegisterType((*InstallSnapshotReply)(nil), "raftpb.InstallSnapshotReply")
	proto.RegisterType((*ProposeArgs)(nil), "raftpb.ProposeArgs")
	proto.RegisterType((*ProposeReply)(nil), "raftpb.ProposeReply")
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
	InstallSnapshot(ctx context.Context, in *InstallSnapshotArgs, opts ...grpc.CallOption) (*InstallSnapshotReply, error)
	Propose(ctx context.Context, in *ProposeArgs, opts ...grpc.CallOption) (*ProposeReply, error)
	ProposeConfChange(ctx context.Context, in *ProposeArgs, opts ...grpc.CallOption) (*ProposeReply, error)
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

func (c *raftClient) InstallSnapshot(ctx context.Context, in *InstallSnapshotArgs, opts ...grpc.CallOption) (*InstallSnapshotReply, error) {
	out := new(InstallSnapshotReply)
	err := grpc.Invoke(ctx, "/raftpb.Raft/InstallSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Propose(ctx context.Context, in *ProposeArgs, opts ...grpc.CallOption) (*ProposeReply, error) {
	out := new(ProposeReply)
	err := grpc.Invoke(ctx, "/raftpb.Raft/Propose", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) ProposeConfChange(ctx context.Context, in *ProposeArgs, opts ...grpc.CallOption) (*ProposeReply, error) {
	out := new(ProposeReply)
	err := grpc.Invoke(ctx, "/raftpb.Raft/ProposeConfChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	AppendEntries(context.Context, *AppendEntriesArgs) (*AppendEntriesReply, error)
	RequestVotes(context.Context, *RequestVotesArgs) (*RequestVotesReply, error)
	InstallSnapshot(context.Context, *InstallSnapshotArgs) (*InstallSnapshotReply, error)
	Propose(context.Context, *ProposeArgs) (*ProposeReply, error)
	ProposeConfChange(context.Context, *ProposeArgs) (*ProposeReply, error)
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

func _Raft_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallSnapshotArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).InstallSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftpb.Raft/InstallSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).InstallSnapshot(ctx, req.(*InstallSnapshotArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Propose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Propose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftpb.Raft/Propose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Propose(ctx, req.(*ProposeArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_ProposeConfChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).ProposeConfChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftpb.Raft/ProposeConfChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).ProposeConfChange(ctx, req.(*ProposeArgs))
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
		{
			MethodName: "InstallSnapshot",
			Handler:    _Raft_InstallSnapshot_Handler,
		},
		{
			MethodName: "Propose",
			Handler:    _Raft_Propose_Handler,
		},
		{
			MethodName: "ProposeConfChange",
			Handler:    _Raft_ProposeConfChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x9d, 0xb7, 0xf4, 0xdf, 0x6d, 0xf7, 0x5b, 0xeb, 0x55, 0x3f, 0x65, 0x01, 0xa1, 0x10, 0x09,
	0x51, 0x55, 0xa2, 0x0f, 0x45, 0xe2, 0x99, 0xd2, 0x55, 0x50, 0x69, 0x54, 0xc8, 0x9b, 0x78, 0xf7,
	0x1a, 0xb7, 0xab, 0x94, 0xda, 0xc1, 0x71, 0x11, 0xfd, 0x12, 0x7c, 0x35, 0xde, 0x78, 0xe7, 0x9b,
	0x20, 0x3b, 0x71, 0xe5, 0x66, 0x2d, 0x12, 0xbc, 0xf9, 0x9e, 0x7b, 0x73, 0xee, 0x3d, 0xc7, 0xd7,
	0x01, 0x90, 0x74, 0xa1, 0x06, 0xa9, 0x14, 0x4a, 0xe0, 0xaa, 0x3e, 0xa7, 0xf7, 0x51, 0x02, 0x95,
	0x09, 0x57, 0x72, 0x8b, 0x5f, 0x80, 0x77, 0xb7, 0x4d, 0x99, 0x8f, 0x42, 0xd4, 0xfb, 0x6f, 0xd8,
	0x19, 0xe4, 0xf9, 0x81, 0x49, 0xea, 0x04, 0x31, 0x69, 0x8c, 0xc1, 0xbb, 0x63, 0x72, 0xed, 0x9f,
	0x86, 0xa8, 0xe7, 0x11, 0x73, 0xc6, 0x5d, 0xa8, 0x4c, 0x79, 0xcc, 0xbe, 0xf9, 0x67, 0x06, 0xcc,
	0x03, 0x5d, 0x79, 0x4d, 0x15, 0xf5, 0xbd, 0x10, 0xf5, 0x5a, 0xc4, 0x9c, 0xa3, 0x9f, 0x08, 0x3a,
	0xa3, 0x34, 0x65, 0x3c, 0xd6, 0xbc, 0x2b, 0x96, 0x8d, 0xe4, 0x32, 0xdb, 0x71, 0x22, 0x87, 0x33,
	0x80, 0xfa, 0x0d, 0xa3, 0x31, 0x93, 0xd3, 0xeb, 0xa2, 0xd7, 0x2e, 0xc6, 0x11, 0xb4, 0x3e, 0x49,
	0xf6, 0xf5, 0x46, 0x2c, 0xdd, 0xb6, 0x7b, 0x18, 0x0e, 0xa1, 0x59, 0xc4, 0x86, 0xda, 0x33, 0x25,
	0x2e, 0x84, 0x5f, 0x42, 0xad, 0x18, 0xc2, 0xaf, 0x84, 0x67, 0xbd, 0xe6, 0xf0, 0x7c, 0x4f, 0x33,
	0xb1, 0x59, 0xdd, 0x2e, 0x6f, 0x3d, 0x16, 0xeb, 0xf5, 0x4a, 0xf9, 0xd5, 0xbc, 0x9d, 0x8b, 0x45,
	0xef, 0x00, 0xef, 0xe9, 0x22, 0x2c, 0x4d, 0xb6, 0x07, 0x85, 0xf9, 0x50, 0xbb, 0xdd, 0xcc, 0xe7,
	0x2c, 0xcb, 0x8c, 0xae, 0x3a, 0xb1, 0x61, 0xf4, 0x1d, 0x41, 0x9b, 0xb0, 0x2f, 0x1b, 0x96, 0xa9,
	0xcf, 0x42, 0xfd, 0xc1, 0x9b, 0x10, 0x9a, 0x63, 0xca, 0xe3, 0x55, 0x4c, 0x15, 0xdb, 0xd9, 0xe3,
	0x42, 0x66, 0x64, 0x9a, 0xa9, 0xb2, 0x43, 0x2e, 0xa6, 0x59, 0x8a, 0xd8, 0x75, 0xc8, 0x81, 0xa2,
	0x29, 0x74, 0xdc, 0x79, 0x8e, 0x6b, 0x0a, 0xa1, 0xa9, 0x2b, 0xde, 0x4b, 0xca, 0x15, 0x8b, 0x0b,
	0x5d, 0x2e, 0x14, 0xfd, 0x40, 0x70, 0x39, 0xe5, 0x99, 0xa2, 0x49, 0x72, 0xcb, 0x69, 0x9a, 0x3d,
	0x08, 0xf5, 0x4f, 0x57, 0xff, 0x0c, 0x40, 0x4f, 0x38, 0x16, 0x7c, 0xb1, 0x5a, 0x1a, 0x59, 0x2d,
	0xe2, 0x20, 0x87, 0x96, 0x0e, 0xf7, 0xa1, 0xad, 0x2b, 0xa6, 0x7c, 0x9e, 0x6c, 0x62, 0x96, 0x1b,
	0x52, 0x31, 0xbc, 0x8f, 0x70, 0xdc, 0x83, 0x0b, 0x07, 0x33, 0xa3, 0xe5, 0xd7, 0x5d, 0x86, 0xa3,
	0x3e, 0x74, 0x4b, 0x82, 0x8e, 0xfa, 0x13, 0x3d, 0xd7, 0xcb, 0x28, 0x52, 0x91, 0x31, 0x2b, 0xda,
	0x0c, 0x89, 0x9c, 0x97, 0xc1, 0xf5, 0x4e, 0x9b, 0x92, 0x9c, 0xc6, 0x59, 0x13, 0xb4, 0xb7, 0x26,
	0x3a, 0x33, 0x91, 0x72, 0x2c, 0x62, 0x56, 0xb8, 0x63, 0x43, 0xfc, 0x3f, 0x54, 0x27, 0x52, 0x7e,
	0xcc, 0x72, 0x63, 0x1a, 0xa4, 0x88, 0x74, 0xbf, 0x51, 0x1c, 0x4b, 0x63, 0x4a, 0x83, 0x98, 0x73,
	0xff, 0x15, 0x34, 0x76, 0x4f, 0x1b, 0x03, 0x54, 0x67, 0x42, 0xae, 0x69, 0xd2, 0x3e, 0xd1, 0xe7,
	0xdc, 0xcb, 0x36, 0xc2, 0x75, 0xf0, 0x66, 0x82, 0xb3, 0xf6, 0xe9, 0xf0, 0xd7, 0x29, 0x78, 0x84,
	0x2e, 0x14, 0xfe, 0x00, 0xe7, 0x7b, 0x8b, 0x8e, 0xaf, 0xec, 0xab, 0x79, 0xf4, 0xae, 0x83, 0xe0,
	0x60, 0xca, 0xe8, 0x8b, 0x4e, 0xf0, 0x04, 0x5a, 0xee, 0x76, 0x61, 0xdf, 0x56, 0x97, 0xdf, 0x40,
	0x70, 0x75, 0x28, 0x63, 0x69, 0x66, 0x70, 0x51, 0xba, 0x07, 0xfc, 0xc4, 0xd6, 0x1f, 0xd8, 0xb8,
	0xe0, 0xe9, 0x91, 0xa4, 0xe5, 0x7b, 0x03, 0xb5, 0xe2, 0x22, 0xf0, 0xa5, 0x2d, 0x75, 0x2e, 0x2f,
	0xe8, 0x96, 0x40, 0xfb, 0xdd, 0x5b, 0xe8, 0x14, 0x88, 0xb6, 0x6f, 0xfc, 0x40, 0xf9, 0xf2, 0xef,
	0x18, 0xee, 0xab, 0xe6, 0xcf, 0xfc, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x19, 0x39,
	0x5c, 0xa7, 0x05, 0x00, 0x00,
}
