# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: oracles/v1/oracle_spec.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='oracles/v1/oracle_spec.proto',
  package='oracles.v1',
  syntax='proto3',
  serialized_options=b'\n\037io.vegaprotocol.vega.oracles.v1Z*code.vegaprotocol.io/vega/proto/oracles/v1',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1coracles/v1/oracle_spec.proto\x12\noracles.v1\"b\n\x17OracleSpecConfiguration\x12\x19\n\x08pub_keys\x18\x01 \x03(\tR\x07pubKeys\x12,\n\x07\x66ilters\x18\x02 \x03(\x0b\x32\x12.oracles.v1.FilterR\x07\x66ilters\"\xa7\x02\n\nOracleSpec\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n\ncreated_at\x18\x02 \x01(\x03R\tcreatedAt\x12\x1d\n\nupdated_at\x18\x03 \x01(\x03R\tupdatedAt\x12\x19\n\x08pub_keys\x18\x04 \x03(\tR\x07pubKeys\x12,\n\x07\x66ilters\x18\x05 \x03(\x0b\x32\x12.oracles.v1.FilterR\x07\x66ilters\x12\x35\n\x06status\x18\x06 \x01(\x0e\x32\x1d.oracles.v1.OracleSpec.StatusR\x06status\"K\n\x06Status\x12\x16\n\x12STATUS_UNSPECIFIED\x10\x00\x12\x11\n\rSTATUS_ACTIVE\x10\x01\x12\x16\n\x12STATUS_DEACTIVATED\x10\x02\"j\n\x06\x46ilter\x12)\n\x03key\x18\x01 \x01(\x0b\x32\x17.oracles.v1.PropertyKeyR\x03key\x12\x35\n\nconditions\x18\x02 \x03(\x0b\x32\x15.oracles.v1.ConditionR\nconditions\"\xdd\x01\n\x0bPropertyKey\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x30\n\x04type\x18\x02 \x01(\x0e\x32\x1c.oracles.v1.PropertyKey.TypeR\x04type\"\x87\x01\n\x04Type\x12\x14\n\x10TYPE_UNSPECIFIED\x10\x00\x12\x0e\n\nTYPE_EMPTY\x10\x01\x12\x10\n\x0cTYPE_INTEGER\x10\x02\x12\x0f\n\x0bTYPE_STRING\x10\x03\x12\x10\n\x0cTYPE_BOOLEAN\x10\x04\x12\x10\n\x0cTYPE_DECIMAL\x10\x05\x12\x12\n\x0eTYPE_TIMESTAMP\x10\x06\"\x91\x02\n\tCondition\x12:\n\x08operator\x18\x01 \x01(\x0e\x32\x1e.oracles.v1.Condition.OperatorR\x08operator\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\"\xb1\x01\n\x08Operator\x12\x18\n\x14OPERATOR_UNSPECIFIED\x10\x00\x12\x13\n\x0fOPERATOR_EQUALS\x10\x01\x12\x19\n\x15OPERATOR_GREATER_THAN\x10\x02\x12\"\n\x1eOPERATOR_GREATER_THAN_OR_EQUAL\x10\x03\x12\x16\n\x12OPERATOR_LESS_THAN\x10\x04\x12\x1f\n\x1bOPERATOR_LESS_THAN_OR_EQUAL\x10\x05\x42M\n\x1fio.vegaprotocol.vega.oracles.v1Z*code.vegaprotocol.io/vega/proto/oracles/v1b\x06proto3'
)



_ORACLESPEC_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='oracles.v1.OracleSpec.Status',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STATUS_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='STATUS_ACTIVE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='STATUS_DEACTIVATED', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=365,
  serialized_end=440,
)
_sym_db.RegisterEnumDescriptor(_ORACLESPEC_STATUS)

_PROPERTYKEY_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='oracles.v1.PropertyKey.Type',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TYPE_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_EMPTY', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_INTEGER', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_STRING', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_BOOLEAN', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_DECIMAL', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TYPE_TIMESTAMP', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=637,
  serialized_end=772,
)
_sym_db.RegisterEnumDescriptor(_PROPERTYKEY_TYPE)

_CONDITION_OPERATOR = _descriptor.EnumDescriptor(
  name='Operator',
  full_name='oracles.v1.Condition.Operator',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_EQUALS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_GREATER_THAN', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_GREATER_THAN_OR_EQUAL', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_LESS_THAN', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OPERATOR_LESS_THAN_OR_EQUAL', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=871,
  serialized_end=1048,
)
_sym_db.RegisterEnumDescriptor(_CONDITION_OPERATOR)


_ORACLESPECCONFIGURATION = _descriptor.Descriptor(
  name='OracleSpecConfiguration',
  full_name='oracles.v1.OracleSpecConfiguration',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='pub_keys', full_name='oracles.v1.OracleSpecConfiguration.pub_keys', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubKeys', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='filters', full_name='oracles.v1.OracleSpecConfiguration.filters', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='filters', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=44,
  serialized_end=142,
)


