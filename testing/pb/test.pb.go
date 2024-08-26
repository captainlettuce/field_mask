// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: test.proto

package pb

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

type RecursiveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A      string            `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B      int32             `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
	Nested *RecursiveMessage `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *RecursiveMessage) Reset() {
	*x = RecursiveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecursiveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecursiveMessage) ProtoMessage() {}

func (x *RecursiveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecursiveMessage.ProtoReflect.Descriptor instead.
func (*RecursiveMessage) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *RecursiveMessage) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *RecursiveMessage) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *RecursiveMessage) GetNested() *RecursiveMessage {
	if x != nil {
		return x.Nested
	}
	return nil
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A             string  `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B             int32   `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
	C             *string `protobuf:"bytes,3,opt,name=c,proto3,oneof" json:"c,omitempty"`
	UntaggedField string  `protobuf:"bytes,4,opt,name=UntaggedField,proto3" json:"UntaggedField,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Base) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Base) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Base) GetC() string {
	if x != nil && x.C != nil {
		return *x.C
	}
	return ""
}

func (x *Base) GetUntaggedField() string {
	if x != nil {
		return x.UntaggedField
	}
	return ""
}

type OptionalFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *string `protobuf:"bytes,1,opt,name=a,proto3,oneof" json:"a,omitempty"`
	B *int32  `protobuf:"zigzag32,2,opt,name=b,proto3,oneof" json:"b,omitempty"`
}

func (x *OptionalFields) Reset() {
	*x = OptionalFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalFields) ProtoMessage() {}

func (x *OptionalFields) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalFields.ProtoReflect.Descriptor instead.
func (*OptionalFields) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *OptionalFields) GetA() string {
	if x != nil && x.A != nil {
		return *x.A
	}
	return ""
}

func (x *OptionalFields) GetB() int32 {
	if x != nil && x.B != nil {
		return *x.B
	}
	return 0
}

type ArrayMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Array []int32          `protobuf:"zigzag32,1,rep,packed,name=array,proto3" json:"array,omitempty"`
	Map   map[int32]string `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ArrayMap) Reset() {
	*x = ArrayMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayMap) ProtoMessage() {}

func (x *ArrayMap) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayMap.ProtoReflect.Descriptor instead.
func (*ArrayMap) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *ArrayMap) GetArray() []int32 {
	if x != nil {
		return x.Array
	}
	return nil
}

func (x *ArrayMap) GetMap() map[int32]string {
	if x != nil {
		return x.Map
	}
	return nil
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A    string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	Base *Base  `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *Nested) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Nested) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type NestedRecursive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A      string           `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B      int32            `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
	Nested *NestedRecursive `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *NestedRecursive) Reset() {
	*x = NestedRecursive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRecursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRecursive) ProtoMessage() {}

func (x *NestedRecursive) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRecursive.ProtoReflect.Descriptor instead.
func (*NestedRecursive) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *NestedRecursive) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *NestedRecursive) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *NestedRecursive) GetNested() *NestedRecursive {
	if x != nil {
		return x.Nested
	}
	return nil
}

type NestedRecursiveVariantA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C string                   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	A *NestedRecursiveVariantA `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *NestedRecursiveVariantB `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *NestedRecursiveVariantA) Reset() {
	*x = NestedRecursiveVariantA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRecursiveVariantA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRecursiveVariantA) ProtoMessage() {}

func (x *NestedRecursiveVariantA) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRecursiveVariantA.ProtoReflect.Descriptor instead.
func (*NestedRecursiveVariantA) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6}
}

func (x *NestedRecursiveVariantA) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *NestedRecursiveVariantA) GetA() *NestedRecursiveVariantA {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *NestedRecursiveVariantA) GetB() *NestedRecursiveVariantB {
	if x != nil {
		return x.B
	}
	return nil
}

type NestedRecursiveVariantB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C string                   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	B *NestedRecursiveVariantB `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	A *NestedRecursiveVariantA `protobuf:"bytes,3,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *NestedRecursiveVariantB) Reset() {
	*x = NestedRecursiveVariantB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRecursiveVariantB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRecursiveVariantB) ProtoMessage() {}

func (x *NestedRecursiveVariantB) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRecursiveVariantB.ProtoReflect.Descriptor instead.
func (*NestedRecursiveVariantB) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{7}
}

func (x *NestedRecursiveVariantB) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *NestedRecursiveVariantB) GetB() *NestedRecursiveVariantB {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *NestedRecursiveVariantB) GetA() *NestedRecursiveVariantA {
	if x != nil {
		return x.A
	}
	return nil
}

type BenchmarkTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string  `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	C int32   `protobuf:"zigzag32,3,opt,name=c,proto3" json:"c,omitempty"`
	D *bool   `protobuf:"varint,4,opt,name=d,proto3,oneof" json:"d,omitempty"`
}

