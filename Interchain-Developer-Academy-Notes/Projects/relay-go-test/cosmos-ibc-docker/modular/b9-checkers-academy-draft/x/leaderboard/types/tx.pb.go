// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgUpdateBoard struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgUpdateBoard) Reset()         { *m = MsgUpdateBoard{} }
func (m *MsgUpdateBoard) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBoard) ProtoMessage()    {}
func (*MsgUpdateBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{0}
}
func (m *MsgUpdateBoard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBoard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBoard.Merge(m, src)
}
func (m *MsgUpdateBoard) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBoard.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBoard proto.InternalMessageInfo

func (m *MsgUpdateBoard) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgUpdateBoardResponse struct {
}

func (m *MsgUpdateBoardResponse) Reset()         { *m = MsgUpdateBoardResponse{} }
func (m *MsgUpdateBoardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBoardResponse) ProtoMessage()    {}
func (*MsgUpdateBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{1}
}
func (m *MsgUpdateBoardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBoardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBoardResponse.Merge(m, src)
}
func (m *MsgUpdateBoardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBoardResponse proto.InternalMessageInfo

type MsgSendCandidate struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
}

func (m *MsgSendCandidate) Reset()         { *m = MsgSendCandidate{} }
func (m *MsgSendCandidate) String() string { return proto.CompactTextString(m) }
func (*MsgSendCandidate) ProtoMessage()    {}
func (*MsgSendCandidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{2}
}
func (m *MsgSendCandidate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendCandidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendCandidate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendCandidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendCandidate.Merge(m, src)
}
func (m *MsgSendCandidate) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendCandidate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendCandidate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendCandidate proto.InternalMessageInfo

func (m *MsgSendCandidate) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendCandidate) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendCandidate) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendCandidate) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

type MsgSendCandidateResponse struct {
}

func (m *MsgSendCandidateResponse) Reset()         { *m = MsgSendCandidateResponse{} }
func (m *MsgSendCandidateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendCandidateResponse) ProtoMessage()    {}
func (*MsgSendCandidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{3}
}
func (m *MsgSendCandidateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendCandidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendCandidateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendCandidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendCandidateResponse.Merge(m, src)
}
func (m *MsgSendCandidateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendCandidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendCandidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendCandidateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateBoard)(nil), "b9lab.checkers.leaderboard.MsgUpdateBoard")
	proto.RegisterType((*MsgUpdateBoardResponse)(nil), "b9lab.checkers.leaderboard.MsgUpdateBoardResponse")
	proto.RegisterType((*MsgSendCandidate)(nil), "b9lab.checkers.leaderboard.MsgSendCandidate")
	proto.RegisterType((*MsgSendCandidateResponse)(nil), "b9lab.checkers.leaderboard.MsgSendCandidateResponse")
}

func init() { proto.RegisterFile("leaderboard/tx.proto", fileDescriptor_abcbe4eb090e075c) }

