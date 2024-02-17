// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: payment-processing.proto

package api

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

// The request message containing the user's document and month
type UserMonthAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Month    string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *UserMonthAverageRequest) Reset() {
	*x = UserMonthAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_processing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMonthAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMonthAverageRequest) ProtoMessage() {}

func (x *UserMonthAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_processing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMonthAverageRequest.ProtoReflect.Descriptor instead.
func (*UserMonthAverageRequest) Descriptor() ([]byte, []int) {
	return file_payment_processing_proto_rawDescGZIP(), []int{0}
}

func (x *UserMonthAverageRequest) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *UserMonthAverageRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

// The response message containing month's total
type UserMonthAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month    string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	Document string `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	Total    string `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *UserMonthAverageResponse) Reset() {
	*x = UserMonthAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_processing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMonthAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMonthAverageResponse) ProtoMessage() {}

func (x *UserMonthAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_processing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMonthAverageResponse.ProtoReflect.Descriptor instead.
func (*UserMonthAverageResponse) Descriptor() ([]byte, []int) {
	return file_payment_processing_proto_rawDescGZIP(), []int{1}
}

func (x *UserMonthAverageResponse) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *UserMonthAverageResponse) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *UserMonthAverageResponse) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

type LastUserTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *LastUserTransactionRequest) Reset() {
	*x = LastUserTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_processing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastUserTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastUserTransactionRequest) ProtoMessage() {}

func (x *LastUserTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_processing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastUserTransactionRequest.ProtoReflect.Descriptor instead.
func (*LastUserTransactionRequest) Descriptor() ([]byte, []int) {
	return file_payment_processing_proto_rawDescGZIP(), []int{2}
}

func (x *LastUserTransactionRequest) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type LastUserTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	SellerId string `protobuf:"bytes,2,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Value    string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LastUserTransactionResponse) Reset() {
	*x = LastUserTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_processing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastUserTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastUserTransactionResponse) ProtoMessage() {}

func (x *LastUserTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_processing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastUserTransactionResponse.ProtoReflect.Descriptor instead.
func (*LastUserTransactionResponse) Descriptor() ([]byte, []int) {
	return file_payment_processing_proto_rawDescGZIP(), []int{3}
}

func (x *LastUserTransactionResponse) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *LastUserTransactionResponse) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

func (x *LastUserTransactionResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LastUserTransactionResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_payment_processing_proto protoreflect.FileDescriptor

var file_payment_processing_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x4b, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x62, 0x0a,
	0x18, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x38, 0x0a, 0x1a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x1b,
	0x4c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xd2, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x50, 0x01, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x2f, 0x66, 0x72, 0x61, 0x75, 0x64, 0x2d, 0x73, 0x63, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_processing_proto_rawDescOnce sync.Once
	file_payment_processing_proto_rawDescData = file_payment_processing_proto_rawDesc
)

func file_payment_processing_proto_rawDescGZIP() []byte {
	file_payment_processing_proto_rawDescOnce.Do(func() {
		file_payment_processing_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_processing_proto_rawDescData)
	})
	return file_payment_processing_proto_rawDescData
}

var file_payment_processing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payment_processing_proto_goTypes = []interface{}{
	(*UserMonthAverageRequest)(nil),     // 0: user.UserMonthAverageRequest
	(*UserMonthAverageResponse)(nil),    // 1: user.UserMonthAverageResponse
	(*LastUserTransactionRequest)(nil),  // 2: user.LastUserTransactionRequest
	(*LastUserTransactionResponse)(nil), // 3: user.LastUserTransactionResponse
}
var file_payment_processing_proto_depIdxs = []int32{
	0, // 0: user.UserTransactionsService.GetUserMonthAverage:input_type -> user.UserMonthAverageRequest
	2, // 1: user.UserTransactionsService.GetLastUserTransaction:input_type -> user.LastUserTransactionRequest
	1, // 2: user.UserTransactionsService.GetUserMonthAverage:output_type -> user.UserMonthAverageResponse
	3, // 3: user.UserTransactionsService.GetLastUserTransaction:output_type -> user.LastUserTransactionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_processing_proto_init() }
func file_payment_processing_proto_init() {
	if File_payment_processing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_processing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMonthAverageRequest); i {
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
		file_payment_processing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMonthAverageResponse); i {
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
		file_payment_processing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastUserTransactionRequest); i {
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
		file_payment_processing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastUserTransactionResponse); i {
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
			RawDescriptor: file_payment_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_processing_proto_goTypes,
		DependencyIndexes: file_payment_processing_proto_depIdxs,
		MessageInfos:      file_payment_processing_proto_msgTypes,
	}.Build()
	File_payment_processing_proto = out.File
	file_payment_processing_proto_rawDesc = nil
	file_payment_processing_proto_goTypes = nil
	file_payment_processing_proto_depIdxs = nil
}