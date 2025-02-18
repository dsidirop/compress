// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pbcurvegenresponsev1.proto

package arena

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PBCurveGenReplyV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string                  `protobuf:"bytes,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	Spec     *PBCurveSpecificationV1 `protobuf:"bytes,2,opt,name=Spec,proto3" json:"Spec,omitempty"`
	Lead1    []int64                 `protobuf:"varint,3,rep,packed,name=Lead1,proto3" json:"Lead1,omitempty"`
	Lead2    []int64                 `protobuf:"varint,4,rep,packed,name=Lead2,proto3" json:"Lead2,omitempty"`
	Lead3    []int64                 `protobuf:"varint,5,rep,packed,name=Lead3,proto3" json:"Lead3,omitempty"`
	Lead4    []int64                 `protobuf:"varint,6,rep,packed,name=Lead4,proto3" json:"Lead4,omitempty"`
	Lead5    []int64                 `protobuf:"varint,7,rep,packed,name=Lead5,proto3" json:"Lead5,omitempty"`
	Lead6    []int64                 `protobuf:"varint,8,rep,packed,name=Lead6,proto3" json:"Lead6,omitempty"`
	Lead7    []int64                 `protobuf:"varint,9,rep,packed,name=Lead7,proto3" json:"Lead7,omitempty"`
	Lead8    []int64                 `protobuf:"varint,10,rep,packed,name=Lead8,proto3" json:"Lead8,omitempty"`
	Lead9    []int64                 `protobuf:"varint,11,rep,packed,name=Lead9,proto3" json:"Lead9,omitempty"`
	Lead10   []int64                 `protobuf:"varint,12,rep,packed,name=Lead10,proto3" json:"Lead10,omitempty"`
	Lead11   []int64                 `protobuf:"varint,13,rep,packed,name=Lead11,proto3" json:"Lead11,omitempty"`
	Lead12   []int64                 `protobuf:"varint,14,rep,packed,name=Lead12,proto3" json:"Lead12,omitempty"`
	Abp      []int64                 `protobuf:"varint,15,rep,packed,name=Abp,proto3" json:"Abp,omitempty"`
	Cvp      []int64                 `protobuf:"varint,16,rep,packed,name=Cvp,proto3" json:"Cvp,omitempty"`
	Pap      []int64                 `protobuf:"varint,17,rep,packed,name=Pap,proto3" json:"Pap,omitempty"`
	Spo2     []int64                 `protobuf:"varint,18,rep,packed,name=Spo2,proto3" json:"Spo2,omitempty"`
	Wp       []int64                 `protobuf:"varint,19,rep,packed,name=Wp,proto3" json:"Wp,omitempty"`
	Tags     []*PBTag                `protobuf:"bytes,20,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *PBCurveGenReplyV1) Reset() {
	*x = PBCurveGenReplyV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcurvegenresponsev1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBCurveGenReplyV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBCurveGenReplyV1) ProtoMessage() {}