_ORACLESPEC = _descriptor.Descriptor(
  name='OracleSpec',
  full_name='oracles.v1.OracleSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='oracles.v1.OracleSpec.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='oracles.v1.OracleSpec.created_at', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='createdAt', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='oracles.v1.OracleSpec.updated_at', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='updatedAt', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pub_keys', full_name='oracles.v1.OracleSpec.pub_keys', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubKeys', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='filters', full_name='oracles.v1.OracleSpec.filters', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='filters', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='oracles.v1.OracleSpec.status', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='status', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _ORACLESPEC_STATUS,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=145,
  serialized_end=440,
)


_FILTER = _descriptor.Descriptor(
  name='Filter',
  full_name='oracles.v1.Filter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='oracles.v1.Filter.key', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='key', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='conditions', full_name='oracles.v1.Filter.conditions', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='conditions', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=442,
  serialized_end=548,
)


_PROPERTYKEY = _descriptor.Descriptor(
  name='PropertyKey',
  full_name='oracles.v1.PropertyKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='oracles.v1.PropertyKey.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='oracles.v1.PropertyKey.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='type', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PROPERTYKEY_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=551,
  serialized_end=772,
)


_CONDITION = _descriptor.Descriptor(
  name='Condition',
  full_name='oracles.v1.Condition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='operator', full_name='oracles.v1.Condition.operator', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='operator', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='oracles.v1.Condition.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _CONDITION_OPERATOR,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=775,
  serialized_end=1048,
)

_ORACLESPECCONFIGURATION.fields_by_name['filters'].message_type = _FILTER
_ORACLESPEC.fields_by_name['filters'].message_type = _FILTER
_ORACLESPEC.fields_by_name['status'].enum_type = _ORACLESPEC_STATUS
_ORACLESPEC_STATUS.containing_type = _ORACLESPEC
_FILTER.fields_by_name['key'].message_type = _PROPERTYKEY
_FILTER.fields_by_name['conditions'].message_type = _CONDITION
_PROPERTYKEY.fields_by_name['type'].enum_type = _PROPERTYKEY_TYPE
_PROPERTYKEY_TYPE.containing_type = _PROPERTYKEY
_CONDITION.fields_by_name['operator'].enum_type = _CONDITION_OPERATOR
_CONDITION_OPERATOR.containing_type = _CONDITION
DESCRIPTOR.message_types_by_name['OracleSpecConfiguration'] = _ORACLESPECCONFIGURATION
DESCRIPTOR.message_types_by_name['OracleSpec'] = _ORACLESPEC
DESCRIPTOR.message_types_by_name['Filter'] = _FILTER
DESCRIPTOR.message_types_by_name['PropertyKey'] = _PROPERTYKEY
DESCRIPTOR.message_types_by_name['Condition'] = _CONDITION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

OracleSpecConfiguration = _reflection.GeneratedProtocolMessageType('OracleSpecConfiguration', (_message.Message,), {
  'DESCRIPTOR' : _ORACLESPECCONFIGURATION,
  '__module__' : 'oracles.v1.oracle_spec_pb2'
  # @@protoc_insertion_point(class_scope:oracles.v1.OracleSpecConfiguration)
  })
_sym_db.RegisterMessage(OracleSpecConfiguration)

OracleSpec = _reflection.GeneratedProtocolMessageType('OracleSpec', (_message.Message,), {
  'DESCRIPTOR' : _ORACLESPEC,
  '__module__' : 'oracles.v1.oracle_spec_pb2'
  # @@protoc_insertion_point(class_scope:oracles.v1.OracleSpec)
  })
_sym_db.RegisterMessage(OracleSpec)

Filter = _reflection.GeneratedProtocolMessageType('Filter', (_message.Message,), {
  'DESCRIPTOR' : _FILTER,
  '__module__' : 'oracles.v1.oracle_spec_pb2'
  # @@protoc_insertion_point(class_scope:oracles.v1.Filter)
  })
_sym_db.RegisterMessage(Filter)

PropertyKey = _reflection.GeneratedProtocolMessageType('PropertyKey', (_message.Message,), {
  'DESCRIPTOR' : _PROPERTYKEY,
  '__module__' : 'oracles.v1.oracle_spec_pb2'
  # @@protoc_insertion_point(class_scope:oracles.v1.PropertyKey)
  })
_sym_db.RegisterMessage(PropertyKey)

Condition = _reflection.GeneratedProtocolMessageType('Condition', (_message.Message,), {
  'DESCRIPTOR' : _CONDITION,
  '__module__' : 'oracles.v1.oracle_spec_pb2'
  # @@protoc_insertion_point(class_scope:oracles.v1.Condition)
  })
_sym_db.RegisterMessage(Condition)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
