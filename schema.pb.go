// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package sansigmabuffers

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderPanel_OrderSide int32

const (
	OrderPanel_Buy  OrderPanel_OrderSide = 0
	OrderPanel_Sell OrderPanel_OrderSide = 1
)

var OrderPanel_OrderSide_name = map[int32]string{
	0: "Buy",
	1: "Sell",
}

var OrderPanel_OrderSide_value = map[string]int32{
	"Buy":  0,
	"Sell": 1,
}

func (x OrderPanel_OrderSide) String() string {
	return proto.EnumName(OrderPanel_OrderSide_name, int32(x))
}

func (OrderPanel_OrderSide) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4, 0}
}

type Arbitrage_ArbitrageType int32

const (
	Arbitrage_Simple     Arbitrage_ArbitrageType = 0
	Arbitrage_Triangular Arbitrage_ArbitrageType = 1
	Arbitrage_Circle     Arbitrage_ArbitrageType = 2
)

var Arbitrage_ArbitrageType_name = map[int32]string{
	0: "Simple",
	1: "Triangular",
	2: "Circle",
}

var Arbitrage_ArbitrageType_value = map[string]int32{
	"Simple":     0,
	"Triangular": 1,
	"Circle":     2,
}

func (x Arbitrage_ArbitrageType) String() string {
	return proto.EnumName(Arbitrage_ArbitrageType_name, int32(x))
}

func (Arbitrage_ArbitrageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5, 0}
}

type BidAskSchema struct {
	Price                float64  `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidAskSchema) Reset()         { *m = BidAskSchema{} }
func (m *BidAskSchema) String() string { return proto.CompactTextString(m) }
func (*BidAskSchema) ProtoMessage()    {}
func (*BidAskSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *BidAskSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidAskSchema.Unmarshal(m, b)
}
func (m *BidAskSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidAskSchema.Marshal(b, m, deterministic)
}
func (m *BidAskSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidAskSchema.Merge(m, src)
}
func (m *BidAskSchema) XXX_Size() int {
	return xxx_messageInfo_BidAskSchema.Size(m)
}
func (m *BidAskSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_BidAskSchema.DiscardUnknown(m)
}

var xxx_messageInfo_BidAskSchema proto.InternalMessageInfo

func (m *BidAskSchema) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BidAskSchema) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type DepthSchema struct {
	Exchange             string          `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Base                 string          `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string          `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	Bids                 []*BidAskSchema `protobuf:"bytes,4,rep,name=bids,proto3" json:"bids,omitempty"`
	Asks                 []*BidAskSchema `protobuf:"bytes,5,rep,name=asks,proto3" json:"asks,omitempty"`
	Timestamp            int64           `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Key                  string          `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DepthSchema) Reset()         { *m = DepthSchema{} }
func (m *DepthSchema) String() string { return proto.CompactTextString(m) }
func (*DepthSchema) ProtoMessage()    {}
func (*DepthSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *DepthSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepthSchema.Unmarshal(m, b)
}
func (m *DepthSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepthSchema.Marshal(b, m, deterministic)
}
func (m *DepthSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepthSchema.Merge(m, src)
}
func (m *DepthSchema) XXX_Size() int {
	return xxx_messageInfo_DepthSchema.Size(m)
}
func (m *DepthSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_DepthSchema.DiscardUnknown(m)
}

var xxx_messageInfo_DepthSchema proto.InternalMessageInfo

func (m *DepthSchema) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *DepthSchema) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *DepthSchema) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *DepthSchema) GetBids() []*BidAskSchema {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *DepthSchema) GetAsks() []*BidAskSchema {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *DepthSchema) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DepthSchema) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SpreadSchema struct {
	LongExchange         string   `protobuf:"bytes,1,opt,name=long_exchange,json=longExchange,proto3" json:"long_exchange,omitempty"`
	ShortExchange        string   `protobuf:"bytes,2,opt,name=short_exchange,json=shortExchange,proto3" json:"short_exchange,omitempty"`
	Base                 string   `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string   `protobuf:"bytes,4,opt,name=quote,proto3" json:"quote,omitempty"`
	ProfitRate           float64  `protobuf:"fixed64,5,opt,name=profit_rate,json=profitRate,proto3" json:"profit_rate,omitempty"`
	Profit               float64  `protobuf:"fixed64,6,opt,name=profit,proto3" json:"profit,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	InvestedCapital      float64  `protobuf:"fixed64,8,opt,name=invested_capital,json=investedCapital,proto3" json:"invested_capital,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpreadSchema) Reset()         { *m = SpreadSchema{} }
