// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquidityincentive/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the liquidityincentive module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params     Params  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	EpochList  []Epoch `protobuf:"bytes,2,rep,name=epochList,proto3" json:"epochList"`
	EpochCount uint64  `protobuf:"varint,3,opt,name=epochCount,proto3" json:"epochCount,omitempty"`
	GaugeList  []Gauge `protobuf:"bytes,4,rep,name=gaugeList,proto3" json:"gaugeList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f3633f334ba0bf3, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEpochList() []Epoch {
	if m != nil {
		return m.EpochList
	}
	return nil
}

func (m *GenesisState) GetEpochCount() uint64 {
	if m != nil {
		return m.EpochCount
	}
	return 0
}

func (m *GenesisState) GetGaugeList() []Gauge {
	if m != nil {
		return m.GaugeList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sunrise.liquidityincentive.GenesisState")
}

func init() {
	proto.RegisterFile("sunrise/liquidityincentive/genesis.proto", fileDescriptor_5f3633f334ba0bf3)
}

var fileDescriptor_5f3633f334ba0bf3 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0xc9, 0x2c, 0xa9, 0xcc, 0xcc, 0x4b,
	0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0xaa, 0xd4, 0xc3, 0x54, 0x29, 0x25, 0x98, 0x98, 0x9b,
	0x99, 0x97, 0xaf, 0x0f, 0x26, 0x21, 0xca, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d,
	0x10, 0x0b, 0x2a, 0xaa, 0x8e, 0xc7, 0xba, 0x82, 0xc4, 0xa2, 0xc4, 0x5c, 0xa8, 0x6d, 0x52, 0x6a,
	0x78, 0x14, 0xa6, 0x16, 0xe4, 0x27, 0x67, 0x10, 0xa1, 0x2e, 0x3d, 0xb1, 0x34, 0x3d, 0x15, 0xa2,
	0x4e, 0xa9, 0x99, 0x89, 0x8b, 0xc7, 0x1d, 0xe2, 0x9f, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x57,
	0x2e, 0x36, 0x88, 0x85, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x4a, 0x7a, 0xb8, 0xfd, 0xa7,
	0x17, 0x00, 0x56, 0xe9, 0xc4, 0x79, 0xe2, 0x9e, 0x3c, 0xc3, 0x8a, 0xe7, 0x1b, 0xb4, 0x18, 0x83,
	0xa0, 0x9a, 0x85, 0x5c, 0xb9, 0x38, 0xc1, 0xce, 0xf1, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x52, 0x60,
	0xd6, 0xe0, 0x36, 0x52, 0xc4, 0x67, 0x92, 0x2b, 0x48, 0xb1, 0x13, 0x0b, 0xc8, 0xa0, 0x20, 0x84,
	0x4e, 0x21, 0x39, 0x2e, 0x2e, 0x30, 0xc7, 0x39, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x25, 0x08, 0x49, 0x04, 0x64, 0x0d, 0xd8, 0x37, 0x60, 0x6b, 0x58, 0x08, 0x5b, 0xe3, 0x0e,
	0x52, 0x0c, 0xb3, 0x06, 0xae, 0xd3, 0x29, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0xac, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0xe6,
	0xe6, 0x24, 0x56, 0xa6, 0x16, 0xc1, 0x38, 0xfa, 0x15, 0xd8, 0x42, 0xb8, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0x1c, 0xc4, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xcf, 0x09, 0x67,
	0x4c, 0x02, 0x00, 0x00,
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
	if len(m.GaugeList) > 0 {
		for iNdEx := len(m.GaugeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GaugeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EpochCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EpochList) > 0 {
		for iNdEx := len(m.EpochList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.EpochList) > 0 {
		for _, e := range m.EpochList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EpochCount != 0 {
		n += 1 + sovGenesis(uint64(m.EpochCount))
	}
	if len(m.GaugeList) > 0 {
		for _, e := range m.GaugeList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochList", wireType)
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
			m.EpochList = append(m.EpochList, Epoch{})
			if err := m.EpochList[len(m.EpochList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCount", wireType)
			}
			m.EpochCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeList", wireType)
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
			m.GaugeList = append(m.GaugeList, Gauge{})
			if err := m.GaugeList[len(m.GaugeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
