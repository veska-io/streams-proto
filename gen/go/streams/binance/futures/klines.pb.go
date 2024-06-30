// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: binance/futures/klines.proto

package binancepb

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

type Kline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenTime                int64  `protobuf:"varint,1,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	Open                    string `protobuf:"bytes,2,opt,name=open,proto3" json:"open,omitempty"`
	High                    string `protobuf:"bytes,3,opt,name=high,proto3" json:"high,omitempty"`
	Low                     string `protobuf:"bytes,4,opt,name=low,proto3" json:"low,omitempty"`
	Close                   string `protobuf:"bytes,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume                  string `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
	CloseTime               int64  `protobuf:"varint,7,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	QuotAssetVolume         string `protobuf:"bytes,8,opt,name=quot_asset_volume,json=quotAssetVolume,proto3" json:"quot_asset_volume,omitempty"`
	TradeNum                int64  `protobuf:"varint,9,opt,name=trade_num,json=tradeNum,proto3" json:"trade_num,omitempty"`
	TakerBuyBaseAssetVolume string `protobuf:"bytes,10,opt,name=taker_buy_base_asset_volume,json=takerBuyBaseAssetVolume,proto3" json:"taker_buy_base_asset_volume,omitempty"`
	TakerBuyQuotAssetVolume string `protobuf:"bytes,11,opt,name=taker_buy_quot_asset_volume,json=takerBuyQuotAssetVolume,proto3" json:"taker_buy_quot_asset_volume,omitempty"`
	Symbol                  string `protobuf:"bytes,12,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Base                    string `protobuf:"bytes,13,opt,name=base,proto3" json:"base,omitempty"`
	Quot                    string `protobuf:"bytes,14,opt,name=quot,proto3" json:"quot,omitempty"`
}

func (x *Kline) Reset() {
	*x = Kline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binance_futures_klines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kline) ProtoMessage() {}

func (x *Kline) ProtoReflect() protoreflect.Message {
	mi := &file_binance_futures_klines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kline.ProtoReflect.Descriptor instead.
func (*Kline) Descriptor() ([]byte, []int) {
	return file_binance_futures_klines_proto_rawDescGZIP(), []int{0}
}

func (x *Kline) GetOpenTime() int64 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *Kline) GetOpen() string {
	if x != nil {
		return x.Open
	}
	return ""
}

func (x *Kline) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *Kline) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *Kline) GetClose() string {
	if x != nil {
		return x.Close
	}
	return ""
}

func (x *Kline) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *Kline) GetCloseTime() int64 {
	if x != nil {
		return x.CloseTime
	}
	return 0
}

func (x *Kline) GetQuotAssetVolume() string {
	if x != nil {
		return x.QuotAssetVolume
	}
	return ""
}

func (x *Kline) GetTradeNum() int64 {
	if x != nil {
		return x.TradeNum
	}
	return 0
}

func (x *Kline) GetTakerBuyBaseAssetVolume() string {
	if x != nil {
		return x.TakerBuyBaseAssetVolume
	}
	return ""
}

func (x *Kline) GetTakerBuyQuotAssetVolume() string {
	if x != nil {
		return x.TakerBuyQuotAssetVolume
	}
	return ""
}

func (x *Kline) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Kline) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Kline) GetQuot() string {
	if x != nil {
		return x.Quot
	}
	return ""
}

var File_binance_futures_klines_proto protoreflect.FileDescriptor

var file_binance_futures_klines_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x2f, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xb0, 0x03, 0x0a, 0x05, 0x4b, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x71, 0x75, 0x6f, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x3c,
	0x0a, 0x1b, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x1b,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x75, 0x6f, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x71, 0x75, 0x6f, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x73, 0x6b, 0x61, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x3b, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_binance_futures_klines_proto_rawDescOnce sync.Once
	file_binance_futures_klines_proto_rawDescData = file_binance_futures_klines_proto_rawDesc
)

func file_binance_futures_klines_proto_rawDescGZIP() []byte {
	file_binance_futures_klines_proto_rawDescOnce.Do(func() {
		file_binance_futures_klines_proto_rawDescData = protoimpl.X.CompressGZIP(file_binance_futures_klines_proto_rawDescData)
	})
	return file_binance_futures_klines_proto_rawDescData
}

var file_binance_futures_klines_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_binance_futures_klines_proto_goTypes = []interface{}{
	(*Kline)(nil), // 0: binance.Kline
}
var file_binance_futures_klines_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_binance_futures_klines_proto_init() }
func file_binance_futures_klines_proto_init() {
	if File_binance_futures_klines_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_binance_futures_klines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kline); i {
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
			RawDescriptor: file_binance_futures_klines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binance_futures_klines_proto_goTypes,
		DependencyIndexes: file_binance_futures_klines_proto_depIdxs,
		MessageInfos:      file_binance_futures_klines_proto_msgTypes,
	}.Build()
	File_binance_futures_klines_proto = out.File
	file_binance_futures_klines_proto_rawDesc = nil
	file_binance_futures_klines_proto_goTypes = nil
	file_binance_futures_klines_proto_depIdxs = nil
}
