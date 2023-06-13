// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/module/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// Ticker is the standard representation of a pair of assets, where one (Asset1)
// is priced in terms of the other (Asset2)
type Ticker struct {
	Asset1 string `protobuf:"bytes,1,opt,name=Asset1,proto3" json:"Asset1,omitempty"`
	Asset2 string `protobuf:"bytes,2,opt,name=Asset2,proto3" json:"Asset2,omitempty"`
}

func (m *Ticker) Reset()         { *m = Ticker{} }
func (m *Ticker) String() string { return proto.CompactTextString(m) }
func (*Ticker) ProtoMessage()    {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f99f7f904ffca77, []int{0}
}
func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(m, src)
}
func (m *Ticker) XXX_Size() int {
	return m.Size()
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetAsset1() string {
	if m != nil {
		return m.Asset1
	}
	return ""
}

func (m *Ticker) GetAsset2() string {
	if m != nil {
		return m.Asset2
	}
	return ""
}

// TickerPrice is the representation of a SpotPrice for a Ticker, where price
// represents the price of Asset1 in terms of Asset2
type TickerPrice struct {
	Price     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"price"`
	Timestamp time.Time                              `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *TickerPrice) Reset()         { *m = TickerPrice{} }
func (m *TickerPrice) String() string { return proto.CompactTextString(m) }
func (*TickerPrice) ProtoMessage()    {}
func (*TickerPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f99f7f904ffca77, []int{1}
}
func (m *TickerPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickerPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickerPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickerPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickerPrice.Merge(m, src)
}
func (m *TickerPrice) XXX_Size() int {
	return m.Size()
}
func (m *TickerPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_TickerPrice.DiscardUnknown(m)
}

var xxx_messageInfo_TickerPrice proto.InternalMessageInfo

func (m *TickerPrice) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// TickerGenesis is the information necessary for initialization of a Ticker,
// both a Ticker + TickerPrice object is required
type TickerGenesis struct {
	Ticker      *Ticker      `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	TickerPrice *TickerPrice `protobuf:"bytes,2,opt,name=ticker_price,json=tickerPrice,proto3" json:"ticker_price,omitempty"`
}

func (m *TickerGenesis) Reset()         { *m = TickerGenesis{} }
func (m *TickerGenesis) String() string { return proto.CompactTextString(m) }
func (*TickerGenesis) ProtoMessage()    {}
func (*TickerGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f99f7f904ffca77, []int{2}
}
func (m *TickerGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickerGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickerGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickerGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickerGenesis.Merge(m, src)
}
func (m *TickerGenesis) XXX_Size() int {
	return m.Size()
}
func (m *TickerGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_TickerGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_TickerGenesis proto.InternalMessageInfo

func (m *TickerGenesis) GetTicker() *Ticker {
	if m != nil {
		return m.Ticker
	}
	return nil
}

func (m *TickerGenesis) GetTickerPrice() *TickerPrice {
	if m != nil {
		return m.TickerPrice
	}
	return nil
}

// GenesisState is the genesis-state for the x/oracle module, it takes a set of
// predefined Tickers + prices
type GenesisState struct {
	TickerGenesis []*TickerGenesis `protobuf:"bytes,1,rep,name=ticker_genesis,json=tickerGenesis,proto3" json:"ticker_genesis,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f99f7f904ffca77, []int{3}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTickerGenesis() []*TickerGenesis {
	if m != nil {
		return m.TickerGenesis
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticker)(nil), "Ticker")
	proto.RegisterType((*TickerPrice)(nil), "TickerPrice")
	proto.RegisterType((*TickerGenesis)(nil), "TickerGenesis")
	proto.RegisterType((*GenesisState)(nil), "GenesisState")
}

func init() { proto.RegisterFile("slinky/module/v1/genesis.proto", fileDescriptor_4f99f7f904ffca77) }

