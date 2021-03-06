// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/Token.proto

package types

import (
	fmt "fmt"
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

type Token struct {
	Creator        string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id             uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Symbol         string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OriginalSymbol string `protobuf:"bytes,4,opt,name=originalSymbol,proto3" json:"originalSymbol,omitempty"`
	Desc           string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	WholeName      string `protobuf:"bytes,6,opt,name=wholeName,proto3" json:"wholeName,omitempty"`
	TotalSupply    int64  `protobuf:"varint,7,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	Own            string `protobuf:"bytes,8,opt,name=own,proto3" json:"own,omitempty"`
	Mintable       bool   `protobuf:"varint,9,opt,name=mintable,proto3" json:"mintable,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b24d611b3d4c8e, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Token) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Token) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Token) GetOriginalSymbol() string {
	if m != nil {
		return m.OriginalSymbol
	}
	return ""
}

func (m *Token) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Token) GetWholeName() string {
	if m != nil {
		return m.WholeName
	}
	return ""
}

func (m *Token) GetTotalSupply() int64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *Token) GetOwn() string {
	if m != nil {
		return m.Own
	}
	return ""
}

func (m *Token) GetMintable() bool {
	if m != nil {
		return m.Mintable
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "oracleNetworkProtocol.token.token.Token")
}

func init() { proto.RegisterFile("token/Token.proto", fileDescriptor_73b24d611b3d4c8e) }

var fileDescriptor_73b24d611b3d4c8e = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4b, 0x4e, 0x84, 0x30,
	0x18, 0x9e, 0xce, 0x9b, 0x9a, 0x4c, 0xb4, 0x31, 0xa6, 0x99, 0x98, 0x06, 0x5d, 0x18, 0x56, 0xb0,
	0xd0, 0x13, 0x78, 0x00, 0x62, 0x18, 0x57, 0xee, 0x0a, 0x34, 0x4c, 0x33, 0x85, 0x9f, 0x94, 0x4e,
	0x90, 0x5b, 0x78, 0x2c, 0x97, 0xb3, 0x74, 0x69, 0xe0, 0x08, 0x5e, 0xc0, 0x50, 0xc6, 0x47, 0x8c,
	0x9b, 0xbf, 0xdf, 0xb3, 0x8b, 0x0f, 0x9f, 0x19, 0xd8, 0x89, 0x22, 0x78, 0xec, 0xaf, 0x5f, 0x6a,
	0x30, 0x40, 0xae, 0x40, 0xf3, 0x44, 0x89, 0x50, 0x98, 0x1a, 0xf4, 0xee, 0xa1, 0xd7, 0x12, 0x50,
	0xbe, 0x0d, 0x0e, 0x77, 0x7d, 0x9e, 0x41, 0x06, 0x36, 0x1d, 0xf4, 0x68, 0x28, 0x5e, 0x7f, 0x20,
	0x3c, 0xb3, 0x1f, 0x11, 0x8a, 0x17, 0x89, 0x16, 0xdc, 0x80, 0xa6, 0xc8, 0x45, 0x9e, 0x13, 0x7d,
	0x51, 0xb2, 0xc2, 0x63, 0x99, 0xd2, 0xb1, 0x8b, 0xbc, 0x69, 0x34, 0x96, 0x29, 0xb9, 0xc0, 0xf3,
	0xaa, 0xc9, 0x63, 0x50, 0x74, 0x62, 0x83, 0x47, 0x46, 0x6e, 0xf0, 0x0a, 0xb4, 0xcc, 0x64, 0xc1,
	0xd5, 0x66, 0xf0, 0xa7, 0xd6, 0xff, 0xa3, 0x12, 0x82, 0xa7, 0xa9, 0xa8, 0x12, 0x3a, 0xb3, 0xae,
	0xc5, 0xe4, 0x12, 0x3b, 0xf5, 0x16, 0x94, 0x08, 0x79, 0x2e, 0xe8, 0xdc, 0x1a, 0x3f, 0x02, 0x71,
	0xf1, 0x89, 0x01, 0xc3, 0xd5, 0x66, 0x5f, 0x96, 0xaa, 0xa1, 0x0b, 0x17, 0x79, 0xb3, 0xe8, 0xb7,
	0x44, 0x4e, 0xf1, 0x04, 0xea, 0x82, 0x2e, 0x6d, 0xb3, 0x87, 0x64, 0x8d, 0x97, 0xb9, 0x2c, 0x0c,
	0x8f, 0x95, 0xa0, 0x8e, 0x8b, 0xbc, 0x65, 0xf4, 0xcd, 0xef, 0xc3, 0xd7, 0x96, 0xa1, 0x43, 0xcb,
	0xd0, 0x7b, 0xcb, 0xd0, 0x4b, 0xc7, 0x46, 0x87, 0x8e, 0x8d, 0xde, 0x3a, 0x36, 0x7a, 0xba, 0xcb,
	0xa4, 0xd9, 0xee, 0x63, 0x3f, 0x81, 0x3c, 0xf8, 0x77, 0xd3, 0x60, 0x18, 0xff, 0xf9, 0xf8, 0x9a,
	0xa6, 0x14, 0x55, 0x3c, 0xb7, 0x63, 0xde, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xd3, 0x47,
	0x46, 0x9a, 0x01, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.Own) > 0 {
		i -= len(m.Own)
		copy(dAtA[i:], m.Own)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Own)))
		i--
		dAtA[i] = 0x42
	}
	if m.TotalSupply != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.TotalSupply))
		i--
		dAtA[i] = 0x38
	}
	if len(m.WholeName) > 0 {
		i -= len(m.WholeName)
		copy(dAtA[i:], m.WholeName)
		i = encodeVarintToken(dAtA, i, uint64(len(m.WholeName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OriginalSymbol) > 0 {
		i -= len(m.OriginalSymbol)
		copy(dAtA[i:], m.OriginalSymbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.OriginalSymbol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovToken(uint64(m.Id))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.OriginalSymbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.WholeName)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.TotalSupply != 0 {
		n += 1 + sovToken(uint64(m.TotalSupply))
	}
	l = len(m.Own)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Mintable {
		n += 2
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
		fieldNum := int64(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WholeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WholeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			m.TotalSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSupply |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Own", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Own = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mintable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