var fileDescriptor_abcbe4eb090e075c = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xb6, 0x28, 0x1d, 0x51, 0xca, 0x22, 0xb2, 0x04, 0x0d, 0x25, 0xa7, 0x52, 0x34,
	0x81, 0xea, 0xc5, 0x6b, 0x15, 0xc4, 0x43, 0x2f, 0x55, 0x2f, 0x5e, 0x64, 0x93, 0x8c, 0x69, 0x30,
	0xd9, 0x5d, 0x76, 0xb7, 0xd0, 0xbe, 0x82, 0x27, 0x1f, 0xcb, 0x63, 0x8f, 0x1e, 0xa5, 0x3d, 0xf9,
	0x16, 0xd2, 0x48, 0xb0, 0x51, 0x8a, 0x7a, 0xdb, 0xfd, 0xe6, 0x37, 0xdf, 0xfc, 0x61, 0x60, 0x2f,
	0x43, 0x1e, 0xa3, 0x0e, 0x25, 0xd7, 0x71, 0x60, 0x27, 0xbe, 0xd2, 0xd2, 0x4a, 0xea, 0x84, 0x67,
	0x19, 0x0f, 0xfd, 0x68, 0x84, 0xd1, 0x23, 0x6a, 0xe3, 0xaf, 0x40, 0xce, 0xe1, 0x6a, 0x86, 0xca,
	0xf8, 0x14, 0xf5, 0x7d, 0x2a, 0x1e, 0xe4, 0x67, 0xaa, 0xd7, 0x85, 0xdd, 0x81, 0x49, 0x6e, 0x55,
	0xcc, 0x2d, 0xf6, 0x97, 0x0c, 0x65, 0xb0, 0x15, 0x69, 0xe4, 0x56, 0x6a, 0x46, 0xda, 0xa4, 0xd3,
	0x1c, 0x96, 0x5f, 0x8f, 0xc1, 0x7e, 0x95, 0x1d, 0xa2, 0x51, 0x52, 0x18, 0xf4, 0x9e, 0x08, 0xb4,
	0x06, 0x26, 0xb9, 0x46, 0x11, 0x9f, 0x73, 0x11, 0xa7, 0x4b, 0x62, 0xbd, 0x11, 0xa5, 0xd0, 0x50,
	0x52, 0x5b, 0xb6, 0x51, 0xc8, 0xc5, 0x9b, 0x1e, 0x40, 0x33, 0x1a, 0x71, 0x21, 0x30, 0xbb, 0xba,
	0x60, 0xf5, 0x22, 0xf0, 0x25, 0xd0, 0x2e, 0xb4, 0x6c, 0x9a, 0xa3, 0x1c, 0xdb, 0x9b, 0x34, 0x47,
	0x63, 0x79, 0xae, 0x58, 0xa3, 0x4d, 0x3a, 0x8d, 0xe1, 0x0f, 0xdd, 0x73, 0x80, 0x7d, 0xef, 0xa5,
	0x6c, 0xb4, 0xf7, 0x4e, 0xa0, 0x3e, 0x30, 0x09, 0xcd, 0x61, 0x7b, 0x75, 0xe6, 0xae, 0xbf, 0x7e,
	0x83, 0x7e, 0x75, 0x66, 0xa7, 0xf7, 0x77, 0xb6, 0x2c, 0x4b, 0x0d, 0xec, 0x54, 0x77, 0x73, 0xf4,
	0x8b, 0x49, 0x85, 0x76, 0x4e, 0xff, 0x43, 0x97, 0x45, 0xfb, 0x97, 0x2f, 0x73, 0x97, 0xcc, 0xe6,
	0x2e, 0x79, 0x9b, 0xbb, 0xe4, 0x79, 0xe1, 0xd6, 0x66, 0x0b, 0xb7, 0xf6, 0xba, 0x70, 0x6b, 0x77,
	0xc7, 0x49, 0x6a, 0x47, 0xe3, 0xd0, 0x8f, 0x64, 0x1e, 0x14, 0xce, 0x41, 0xe9, 0x1c, 0x4c, 0x82,
	0xca, 0x85, 0x4d, 0x15, 0x9a, 0x70, 0xb3, 0x38, 0x95, 0x93, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x89, 0x48, 0x36, 0xcc, 0x7d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateBoard(ctx context.Context, in *MsgUpdateBoard, opts ...grpc.CallOption) (*MsgUpdateBoardResponse, error)
	SendCandidate(ctx context.Context, in *MsgSendCandidate, opts ...grpc.CallOption) (*MsgSendCandidateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateBoard(ctx context.Context, in *MsgUpdateBoard, opts ...grpc.CallOption) (*MsgUpdateBoardResponse, error) {
	out := new(MsgUpdateBoardResponse)
	err := c.cc.Invoke(ctx, "/b9lab.checkers.leaderboard.Msg/UpdateBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendCandidate(ctx context.Context, in *MsgSendCandidate, opts ...grpc.CallOption) (*MsgSendCandidateResponse, error) {
	out := new(MsgSendCandidateResponse)
	err := c.cc.Invoke(ctx, "/b9lab.checkers.leaderboard.Msg/SendCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateBoard(context.Context, *MsgUpdateBoard) (*MsgUpdateBoardResponse, error)
	SendCandidate(context.Context, *MsgSendCandidate) (*MsgSendCandidateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateBoard(ctx context.Context, req *MsgUpdateBoard) (*MsgUpdateBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBoard not implemented")
}
func (*UnimplementedMsgServer) SendCandidate(ctx context.Context, req *MsgSendCandidate) (*MsgSendCandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCandidate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateBoard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/b9lab.checkers.leaderboard.Msg/UpdateBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateBoard(ctx, req.(*MsgUpdateBoard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendCandidate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/b9lab.checkers.leaderboard.Msg/SendCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendCandidate(ctx, req.(*MsgSendCandidate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "b9lab.checkers.leaderboard.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBoard",
			Handler:    _Msg_UpdateBoard_Handler,
		},
		{
			MethodName: "SendCandidate",
			Handler:    _Msg_SendCandidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaderboard/tx.proto",
}

func (m *MsgUpdateBoard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBoard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBoard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBoardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBoardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBoardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendCandidate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendCandidate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendCandidate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendCandidateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendCandidateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendCandidateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateBoard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateBoardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendCandidate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgSendCandidateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateBoard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBoard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBoard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateBoardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBoardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBoardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendCandidate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendCandidate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendCandidate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendCandidateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendCandidateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendCandidateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
