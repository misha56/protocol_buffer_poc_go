// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/v1/api.proto

package apiv1

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

type GameState int32

const (
	GameState_GAME_STATE_PENDING_UNSPECIFIED GameState = 0
	GameState_GAME_STATE_ONGOING             GameState = 1
	GameState_GAME_STATE_PAUSED              GameState = 2
	GameState_GAME_STATE_ENDED               GameState = 3
)

// Enum value maps for GameState.
var (
	GameState_name = map[int32]string{
		0: "GAME_STATE_PENDING_UNSPECIFIED",
		1: "GAME_STATE_ONGOING",
		2: "GAME_STATE_PAUSED",
		3: "GAME_STATE_ENDED",
	}
	GameState_value = map[string]int32{
		"GAME_STATE_PENDING_UNSPECIFIED": 0,
		"GAME_STATE_ONGOING":             1,
		"GAME_STATE_PAUSED":              2,
		"GAME_STATE_ENDED":               3,
	}
)

func (x GameState) Enum() *GameState {
	p := new(GameState)
	*p = x
	return p
}

func (x GameState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_api_proto_enumTypes[0].Descriptor()
}

func (GameState) Type() protoreflect.EnumType {
	return &file_api_v1_api_proto_enumTypes[0]
}

func (x GameState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameState.Descriptor instead.
func (GameState) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{0}
}

type RegisterArenaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArenaUuid  string `protobuf:"bytes,1,opt,name=arena_uuid,json=arenaUuid,proto3" json:"arena_uuid,omitempty"`
	ArenaName  string `protobuf:"bytes,2,opt,name=arena_name,json=arenaName,proto3" json:"arena_name,omitempty"`
	ArenaEmail string `protobuf:"bytes,3,opt,name=arena_email,json=arenaEmail,proto3" json:"arena_email,omitempty"`
}

func (x *RegisterArenaRequest) Reset() {
	*x = RegisterArenaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterArenaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterArenaRequest) ProtoMessage() {}

func (x *RegisterArenaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterArenaRequest.ProtoReflect.Descriptor instead.
func (*RegisterArenaRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterArenaRequest) GetArenaUuid() string {
	if x != nil {
		return x.ArenaUuid
	}
	return ""
}

func (x *RegisterArenaRequest) GetArenaName() string {
	if x != nil {
		return x.ArenaName
	}
	return ""
}

func (x *RegisterArenaRequest) GetArenaEmail() string {
	if x != nil {
		return x.ArenaEmail
	}
	return ""
}

type GameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArenaUuid   string    `protobuf:"bytes,1,opt,name=arena_uuid,json=arenaUuid,proto3" json:"arena_uuid,omitempty"`
	GameUuid    string    `protobuf:"bytes,2,opt,name=game_uuid,json=gameUuid,proto3" json:"game_uuid,omitempty"`
	GameTitle   string    `protobuf:"bytes,3,opt,name=game_title,json=gameTitle,proto3" json:"game_title,omitempty"`
	GameState   GameState `protobuf:"varint,4,opt,name=game_state,json=gameState,proto3,enum=api.v1.GameState" json:"game_state,omitempty"`
	PlayerCount uint32    `protobuf:"varint,5,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
}

func (x *GameEvent) Reset() {
	*x = GameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEvent) ProtoMessage() {}

func (x *GameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEvent.ProtoReflect.Descriptor instead.
func (*GameEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *GameEvent) GetArenaUuid() string {
	if x != nil {
		return x.ArenaUuid
	}
	return ""
}

func (x *GameEvent) GetGameUuid() string {
	if x != nil {
		return x.GameUuid
	}
	return ""
}

func (x *GameEvent) GetGameTitle() string {
	if x != nil {
		return x.GameTitle
	}
	return ""
}

func (x *GameEvent) GetGameState() GameState {
	if x != nil {
		return x.GameState
	}
	return GameState_GAME_STATE_PENDING_UNSPECIFIED
}

func (x *GameEvent) GetPlayerCount() uint32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

var File_api_v1_api_proto protoreflect.FileDescriptor

var file_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x0a, 0x14, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x2a, 0x74, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x1e, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x4f, 0x4e, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x32, 0xa5, 0x02, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12,
	0x5a, 0x10, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_api_proto_rawDescOnce sync.Once
	file_api_v1_api_proto_rawDescData = file_api_v1_api_proto_rawDesc
)

func file_api_v1_api_proto_rawDescGZIP() []byte {
	file_api_v1_api_proto_rawDescOnce.Do(func() {
		file_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_api_proto_rawDescData)
	})
	return file_api_v1_api_proto_rawDescData
}

var file_api_v1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_api_proto_goTypes = []interface{}{
	(GameState)(0),               // 0: api.v1.GameState
	(*EmptyResponse)(nil),        // 1: api.v1.EmptyResponse
	(*RegisterArenaRequest)(nil), // 2: api.v1.RegisterArenaRequest
	(*GameEvent)(nil),            // 3: api.v1.GameEvent
}
var file_api_v1_api_proto_depIdxs = []int32{
	0, // 0: api.v1.GameEvent.game_state:type_name -> api.v1.GameState
	2, // 1: api.v1.InstallationService.RegisterInstallation:input_type -> api.v1.RegisterArenaRequest
	3, // 2: api.v1.InstallationService.NotifyGameCreated:input_type -> api.v1.GameEvent
	3, // 3: api.v1.InstallationService.NotifyGameStarted:input_type -> api.v1.GameEvent
	3, // 4: api.v1.InstallationService.NotifyGameEnded:input_type -> api.v1.GameEvent
	1, // 5: api.v1.InstallationService.RegisterInstallation:output_type -> api.v1.EmptyResponse
	1, // 6: api.v1.InstallationService.NotifyGameCreated:output_type -> api.v1.EmptyResponse
	1, // 7: api.v1.InstallationService.NotifyGameStarted:output_type -> api.v1.EmptyResponse
	1, // 8: api.v1.InstallationService.NotifyGameEnded:output_type -> api.v1.EmptyResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_api_proto_init() }
func file_api_v1_api_proto_init() {
	if File_api_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterArenaRequest); i {
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
		file_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEvent); i {
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
			RawDescriptor: file_api_v1_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_proto_goTypes,
		DependencyIndexes: file_api_v1_api_proto_depIdxs,
		EnumInfos:         file_api_v1_api_proto_enumTypes,
		MessageInfos:      file_api_v1_api_proto_msgTypes,
	}.Build()
	File_api_v1_api_proto = out.File
	file_api_v1_api_proto_rawDesc = nil
	file_api_v1_api_proto_goTypes = nil
	file_api_v1_api_proto_depIdxs = nil
}
