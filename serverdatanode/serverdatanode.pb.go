// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: serverdatanode.proto

package serverdatanode

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MensajeTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *MensajeTest) Reset() {
	*x = MensajeTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverdatanode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeTest) ProtoMessage() {}

func (x *MensajeTest) ProtoReflect() protoreflect.Message {
	mi := &file_serverdatanode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeTest.ProtoReflect.Descriptor instead.
func (*MensajeTest) Descriptor() ([]byte, []int) {
	return file_serverdatanode_proto_rawDescGZIP(), []int{0}
}

func (x *MensajeTest) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type ChunkLibro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=Nombre,proto3" json:"Nombre,omitempty"`
	Chunk  []byte `protobuf:"bytes,2,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
}

func (x *ChunkLibro) Reset() {
	*x = ChunkLibro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverdatanode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLibro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLibro) ProtoMessage() {}

func (x *ChunkLibro) ProtoReflect() protoreflect.Message {
	mi := &file_serverdatanode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkLibro.ProtoReflect.Descriptor instead.
func (*ChunkLibro) Descriptor() ([]byte, []int) {
	return file_serverdatanode_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkLibro) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *ChunkLibro) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type Propuestagrpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibroSubido string `protobuf:"bytes,1,opt,name=NombreLibroSubido,proto3" json:"NombreLibroSubido,omitempty"`
	PartesDN1         string `protobuf:"bytes,2,opt,name=PartesDN1,proto3" json:"PartesDN1,omitempty"`
	PartesDN2         string `protobuf:"bytes,3,opt,name=PartesDN2,proto3" json:"PartesDN2,omitempty"`
	PartesDN3         string `protobuf:"bytes,4,opt,name=PartesDN3,proto3" json:"PartesDN3,omitempty"`
}

func (x *Propuestagrpc) Reset() {
	*x = Propuestagrpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverdatanode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuestagrpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuestagrpc) ProtoMessage() {}

func (x *Propuestagrpc) ProtoReflect() protoreflect.Message {
	mi := &file_serverdatanode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuestagrpc.ProtoReflect.Descriptor instead.
func (*Propuestagrpc) Descriptor() ([]byte, []int) {
	return file_serverdatanode_proto_rawDescGZIP(), []int{2}
}

func (x *Propuestagrpc) GetNombreLibroSubido() string {
	if x != nil {
		return x.NombreLibroSubido
	}
	return ""
}

func (x *Propuestagrpc) GetPartesDN1() string {
	if x != nil {
		return x.PartesDN1
	}
	return ""
}

func (x *Propuestagrpc) GetPartesDN2() string {
	if x != nil {
		return x.PartesDN2
	}
	return ""
}

func (x *Propuestagrpc) GetPartesDN3() string {
	if x != nil {
		return x.PartesDN3
	}
	return ""
}

type Booleano struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booleano bool `protobuf:"varint,1,opt,name=Booleano,proto3" json:"Booleano,omitempty"`
}

func (x *Booleano) Reset() {
	*x = Booleano{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverdatanode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booleano) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booleano) ProtoMessage() {}

func (x *Booleano) ProtoReflect() protoreflect.Message {
	mi := &file_serverdatanode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booleano.ProtoReflect.Descriptor instead.
func (*Booleano) Descriptor() ([]byte, []int) {
	return file_serverdatanode_proto_rawDescGZIP(), []int{3}
}

func (x *Booleano) GetBooleano() bool {
	if x != nil {
		return x.Booleano
	}
	return false
}

var File_serverdatanode_proto protoreflect.FileDescriptor

var file_serverdatanode_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22,
	0x3a, 0x0a, 0x0a, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x97, 0x01, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x67, 0x72, 0x70, 0x63, 0x12, 0x2c, 0x0a,
	0x11, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x53, 0x75, 0x62, 0x69,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x4c, 0x69, 0x62, 0x72, 0x6f, 0x53, 0x75, 0x62, 0x69, 0x64, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x61, 0x72, 0x74, 0x65, 0x73, 0x44, 0x4e, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x44, 0x4e, 0x31, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72,
	0x74, 0x65, 0x73, 0x44, 0x4e, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61,
	0x72, 0x74, 0x65, 0x73, 0x44, 0x4e, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x65,
	0x73, 0x44, 0x4e, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x72, 0x74,
	0x65, 0x73, 0x44, 0x4e, 0x33, 0x22, 0x26, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x6f, 0x32, 0xb6, 0x03,
	0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64,
	0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x62,
	0x72, 0x6f, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x1b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x6f, 0x44, 0x65, 0x53, 0x75, 0x62, 0x69, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x15,
	0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x5f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x69, 0x64, 0x6f, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74,
	0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x6f, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x61, 0x72, 0x67, 0x61, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serverdatanode_proto_rawDescOnce sync.Once
	file_serverdatanode_proto_rawDescData = file_serverdatanode_proto_rawDesc
)

func file_serverdatanode_proto_rawDescGZIP() []byte {
	file_serverdatanode_proto_rawDescOnce.Do(func() {
		file_serverdatanode_proto_rawDescData = protoimpl.X.CompressGZIP(file_serverdatanode_proto_rawDescData)
	})
	return file_serverdatanode_proto_rawDescData
}

var file_serverdatanode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_serverdatanode_proto_goTypes = []interface{}{
	(*MensajeTest)(nil),   // 0: serverdatanode.MensajeTest
	(*ChunkLibro)(nil),    // 1: serverdatanode.ChunkLibro
	(*Propuestagrpc)(nil), // 2: serverdatanode.Propuestagrpc
	(*Booleano)(nil),      // 3: serverdatanode.Booleano
}
var file_serverdatanode_proto_depIdxs = []int32{
	0, // 0: serverdatanode.DataNodeService.EnvioMensajeTest:input_type -> serverdatanode.MensajeTest
	1, // 1: serverdatanode.DataNodeService.UploaderSubeLibro:input_type -> serverdatanode.ChunkLibro
	0, // 2: serverdatanode.DataNodeService.UploaderTerminoDeSubirLibro:input_type -> serverdatanode.MensajeTest
	2, // 3: serverdatanode.DataNodeService.Propuesta_Distribuido:input_type -> serverdatanode.Propuestagrpc
	0, // 4: serverdatanode.DataNodeService.DownloaderDescargaLibro:input_type -> serverdatanode.MensajeTest
	0, // 5: serverdatanode.DataNodeService.EnvioMensajeTest:output_type -> serverdatanode.MensajeTest
	0, // 6: serverdatanode.DataNodeService.UploaderSubeLibro:output_type -> serverdatanode.MensajeTest
	0, // 7: serverdatanode.DataNodeService.UploaderTerminoDeSubirLibro:output_type -> serverdatanode.MensajeTest
	3, // 8: serverdatanode.DataNodeService.Propuesta_Distribuido:output_type -> serverdatanode.Booleano
	1, // 9: serverdatanode.DataNodeService.DownloaderDescargaLibro:output_type -> serverdatanode.ChunkLibro
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serverdatanode_proto_init() }
func file_serverdatanode_proto_init() {
	if File_serverdatanode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serverdatanode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeTest); i {
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
		file_serverdatanode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkLibro); i {
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
		file_serverdatanode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuestagrpc); i {
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
		file_serverdatanode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booleano); i {
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
			RawDescriptor: file_serverdatanode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serverdatanode_proto_goTypes,
		DependencyIndexes: file_serverdatanode_proto_depIdxs,
		MessageInfos:      file_serverdatanode_proto_msgTypes,
	}.Build()
	File_serverdatanode_proto = out.File
	file_serverdatanode_proto_rawDesc = nil
	file_serverdatanode_proto_goTypes = nil
	file_serverdatanode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataNodeServiceClient is the client API for DataNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataNodeServiceClient interface {
	EnvioMensajeTest(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*MensajeTest, error)
	UploaderSubeLibro(ctx context.Context, in *ChunkLibro, opts ...grpc.CallOption) (*MensajeTest, error)
	UploaderTerminoDeSubirLibro(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*MensajeTest, error)
	Propuesta_Distribuido(ctx context.Context, in *Propuestagrpc, opts ...grpc.CallOption) (*Booleano, error)
	DownloaderDescargaLibro(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*ChunkLibro, error)
}

type dataNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeServiceClient(cc grpc.ClientConnInterface) DataNodeServiceClient {
	return &dataNodeServiceClient{cc}
}

func (c *dataNodeServiceClient) EnvioMensajeTest(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*MensajeTest, error) {
	out := new(MensajeTest)
	err := c.cc.Invoke(ctx, "/serverdatanode.DataNodeService/EnvioMensajeTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) UploaderSubeLibro(ctx context.Context, in *ChunkLibro, opts ...grpc.CallOption) (*MensajeTest, error) {
	out := new(MensajeTest)
	err := c.cc.Invoke(ctx, "/serverdatanode.DataNodeService/UploaderSubeLibro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) UploaderTerminoDeSubirLibro(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*MensajeTest, error) {
	out := new(MensajeTest)
	err := c.cc.Invoke(ctx, "/serverdatanode.DataNodeService/UploaderTerminoDeSubirLibro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Propuesta_Distribuido(ctx context.Context, in *Propuestagrpc, opts ...grpc.CallOption) (*Booleano, error) {
	out := new(Booleano)
	err := c.cc.Invoke(ctx, "/serverdatanode.DataNodeService/Propuesta_Distribuido", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) DownloaderDescargaLibro(ctx context.Context, in *MensajeTest, opts ...grpc.CallOption) (*ChunkLibro, error) {
	out := new(ChunkLibro)
	err := c.cc.Invoke(ctx, "/serverdatanode.DataNodeService/DownloaderDescargaLibro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServiceServer is the server API for DataNodeService service.
type DataNodeServiceServer interface {
	EnvioMensajeTest(context.Context, *MensajeTest) (*MensajeTest, error)
	UploaderSubeLibro(context.Context, *ChunkLibro) (*MensajeTest, error)
	UploaderTerminoDeSubirLibro(context.Context, *MensajeTest) (*MensajeTest, error)
	Propuesta_Distribuido(context.Context, *Propuestagrpc) (*Booleano, error)
	DownloaderDescargaLibro(context.Context, *MensajeTest) (*ChunkLibro, error)
}

// UnimplementedDataNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataNodeServiceServer struct {
}

func (*UnimplementedDataNodeServiceServer) EnvioMensajeTest(context.Context, *MensajeTest) (*MensajeTest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvioMensajeTest not implemented")
}
func (*UnimplementedDataNodeServiceServer) UploaderSubeLibro(context.Context, *ChunkLibro) (*MensajeTest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploaderSubeLibro not implemented")
}
func (*UnimplementedDataNodeServiceServer) UploaderTerminoDeSubirLibro(context.Context, *MensajeTest) (*MensajeTest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploaderTerminoDeSubirLibro not implemented")
}
func (*UnimplementedDataNodeServiceServer) Propuesta_Distribuido(context.Context, *Propuestagrpc) (*Booleano, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Propuesta_Distribuido not implemented")
}
func (*UnimplementedDataNodeServiceServer) DownloaderDescargaLibro(context.Context, *MensajeTest) (*ChunkLibro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloaderDescargaLibro not implemented")
}

func RegisterDataNodeServiceServer(s *grpc.Server, srv DataNodeServiceServer) {
	s.RegisterService(&_DataNodeService_serviceDesc, srv)
}

func _DataNodeService_EnvioMensajeTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).EnvioMensajeTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverdatanode.DataNodeService/EnvioMensajeTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).EnvioMensajeTest(ctx, req.(*MensajeTest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_UploaderSubeLibro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkLibro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).UploaderSubeLibro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverdatanode.DataNodeService/UploaderSubeLibro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).UploaderSubeLibro(ctx, req.(*ChunkLibro))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_UploaderTerminoDeSubirLibro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).UploaderTerminoDeSubirLibro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverdatanode.DataNodeService/UploaderTerminoDeSubirLibro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).UploaderTerminoDeSubirLibro(ctx, req.(*MensajeTest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Propuesta_Distribuido_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuestagrpc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Propuesta_Distribuido(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverdatanode.DataNodeService/Propuesta_Distribuido",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Propuesta_Distribuido(ctx, req.(*Propuestagrpc))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_DownloaderDescargaLibro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).DownloaderDescargaLibro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverdatanode.DataNodeService/DownloaderDescargaLibro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).DownloaderDescargaLibro(ctx, req.(*MensajeTest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverdatanode.DataNodeService",
	HandlerType: (*DataNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnvioMensajeTest",
			Handler:    _DataNodeService_EnvioMensajeTest_Handler,
		},
		{
			MethodName: "UploaderSubeLibro",
			Handler:    _DataNodeService_UploaderSubeLibro_Handler,
		},
		{
			MethodName: "UploaderTerminoDeSubirLibro",
			Handler:    _DataNodeService_UploaderTerminoDeSubirLibro_Handler,
		},
		{
			MethodName: "Propuesta_Distribuido",
			Handler:    _DataNodeService_Propuesta_Distribuido_Handler,
		},
		{
			MethodName: "DownloaderDescargaLibro",
			Handler:    _DataNodeService_DownloaderDescargaLibro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverdatanode.proto",
}
