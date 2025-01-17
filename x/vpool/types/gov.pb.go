// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vpool/v1/gov.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type CreatePoolProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// pair represents the pair of the vpool.
	Pair string `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	// trade_limit_ratio represents the limit on trading amounts.
	TradeLimitRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=trade_limit_ratio,json=tradeLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"trade_limit_ratio"`
	// quote_asset_reserve is the amount of quote asset the pool will be initialized with.
	QuoteAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_asset_reserve"`
	// base_asset_reserve is the amount of base asset the pool will be initialized with.
	BaseAssetReserve github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=base_asset_reserve,json=baseAssetReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_asset_reserve"`
	// fluctuation_limit_ratio represents the maximum price
	// percentage difference a trade can create on the pool.
	FluctuationLimitRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=fluctuation_limit_ratio,json=fluctuationLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fluctuation_limit_ratio"`
	// max_oracle_spread_ratio represents the maximum price percentage
	// difference that can exist between oracle price and vpool prices after a trade.
	MaxOracleSpreadRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=max_oracle_spread_ratio,json=maxOracleSpreadRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_oracle_spread_ratio"`
	// maintenance_margin_ratio
	MaintenanceMarginRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=maintenance_margin_ratio,json=maintenanceMarginRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maintenance_margin_ratio"`
	// max_leverage
	MaxLeverage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=max_leverage,json=maxLeverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_leverage"`
}

func (m *CreatePoolProposal) Reset()         { *m = CreatePoolProposal{} }
func (m *CreatePoolProposal) String() string { return proto.CompactTextString(m) }
func (*CreatePoolProposal) ProtoMessage()    {}
func (*CreatePoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a393460ab414204, []int{0}
}
func (m *CreatePoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePoolProposal.Merge(m, src)
}
func (m *CreatePoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreatePoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePoolProposal proto.InternalMessageInfo

func (m *CreatePoolProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePoolProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreatePoolProposal) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePoolProposal)(nil), "nibiru.vpool.v1.CreatePoolProposal")
}

func init() { proto.RegisterFile("vpool/v1/gov.proto", fileDescriptor_8a393460ab414204) }

