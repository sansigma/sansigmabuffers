# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema.proto',
  package='sansigmabuffers',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0cschema.proto\x12\x0fsansigmabuffers\"-\n\x0c\x42idAskSchema\x12\r\n\x05price\x18\x01 \x01(\x01\x12\x0e\n\x06\x61mount\x18\x02 \x01(\x01\"\xb6\x01\n\x0b\x44\x65pthSchema\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0c\n\x04\x62\x61se\x18\x02 \x01(\t\x12\r\n\x05quote\x18\x03 \x01(\t\x12+\n\x04\x62ids\x18\x04 \x03(\x0b\x32\x1d.sansigmabuffers.BidAskSchema\x12+\n\x04\x61sks\x18\x05 \x03(\x0b\x32\x1d.sansigmabuffers.BidAskSchema\x12\x11\n\ttimestamp\x18\x06 \x01(\x03\x12\x0b\n\x03key\x18\x07 \x01(\t\"\xac\x01\n\x0cSpreadSchema\x12\x15\n\rlong_exchange\x18\x01 \x01(\t\x12\x16\n\x0eshort_exchange\x18\x02 \x01(\t\x12\x0c\n\x04\x62\x61se\x18\x03 \x01(\t\x12\r\n\x05quote\x18\x04 \x01(\t\x12\x13\n\x0bprofit_rate\x18\x05 \x01(\x01\x12\x0e\n\x06profit\x18\x06 \x01(\x01\x12\x11\n\ttimestamp\x18\x07 \x01(\x03\x12\x18\n\x10invested_capital\x18\x08 \x01(\x01\"\x9b\x01\n\x0bTradeSchema\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x10\n\x08trade_id\x18\x02 \x01(\t\x12\x10\n\x08\x65xchange\x18\x03 \x01(\t\x12\x0c\n\x04\x62\x61se\x18\x04 \x01(\t\x12\r\n\x05quote\x18\x05 \x01(\t\x12\r\n\x05price\x18\x06 \x01(\x01\x12\x0e\n\x06\x61mount\x18\x07 \x01(\x01\x12\x11\n\ttimestamp\x18\x08 \x01(\x03\x12\x0c\n\x04side\x18\t \x01(\t\"\xb2\x01\n\x10OrderPanelSchema\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0c\n\x04\x62\x61se\x18\x02 \x01(\t\x12\r\n\x05quote\x18\x03 \x01(\t\x12\x34\n\norder_side\x18\x04 \x01(\x0e\x32 .sansigmabuffers.OrderSideSchema\x12\r\n\x05price\x18\x05 \x01(\x01\x12\x14\n\x0c\x64\x65pth_amount\x18\x06 \x01(\x01\x12\x14\n\x0corder_amount\x18\x07 \x01(\x02\"\xca\x01\n\x0f\x41rbitrageSchema\x12\x32\n\x04type\x18\x01 \x01(\x0e\x32$.sansigmabuffers.ArbitrageTypeSchema\x12\x31\n\x06orders\x18\x02 \x03(\x0b\x32!.sansigmabuffers.OrderPanelSchema\x12\x0e\n\x06profit\x18\x03 \x01(\x01\x12\x13\n\x0bprofit_rate\x18\x04 \x01(\x01\x12\x12\n\ninvestment\x18\x05 \x01(\x01\x12\x17\n\x0fprofit_currency\x18\x06 \x01(\x01*$\n\x0fOrderSideSchema\x12\x07\n\x03\x42uy\x10\x00\x12\x08\n\x04Sell\x10\x01*=\n\x13\x41rbitrageTypeSchema\x12\n\n\x06Simple\x10\x00\x12\x0e\n\nTriangular\x10\x01\x12\n\n\x06\x43ircle\x10\x02\x62\x06proto3')
)

_ORDERSIDESCHEMA = _descriptor.EnumDescriptor(
  name='OrderSideSchema',
  full_name='sansigmabuffers.OrderSideSchema',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Buy', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Sell', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=984,
  serialized_end=1020,
)
_sym_db.RegisterEnumDescriptor(_ORDERSIDESCHEMA)

