// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lockup/v1/lock.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Lock represents a users locked tokens for a period of time.
// It stores owner, duration, unlock time and the amount of coins locked.
type Lock struct {
	// unique autoincrementing numeric lock id
	LockId uint64 `protobuf:"varint,1,opt,name=lockId,proto3" json:"lockId,omitempty"`
	// the user's address who owns the tokens that are locked
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// the duration that the lock is locked for
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// when the lock was unlocked
	EndTime time.Time `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// the coins locked in this Lock
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_961d2c254326faef, []int{0}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetLockId() uint64 {
	if m != nil {
		return m.LockId
	}
	return 0
}

func (m *Lock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Lock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Lock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *Lock) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*Lock)(nil), "matrix.lockup.v1.Lock")
}

func init() { proto.RegisterFile("lockup/v1/lock.proto", fileDescriptor_961d2c254326faef) }

var fileDescriptor_961d2c254326faef = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0xbb, 0xf6, 0x38, 0x02, 0x12, 0xa7, 0xe8, 0x84, 0x42, 0x91, 0x92, 0x28, 0x03,
	0x8a, 0x10, 0xd8, 0x04, 0x36, 0xc6, 0xd0, 0x05, 0x09, 0x96, 0x88, 0x89, 0x05, 0x39, 0x89, 0xc9,
	0x45, 0xad, 0xf3, 0xa2, 0xd8, 0x29, 0xd7, 0x6f, 0x71, 0x23, 0x9f, 0x81, 0x4f, 0xd2, 0xb1, 0x23,
	0x53, 0x8a, 0xda, 0x01, 0x89, 0xb1, 0x9f, 0x00, 0xd9, 0x4e, 0x00, 0x71, 0x53, 0xec, 0xfc, 0xdf,
	0xfb, 0xbd, 0xbf, 0xff, 0x7a, 0xf6, 0xe5, 0x12, 0xf2, 0x45, 0xd7, 0x90, 0x55, 0x4c, 0xd4, 0x09,
	0x37, 0x2d, 0x48, 0x70, 0x2e, 0x38, 0x95, 0x6d, 0x75, 0x8d, 0x8d, 0x88, 0x57, 0xf1, 0xec, 0xb2,
	0x84, 0x12, 0xb4, 0x48, 0xd4, 0xc9, 0xd4, 0xcd, 0xbc, 0x12, 0xa0, 0x5c, 0x32, 0xa2, 0x6f, 0x59,
	0xf7, 0x99, 0x14, 0x5d, 0x4b, 0x65, 0x05, 0xf5, 0xa0, 0xfb, 0xff, 0xeb, 0xb2, 0xe2, 0x4c, 0x48,
	0xca, 0x9b, 0x11, 0x90, 0x83, 0xe0, 0x20, 0x48, 0x46, 0x05, 0x23, 0xab, 0x38, 0x63, 0x92, 0xc6,
	0x24, 0x87, 0x6a, 0x00, 0x84, 0x3f, 0x4f, 0xec, 0xc9, 0x3b, 0xc8, 0x17, 0xce, 0x43, 0xfb, 0x4c,
	0x99, 0x79, 0x5b, 0xb8, 0x28, 0x40, 0xd1, 0x24, 0x1d, 0x6e, 0xce, 0x13, 0x7b, 0x0a, 0x5f, 0x6a,
	0xd6, 0xba, 0x27, 0x01, 0x8a, 0xee, 0x26, 0x17, 0xc7, 0xde, 0xbf, 0xbf, 0xa6, 0x7c, 0xf9, 0x3a,
	0xd4, 0xbf, 0xc3, 0xd4, 0xc8, 0xce, 0x95, 0x7d, 0x3e, 0x7a, 0x73, 0x4f, 0x03, 0x14, 0xdd, 0x7b,
	0xf9, 0x08, 0x1b, 0x73, 0x78, 0x34, 0x87, 0xe7, 0x43, 0x41, 0x12, 0x6f, 0x7a, 0xdf, 0xfa, 0xd5,
	0xfb, 0xce, 0xd8, 0xf2, 0x0c, 0x78, 0x25, 0x19, 0x6f, 0xe4, 0xfa, 0xd8, 0xfb, 0x0f, 0x0c, 0x7f,
	0xd4, 0xc2, 0xaf, 0x3b, 0x1f, 0xa5, 0x7f, 0xe8, 0x4e, 0x6a, 0x9f, 0xb3, 0xba, 0xf8, 0xa4, 0x5e,
	0xea, 0x4e, 0xf4, 0xa4, 0xd9, 0xad, 0x49, 0x1f, 0xc6, 0x18, 0x92, 0xc7, 0x6a, 0xd4, 0x5f, 0xe8,
	0xd8, 0x19, 0xde, 0x28, 0xe8, 0x1d, 0x56, 0x17, 0xaa, 0xd4, 0xa1, 0xf6, 0x54, 0x85, 0x22, 0xdc,
	0x69, 0x70, 0xaa, 0xad, 0x9b, 0xd8, 0xb0, 0x8a, 0x0d, 0x0f, 0xb1, 0xe1, 0x37, 0x50, 0xd5, 0xc9,
	0x0b, 0xc5, 0xfb, 0xb6, 0xf3, 0xa3, 0xb2, 0x92, 0x57, 0x5d, 0x86, 0x73, 0xe0, 0x64, 0xc8, 0xd8,
	0x7c, 0x9e, 0x8b, 0x62, 0x41, 0xe4, 0xba, 0x61, 0x42, 0x37, 0x88, 0xd4, 0x90, 0x93, 0xf9, 0x66,
	0xef, 0xa1, 0xed, 0xde, 0x43, 0x3f, 0xf6, 0x1e, 0xba, 0x39, 0x78, 0xd6, 0xf6, 0xe0, 0x59, 0xdf,
	0x0f, 0x9e, 0xf5, 0xf1, 0xe9, 0x3f, 0xa8, 0xf7, 0x7a, 0x2f, 0xe6, 0x14, 0x88, 0xd9, 0x10, 0x72,
	0x4d, 0x86, 0x05, 0xd2, 0xc8, 0xec, 0x4c, 0x3f, 0xf1, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0xdb, 0x55, 0xcf, 0x57, 0x02, 0x00, 0x00,
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLock(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLock(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.LockId != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.LockId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockId != 0 {
		n += 1 + sovLock(uint64(m.LockId))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovLock(uint64(l))
		}
	}
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockId", wireType)
			}
			m.LockId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)