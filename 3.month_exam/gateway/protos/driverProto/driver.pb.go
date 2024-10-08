// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: protos/driverProto/driver.proto

package driverProto

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

type DriverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string         `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string         `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	WorkingRegion string         `protobuf:"bytes,5,opt,name=working_region,json=workingRegion,proto3" json:"working_region,omitempty"`
	Vehicle       string         `protobuf:"bytes,6,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	Status        string         `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	HiredAt       string         `protobuf:"bytes,8,opt,name=hired_at,json=hiredAt,proto3" json:"hired_at,omitempty"`
	DriverAddress *DriverAddress `protobuf:"bytes,9,opt,name=driver_address,json=driverAddress,proto3" json:"driver_address,omitempty"`
}

func (x *DriverInfo) Reset() {
	*x = DriverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverInfo) ProtoMessage() {}

func (x *DriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverInfo.ProtoReflect.Descriptor instead.
func (*DriverInfo) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{0}
}

func (x *DriverInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DriverInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DriverInfo) GetWorkingRegion() string {
	if x != nil {
		return x.WorkingRegion
	}
	return ""
}

func (x *DriverInfo) GetVehicle() string {
	if x != nil {
		return x.Vehicle
	}
	return ""
}

func (x *DriverInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DriverInfo) GetHiredAt() string {
	if x != nil {
		return x.HiredAt
	}
	return ""
}

func (x *DriverInfo) GetDriverAddress() *DriverAddress {
	if x != nil {
		return x.DriverAddress
	}
	return nil
}

type DriverAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City        string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Region      string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	HomeAddress string `protobuf:"bytes,3,opt,name=home_address,json=homeAddress,proto3" json:"home_address,omitempty"`
}

func (x *DriverAddress) Reset() {
	*x = DriverAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverAddress) ProtoMessage() {}

func (x *DriverAddress) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverAddress.ProtoReflect.Descriptor instead.
func (*DriverAddress) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{1}
}

func (x *DriverAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *DriverAddress) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DriverAddress) GetHomeAddress() string {
	if x != nil {
		return x.HomeAddress
	}
	return ""
}

type DriverID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DriverID) Reset() {
	*x = DriverID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverID) ProtoMessage() {}

func (x *DriverID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverID.ProtoReflect.Descriptor instead.
func (*DriverID) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{2}
}

func (x *DriverID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DriverResponse) Reset() {
	*x = DriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverResponse) ProtoMessage() {}

func (x *DriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverResponse.ProtoReflect.Descriptor instead.
func (*DriverResponse) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{3}
}

func (x *DriverResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{4}
}

func (x *GetLocationRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Vehicle string `protobuf:"bytes,5,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
}

func (x *GetLocationResponse) Reset() {
	*x = GetLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driverProto_driver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationResponse) ProtoMessage() {}

func (x *GetLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driverProto_driver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationResponse.ProtoReflect.Descriptor instead.
func (*GetLocationResponse) Descriptor() ([]byte, []int) {
	return file_protos_driverProto_driver_proto_rawDescGZIP(), []int{5}
}

func (x *GetLocationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLocationResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLocationResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetLocationResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetLocationResponse) GetVehicle() string {
	if x != nil {
		return x.Vehicle
	}
	return ""
}

var File_protos_driverProto_driver_proto protoreflect.FileDescriptor

var file_protos_driverProto_driver_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x02, 0x0a, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x69, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x0d, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x32, 0xaa, 0x01, 0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x09, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_driverProto_driver_proto_rawDescOnce sync.Once
	file_protos_driverProto_driver_proto_rawDescData = file_protos_driverProto_driver_proto_rawDesc
)

func file_protos_driverProto_driver_proto_rawDescGZIP() []byte {
	file_protos_driverProto_driver_proto_rawDescOnce.Do(func() {
		file_protos_driverProto_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_driverProto_driver_proto_rawDescData)
	})
	return file_protos_driverProto_driver_proto_rawDescData
}

var file_protos_driverProto_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_driverProto_driver_proto_goTypes = []interface{}{
	(*DriverInfo)(nil),          // 0: DriverInfo
	(*DriverAddress)(nil),       // 1: DriverAddress
	(*DriverID)(nil),            // 2: DriverID
	(*DriverResponse)(nil),      // 3: DriverResponse
	(*GetLocationRequest)(nil),  // 4: GetLocationRequest
	(*GetLocationResponse)(nil), // 5: GetLocationResponse
}
var file_protos_driverProto_driver_proto_depIdxs = []int32{
	1, // 0: DriverInfo.driver_address:type_name -> DriverAddress
	0, // 1: DriverService.CreateDriver:input_type -> DriverInfo
	2, // 2: DriverService.DeleteDriver:input_type -> DriverID
	4, // 3: DriverService.GetAvailableDriver:input_type -> GetLocationRequest
	2, // 4: DriverService.CreateDriver:output_type -> DriverID
	3, // 5: DriverService.DeleteDriver:output_type -> DriverResponse
	5, // 6: DriverService.GetAvailableDriver:output_type -> GetLocationResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_driverProto_driver_proto_init() }
func file_protos_driverProto_driver_proto_init() {
	if File_protos_driverProto_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_driverProto_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverInfo); i {
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
		file_protos_driverProto_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverAddress); i {
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
		file_protos_driverProto_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverID); i {
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
		file_protos_driverProto_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverResponse); i {
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
		file_protos_driverProto_driver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationRequest); i {
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
		file_protos_driverProto_driver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationResponse); i {
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
			RawDescriptor: file_protos_driverProto_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_driverProto_driver_proto_goTypes,
		DependencyIndexes: file_protos_driverProto_driver_proto_depIdxs,
		MessageInfos:      file_protos_driverProto_driver_proto_msgTypes,
	}.Build()
	File_protos_driverProto_driver_proto = out.File
	file_protos_driverProto_driver_proto_rawDesc = nil
	file_protos_driverProto_driver_proto_goTypes = nil
	file_protos_driverProto_driver_proto_depIdxs = nil
}
