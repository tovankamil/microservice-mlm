// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: lib/protos/v1/payment/payment.proto

package payment

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for updating user balance
type TransferBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SourceUserId  string                 `protobuf:"bytes,1,opt,name=source_user_id,json=sourceUserId,proto3" json:"source_user_id,omitempty"` // Source user ID or account
	Destination   string                 `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`                         // Destination user ID or account
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`                                 // Amount to add or deduct
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferBalanceRequest) Reset() {
	*x = TransferBalanceRequest{}
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBalanceRequest) ProtoMessage() {}

func (x *TransferBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBalanceRequest.ProtoReflect.Descriptor instead.
func (*TransferBalanceRequest) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *TransferBalanceRequest) GetSourceUserId() string {
	if x != nil {
		return x.SourceUserId
	}
	return ""
}

func (x *TransferBalanceRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *TransferBalanceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferBalanceResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Success          bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	DestinationAmout float64                `protobuf:"fixed64,2,opt,name=destination_amout,json=destinationAmout,proto3" json:"destination_amout,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TransferBalanceResponse) Reset() {
	*x = TransferBalanceResponse{}
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBalanceResponse) ProtoMessage() {}

func (x *TransferBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_protos_v1_payment_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBalanceResponse.ProtoReflect.Descriptor instead.
func (*TransferBalanceResponse) Descriptor() ([]byte, []int) {
	return file_lib_protos_v1_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *TransferBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TransferBalanceResponse) GetDestinationAmout() float64 {
	if x != nil {
		return x.DestinationAmout
	}
	return 0
}

var File_lib_protos_v1_payment_payment_proto protoreflect.FileDescriptor

const file_lib_protos_v1_payment_payment_proto_rawDesc = "" +
	"\n" +
	"#lib/protos/v1/payment/payment.proto\x12\x19e_wallet_microservices_v2\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/descriptor.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a,protoc-gen-openapiv2/options/openapiv2.proto\"x\n" +
	"\x16TransferBalanceRequest\x12$\n" +
	"\x0esource_user_id\x18\x01 \x01(\tR\fsourceUserId\x12 \n" +
	"\vdestination\x18\x02 \x01(\tR\vdestination\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\"`\n" +
	"\x17TransferBalanceResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12+\n" +
	"\x11destination_amout\x18\x02 \x01(\x01R\x10destinationAmout2\xcb\x02\n" +
	"\x0ePaymentService\x12\xb8\x02\n" +
	"\x16TransferBalanceService\x121.e_wallet_microservices_v2.TransferBalanceRequest\x1a2.e_wallet_microservices_v2.TransferBalanceResponse\"\xb6\x01\x92A\x93\x01\n" +
	"\apayment\x12\x1aTransfer User Balance info*\x16TransferBalanceServiceJT\n" +
	"\x03200\x12M\n" +
	"\x13Successful response\x126\n" +
	"4\x1a2.e_wallet_microservices_v2.TransferBalanceResponse\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v1/payment/transferB\xe0\x01\x92A\x9f\x01\x12u\n" +
	"%Aplikasimlm.com (Wallet API Endpoint)\"G\n" +
	"\x0fAplikasimlm.com\x12\x1dhttps://github.com/tovankamil\x1a\x15tovan.kamil@gmail.com2\x031.0*\x02\x01\x022\x10application/json:\x10application/jsonZ;e_wallet_microservices_v2/lib/protos/public/payment;paymentb\x06proto3"

var (
	file_lib_protos_v1_payment_payment_proto_rawDescOnce sync.Once
	file_lib_protos_v1_payment_payment_proto_rawDescData []byte
)

func file_lib_protos_v1_payment_payment_proto_rawDescGZIP() []byte {
	file_lib_protos_v1_payment_payment_proto_rawDescOnce.Do(func() {
		file_lib_protos_v1_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lib_protos_v1_payment_payment_proto_rawDesc), len(file_lib_protos_v1_payment_payment_proto_rawDesc)))
	})
	return file_lib_protos_v1_payment_payment_proto_rawDescData
}

var file_lib_protos_v1_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lib_protos_v1_payment_payment_proto_goTypes = []any{
	(*TransferBalanceRequest)(nil),  // 0: e_wallet_microservices_v2.TransferBalanceRequest
	(*TransferBalanceResponse)(nil), // 1: e_wallet_microservices_v2.TransferBalanceResponse
}
var file_lib_protos_v1_payment_payment_proto_depIdxs = []int32{
	0, // 0: e_wallet_microservices_v2.PaymentService.TransferBalanceService:input_type -> e_wallet_microservices_v2.TransferBalanceRequest
	1, // 1: e_wallet_microservices_v2.PaymentService.TransferBalanceService:output_type -> e_wallet_microservices_v2.TransferBalanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lib_protos_v1_payment_payment_proto_init() }
func file_lib_protos_v1_payment_payment_proto_init() {
	if File_lib_protos_v1_payment_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lib_protos_v1_payment_payment_proto_rawDesc), len(file_lib_protos_v1_payment_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_protos_v1_payment_payment_proto_goTypes,
		DependencyIndexes: file_lib_protos_v1_payment_payment_proto_depIdxs,
		MessageInfos:      file_lib_protos_v1_payment_payment_proto_msgTypes,
	}.Build()
	File_lib_protos_v1_payment_payment_proto = out.File
	file_lib_protos_v1_payment_payment_proto_goTypes = nil
	file_lib_protos_v1_payment_payment_proto_depIdxs = nil
}