func (m *SpreadSchema) String() string { return proto.CompactTextString(m) }
func (*SpreadSchema) ProtoMessage()    {}
func (*SpreadSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}

func (m *SpreadSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpreadSchema.Unmarshal(m, b)
}
func (m *SpreadSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpreadSchema.Marshal(b, m, deterministic)
}
func (m *SpreadSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpreadSchema.Merge(m, src)
}
func (m *SpreadSchema) XXX_Size() int {
	return xxx_messageInfo_SpreadSchema.Size(m)
}
func (m *SpreadSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_SpreadSchema.DiscardUnknown(m)
}

var xxx_messageInfo_SpreadSchema proto.InternalMessageInfo

func (m *SpreadSchema) GetLongExchange() string {
	if m != nil {
		return m.LongExchange
	}
	return ""
}

func (m *SpreadSchema) GetShortExchange() string {
	if m != nil {
		return m.ShortExchange
	}
	return ""
}

func (m *SpreadSchema) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *SpreadSchema) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *SpreadSchema) GetProfitRate() float64 {
	if m != nil {
		return m.ProfitRate
	}
	return 0
}

func (m *SpreadSchema) GetProfit() float64 {
	if m != nil {
		return m.Profit
	}
	return 0
}

func (m *SpreadSchema) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SpreadSchema) GetInvestedCapital() float64 {
	if m != nil {
		return m.InvestedCapital
	}
	return 0
}

type TradeSchema struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TradeId              string   `protobuf:"bytes,2,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Exchange             string   `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Base                 string   `protobuf:"bytes,4,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string   `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Amount               float64  `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp            int64    `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Side                 string   `protobuf:"bytes,9,opt,name=side,proto3" json:"side,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeSchema) Reset()         { *m = TradeSchema{} }
func (m *TradeSchema) String() string { return proto.CompactTextString(m) }
func (*TradeSchema) ProtoMessage()    {}
func (*TradeSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}

func (m *TradeSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSchema.Unmarshal(m, b)
}
func (m *TradeSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSchema.Marshal(b, m, deterministic)
}
func (m *TradeSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSchema.Merge(m, src)
}
func (m *TradeSchema) XXX_Size() int {
	return xxx_messageInfo_TradeSchema.Size(m)
}
func (m *TradeSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSchema.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSchema proto.InternalMessageInfo

func (m *TradeSchema) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TradeSchema) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *TradeSchema) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *TradeSchema) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *TradeSchema) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *TradeSchema) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TradeSchema) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TradeSchema) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TradeSchema) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

type OrderPanel struct {
	Exchange             string               `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Base                 string               `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string               `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	Side                 OrderPanel_OrderSide `protobuf:"varint,4,opt,name=side,proto3,enum=sansigmabuffers.OrderPanel_OrderSide" json:"side,omitempty"`
	Price                float64              `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	DepthAmount          float64              `protobuf:"fixed64,6,opt,name=depth_amount,json=depthAmount,proto3" json:"depth_amount,omitempty"`
	OrderAmount          float64              `protobuf:"fixed64,7,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderPanel) Reset()         { *m = OrderPanel{} }
func (m *OrderPanel) String() string { return proto.CompactTextString(m) }
func (*OrderPanel) ProtoMessage()    {}
func (*OrderPanel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4}
}

func (m *OrderPanel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPanel.Unmarshal(m, b)
}
func (m *OrderPanel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPanel.Marshal(b, m, deterministic)
}
func (m *OrderPanel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPanel.Merge(m, src)
}
func (m *OrderPanel) XXX_Size() int {
	return xxx_messageInfo_OrderPanel.Size(m)
}
func (m *OrderPanel) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPanel.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPanel proto.InternalMessageInfo

