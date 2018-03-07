// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beeswax/augment/augmentor.proto

/*
	Package augment is a generated protocol buffer package.

	It is generated from these files:
		beeswax/augment/augmentor.proto

	It has these top-level messages:
		AugmentorRequest
		AugmentorResponse
*/
package augment

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import openrtb2 "beeswax/openrtb"

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

// Stinger will issue a HTTP request to the augmentor.
// The following message will be present in the body of the HTTP request.
type AugmentorRequest struct {
	// Bid request received from the exchange after it has been processed by Beeswax.
	BidRequest       *openrtb2.BidRequest `protobuf:"bytes,1,opt,name=bid_request,json=bidRequest" json:"bid_request,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AugmentorRequest) Reset()                    { *m = AugmentorRequest{} }
func (m *AugmentorRequest) String() string            { return proto.CompactTextString(m) }
func (*AugmentorRequest) ProtoMessage()               {}
func (*AugmentorRequest) Descriptor() ([]byte, []int) { return fileDescriptorAugmentor, []int{0} }

func (m *AugmentorRequest) GetBidRequest() *openrtb2.BidRequest {
	if m != nil {
		return m.BidRequest
	}
	return nil
}

// Augmentor will respond to Stinger with a HTTP response.
// The following message will be present in the body of the HTTP response.
// Augmentor can choose not to bid at all in which case it must still
// respond to the HTTP request with a 204 status code.
//
// Also custom bidding agents can receive this data through the augmentor_data field in
// the BidRequestExtension.
type AugmentorResponse struct {
	// Array of Augmentor segments
	Segments         []*AugmentorResponse_Segment `protobuf:"bytes,1,rep,name=segments" json:"segments,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *AugmentorResponse) Reset()                    { *m = AugmentorResponse{} }
func (m *AugmentorResponse) String() string            { return proto.CompactTextString(m) }
func (*AugmentorResponse) ProtoMessage()               {}
func (*AugmentorResponse) Descriptor() ([]byte, []int) { return fileDescriptorAugmentor, []int{1} }

func (m *AugmentorResponse) GetSegments() []*AugmentorResponse_Segment {
	if m != nil {
		return m.Segments
	}
	return nil
}

type AugmentorResponse_Segment struct {
	// The Id of an augmentor segment.
	// This field is primarily used for line item targeting.
	// For targeting to work with the segment ids, it is necessary
	// to register the segment ids with Beeswax via the Buzz Segments API.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// This field can be used to pass additional information about the segment
	// that Stinger will forward to the bidding agent.
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AugmentorResponse_Segment) Reset()         { *m = AugmentorResponse_Segment{} }
func (m *AugmentorResponse_Segment) String() string { return proto.CompactTextString(m) }
func (*AugmentorResponse_Segment) ProtoMessage()    {}
func (*AugmentorResponse_Segment) Descriptor() ([]byte, []int) {
	return fileDescriptorAugmentor, []int{1, 0}
}

func (m *AugmentorResponse_Segment) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AugmentorResponse_Segment) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*AugmentorRequest)(nil), "augment.AugmentorRequest")
	proto.RegisterType((*AugmentorResponse)(nil), "augment.AugmentorResponse")
	proto.RegisterType((*AugmentorResponse_Segment)(nil), "augment.AugmentorResponse.Segment")
}
func (m *AugmentorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AugmentorRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BidRequest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAugmentor(dAtA, i, uint64(m.BidRequest.Size()))
		n1, err := m.BidRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AugmentorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AugmentorResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Segments) > 0 {
		for _, msg := range m.Segments {
			dAtA[i] = 0xa
			i++
			i = encodeVarintAugmentor(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AugmentorResponse_Segment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AugmentorResponse_Segment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAugmentor(dAtA, i, uint64(len(*m.Id)))
		i += copy(dAtA[i:], *m.Id)
	}
	if m.Value != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAugmentor(dAtA, i, uint64(len(*m.Value)))
		i += copy(dAtA[i:], *m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAugmentor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AugmentorRequest) Size() (n int) {
	var l int
	_ = l
	if m.BidRequest != nil {
		l = m.BidRequest.Size()
		n += 1 + l + sovAugmentor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AugmentorResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Segments) > 0 {
		for _, e := range m.Segments {
			l = e.Size()
			n += 1 + l + sovAugmentor(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AugmentorResponse_Segment) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = len(*m.Id)
		n += 1 + l + sovAugmentor(uint64(l))
	}
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovAugmentor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAugmentor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAugmentor(x uint64) (n int) {
	return sovAugmentor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AugmentorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAugmentor
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
			return fmt.Errorf("proto: AugmentorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AugmentorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAugmentor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAugmentor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BidRequest == nil {
				m.BidRequest = &openrtb2.BidRequest{}
			}
			if err := m.BidRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAugmentor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAugmentor
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
func (m *AugmentorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAugmentor
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
			return fmt.Errorf("proto: AugmentorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AugmentorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Segments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAugmentor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAugmentor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Segments = append(m.Segments, &AugmentorResponse_Segment{})
			if err := m.Segments[len(m.Segments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAugmentor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAugmentor
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
func (m *AugmentorResponse_Segment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAugmentor
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
			return fmt.Errorf("proto: Segment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Segment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAugmentor
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
				return ErrInvalidLengthAugmentor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Id = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAugmentor
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
				return ErrInvalidLengthAugmentor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Value = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAugmentor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAugmentor
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
func skipAugmentor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAugmentor
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
					return 0, ErrIntOverflowAugmentor
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
					return 0, ErrIntOverflowAugmentor
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
				return 0, ErrInvalidLengthAugmentor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAugmentor
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
				next, err := skipAugmentor(dAtA[start:])
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
	ErrInvalidLengthAugmentor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAugmentor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("beeswax/augment/augmentor.proto", fileDescriptorAugmentor) }

var fileDescriptorAugmentor = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4a, 0x4d, 0x2d,
	0x2e, 0x4f, 0xac, 0xd0, 0x4f, 0x2c, 0x4d, 0xcf, 0x4d, 0xcd, 0x2b, 0x81, 0xd1, 0xf9, 0x45, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0x01, 0x29, 0x59, 0x98, 0xca, 0xfc, 0x82, 0xd4,
	0xbc, 0xa2, 0x92, 0x24, 0x18, 0x0d, 0x51, 0xa7, 0xe4, 0xc1, 0x25, 0xe0, 0x08, 0xd3, 0x1a, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x64, 0xc2, 0xc5, 0x9d, 0x94, 0x99, 0x12, 0x5f, 0x04, 0xe1,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xeb, 0xc1, 0x34, 0x3a, 0x65, 0xa6, 0x40, 0x55,
	0x06, 0x71, 0x25, 0xc1, 0xd9, 0x4a, 0x2d, 0x8c, 0x5c, 0x82, 0x48, 0x46, 0x15, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0xd9, 0x71, 0x71, 0x14, 0xa7, 0x82, 0x05, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35,
	0xb8, 0x8d, 0x94, 0xf4, 0xa0, 0x4e, 0xd3, 0xc3, 0x50, 0xad, 0x17, 0x0c, 0x51, 0x1a, 0x04, 0xd7,
	0x23, 0xa5, 0xcf, 0xc5, 0x0e, 0x15, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x01, 0xbb, 0x86, 0x33,
	0x88, 0x29, 0x33, 0x45, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c,
	0x04, 0xe1, 0x38, 0x29, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0x7a, 0xd0, 0x10, 0x80, 0xd9, 0x0b, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x14, 0xb3, 0x32, 0x3d, 0x01, 0x00, 0x00,
}