OrderSideSchema = enum_type_wrapper.EnumTypeWrapper(_ORDERSIDESCHEMA)
_ARBITRAGETYPESCHEMA = _descriptor.EnumDescriptor(
  name='ArbitrageTypeSchema',
  full_name='sansigmabuffers.ArbitrageTypeSchema',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Simple', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Triangular', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Circle', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1022,
  serialized_end=1083,
)
_sym_db.RegisterEnumDescriptor(_ARBITRAGETYPESCHEMA)

ArbitrageTypeSchema = enum_type_wrapper.EnumTypeWrapper(_ARBITRAGETYPESCHEMA)
Buy = 0
Sell = 1
Simple = 0
Triangular = 1
Circle = 2



_BIDASKSCHEMA = _descriptor.Descriptor(
  name='BidAskSchema',
  full_name='sansigmabuffers.BidAskSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='price', full_name='sansigmabuffers.BidAskSchema.price', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='amount', full_name='sansigmabuffers.BidAskSchema.amount', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=33,
  serialized_end=78,
)


_DEPTHSCHEMA = _descriptor.Descriptor(
  name='DepthSchema',
  full_name='sansigmabuffers.DepthSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='exchange', full_name='sansigmabuffers.DepthSchema.exchange', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='base', full_name='sansigmabuffers.DepthSchema.base', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quote', full_name='sansigmabuffers.DepthSchema.quote', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bids', full_name='sansigmabuffers.DepthSchema.bids', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='asks', full_name='sansigmabuffers.DepthSchema.asks', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='sansigmabuffers.DepthSchema.timestamp', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key', full_name='sansigmabuffers.DepthSchema.key', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=81,
  serialized_end=263,
)


_SPREADSCHEMA = _descriptor.Descriptor(
  name='SpreadSchema',
  full_name='sansigmabuffers.SpreadSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='long_exchange', full_name='sansigmabuffers.SpreadSchema.long_exchange', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='short_exchange', full_name='sansigmabuffers.SpreadSchema.short_exchange', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='base', full_name='sansigmabuffers.SpreadSchema.base', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quote', full_name='sansigmabuffers.SpreadSchema.quote', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='profit_rate', full_name='sansigmabuffers.SpreadSchema.profit_rate', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='profit', full_name='sansigmabuffers.SpreadSchema.profit', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='sansigmabuffers.SpreadSchema.timestamp', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='invested_capital', full_name='sansigmabuffers.SpreadSchema.invested_capital', index=7,
      number=8, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=266,
  serialized_end=438,
)


_TRADESCHEMA = _descriptor.Descriptor(
  name='TradeSchema',
  full_name='sansigmabuffers.TradeSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='sansigmabuffers.TradeSchema.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='trade_id', full_name='sansigmabuffers.TradeSchema.trade_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='exchange', full_name='sansigmabuffers.TradeSchema.exchange', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='base', full_name='sansigmabuffers.TradeSchema.base', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quote', full_name='sansigmabuffers.TradeSchema.quote', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sansigmabuffers.TradeSchema.price', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='amount', full_name='sansigmabuffers.TradeSchema.amount', index=6,
      number=7, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='sansigmabuffers.TradeSchema.timestamp', index=7,
      number=8, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='side', full_name='sansigmabuffers.TradeSchema.side', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=441,
  serialized_end=596,
)


_ORDERPANELSCHEMA = _descriptor.Descriptor(
  name='OrderPanelSchema',
  full_name='sansigmabuffers.OrderPanelSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='exchange', full_name='sansigmabuffers.OrderPanelSchema.exchange', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='base', full_name='sansigmabuffers.OrderPanelSchema.base', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quote', full_name='sansigmabuffers.OrderPanelSchema.quote', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order_side', full_name='sansigmabuffers.OrderPanelSchema.order_side', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sansigmabuffers.OrderPanelSchema.price', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='depth_amount', full_name='sansigmabuffers.OrderPanelSchema.depth_amount', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order_amount', full_name='sansigmabuffers.OrderPanelSchema.order_amount', index=6,
      number=7, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=599,
  serialized_end=777,
)