var fileDescriptor_8a393460ab414204 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd3, 0xb1, 0x6e, 0xd4, 0x30,
	0x18, 0x07, 0xf0, 0x0b, 0xb4, 0x07, 0x75, 0x91, 0x4a, 0xcd, 0x41, 0x23, 0x86, 0xb4, 0x62, 0x40,
	0x48, 0x88, 0x58, 0x15, 0x4f, 0x40, 0x0b, 0x5b, 0x81, 0x72, 0x6c, 0x15, 0x22, 0xfa, 0x2e, 0xf7,
	0x35, 0x67, 0x61, 0xfb, 0x0b, 0xb6, 0x13, 0x1d, 0x2f, 0xc0, 0xcc, 0x63, 0x75, 0xec, 0x88, 0x18,
	0x2a, 0x74, 0xf7, 0x22, 0xc8, 0xce, 0x0d, 0x29, 0x63, 0xa6, 0x38, 0xfe, 0xac, 0xdf, 0x5f, 0xb6,
	0xf4, 0x67, 0xbc, 0xad, 0x89, 0x94, 0x68, 0x8f, 0x45, 0x45, 0x6d, 0x5e, 0x5b, 0xf2, 0xc4, 0xf7,
	0x8c, 0x9c, 0x49, 0xdb, 0xe4, 0x71, 0x94, 0xb7, 0xc7, 0x4f, 0x27, 0x15, 0x55, 0x14, 0x67, 0x22,
	0xac, 0xba, 0x63, 0xcf, 0x7e, 0x8e, 0x19, 0x3f, 0xb5, 0x08, 0x1e, 0xcf, 0x89, 0xd4, 0xb9, 0xa5,
	0x9a, 0x1c, 0x28, 0x3e, 0x61, 0xdb, 0x5e, 0x7a, 0x85, 0x69, 0x72, 0x94, 0xbc, 0xd8, 0x99, 0x76,
	0x3f, 0xfc, 0x88, 0xed, 0xce, 0xd1, 0x95, 0x56, 0xd6, 0x5e, 0x92, 0x49, 0xef, 0xc4, 0x59, 0x7f,
	0x8b, 0x73, 0xb6, 0x55, 0x83, 0xb4, 0xe9, 0xdd, 0x38, 0x8a, 0x6b, 0x7e, 0xc1, 0xf6, 0xbd, 0x85,
	0x39, 0x16, 0x4a, 0x6a, 0xe9, 0x0b, 0x0b, 0x5e, 0x52, 0xba, 0x15, 0x0e, 0x9c, 0xe4, 0x57, 0x37,
	0x87, 0xa3, 0x3f, 0x37, 0x87, 0xcf, 0x2b, 0xe9, 0x17, 0xcd, 0x2c, 0x2f, 0x49, 0x8b, 0x92, 0x9c,
	0x26, 0xb7, 0xf9, 0xbc, 0x72, 0xf3, 0x6f, 0xc2, 0xff, 0xa8, 0xd1, 0xe5, 0x6f, 0xb1, 0x9c, 0xee,
	0x45, 0xe8, 0x2c, 0x38, 0xd3, 0xc0, 0xf0, 0xaf, 0xec, 0xd1, 0xf7, 0x86, 0x3c, 0x16, 0xe0, 0x1c,
	0xfa, 0xc2, 0xa2, 0x43, 0xdb, 0x62, 0xba, 0x3d, 0x48, 0xdf, 0x8f, 0xd4, 0x9b, 0x20, 0x4d, 0x3b,
	0x88, 0x7f, 0x61, 0x7c, 0x06, 0xee, 0x7f, 0x7e, 0x3c, 0x88, 0x7f, 0x18, 0xa4, 0x5b, 0xfa, 0x25,
	0x3b, 0xb8, 0x54, 0x4d, 0xe9, 0x9b, 0x70, 0x17, 0x73, 0xeb, 0x7d, 0xee, 0x0d, 0x8a, 0x78, 0xdc,
	0xe3, 0x7a, 0xaf, 0x84, 0xec, 0x40, 0xc3, 0xb2, 0x20, 0x0b, 0xa5, 0xc2, 0xc2, 0xd5, 0x16, 0x61,
	0xbe, 0xc9, 0xb9, 0x3f, 0x28, 0x67, 0xa2, 0x61, 0xf9, 0x31, 0x6a, 0x9f, 0x23, 0xd6, 0xc5, 0x2c,
	0x58, 0xaa, 0x41, 0x1a, 0x8f, 0x06, 0x4c, 0x89, 0x85, 0x06, 0x5b, 0x49, 0xb3, 0xc9, 0xd9, 0x19,
	0x94, 0xf3, 0xa4, 0xe7, 0xbd, 0x8f, 0x5c, 0x97, 0xf4, 0x89, 0x3d, 0x08, 0x17, 0x52, 0xd8, 0xa2,
	0x85, 0x0a, 0x53, 0x36, 0x48, 0xdf, 0xd5, 0xb0, 0x3c, 0xdb, 0x10, 0x27, 0xef, 0xae, 0x56, 0x59,
	0x72, 0xbd, 0xca, 0x92, 0xbf, 0xab, 0x2c, 0xf9, 0xb5, 0xce, 0x46, 0xd7, 0xeb, 0x6c, 0xf4, 0x7b,
	0x9d, 0x8d, 0x2e, 0x5e, 0xf6, 0xb8, 0x0f, 0xb1, 0x54, 0xa7, 0x0b, 0x90, 0x46, 0x74, 0x05, 0x13,
	0x4b, 0xd1, 0xb5, 0x2f, 0xba, 0xb3, 0x71, 0xac, 0xd5, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0xe9, 0xb5, 0x83, 0x93, 0x03, 0x00, 0x00,
}

func (m *CreatePoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLeverage.Size()
		i -= size
		if _, err := m.MaxLeverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.MaintenanceMarginRatio.Size()
		i -= size
		if _, err := m.MaintenanceMarginRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.MaxOracleSpreadRatio.Size()
		i -= size
		if _, err := m.MaxOracleSpreadRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.FluctuationLimitRatio.Size()
		i -= size
		if _, err := m.FluctuationLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.BaseAssetReserve.Size()
		i -= size
		if _, err := m.BaseAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.QuoteAssetReserve.Size()
		i -= size
		if _, err := m.QuoteAssetReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TradeLimitRatio.Size()
		i -= size
		if _, err := m.TradeLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreatePoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.TradeLimitRatio.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.QuoteAssetReserve.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.BaseAssetReserve.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.FluctuationLimitRatio.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.MaxOracleSpreadRatio.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.MaintenanceMarginRatio.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.MaxLeverage.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreatePoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseAssetReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FluctuationLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FluctuationLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOracleSpreadRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxOracleSpreadRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintenanceMarginRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaintenanceMarginRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLeverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLeverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