func (x *BenchmarkTest) Reset() {
	*x = BenchmarkTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkTest) ProtoMessage() {}

func (x *BenchmarkTest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkTest.ProtoReflect.Descriptor instead.
func (*BenchmarkTest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{8}
}

func (x *BenchmarkTest) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *BenchmarkTest) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *BenchmarkTest) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *BenchmarkTest) GetD() bool {
	if x != nil && x.D != nil {
		return *x.D
	}
	return false
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x10,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x62, 0x12, 0x29, 0x0a, 0x06,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52,
	0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x61, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x62, 0x12, 0x11, 0x0a, 0x01, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x63, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0d, 0x55, 0x6e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x6e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x63, 0x22, 0x42, 0x0a, 0x0e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x11, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x61, 0x88, 0x01, 0x01, 0x12,
	0x11, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x48, 0x01, 0x52, 0x01, 0x62, 0x88,
	0x01, 0x01, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x61, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x62, 0x22, 0x7e,
	0x0a, 0x08, 0x41, 0x72, 0x72, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x11, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x24, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31,
	0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x22, 0x57, 0x0a, 0x0f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72,
	0x73, 0x69, 0x76, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x62,
	0x12, 0x28, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x77, 0x0a, 0x17, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x63, 0x12, 0x26, 0x0a, 0x01, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x41, 0x52, 0x01, 0x61, 0x12, 0x26, 0x0a, 0x01, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42,
	0x52, 0x01, 0x62, 0x22, 0x77, 0x0a, 0x17, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63,
	0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x12, 0x0c,
	0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x63, 0x12, 0x26, 0x0a, 0x01,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x42, 0x52, 0x01, 0x62, 0x12, 0x26, 0x0a, 0x01, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x41, 0x52, 0x01, 0x61, 0x22, 0x52, 0x0a, 0x0d,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x63, 0x12, 0x11, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x01, 0x64, 0x88, 0x01, 0x01, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x64,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_test_proto_goTypes = []interface{}{
	(*RecursiveMessage)(nil),        // 0: RecursiveMessage
	(*Base)(nil),                    // 1: Base
	(*OptionalFields)(nil),          // 2: OptionalFields
	(*ArrayMap)(nil),                // 3: ArrayMap
	(*Nested)(nil),                  // 4: Nested
	(*NestedRecursive)(nil),         // 5: NestedRecursive
	(*NestedRecursiveVariantA)(nil), // 6: NestedRecursiveVariantA
	(*NestedRecursiveVariantB)(nil), // 7: NestedRecursiveVariantB
	(*BenchmarkTest)(nil),           // 8: benchmarkTest
	nil,                             // 9: ArrayMap.MapEntry
}
var file_test_proto_depIdxs = []int32{
	0, // 0: RecursiveMessage.nested:type_name -> RecursiveMessage
	9, // 1: ArrayMap.map:type_name -> ArrayMap.MapEntry
	1, // 2: Nested.base:type_name -> Base
	5, // 3: NestedRecursive.nested:type_name -> NestedRecursive
	6, // 4: NestedRecursiveVariantA.a:type_name -> NestedRecursiveVariantA
	7, // 5: NestedRecursiveVariantA.b:type_name -> NestedRecursiveVariantB
	7, // 6: NestedRecursiveVariantB.b:type_name -> NestedRecursiveVariantB
	6, // 7: NestedRecursiveVariantB.a:type_name -> NestedRecursiveVariantA
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecursiveMessage); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionalFields); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayMap); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRecursive); i {
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
		file_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRecursiveVariantA); i {
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
		file_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRecursiveVariantB); i {
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
		file_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkTest); i {
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
	file_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_test_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_test_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
