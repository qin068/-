// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/denom_trace.proto

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

type DenomTrace struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Port    string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	Channel string `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	Origin  string `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (m *DenomTrace) Reset()         { *m = DenomTrace{} }
func (m *DenomTrace) String() string { return proto.CompactTextString(m) }
func (*DenomTrace) ProtoMessage()    {}
func (*DenomTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_d98bc4572f53ce26, []int{0}
}
func (m *DenomTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomTrace.Merge(m, src)
}
func (m *DenomTrace) XXX_Size() int {
	return m.Size()
}
func (m *DenomTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomTrace.DiscardUnknown(m)
}

var xxx_messageInfo_DenomTrace proto.InternalMessageInfo

func (m *DenomTrace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DenomTrace) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DenomTrace) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *DenomTrace) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *DenomTrace) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomTrace)(nil), "user.interchange.ibcdex.DenomTrace")
}

func init() { proto.RegisterFile("ibcdex/denom_trace.proto", fileDescriptor_d98bc4572f53ce26) }

var fileDescriptor_d98bc4572f53ce26 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x1b, 0x6d, 0x2b, 0x66, 0x0c, 0xa2, 0x99, 0x82, 0x38, 0x89, 0x43, 0x32, 0xf8, 0x06,
	0xd2, 0x27, 0x10, 0x27, 0x17, 0x49, 0xd3, 0x43, 0x1b, 0xb0, 0x49, 0x39, 0x4d, 0xa1, 0xae, 0x3e,
	0x81, 0x8f, 0xe5, 0xd8, 0xf1, 0x8e, 0x97, 0xf6, 0x45, 0x2e, 0x4d, 0x7b, 0xe1, 0x6e, 0xe7, 0xfb,
	0xcf, 0xcf, 0x0f, 0x1f, 0xe5, 0xb6, 0x34, 0x15, 0x8c, 0xaa, 0x02, 0xe7, 0xdb, 0xaf, 0x80, 0xda,
	0x80, 0xec, 0xd0, 0x07, 0xcf, 0x1e, 0x86, 0x1e, 0x50, 0x5a, 0x17, 0x00, 0x4d, 0xa3, 0x5d, 0x0d,
	0x72, 0xab, 0x3e, 0xfd, 0x12, 0x4a, 0x8b, 0xb5, 0xfe, 0xb1, 0xb6, 0x19, 0xa7, 0x37, 0x06, 0x41,
	0x07, 0x8f, 0x9c, 0x3c, 0x92, 0xe7, 0xdb, 0xf7, 0x33, 0xb2, 0x3b, 0x9a, 0x59, 0x57, 0xc1, 0xc8,
	0xaf, 0x62, 0xbe, 0x01, 0x63, 0x34, 0xed, 0x3c, 0x06, 0x7e, 0x1d, 0xc3, 0x78, 0xc7, 0x8d, 0x46,
	0x3b, 0x07, 0xdf, 0x3c, 0xdd, 0x37, 0x36, 0x64, 0xf7, 0x34, 0xf7, 0x68, 0x6b, 0xeb, 0x78, 0x16,
	0x1f, 0x3b, 0xbd, 0x15, 0xff, 0xb3, 0x20, 0xd3, 0x2c, 0xc8, 0x71, 0x16, 0xe4, 0x6f, 0x11, 0xc9,
	0xb4, 0x88, 0xe4, 0xb0, 0x88, 0xe4, 0xf3, 0xa5, 0xb6, 0xa1, 0x19, 0x4a, 0x69, 0x7c, 0xab, 0x56,
	0x05, 0x75, 0xa1, 0xa0, 0x46, 0xb5, 0xfb, 0x86, 0x9f, 0x0e, 0xfa, 0x32, 0x8f, 0xaa, 0xaf, 0xa7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xd1, 0x95, 0x31, 0x06, 0x01, 0x00, 0x00,
}

func (m *DenomTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenomTrace(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenomTrace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	return n
}

func sovDenomTrace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenomTrace(x uint64) (n int) {
	return sovDenomTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenomTrace
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
			return fmt.Errorf("proto: DenomTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
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
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
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
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
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
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
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
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
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
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenomTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenomTrace
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
func skipDenomTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenomTrace
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
					return 0, ErrIntOverflowDenomTrace
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
					return 0, ErrIntOverflowDenomTrace
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
				return 0, ErrInvalidLengthDenomTrace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenomTrace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenomTrace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenomTrace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenomTrace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenomTrace = fmt.Errorf("proto: unexpected end of group")
)