_ARBITRAGESCHEMA = _descriptor.Descriptor(
  name='ArbitrageSchema',
  full_name='sansigmabuffers.ArbitrageSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='sansigmabuffers.ArbitrageSchema.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='orders', full_name='sansigmabuffers.ArbitrageSchema.orders', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='profit', full_name='sansigmabuffers.ArbitrageSchema.profit', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='profit_rate', full_name='sansigmabuffers.ArbitrageSchema.profit_rate', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='investment', full_name='sansigmabuffers.ArbitrageSchema.investment', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='profit_currency', full_name='sansigmabuffers.ArbitrageSchema.profit_currency', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=780,
  serialized_end=982,
)

_DEPTHSCHEMA.fields_by_name['bids'].message_type = _BIDASKSCHEMA
_DEPTHSCHEMA.fields_by_name['asks'].message_type = _BIDASKSCHEMA
_ORDERPANELSCHEMA.fields_by_name['order_side'].enum_type = _ORDERSIDESCHEMA
_ARBITRAGESCHEMA.fields_by_name['type'].enum_type = _ARBITRAGETYPESCHEMA
_ARBITRAGESCHEMA.fields_by_name['orders'].message_type = _ORDERPANELSCHEMA
DESCRIPTOR.message_types_by_name['BidAskSchema'] = _BIDASKSCHEMA
DESCRIPTOR.message_types_by_name['DepthSchema'] = _DEPTHSCHEMA
DESCRIPTOR.message_types_by_name['SpreadSchema'] = _SPREADSCHEMA
DESCRIPTOR.message_types_by_name['TradeSchema'] = _TRADESCHEMA
DESCRIPTOR.message_types_by_name['OrderPanelSchema'] = _ORDERPANELSCHEMA
DESCRIPTOR.message_types_by_name['ArbitrageSchema'] = _ARBITRAGESCHEMA
DESCRIPTOR.enum_types_by_name['OrderSideSchema'] = _ORDERSIDESCHEMA
DESCRIPTOR.enum_types_by_name['ArbitrageTypeSchema'] = _ARBITRAGETYPESCHEMA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

BidAskSchema = _reflection.GeneratedProtocolMessageType('BidAskSchema', (_message.Message,), dict(
  DESCRIPTOR = _BIDASKSCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.BidAskSchema)
  ))
_sym_db.RegisterMessage(BidAskSchema)

DepthSchema = _reflection.GeneratedProtocolMessageType('DepthSchema', (_message.Message,), dict(
  DESCRIPTOR = _DEPTHSCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.DepthSchema)
  ))
_sym_db.RegisterMessage(DepthSchema)

SpreadSchema = _reflection.GeneratedProtocolMessageType('SpreadSchema', (_message.Message,), dict(
  DESCRIPTOR = _SPREADSCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.SpreadSchema)
  ))
_sym_db.RegisterMessage(SpreadSchema)

TradeSchema = _reflection.GeneratedProtocolMessageType('TradeSchema', (_message.Message,), dict(
  DESCRIPTOR = _TRADESCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.TradeSchema)
  ))
_sym_db.RegisterMessage(TradeSchema)

OrderPanelSchema = _reflection.GeneratedProtocolMessageType('OrderPanelSchema', (_message.Message,), dict(
  DESCRIPTOR = _ORDERPANELSCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.OrderPanelSchema)
  ))
_sym_db.RegisterMessage(OrderPanelSchema)

ArbitrageSchema = _reflection.GeneratedProtocolMessageType('ArbitrageSchema', (_message.Message,), dict(
  DESCRIPTOR = _ARBITRAGESCHEMA,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:sansigmabuffers.ArbitrageSchema)
  ))
_sym_db.RegisterMessage(ArbitrageSchema)


# @@protoc_insertion_point(module_scope)
