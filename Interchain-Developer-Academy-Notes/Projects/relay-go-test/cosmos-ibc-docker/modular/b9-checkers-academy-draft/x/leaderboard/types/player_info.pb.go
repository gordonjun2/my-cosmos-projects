// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/player_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type PlayerInfo struct {
	Index          string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	WonCount       uint64 `protobuf:"varint,2,opt,name=wonCount,proto3" json:"wonCount,omitempty"`
	LostCount      uint64 `protobuf:"varint,3,opt,name=lostCount,proto3" json:"lostCount,omitempty"`
	ForfeitedCount uint64 `protobuf:"varint,4,opt,name=forfeitedCount,proto3" json:"forfeitedCount,omitempty"`
	DateUpdated    string `protobuf:"bytes,5,opt,name=dateUpdated,proto3" json:"dateUpdated,omitempty"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2746532f25366801, []int{0}
}
func (m *PlayerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInfo.Merge(m, src)
}
func (m *PlayerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInfo proto.InternalMessageInfo

func (m *PlayerInfo) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *PlayerInfo) GetWonCount() uint64 {
	if m != nil {
		return m.WonCount
	}
	return 0
}

func (m *PlayerInfo) GetLostCount() uint64 {
	if m != nil {
		return m.LostCount
	}
	return 0
}

func (m *PlayerInfo) GetForfeitedCount() uint64 {
	if m != nil {
		return m.ForfeitedCount
	}
	return 0
}

func (m *PlayerInfo) GetDateUpdated() string {
	if m != nil {
		return m.DateUpdated
	}
	return ""
}

func init() {
	proto.RegisterType((*PlayerInfo)(nil), "b9lab.checkers.leaderboard.PlayerInfo")
}

func init() { proto.RegisterFile("leaderboard/player_info.proto", fileDescriptor_2746532f25366801) }

var fileDescriptor_2746532f25366801 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x49, 0x4d, 0x4c,
	0x49, 0x2d, 0x4a, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2f, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x8a, 0xcf,
	0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0xb2, 0xcc, 0x49, 0x4c,
	0xd2, 0x4b, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0x43, 0x52, 0xad, 0xb4, 0x8c, 0x91,
	0x8b, 0x2b, 0x00, 0xac, 0xc3, 0x33, 0x2f, 0x2d, 0x5f, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25,
	0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0xcf,
	0xcf, 0x73, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x82, 0xf3, 0x85,
	0x64, 0xb8, 0x38, 0x73, 0xf2, 0x8b, 0x4b, 0x20, 0x92, 0xcc, 0x60, 0x49, 0x84, 0x80, 0x90, 0x1a,
	0x17, 0x5f, 0x5a, 0x7e, 0x51, 0x5a, 0x6a, 0x66, 0x49, 0x6a, 0x0a, 0x44, 0x09, 0x0b, 0x58, 0x09,
	0x9a, 0xa8, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x62, 0x49, 0x6a, 0x68, 0x01, 0x88, 0x4c, 0x91, 0x60,
	0x05, 0xdb, 0x8e, 0x2c, 0xe4, 0xe4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x60, 0x9f, 0xea,
	0xc3, 0x7c, 0xaa, 0x5f, 0xa1, 0x8f, 0x1c, 0x32, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0,
	0x40, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x7a, 0x50, 0x05, 0x35, 0x01, 0x00, 0x00,
}

func (m *PlayerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DateUpdated) > 0 {
		i -= len(m.DateUpdated)
		copy(dAtA[i:], m.DateUpdated)
		i = encodeVarintPlayerInfo(dAtA, i, uint64(len(m.DateUpdated)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ForfeitedCount != 0 {
		i = encodeVarintPlayerInfo(dAtA, i, uint64(m.ForfeitedCount))
		i--
		dAtA[i] = 0x20
	}
	if m.LostCount != 0 {
		i = encodeVarintPlayerInfo(dAtA, i, uint64(m.LostCount))
		i--
		dAtA[i] = 0x18
	}
	if m.WonCount != 0 {
		i = encodeVarintPlayerInfo(dAtA, i, uint64(m.WonCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPlayerInfo(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlayerInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayerInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlayerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPlayerInfo(uint64(l))
	}
	if m.WonCount != 0 {
		n += 1 + sovPlayerInfo(uint64(m.WonCount))
	}
	if m.LostCount != 0 {
		n += 1 + sovPlayerInfo(uint64(m.LostCount))
	}
	if m.ForfeitedCount != 0 {
		n += 1 + sovPlayerInfo(uint64(m.ForfeitedCount))
	}
	l = len(m.DateUpdated)
	if l > 0 {
		n += 1 + l + sovPlayerInfo(uint64(l))
	}
	return n
}

func sovPlayerInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayerInfo(x uint64) (n int) {
	return sovPlayerInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerInfo
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
			return fmt.Errorf("proto: PlayerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
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
				return ErrInvalidLengthPlayerInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WonCount", wireType)
			}
			m.WonCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WonCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LostCount", wireType)
			}
			m.LostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForfeitedCount", wireType)
			}
			m.ForfeitedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForfeitedCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateUpdated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
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
				return ErrInvalidLengthPlayerInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DateUpdated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlayerInfo
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
func skipPlayerInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayerInfo
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
					return 0, ErrIntOverflowPlayerInfo
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
					return 0, ErrIntOverflowPlayerInfo
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
				return 0, ErrInvalidLengthPlayerInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayerInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayerInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayerInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayerInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayerInfo = fmt.Errorf("proto: unexpected end of group")
)
