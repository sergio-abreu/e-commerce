// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: discount_calculator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestForDiscount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *RequestForDiscount) Reset() {
	*x = RequestForDiscount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestForDiscount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestForDiscount) ProtoMessage() {}

func (x *RequestForDiscount) ProtoReflect() protoreflect.Message {
	mi := &file_discount_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestForDiscount.ProtoReflect.Descriptor instead.
func (*RequestForDiscount) Descriptor() ([]byte, []int) {
	return file_discount_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *RequestForDiscount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RequestForDiscount) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percentage   float64 `protobuf:"fixed64,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	ValueInCents int32   `protobuf:"varint,2,opt,name=value_in_cents,json=valueInCents,proto3" json:"value_in_cents,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_discount_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_discount_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *Discount) GetPercentage() float64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Discount) GetValueInCents() int32 {
	if x != nil {
		return x.ValueInCents
	}
	return 0
}

var File_discount_calculator_proto protoreflect.FileDescriptor

var file_discount_calculator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x08, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x5a, 0x0a, 0x12, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x44, 0x0a, 0x20, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x6f, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x09, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_calculator_proto_rawDescOnce sync.Once
	file_discount_calculator_proto_rawDescData = file_discount_calculator_proto_rawDesc
)

func file_discount_calculator_proto_rawDescGZIP() []byte {
	file_discount_calculator_proto_rawDescOnce.Do(func() {
		file_discount_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_calculator_proto_rawDescData)
	})
	return file_discount_calculator_proto_rawDescData
}

var file_discount_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discount_calculator_proto_goTypes = []interface{}{
	(*RequestForDiscount)(nil), // 0: RequestForDiscount
	(*Discount)(nil),           // 1: Discount
}
var file_discount_calculator_proto_depIdxs = []int32{
	0, // 0: DiscountCalculator.CalculateUsersDiscountForProduct:input_type -> RequestForDiscount
	1, // 1: DiscountCalculator.CalculateUsersDiscountForProduct:output_type -> Discount
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discount_calculator_proto_init() }
func file_discount_calculator_proto_init() {
	if File_discount_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestForDiscount); i {
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
		file_discount_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discount); i {
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
			RawDescriptor: file_discount_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_calculator_proto_goTypes,
		DependencyIndexes: file_discount_calculator_proto_depIdxs,
		MessageInfos:      file_discount_calculator_proto_msgTypes,
	}.Build()
	File_discount_calculator_proto = out.File
	file_discount_calculator_proto_rawDesc = nil
	file_discount_calculator_proto_goTypes = nil
	file_discount_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiscountCalculatorClient is the client API for DiscountCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscountCalculatorClient interface {
	CalculateUsersDiscountForProduct(ctx context.Context, in *RequestForDiscount, opts ...grpc.CallOption) (*Discount, error)
}

type discountCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountCalculatorClient(cc grpc.ClientConnInterface) DiscountCalculatorClient {
	return &discountCalculatorClient{cc}
}

func (c *discountCalculatorClient) CalculateUsersDiscountForProduct(ctx context.Context, in *RequestForDiscount, opts ...grpc.CallOption) (*Discount, error) {
	out := new(Discount)
	err := c.cc.Invoke(ctx, "/DiscountCalculator/CalculateUsersDiscountForProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountCalculatorServer is the server API for DiscountCalculator service.
type DiscountCalculatorServer interface {
	CalculateUsersDiscountForProduct(context.Context, *RequestForDiscount) (*Discount, error)
}

// UnimplementedDiscountCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedDiscountCalculatorServer struct {
}

func (*UnimplementedDiscountCalculatorServer) CalculateUsersDiscountForProduct(context.Context, *RequestForDiscount) (*Discount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateUsersDiscountForProduct not implemented")
}

func RegisterDiscountCalculatorServer(s *grpc.Server, srv DiscountCalculatorServer) {
	s.RegisterService(&_DiscountCalculator_serviceDesc, srv)
}

func _DiscountCalculator_CalculateUsersDiscountForProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForDiscount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountCalculatorServer).CalculateUsersDiscountForProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DiscountCalculator/CalculateUsersDiscountForProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountCalculatorServer).CalculateUsersDiscountForProduct(ctx, req.(*RequestForDiscount))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscountCalculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DiscountCalculator",
	HandlerType: (*DiscountCalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateUsersDiscountForProduct",
			Handler:    _DiscountCalculator_CalculateUsersDiscountForProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discount_calculator.proto",
}
