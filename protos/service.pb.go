// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/service.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TestHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestHelloRequest) Reset() {
	*x = TestHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestHelloRequest) ProtoMessage() {}

func (x *TestHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestHelloRequest.ProtoReflect.Descriptor instead.
func (*TestHelloRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{0}
}

func (x *TestHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TestHelloResponse) Reset() {
	*x = TestHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestHelloResponse) ProtoMessage() {}

func (x *TestHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestHelloResponse.ProtoReflect.Descriptor instead.
func (*TestHelloResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{1}
}

func (x *TestHelloResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{2}
}

type Veterinary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VetId     string        `protobuf:"bytes,10,opt,name=vetId,proto3" json:"vetId,omitempty"`
	Title     string        `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	FirstName string        `protobuf:"bytes,12,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string        `protobuf:"bytes,13,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string        `protobuf:"bytes,14,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string        `protobuf:"bytes,15,opt,name=phone,proto3" json:"phone,omitempty"`
	Address   *Location     `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty"`
	Summary   string        `protobuf:"bytes,17,opt,name=summary,proto3" json:"summary,omitempty"`
	Services  []*VetService `protobuf:"bytes,19,rep,name=services,proto3" json:"services,omitempty"`
	Photo     []byte        `protobuf:"bytes,20,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *Veterinary) Reset() {
	*x = Veterinary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Veterinary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Veterinary) ProtoMessage() {}

func (x *Veterinary) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Veterinary.ProtoReflect.Descriptor instead.
func (*Veterinary) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{3}
}

func (x *Veterinary) GetVetId() string {
	if x != nil {
		return x.VetId
	}
	return ""
}

func (x *Veterinary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Veterinary) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Veterinary) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Veterinary) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Veterinary) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Veterinary) GetAddress() *Location {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Veterinary) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Veterinary) GetServices() []*VetService {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Veterinary) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat     float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long    float64 `protobuf:"fixed64,2,opt,name=long,proto3" json:"long,omitempty"`
	Address string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type VetService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId   string  `protobuf:"bytes,10,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Title       string  `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	Unit        string  `protobuf:"bytes,13,opt,name=unit,proto3" json:"unit,omitempty"`
	CostPerUnit float64 `protobuf:"fixed64,14,opt,name=costPerUnit,proto3" json:"costPerUnit,omitempty"`
}

func (x *VetService) Reset() {
	*x = VetService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VetService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VetService) ProtoMessage() {}

func (x *VetService) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VetService.ProtoReflect.Descriptor instead.
func (*VetService) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{5}
}

func (x *VetService) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *VetService) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VetService) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VetService) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *VetService) GetCostPerUnit() float64 {
	if x != nil {
		return x.CostPerUnit
	}
	return 0
}

type VetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VetId       string    `protobuf:"bytes,10,opt,name=vetId,proto3" json:"vetId,omitempty"`
	SearchQuery string    `protobuf:"bytes,11,opt,name=searchQuery,proto3" json:"searchQuery,omitempty"`
	Location    *Location `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *VetRequest) Reset() {
	*x = VetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VetRequest) ProtoMessage() {}

func (x *VetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VetRequest.ProtoReflect.Descriptor instead.
func (*VetRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{6}
}

func (x *VetRequest) GetVetId() string {
	if x != nil {
		return x.VetId
	}
	return ""
}

func (x *VetRequest) GetSearchQuery() string {
	if x != nil {
		return x.SearchQuery
	}
	return ""
}

func (x *VetRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Farmer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerId  int32     `protobuf:"varint,10,opt,name=farmerId,proto3" json:"farmerId,omitempty"`
	FirstName string    `protobuf:"bytes,11,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string    `protobuf:"bytes,12,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string    `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string    `protobuf:"bytes,15,opt,name=phone,proto3" json:"phone,omitempty"`
	Address   *Location `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty"`
	Photo     []byte    `protobuf:"bytes,17,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *Farmer) Reset() {
	*x = Farmer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Farmer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Farmer) ProtoMessage() {}

func (x *Farmer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Farmer.ProtoReflect.Descriptor instead.
func (*Farmer) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{7}
}

func (x *Farmer) GetFarmerId() int32 {
	if x != nil {
		return x.FarmerId
	}
	return 0
}

func (x *Farmer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Farmer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Farmer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Farmer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Farmer) GetAddress() *Location {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Farmer) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

type FarmerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerId int32  `protobuf:"varint,1,opt,name=farmerId,proto3" json:"farmerId,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FarmerRequest) Reset() {
	*x = FarmerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FarmerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FarmerRequest) ProtoMessage() {}