func (x *PBCurveGenReplyV1) ProtoReflect() protoreflect.Message {
	mi := &file_pbcurvegenresponsev1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBCurveGenReplyV1.ProtoReflect.Descriptor instead.
func (*PBCurveGenReplyV1) Descriptor() ([]byte, []int) {
	return file_pbcurvegenresponsev1_proto_rawDescGZIP(), []int{0}
}

func (x *PBCurveGenReplyV1) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *PBCurveGenReplyV1) GetSpec() *PBCurveSpecificationV1 {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead1() []int64 {
	if x != nil {
		return x.Lead1
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead2() []int64 {
	if x != nil {
		return x.Lead2
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead3() []int64 {
	if x != nil {
		return x.Lead3
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead4() []int64 {
	if x != nil {
		return x.Lead4
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead5() []int64 {
	if x != nil {
		return x.Lead5
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead6() []int64 {
	if x != nil {
		return x.Lead6
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead7() []int64 {
	if x != nil {
		return x.Lead7
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead8() []int64 {
	if x != nil {
		return x.Lead8
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead9() []int64 {
	if x != nil {
		return x.Lead9
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead10() []int64 {
	if x != nil {
		return x.Lead10
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead11() []int64 {
	if x != nil {
		return x.Lead11
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetLead12() []int64 {
	if x != nil {
		return x.Lead12
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetAbp() []int64 {
	if x != nil {
		return x.Abp
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetCvp() []int64 {
	if x != nil {
		return x.Cvp
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetPap() []int64 {
	if x != nil {
		return x.Pap
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetSpo2() []int64 {
	if x != nil {
		return x.Spo2
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetWp() []int64 {
	if x != nil {
		return x.Wp
	}
	return nil
}

func (x *PBCurveGenReplyV1) GetTags() []*PBTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PBCurveSpecificationV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimulatorName  string `protobuf:"bytes,1,opt,name=SimulatorName,proto3" json:"SimulatorName,omitempty"`
	Tenant         string `protobuf:"bytes,2,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	StartTime      int64  `protobuf:"varint,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime        int64  `protobuf:"varint,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	SampleInterval int64  `protobuf:"varint,5,opt,name=SampleInterval,proto3" json:"SampleInterval,omitempty"`
	CurveTypes     int64  `protobuf:"varint,6,opt,name=CurveTypes,proto3" json:"CurveTypes,omitempty"`
}

func (x *PBCurveSpecificationV1) Reset() {
	*x = PBCurveSpecificationV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcurvegenresponsev1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBCurveSpecificationV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBCurveSpecificationV1) ProtoMessage() {}

func (x *PBCurveSpecificationV1) ProtoReflect() protoreflect.Message {
	mi := &file_pbcurvegenresponsev1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBCurveSpecificationV1.ProtoReflect.Descriptor instead.
func (*PBCurveSpecificationV1) Descriptor() ([]byte, []int) {
	return file_pbcurvegenresponsev1_proto_rawDescGZIP(), []int{1}
}

func (x *PBCurveSpecificationV1) GetSimulatorName() string {
	if x != nil {
		return x.SimulatorName
	}
	return ""
}

func (x *PBCurveSpecificationV1) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *PBCurveSpecificationV1) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PBCurveSpecificationV1) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *PBCurveSpecificationV1) GetSampleInterval() int64 {
	if x != nil {
		return x.SampleInterval
	}
	return 0
}

func (x *PBCurveSpecificationV1) GetCurveTypes() int64 {
	if x != nil {
		return x.CurveTypes
	}
	return 0
}

type PBTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagTime int64 `protobuf:"varint,1,opt,name=TagTime,proto3" json:"TagTime,omitempty"`
	TagType int64 `protobuf:"varint,2,opt,name=TagType,proto3" json:"TagType,omitempty"`
}

func (x *PBTag) Reset() {
	*x = PBTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcurvegenresponsev1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBTag) ProtoMessage() {}

func (x *PBTag) ProtoReflect() protoreflect.Message {
	mi := &file_pbcurvegenresponsev1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBTag.ProtoReflect.Descriptor instead.
func (*PBTag) Descriptor() ([]byte, []int) {
	return file_pbcurvegenresponsev1_proto_rawDescGZIP(), []int{2}
}

func (x *PBTag) GetTagTime() int64 {
	if x != nil {
		return x.TagTime
	}
	return 0
}

func (x *PBTag) GetTagType() int64 {
	if x != nil {
		return x.TagType
	}
	return 0
}

var File_pbcurvegenresponsev1_proto protoreflect.FileDescriptor

var file_pbcurvegenresponsev1_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x63, 0x75, 0x72, 0x76, 0x65, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x6f,
	0x6f, 0x2e, 0x63, 0x75, 0x72, 0x76, 0x65, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x76,
	0x31, 0x22, 0x88, 0x04, 0x0a, 0x11, 0x50, 0x42, 0x43, 0x75, 0x72, 0x76, 0x65, 0x47, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x56, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x63, 0x75, 0x72, 0x76, 0x65, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x76, 0x31, 0x2e, 0x50, 0x42, 0x43, 0x75, 0x72, 0x76, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x04,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x31, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65,
	0x61, 0x64, 0x32, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x32,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x33, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x05, 0x4c, 0x65, 0x61, 0x64, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x34, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x34, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x61, 0x64, 0x35, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61,
	0x64, 0x35, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x36, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64,
	0x37, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x37, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x38, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x4c,
	0x65, 0x61, 0x64, 0x38, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x39, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x61, 0x64, 0x39, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65,
	0x61, 0x64, 0x31, 0x30, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x4c, 0x65, 0x61, 0x64,
	0x31, 0x30, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x31, 0x31, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x31, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65,
	0x61, 0x64, 0x31, 0x32, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x4c, 0x65, 0x61, 0x64,
	0x31, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x62, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x41, 0x62, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x76, 0x70, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x03, 0x43, 0x76, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x70, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x50, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x70, 0x6f, 0x32,
	0x18, 0x12, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x53, 0x70, 0x6f, 0x32, 0x12, 0x0e, 0x0a, 0x02,
	0x57, 0x70, 0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x57, 0x70, 0x12, 0x2e, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6f, 0x6f,
	0x2e, 0x63, 0x75, 0x72, 0x76, 0x65, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x76, 0x31,
	0x2e, 0x50, 0x42, 0x54, 0x61, 0x67, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x22, 0xd6, 0x01, 0x0a,
	0x16, 0x50, 0x42, 0x43, 0x75, 0x72, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x76, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x72, 0x76, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x05, 0x50, 0x42, 0x54, 0x61, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x61, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x54, 0x61, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbcurvegenresponsev1_proto_rawDescOnce sync.Once
	file_pbcurvegenresponsev1_proto_rawDescData = file_pbcurvegenresponsev1_proto_rawDesc
)

func file_pbcurvegenresponsev1_proto_rawDescGZIP() []byte {
	file_pbcurvegenresponsev1_proto_rawDescOnce.Do(func() {
		file_pbcurvegenresponsev1_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbcurvegenresponsev1_proto_rawDescData)
	})
	return file_pbcurvegenresponsev1_proto_rawDescData
}

var file_pbcurvegenresponsev1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pbcurvegenresponsev1_proto_goTypes = []interface{}{
	(*PBCurveGenReplyV1)(nil),      // 0: foo.curvegenreplyv1.PBCurveGenReplyV1
	(*PBCurveSpecificationV1)(nil), // 1: foo.curvegenreplyv1.PBCurveSpecificationV1
	(*PBTag)(nil),                  // 2: foo.curvegenreplyv1.PBTag
}
var file_pbcurvegenresponsev1_proto_depIdxs = []int32{
	1, // 0: foo.curvegenreplyv1.PBCurveGenReplyV1.Spec:type_name -> foo.curvegenreplyv1.PBCurveSpecificationV1
	2, // 1: foo.curvegenreplyv1.PBCurveGenReplyV1.Tags:type_name -> foo.curvegenreplyv1.PBTag
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbcurvegenresponsev1_proto_init() }
func file_pbcurvegenresponsev1_proto_init() {
	if File_pbcurvegenresponsev1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbcurvegenresponsev1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBCurveGenReplyV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbcurvegenresponsev1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBCurveSpecificationV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbcurvegenresponsev1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbcurvegenresponsev1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbcurvegenresponsev1_proto_goTypes,
		DependencyIndexes: file_pbcurvegenresponsev1_proto_depIdxs,
		MessageInfos:      file_pbcurvegenresponsev1_proto_msgTypes,
	}.Build()
	File_pbcurvegenresponsev1_proto = out.File
	file_pbcurvegenresponsev1_proto_rawDesc = nil
	file_pbcurvegenresponsev1_proto_goTypes = nil
	file_pbcurvegenresponsev1_proto_depIdxs = nil
}