func (m *OrderPanel) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *OrderPanel) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *OrderPanel) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *OrderPanel) GetSide() OrderPanel_OrderSide {
	if m != nil {
		return m.Side
	}
	return OrderPanel_Buy
}

func (m *OrderPanel) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderPanel) GetDepthAmount() float64 {
	if m != nil {
		return m.DepthAmount
	}
	return 0
}

func (m *OrderPanel) GetOrderAmount() float64 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

type Arbitrage struct {
	Type                 Arbitrage_ArbitrageType `protobuf:"varint,1,opt,name=type,proto3,enum=sansigmabuffers.Arbitrage_ArbitrageType" json:"type,omitempty"`
	Orders               []*OrderPanel           `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
	Profit               float64                 `protobuf:"fixed64,3,opt,name=profit,proto3" json:"profit,omitempty"`
	ProfitRate           float64                 `protobuf:"fixed64,4,opt,name=profit_rate,json=profitRate,proto3" json:"profit_rate,omitempty"`
	Investment           float64                 `protobuf:"fixed64,5,opt,name=investment,proto3" json:"investment,omitempty"`
	ProfitAsset          string                  `protobuf:"bytes,6,opt,name=profit_asset,json=profitAsset,proto3" json:"profit_asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Arbitrage) Reset()         { *m = Arbitrage{} }
func (m *Arbitrage) String() string { return proto.CompactTextString(m) }
func (*Arbitrage) ProtoMessage()    {}
func (*Arbitrage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5}
}

func (m *Arbitrage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arbitrage.Unmarshal(m, b)
}
func (m *Arbitrage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arbitrage.Marshal(b, m, deterministic)
}
func (m *Arbitrage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arbitrage.Merge(m, src)
}
func (m *Arbitrage) XXX_Size() int {
	return xxx_messageInfo_Arbitrage.Size(m)
}
func (m *Arbitrage) XXX_DiscardUnknown() {
	xxx_messageInfo_Arbitrage.DiscardUnknown(m)
}

var xxx_messageInfo_Arbitrage proto.InternalMessageInfo

func (m *Arbitrage) GetType() Arbitrage_ArbitrageType {
	if m != nil {
		return m.Type
	}
	return Arbitrage_Simple
}

func (m *Arbitrage) GetOrders() []*OrderPanel {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *Arbitrage) GetProfit() float64 {
	if m != nil {
		return m.Profit
	}
	return 0
}

func (m *Arbitrage) GetProfitRate() float64 {
	if m != nil {
		return m.ProfitRate
	}
	return 0
}

func (m *Arbitrage) GetInvestment() float64 {
	if m != nil {
		return m.Investment
	}
	return 0
}

func (m *Arbitrage) GetProfitAsset() string {
	if m != nil {
		return m.ProfitAsset
	}
	return ""
}