func (x *FarmerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FarmerRequest.ProtoReflect.Descriptor instead.
func (*FarmerRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{8}
}

func (x *FarmerRequest) GetFarmerId() int32 {
	if x != nil {
		return x.FarmerId
	}
	return 0
}

func (x *FarmerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TreatmentSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId    string               `protobuf:"bytes,10,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Time         *timestamp.Timestamp `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`
	Location     *Location            `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	FarmerId     string               `protobuf:"bytes,13,opt,name=farmerId,proto3" json:"farmerId,omitempty"`
	VeterinaryId string               `protobuf:"bytes,14,opt,name=VeterinaryId,proto3" json:"VeterinaryId,omitempty"`
	Services     []*VetService        `protobuf:"bytes,15,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *TreatmentSession) Reset() {
	*x = TreatmentSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreatmentSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreatmentSession) ProtoMessage() {}

func (x *TreatmentSession) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreatmentSession.ProtoReflect.Descriptor instead.
func (*TreatmentSession) Descriptor() ([]byte, []int) {
	return file_protos_service_proto_rawDescGZIP(), []int{9}
}

func (x *TreatmentSession) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TreatmentSession) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *TreatmentSession) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *TreatmentSession) GetFarmerId() string {
	if x != nil {
		return x.FarmerId
	}
	return ""
}

func (x *TreatmentSession) GetVeterinaryId() string {
	if x != nil {
		return x.VeterinaryId
	}
	return ""
}

func (x *TreatmentSession) GetServices() []*VetService {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_protos_service_proto protoreflect.FileDescriptor

var file_protos_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11,
	0x54, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb4, 0x02,
	0x0a, 0x0a, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x65, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2f,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65,
	0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x98, 0x01, 0x0a, 0x0a, 0x56, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x63, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x77, 0x0a, 0x0a, 0x56,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x65, 0x74,
	0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x31, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd1, 0x01, 0x0a, 0x06, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x46, 0x61, 0x72, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x72,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x61, 0x72,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x88, 0x02, 0x0a, 0x10,
	0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x31,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x56, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xae, 0x03, 0x0a, 0x0b, 0x56, 0x65, 0x74, 0x73, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x4a, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x1d, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6e, 0x12, 0x17, 0x2e, 0x76,
	0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x38,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x46, 0x61, 0x72,
	0x6d, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x76, 0x65, 0x74, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x2e, 0x76, 0x65, 0x74, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x64, 0x69, 0x72, 0x61, 0x6e, 0x67, 0x75, 0x67, 0x2f,
	0x76, 0x65, 0x74, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_service_proto_rawDescOnce sync.Once
	file_protos_service_proto_rawDescData = file_protos_service_proto_rawDesc
)

func file_protos_service_proto_rawDescGZIP() []byte {
	file_protos_service_proto_rawDescOnce.Do(func() {
		file_protos_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_service_proto_rawDescData)
	})
	return file_protos_service_proto_rawDescData
}

var file_protos_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_service_proto_goTypes = []interface{}{
	(*TestHelloRequest)(nil),    // 0: vet_backend.TestHelloRequest
	(*TestHelloResponse)(nil),   // 1: vet_backend.TestHelloResponse
	(*EmptyMessage)(nil),        // 2: vet_backend.EmptyMessage
	(*Veterinary)(nil),          // 3: vet_backend.Veterinary
	(*Location)(nil),            // 4: vet_backend.Location
	(*VetService)(nil),          // 5: vet_backend.VetService
	(*VetRequest)(nil),          // 6: vet_backend.VetRequest
	(*Farmer)(nil),              // 7: vet_backend.Farmer
	(*FarmerRequest)(nil),       // 8: vet_backend.FarmerRequest
	(*TreatmentSession)(nil),    // 9: vet_backend.TreatmentSession
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_protos_service_proto_depIdxs = []int32{
	4,  // 0: vet_backend.Veterinary.address:type_name -> vet_backend.Location
	5,  // 1: vet_backend.Veterinary.services:type_name -> vet_backend.VetService
	4,  // 2: vet_backend.VetRequest.location:type_name -> vet_backend.Location
	4,  // 3: vet_backend.Farmer.address:type_name -> vet_backend.Location
	10, // 4: vet_backend.TreatmentSession.time:type_name -> google.protobuf.Timestamp
	4,  // 5: vet_backend.TreatmentSession.location:type_name -> vet_backend.Location
	5,  // 6: vet_backend.TreatmentSession.services:type_name -> vet_backend.VetService
	0,  // 7: vet_backend.VetsBackend.TestHello:input_type -> vet_backend.TestHelloRequest
	6,  // 8: vet_backend.VetsBackend.GetVeterinarians:input_type -> vet_backend.VetRequest
	3,  // 9: vet_backend.VetsBackend.UpdateVeterian:input_type -> vet_backend.Veterinary
	7,  // 10: vet_backend.VetsBackend.UpdateFarmer:input_type -> vet_backend.Farmer
	8,  // 11: vet_backend.VetsBackend.GetFarmer:input_type -> vet_backend.FarmerRequest
	9,  // 12: vet_backend.VetsBackend.ScheduleSession:input_type -> vet_backend.TreatmentSession
	1,  // 13: vet_backend.VetsBackend.TestHello:output_type -> vet_backend.TestHelloResponse
	3,  // 14: vet_backend.VetsBackend.GetVeterinarians:output_type -> vet_backend.Veterinary
	3,  // 15: vet_backend.VetsBackend.UpdateVeterian:output_type -> vet_backend.Veterinary
	7,  // 16: vet_backend.VetsBackend.UpdateFarmer:output_type -> vet_backend.Farmer
	7,  // 17: vet_backend.VetsBackend.GetFarmer:output_type -> vet_backend.Farmer
	9,  // 18: vet_backend.VetsBackend.ScheduleSession:output_type -> vet_backend.TreatmentSession
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_service_proto_init() }
func file_protos_service_proto_init() {
	if File_protos_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestHelloRequest); i {
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
		file_protos_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestHelloResponse); i {
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
		file_protos_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_protos_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Veterinary); i {
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
		file_protos_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_protos_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VetService); i {
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
		file_protos_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VetRequest); i {
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
		file_protos_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Farmer); i {
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
		file_protos_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FarmerRequest); i {
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
		file_protos_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreatmentSession); i {
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
			RawDescriptor: file_protos_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_proto_goTypes,
		DependencyIndexes: file_protos_service_proto_depIdxs,
		MessageInfos:      file_protos_service_proto_msgTypes,
	}.Build()
	File_protos_service_proto = out.File
	file_protos_service_proto_rawDesc = nil
	file_protos_service_proto_goTypes = nil
	file_protos_service_proto_depIdxs = nil
}
