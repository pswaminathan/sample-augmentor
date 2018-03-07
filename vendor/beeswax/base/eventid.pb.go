// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beeswax/base/eventid.proto

/*
	Package base is a generated protocol buffer package.

	It is generated from these files:
		beeswax/base/eventid.proto

	It has these top-level messages:
		EventId
		AdGroupId
*/
package base

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EventId struct {
	Hostid           *uint32 `protobuf:"varint,1,opt,name=hostid" json:"hostid,omitempty"`
	Tid              *uint32 `protobuf:"varint,2,opt,name=tid" json:"tid,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EventId) Reset()                    { *m = EventId{} }
func (m *EventId) String() string            { return proto.CompactTextString(m) }
func (*EventId) ProtoMessage()               {}
func (*EventId) Descriptor() ([]byte, []int) { return fileDescriptorEventid, []int{0} }

func (m *EventId) GetHostid() uint32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *EventId) GetTid() uint32 {
	if m != nil && m.Tid != nil {
		return *m.Tid
	}
	return 0
}

func (m *EventId) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// Next Tag: 5
type AdGroupId struct {
	BuzzKey          *string `protobuf:"bytes,1,opt,name=buzz_key,json=buzzKey" json:"buzz_key,omitempty"`
	Accountid        *uint64 `protobuf:"varint,2,opt,name=accountid" json:"accountid,omitempty"`
	Campaignid       *uint64 `protobuf:"varint,3,opt,name=campaignid" json:"campaignid,omitempty"`
	Lineitemid       *uint64 `protobuf:"varint,4,opt,name=lineitemid" json:"lineitemid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdGroupId) Reset()                    { *m = AdGroupId{} }
func (m *AdGroupId) String() string            { return proto.CompactTextString(m) }
func (*AdGroupId) ProtoMessage()               {}
func (*AdGroupId) Descriptor() ([]byte, []int) { return fileDescriptorEventid, []int{1} }

func (m *AdGroupId) GetBuzzKey() string {
	if m != nil && m.BuzzKey != nil {
		return *m.BuzzKey
	}
	return ""
}

func (m *AdGroupId) GetAccountid() uint64 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *AdGroupId) GetCampaignid() uint64 {
	if m != nil && m.Campaignid != nil {
		return *m.Campaignid
	}
	return 0
}

func (m *AdGroupId) GetLineitemid() uint64 {
	if m != nil && m.Lineitemid != nil {
		return *m.Lineitemid
	}
	return 0
}

func init() {
	proto.RegisterType((*EventId)(nil), "base.EventId")
	proto.RegisterType((*AdGroupId)(nil), "base.AdGroupId")
}
func (m *EventId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hostid != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Hostid))
	}
	if m.Tid != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Tid))
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AdGroupId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdGroupId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BuzzKey != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEventid(dAtA, i, uint64(len(*m.BuzzKey)))
		i += copy(dAtA[i:], *m.BuzzKey)
	}
	if m.Accountid != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Accountid))
	}
	if m.Campaignid != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Campaignid))
	}
	if m.Lineitemid != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEventid(dAtA, i, uint64(*m.Lineitemid))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEventid(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EventId) Size() (n int) {
	var l int
	_ = l
	if m.Hostid != nil {
		n += 1 + sovEventid(uint64(*m.Hostid))
	}
	if m.Tid != nil {
		n += 1 + sovEventid(uint64(*m.Tid))
	}
	if m.Timestamp != nil {
		n += 1 + sovEventid(uint64(*m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AdGroupId) Size() (n int) {
	var l int
	_ = l
	if m.BuzzKey != nil {
		l = len(*m.BuzzKey)
		n += 1 + l + sovEventid(uint64(l))
	}
	if m.Accountid != nil {
		n += 1 + sovEventid(uint64(*m.Accountid))
	}
	if m.Campaignid != nil {
		n += 1 + sovEventid(uint64(*m.Campaignid))
	}
	if m.Lineitemid != nil {
		n += 1 + sovEventid(uint64(*m.Lineitemid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEventid(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEventid(x uint64) (n int) {
	return sovEventid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostid", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Hostid = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tid = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Timestamp = &v
		default:
			iNdEx = preIndex
			skippy, err := skipEventid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdGroupId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdGroupId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdGroupId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuzzKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEventid
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.BuzzKey = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accountid", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Accountid = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Campaignid", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Campaignid = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lineitemid", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lineitemid = &v
		default:
			iNdEx = preIndex
			skippy, err := skipEventid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEventid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventid
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
					return 0, ErrIntOverflowEventid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEventid
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEventid
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEventid
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEventid(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEventid = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventid   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("beeswax/base/eventid.proto", fileDescriptorEventid) }

var fileDescriptorEventid = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0x06, 0x70, 0x62, 0x07, 0xeb, 0x3c, 0x10, 0x4a, 0x16, 0x32, 0x8a, 0x84, 0xd2, 0x55, 0x57,
	0xd3, 0x33, 0x28, 0x88, 0x14, 0x57, 0xce, 0x05, 0x24, 0x93, 0x3c, 0xf4, 0xa1, 0xf9, 0x43, 0x93,
	0x51, 0xdb, 0xb5, 0x87, 0x73, 0xe9, 0x11, 0x64, 0x4e, 0x22, 0x6f, 0x1c, 0x19, 0x77, 0xf9, 0x7e,
	0x1f, 0x7c, 0x21, 0x81, 0x8b, 0x16, 0x31, 0xbd, 0xe9, 0xf7, 0x4d, 0xab, 0x13, 0x6e, 0xf0, 0x15,
	0x7d, 0x26, 0x5b, 0xc7, 0x5d, 0xc8, 0x41, 0x16, 0x6c, 0xab, 0x7b, 0x98, 0xdf, 0x30, 0x6f, 0xad,
	0x3c, 0x83, 0xe3, 0xa7, 0x90, 0x32, 0xd9, 0x4a, 0x2c, 0xc5, 0xfa, 0xb4, 0x19, 0x93, 0x5c, 0xc0,
	0x8c, 0xf1, 0x68, 0x40, 0x3e, 0xca, 0x4b, 0x28, 0x33, 0x39, 0x4c, 0x59, 0xbb, 0x58, 0xcd, 0x96,
	0x62, 0x5d, 0x34, 0x13, 0xac, 0x3e, 0x04, 0x94, 0x57, 0xf6, 0x76, 0x17, 0xba, 0xb8, 0xb5, 0xf2,
	0x1c, 0x4e, 0xda, 0xee, 0x70, 0x78, 0x78, 0xc6, 0xfd, 0xb0, 0x5b, 0x36, 0x73, 0xce, 0x77, 0xb8,
	0xe7, 0x19, 0x6d, 0x4c, 0xe8, 0xfc, 0xdf, 0x7c, 0xd1, 0x4c, 0x20, 0x15, 0x80, 0xd1, 0x2e, 0x6a,
	0x7a, 0xf4, 0x64, 0xc7, 0x5b, 0xfe, 0x09, 0xf7, 0x2f, 0xe4, 0x91, 0x32, 0x3a, 0xb2, 0x55, 0xf1,
	0xdb, 0x4f, 0x72, 0xad, 0x3e, 0x7b, 0x25, 0xbe, 0x7a, 0x25, 0xbe, 0x7b, 0x25, 0x60, 0x61, 0x82,
	0xab, 0xc7, 0xdf, 0xa8, 0xf9, 0xe5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb7, 0x36, 0x09,
	0x1c, 0x01, 0x00, 0x00,
}