type AssetPair struct {
	Base                 string   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string   `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetPair) Reset()         { *m = AssetPair{} }
func (m *AssetPair) String() string { return proto.CompactTextString(m) }
func (*AssetPair) ProtoMessage()    {}
func (*AssetPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{6}
}

func (m *AssetPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetPair.Unmarshal(m, b)
}
func (m *AssetPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetPair.Marshal(b, m, deterministic)
}
func (m *AssetPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetPair.Merge(m, src)
}
func (m *AssetPair) XXX_Size() int {
	return xxx_messageInfo_AssetPair.Size(m)
}
func (m *AssetPair) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetPair.DiscardUnknown(m)
}

var xxx_messageInfo_AssetPair proto.InternalMessageInfo

func (m *AssetPair) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *AssetPair) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func init() {
	proto.RegisterEnum("sansigmabuffers.OrderPanel_OrderSide", OrderPanel_OrderSide_name, OrderPanel_OrderSide_value)
	proto.RegisterEnum("sansigmabuffers.Arbitrage_ArbitrageType", Arbitrage_ArbitrageType_name, Arbitrage_ArbitrageType_value)
	proto.RegisterType((*BidAskSchema)(nil), "sansigmabuffers.BidAskSchema")
	proto.RegisterType((*DepthSchema)(nil), "sansigmabuffers.DepthSchema")
	proto.RegisterType((*SpreadSchema)(nil), "sansigmabuffers.SpreadSchema")
	proto.RegisterType((*TradeSchema)(nil), "sansigmabuffers.TradeSchema")
	proto.RegisterType((*OrderPanel)(nil), "sansigmabuffers.OrderPanel")
	proto.RegisterType((*Arbitrage)(nil), "sansigmabuffers.Arbitrage")
	proto.RegisterType((*AssetPair)(nil), "sansigmabuffers.AssetPair")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6a, 0xdb, 0x4a,
	0x14, 0x8d, 0x2e, 0xbe, 0x68, 0x5b, 0x71, 0xc4, 0x70, 0x38, 0xe8, 0xe4, 0xa4, 0x69, 0xa2, 0x12,
	0x70, 0x5f, 0x42, 0xeb, 0x50, 0x4a, 0x21, 0x2f, 0x4e, 0x1a, 0x4a, 0xa1, 0xd0, 0x20, 0x87, 0x42,
	0x9f, 0xcc, 0xd8, 0xda, 0xb1, 0x07, 0xeb, 0xd6, 0x99, 0x71, 0x89, 0x7f, 0xa0, 0xd0, 0xef, 0xe9,
	0x0f, 0xf5, 0xb5, 0x7f, 0x51, 0x34, 0x23, 0xdb, 0xf2, 0xa5, 0xa5, 0x0f, 0x7d, 0x9b, 0xbd, 0xb4,
	0x2f, 0xb3, 0xd6, 0xac, 0x6d, 0x83, 0x2b, 0x46, 0x13, 0x4c, 0xe8, 0x79, 0xce, 0x33, 0x99, 0x91,
	0x03, 0x41, 0x53, 0xc1, 0xc6, 0x09, 0x1d, 0xce, 0xee, 0xef, 0x91, 0x8b, 0xe0, 0x12, 0xdc, 0x2b,
	0x16, 0xf5, 0xc4, 0xb4, 0xaf, 0xd2, 0xc8, 0x3f, 0x50, 0xcb, 0x39, 0x1b, 0xa1, 0x6f, 0x9c, 0x18,
	0x1d, 0x23, 0xd4, 0x01, 0xf9, 0x17, 0xea, 0x34, 0xc9, 0x66, 0xa9, 0xf4, 0x4d, 0x05, 0x97, 0x51,
	0xf0, 0xc3, 0x80, 0xd6, 0x6b, 0xcc, 0xe5, 0xa4, 0xac, 0x3e, 0x84, 0x26, 0x3e, 0x8c, 0x26, 0x34,
	0x1d, 0xeb, 0x06, 0x4e, 0xb8, 0x8c, 0x09, 0x01, 0x7b, 0x48, 0x05, 0xaa, 0x0e, 0x4e, 0xa8, 0xce,
	0xc5, 0xb4, 0x4f, 0xb3, 0x4c, 0xa2, 0x6f, 0x29, 0x50, 0x07, 0xe4, 0x39, 0xd8, 0x43, 0x16, 0x09,
	0xdf, 0x3e, 0xb1, 0x3a, 0xad, 0xee, 0xa3, 0xf3, 0x8d, 0x3b, 0x9f, 0x57, 0x2f, 0x1c, 0xaa, 0xd4,
	0xa2, 0x84, 0x8a, 0xa9, 0xf0, 0x6b, 0x7f, 0x54, 0x52, 0xa4, 0x92, 0x23, 0x70, 0x24, 0x4b, 0x50,
	0x48, 0x9a, 0xe4, 0x7e, 0xfd, 0xc4, 0xe8, 0x58, 0xe1, 0x0a, 0x20, 0x1e, 0x58, 0x53, 0x9c, 0xfb,
	0x0d, 0x75, 0xaf, 0xe2, 0x18, 0x7c, 0x31, 0xc1, 0xed, 0xe7, 0x1c, 0x69, 0x54, 0x92, 0x7d, 0x02,
	0xfb, 0x71, 0x96, 0x8e, 0x07, 0x1b, 0x8c, 0xdd, 0x02, 0xbc, 0x59, 0xb0, 0x3e, 0x83, 0xb6, 0x98,
	0x64, 0x5c, 0xae, 0xb2, 0x34, 0xff, 0x7d, 0x85, 0xde, 0x6c, 0x8a, 0x63, 0xed, 0x12, 0xc7, 0xae,
	0x8a, 0xf3, 0x18, 0x5a, 0x39, 0xcf, 0xee, 0x99, 0x1c, 0x70, 0x2a, 0xd1, 0xaf, 0xa9, 0xf7, 0x00,
	0x0d, 0x85, 0x54, 0xaa, 0xb7, 0xd2, 0x91, 0x22, 0x65, 0x84, 0x65, 0xb4, 0xce, 0xb7, 0xb1, 0xc9,
	0xf7, 0x29, 0x78, 0x2c, 0xfd, 0x8c, 0x42, 0x62, 0x34, 0x18, 0xd1, 0x9c, 0x49, 0x1a, 0xfb, 0x4d,
	0x55, 0x7f, 0xb0, 0xc0, 0xaf, 0x35, 0x1c, 0x7c, 0x37, 0xa0, 0x75, 0xc7, 0x69, 0x84, 0xa5, 0x0e,
	0xa5, 0x54, 0xc6, 0x52, 0x2a, 0xf2, 0x1f, 0x34, 0x65, 0x91, 0x30, 0x60, 0x51, 0x49, 0xb7, 0xa1,
	0xe2, 0xb7, 0xd1, 0x9a, 0x43, 0xac, 0x5f, 0x38, 0xc4, 0xde, 0x25, 0x42, 0xad, 0x2a, 0xc2, 0xd2,
	0xa5, 0xf5, 0xdd, 0x2e, 0x6d, 0x54, 0x5d, 0xba, 0xce, 0xbc, 0xb9, 0xc9, 0x9c, 0x80, 0x2d, 0x58,
	0x84, 0xbe, 0xa3, 0xa7, 0x16, 0xe7, 0xe0, 0xab, 0x09, 0xf0, 0x9e, 0x47, 0xc8, 0x6f, 0x69, 0x8a,
	0xf1, 0x5f, 0xb2, 0xf5, 0xab, 0x72, 0x50, 0x41, 0xaf, 0xdd, 0x3d, 0xdb, 0xf2, 0xe8, 0x6a, 0xa0,
	0x3e, 0xf6, 0x59, 0x84, 0xfa, 0x3e, 0x2b, 0xbe, 0xb5, 0x2a, 0xdf, 0x53, 0x70, 0xa3, 0x62, 0xf9,
	0x06, 0x25, 0x6b, 0x2d, 0x46, 0x4b, 0x61, 0x3d, 0x4d, 0xfd, 0x14, 0xdc, 0xac, 0xe8, 0x35, 0x58,
	0x13, 0xa6, 0xa5, 0x30, 0x9d, 0x12, 0x1c, 0x83, 0xb3, 0x1c, 0x47, 0x1a, 0x60, 0x5d, 0xcd, 0xe6,
	0xde, 0x1e, 0x69, 0x82, 0xdd, 0xc7, 0x38, 0xf6, 0x8c, 0xe0, 0x9b, 0x09, 0x4e, 0x8f, 0x0f, 0x99,
	0xe4, 0x74, 0x8c, 0xe4, 0x12, 0x6c, 0x39, 0xcf, 0xb5, 0x0c, 0xed, 0x6e, 0x67, 0x8b, 0xc4, 0x32,
	0x73, 0x75, 0xba, 0x9b, 0xe7, 0x18, 0xaa, 0x2a, 0x72, 0x01, 0x75, 0x35, 0x5a, 0xf8, 0xa6, 0x5a,
	0xd4, 0xff, 0x7f, 0x23, 0x42, 0x58, 0xa6, 0x56, 0x0c, 0x6d, 0xad, 0x19, 0x7a, 0x63, 0x13, 0xec,
	0xad, 0x4d, 0x38, 0x06, 0xd0, 0xde, 0x4d, 0x30, 0x95, 0x8b, 0x4d, 0x59, 0x21, 0x85, 0x38, 0x65,
	0x03, 0x2a, 0x04, 0x6a, 0xfd, 0x9c, 0xb0, 0x6c, 0xda, 0x2b, 0xa0, 0xe0, 0x25, 0xec, 0xaf, 0xf1,
	0x20, 0x00, 0xf5, 0x3e, 0x4b, 0xf2, 0x18, 0xbd, 0x3d, 0xd2, 0x06, 0xb8, 0xe3, 0x8c, 0xa6, 0xe3,
	0x59, 0x4c, 0xb9, 0x67, 0x14, 0xdf, 0xae, 0x19, 0x1f, 0xc5, 0xe8, 0x99, 0xc1, 0x0b, 0x70, 0x54,
	0x87, 0x5b, 0xca, 0xf8, 0xd2, 0x23, 0xc6, 0x2e, 0x8f, 0x98, 0x15, 0x8f, 0x74, 0x3f, 0x82, 0xf3,
	0x61, 0x86, 0x09, 0x3e, 0xf4, 0x72, 0x46, 0xde, 0x81, 0xfb, 0x06, 0xa5, 0x52, 0xe4, 0x2a, 0xcb,
	0xa6, 0xe4, 0x70, 0x5b, 0xed, 0xc5, 0x88, 0xc3, 0xa3, 0xad, 0x6f, 0x95, 0xdf, 0xe5, 0x60, 0xef,
	0x99, 0x31, 0xac, 0xab, 0x7f, 0x80, 0x8b, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x10, 0x44,
	0x92, 0x11, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VuemexApiClient is the client API for VuemexApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VuemexApiClient interface {
	GetOrderBook(ctx context.Context, in *AssetPair, opts ...grpc.CallOption) (VuemexApi_GetOrderBookClient, error)
}

type vuemexApiClient struct {
	cc *grpc.ClientConn
}

func NewVuemexApiClient(cc *grpc.ClientConn) VuemexApiClient {
	return &vuemexApiClient{cc}
}

func (c *vuemexApiClient) GetOrderBook(ctx context.Context, in *AssetPair, opts ...grpc.CallOption) (VuemexApi_GetOrderBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VuemexApi_serviceDesc.Streams[0], "/sansigmabuffers.VuemexApi/GetOrderBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &vuemexApiGetOrderBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VuemexApi_GetOrderBookClient interface {
	Recv() (*DepthSchema, error)
	grpc.ClientStream
}

type vuemexApiGetOrderBookClient struct {
	grpc.ClientStream
}

func (x *vuemexApiGetOrderBookClient) Recv() (*DepthSchema, error) {
	m := new(DepthSchema)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VuemexApiServer is the server API for VuemexApi service.
type VuemexApiServer interface {
	GetOrderBook(*AssetPair, VuemexApi_GetOrderBookServer) error
}

func RegisterVuemexApiServer(s *grpc.Server, srv VuemexApiServer) {
	s.RegisterService(&_VuemexApi_serviceDesc, srv)
}

func _VuemexApi_GetOrderBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AssetPair)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VuemexApiServer).GetOrderBook(m, &vuemexApiGetOrderBookServer{stream})
}

type VuemexApi_GetOrderBookServer interface {
	Send(*DepthSchema) error
	grpc.ServerStream
}

type vuemexApiGetOrderBookServer struct {
	grpc.ServerStream
}

func (x *vuemexApiGetOrderBookServer) Send(m *DepthSchema) error {
	return x.ServerStream.SendMsg(m)
}

var _VuemexApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sansigmabuffers.VuemexApi",
	HandlerType: (*VuemexApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrderBook",
			Handler:       _VuemexApi_GetOrderBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema.proto",
}