var fileDescriptor_4f99f7f904ffca77 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcb, 0x4e, 0xf2, 0x40,
	0x14, 0xee, 0xfc, 0xbf, 0x56, 0x99, 0x02, 0x8b, 0xc6, 0x18, 0x64, 0xd1, 0x92, 0x2e, 0x0c, 0x2e,
	0x98, 0x09, 0x35, 0x24, 0x2e, 0xdc, 0x58, 0x63, 0x0c, 0x3b, 0x53, 0x59, 0xb9, 0x21, 0x50, 0xc6,
	0xda, 0xf4, 0x32, 0x4d, 0x67, 0x20, 0xf2, 0x16, 0xbc, 0x84, 0x6f, 0xe0, 0x43, 0xb0, 0x24, 0xae,
	0x8c, 0x0b, 0x34, 0xf0, 0x22, 0x86, 0xce, 0x00, 0x75, 0xd5, 0x73, 0xbe, 0x99, 0x73, 0xbe, 0x4b,
	0x07, 0x1a, 0x2c, 0x0a, 0x92, 0x70, 0x8a, 0x63, 0x3a, 0x1a, 0x47, 0x04, 0x4f, 0xda, 0xd8, 0x27,
	0x09, 0x61, 0x01, 0x43, 0x69, 0x46, 0x39, 0xad, 0x9f, 0xf8, 0xd4, 0xa7, 0x79, 0x89, 0x37, 0x95,
	0x44, 0x4d, 0x9f, 0x52, 0x3f, 0x22, 0x38, 0xef, 0x86, 0xe3, 0x67, 0xcc, 0x83, 0x98, 0x30, 0x3e,
	0x88, 0x53, 0x79, 0xe1, 0xcc, 0xa3, 0x2c, 0xa6, 0xac, 0x2f, 0x26, 0x45, 0x23, 0x8e, 0xac, 0x2b,
	0xa8, 0xf6, 0x02, 0x2f, 0x24, 0x99, 0x7e, 0x0a, 0xd5, 0x1b, 0xc6, 0x08, 0x6f, 0xd7, 0x40, 0x03,
	0x34, 0x4b, 0xae, 0xec, 0x76, 0xb8, 0x5d, 0xfb, 0x57, 0xc0, 0x6d, 0xeb, 0x0d, 0x40, 0x4d, 0x8c,
	0x3e, 0x64, 0x81, 0x47, 0x74, 0x17, 0x1e, 0xa6, 0x9b, 0x42, 0x8c, 0x3b, 0xd7, 0xf3, 0xa5, 0xa9,
	0x7c, 0x2d, 0xcd, 0x73, 0x3f, 0xe0, 0x2f, 0xe3, 0x21, 0xf2, 0x68, 0x2c, 0x99, 0xe5, 0xa7, 0xc5,
	0x46, 0x21, 0xe6, 0xd3, 0x94, 0x30, 0xd4, 0x4d, 0xf8, 0xc7, 0x7b, 0x0b, 0x4a, 0x61, 0xdd, 0x84,
	0xbb, 0x62, 0x95, 0xee, 0xc0, 0xd2, 0xce, 0x4b, 0x4e, 0xaf, 0xd9, 0x75, 0x24, 0xdc, 0xa2, 0xad,
	0x5b, 0xd4, 0xdb, 0xde, 0x70, 0x8e, 0x37, 0x9c, 0xb3, 0x6f, 0x13, 0xb8, 0xfb, 0x31, 0xcb, 0x87,
	0x15, 0x21, 0xf3, 0x5e, 0x44, 0xa9, 0x9b, 0x50, 0xe5, 0x39, 0x90, 0x2b, 0xd5, 0xec, 0x23, 0x24,
	0xce, 0x5d, 0x09, 0xeb, 0x1d, 0x58, 0x16, 0x55, 0x5f, 0x18, 0x12, 0xc4, 0x65, 0x54, 0x70, 0xeb,
	0x1c, 0xcc, 0x97, 0x26, 0x70, 0x35, 0xbe, 0x87, 0xac, 0x3b, 0x58, 0x96, 0x14, 0x8f, 0x7c, 0xc0,
	0x89, 0xde, 0x81, 0x55, 0xb9, 0x46, 0xfe, 0xc4, 0x1a, 0x68, 0xfc, 0x6f, 0x6a, 0x76, 0x15, 0xfd,
	0xd1, 0xe3, 0x56, 0x78, 0xb1, 0x75, 0x6e, 0xe7, 0x2b, 0x03, 0x2c, 0x56, 0x06, 0xf8, 0x59, 0x19,
	0x60, 0xb6, 0x36, 0x94, 0xc5, 0xda, 0x50, 0x3e, 0xd7, 0x86, 0xf2, 0x74, 0x51, 0x88, 0x92, 0x85,
	0x41, 0xda, 0x8a, 0xc9, 0x04, 0xcb, 0x17, 0xf3, 0x8a, 0x69, 0x36, 0xf0, 0x22, 0x22, 0x12, 0x1d,
	0xaa, 0x79, 0x3a, 0x97, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xdd, 0x71, 0x39, 0x51, 0x02,
	0x00, 0x00,
}

func (m *Ticker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ticker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ticker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Asset2) > 0 {
		i -= len(m.Asset2)
		copy(dAtA[i:], m.Asset2)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Asset2)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Asset1) > 0 {
		i -= len(m.Asset1)
		copy(dAtA[i:], m.Asset1)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Asset1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TickerPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickerPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickerPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TickerGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickerGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickerGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TickerPrice != nil {
		{
			size, err := m.TickerPrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Ticker != nil {
		{
			size, err := m.Ticker.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TickerGenesis) > 0 {
		for iNdEx := len(m.TickerGenesis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickerGenesis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ticker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset1)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Asset2)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *TickerPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *TickerGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ticker != nil {
		l = m.Ticker.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TickerPrice != nil {
		l = m.TickerPrice.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TickerGenesis) > 0 {
		for _, e := range m.TickerGenesis {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ticker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Ticker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ticker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *TickerPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: TickerPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickerPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *TickerGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: TickerGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickerGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ticker == nil {
				m.Ticker = &Ticker{}
			}
			if err := m.Ticker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickerPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TickerPrice == nil {
				m.TickerPrice = &TickerPrice{}
			}
			if err := m.TickerPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickerGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickerGenesis = append(m.TickerGenesis, &TickerGenesis{})
			if err := m.TickerGenesis[len(m.TickerGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)