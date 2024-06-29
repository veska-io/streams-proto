// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: main/futures/exchanges_events.proto

package eeventspb

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event                      string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	EventTimestamp             uint64   `protobuf:"varint,2,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
	Exchange                   string   `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Market                     string   `protobuf:"bytes,4,opt,name=market,proto3" json:"market,omitempty"`
	Base                       string   `protobuf:"bytes,5,opt,name=base,proto3" json:"base,omitempty"`
	Quot                       string   `protobuf:"bytes,6,opt,name=quot,proto3" json:"quot,omitempty"`
	PriceOpen                  *float64 `protobuf:"fixed64,7,opt,name=price_open,json=priceOpen,proto3,oneof" json:"price_open,omitempty"`
	PriceClose                 *float64 `protobuf:"fixed64,8,opt,name=price_close,json=priceClose,proto3,oneof" json:"price_close,omitempty"`
	PriceHigh                  *float64 `protobuf:"fixed64,9,opt,name=price_high,json=priceHigh,proto3,oneof" json:"price_high,omitempty"`
	PriceLow                   *float64 `protobuf:"fixed64,10,opt,name=price_low,json=priceLow,proto3,oneof" json:"price_low,omitempty"`
	VolumeQuot                 *float64 `protobuf:"fixed64,11,opt,name=volume_quot,json=volumeQuot,proto3,oneof" json:"volume_quot,omitempty"`
	VolumeBase                 *float64 `protobuf:"fixed64,12,opt,name=volume_base,json=volumeBase,proto3,oneof" json:"volume_base,omitempty"`
	VolumeBaseSellTaker        *float64 `protobuf:"fixed64,13,opt,name=volume_base_sell_taker,json=volumeBaseSellTaker,proto3,oneof" json:"volume_base_sell_taker,omitempty"`
	VolumeBaseBuyTaker         *float64 `protobuf:"fixed64,14,opt,name=volume_base_buy_taker,json=volumeBaseBuyTaker,proto3,oneof" json:"volume_base_buy_taker,omitempty"`
	OiOpen                     *float64 `protobuf:"fixed64,15,opt,name=oi_open,json=oiOpen,proto3,oneof" json:"oi_open,omitempty"`
	TradesCount                *int32   `protobuf:"varint,16,opt,name=trades_count,json=tradesCount,proto3,oneof" json:"trades_count,omitempty"`
	LiquidationsSellCount      *int32   `protobuf:"varint,17,opt,name=liquidations_sell_count,json=liquidationsSellCount,proto3,oneof" json:"liquidations_sell_count,omitempty"`
	LiquidationsBuyCount       *int32   `protobuf:"varint,18,opt,name=liquidations_buy_count,json=liquidationsBuyCount,proto3,oneof" json:"liquidations_buy_count,omitempty"`
	LiquidationsSellBaseVolume *float64 `protobuf:"fixed64,19,opt,name=liquidations_sell_base_volume,json=liquidationsSellBaseVolume,proto3,oneof" json:"liquidations_sell_base_volume,omitempty"`
	LiquidationsBuyBaseVolume  *float64 `protobuf:"fixed64,20,opt,name=liquidations_buy_base_volume,json=liquidationsBuyBaseVolume,proto3,oneof" json:"liquidations_buy_base_volume,omitempty"`
	LiquidationsSellQuotVolume *float64 `protobuf:"fixed64,21,opt,name=liquidations_sell_quot_volume,json=liquidationsSellQuotVolume,proto3,oneof" json:"liquidations_sell_quot_volume,omitempty"`
	LiquidationsBuyQuotVolume  *float64 `protobuf:"fixed64,22,opt,name=liquidations_buy_quot_volume,json=liquidationsBuyQuotVolume,proto3,oneof" json:"liquidations_buy_quot_volume,omitempty"`
	FundingRate                *float64 `protobuf:"fixed64,23,opt,name=funding_rate,json=fundingRate,proto3,oneof" json:"funding_rate,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_futures_exchanges_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_main_futures_exchanges_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_main_futures_exchanges_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Event) GetEventTimestamp() uint64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

func (x *Event) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *Event) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *Event) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Event) GetQuot() string {
	if x != nil {
		return x.Quot
	}
	return ""
}

func (x *Event) GetPriceOpen() float64 {
	if x != nil && x.PriceOpen != nil {
		return *x.PriceOpen
	}
	return 0
}

func (x *Event) GetPriceClose() float64 {
	if x != nil && x.PriceClose != nil {
		return *x.PriceClose
	}
	return 0
}

func (x *Event) GetPriceHigh() float64 {
	if x != nil && x.PriceHigh != nil {
		return *x.PriceHigh
	}
	return 0
}

func (x *Event) GetPriceLow() float64 {
	if x != nil && x.PriceLow != nil {
		return *x.PriceLow
	}
	return 0
}

func (x *Event) GetVolumeQuot() float64 {
	if x != nil && x.VolumeQuot != nil {
		return *x.VolumeQuot
	}
	return 0
}

func (x *Event) GetVolumeBase() float64 {
	if x != nil && x.VolumeBase != nil {
		return *x.VolumeBase
	}
	return 0
}

func (x *Event) GetVolumeBaseSellTaker() float64 {
	if x != nil && x.VolumeBaseSellTaker != nil {
		return *x.VolumeBaseSellTaker
	}
	return 0
}

func (x *Event) GetVolumeBaseBuyTaker() float64 {
	if x != nil && x.VolumeBaseBuyTaker != nil {
		return *x.VolumeBaseBuyTaker
	}
	return 0
}

func (x *Event) GetOiOpen() float64 {
	if x != nil && x.OiOpen != nil {
		return *x.OiOpen
	}
	return 0
}

func (x *Event) GetTradesCount() int32 {
	if x != nil && x.TradesCount != nil {
		return *x.TradesCount
	}
	return 0
}

func (x *Event) GetLiquidationsSellCount() int32 {
	if x != nil && x.LiquidationsSellCount != nil {
		return *x.LiquidationsSellCount
	}
	return 0
}

func (x *Event) GetLiquidationsBuyCount() int32 {
	if x != nil && x.LiquidationsBuyCount != nil {
		return *x.LiquidationsBuyCount
	}
	return 0
}

func (x *Event) GetLiquidationsSellBaseVolume() float64 {
	if x != nil && x.LiquidationsSellBaseVolume != nil {
		return *x.LiquidationsSellBaseVolume
	}
	return 0
}

func (x *Event) GetLiquidationsBuyBaseVolume() float64 {
	if x != nil && x.LiquidationsBuyBaseVolume != nil {
		return *x.LiquidationsBuyBaseVolume
	}
	return 0
}

func (x *Event) GetLiquidationsSellQuotVolume() float64 {
	if x != nil && x.LiquidationsSellQuotVolume != nil {
		return *x.LiquidationsSellQuotVolume
	}
	return 0
}

func (x *Event) GetLiquidationsBuyQuotVolume() float64 {
	if x != nil && x.LiquidationsBuyQuotVolume != nil {
		return *x.LiquidationsBuyQuotVolume
	}
	return 0
}

func (x *Event) GetFundingRate() float64 {
	if x != nil && x.FundingRate != nil {
		return *x.FundingRate
	}
	return 0
}

var File_main_futures_exchanges_events_proto protoreflect.FileDescriptor

var file_main_futures_exchanges_events_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xee, 0x0a, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x75, 0x6f, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x71, 0x75, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0a,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x69, 0x67, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x09, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x48, 0x69, 0x67, 0x68, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x04, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48, 0x05, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x16, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x6b,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x48, 0x06, 0x52, 0x13, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x54, 0x61, 0x6b, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x36, 0x0a, 0x15, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x07, 0x52, 0x12, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x79, 0x54, 0x61, 0x6b, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6f, 0x69,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x48, 0x08, 0x52, 0x06, 0x6f,
	0x69, 0x4f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x09,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x3b, 0x0a, 0x17, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x0a, 0x52, 0x15, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a,
	0x16, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x75,
	0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0b, 0x52,
	0x14, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x75, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x1d, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x0c, 0x52, 0x1a, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x65, 0x6c, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x44, 0x0a, 0x1c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x62, 0x75, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x48, 0x0d, 0x52, 0x19, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73, 0x65, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x1d, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x48, 0x0e, 0x52,
	0x1a, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x6c,
	0x6c, 0x51, 0x75, 0x6f, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x44,
	0x0a, 0x1c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62,
	0x75, 0x79, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x0f, 0x52, 0x19, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x48, 0x10, 0x52, 0x0b, 0x66, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6f, 0x69, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x73, 0x6b, 0x61, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x3b, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_futures_exchanges_events_proto_rawDescOnce sync.Once
	file_main_futures_exchanges_events_proto_rawDescData = file_main_futures_exchanges_events_proto_rawDesc
)

func file_main_futures_exchanges_events_proto_rawDescGZIP() []byte {
	file_main_futures_exchanges_events_proto_rawDescOnce.Do(func() {
		file_main_futures_exchanges_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_futures_exchanges_events_proto_rawDescData)
	})
	return file_main_futures_exchanges_events_proto_rawDescData
}

var file_main_futures_exchanges_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_main_futures_exchanges_events_proto_goTypes = []interface{}{
	(*Event)(nil), // 0: exchanges_events.Event
}
var file_main_futures_exchanges_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_futures_exchanges_events_proto_init() }
func file_main_futures_exchanges_events_proto_init() {
	if File_main_futures_exchanges_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_futures_exchanges_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_main_futures_exchanges_events_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_futures_exchanges_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_main_futures_exchanges_events_proto_goTypes,
		DependencyIndexes: file_main_futures_exchanges_events_proto_depIdxs,
		MessageInfos:      file_main_futures_exchanges_events_proto_msgTypes,
	}.Build()
	File_main_futures_exchanges_events_proto = out.File
	file_main_futures_exchanges_events_proto_rawDesc = nil
	file_main_futures_exchanges_events_proto_goTypes = nil
	file_main_futures_exchanges_events_proto_depIdxs = nil
}
