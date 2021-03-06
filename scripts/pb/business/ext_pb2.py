# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: business.ext.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='business.ext.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12\x62usiness.ext.proto\x12\x02pb\"B\n\tSignInReq\x12\x14\n\x0cphone_number\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\x11\n\tdevice_id\x18\x03 \x01(\x03\"<\n\nSignInResp\x12\x0e\n\x06is_new\x18\x01 \x01(\x08\x12\x0f\n\x07user_id\x18\x02 \x01(\x03\x12\r\n\x05token\x18\x03 \x01(\t\"\x83\x01\n\x04User\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x10\n\x08nickname\x18\x02 \x01(\t\x12\x0b\n\x03sex\x18\x03 \x01(\x05\x12\x12\n\navatar_url\x18\x04 \x01(\t\x12\r\n\x05\x65xtra\x18\x05 \x01(\t\x12\x13\n\x0b\x63reate_time\x18\x06 \x01(\x03\x12\x13\n\x0bupdate_time\x18\x07 \x01(\x03\"\x1d\n\nGetUserReq\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\"%\n\x0bGetUserResp\x12\x16\n\x04user\x18\x01 \x01(\x0b\x32\x08.pb.User\"Q\n\rUpdateUserReq\x12\x10\n\x08nickname\x18\x01 \x01(\t\x12\x0b\n\x03sex\x18\x02 \x01(\x05\x12\x12\n\navatar_url\x18\x03 \x01(\t\x12\r\n\x05\x65xtra\x18\x04 \x01(\t\"\x10\n\x0eUpdateUserResp\"\x1c\n\rSearchUserReq\x12\x0b\n\x03key\x18\x01 \x01(\t\")\n\x0eSearchUserResp\x12\x17\n\x05users\x18\x01 \x03(\x0b\x32\x08.pb.User2\xcc\x01\n\x0b\x42usinessExt\x12\'\n\x06SignIn\x12\r.pb.SignInReq\x1a\x0e.pb.SignInResp\x12*\n\x07GetUser\x12\x0e.pb.GetUserReq\x1a\x0f.pb.GetUserResp\x12\x33\n\nUpdateUser\x12\x11.pb.UpdateUserReq\x1a\x12.pb.UpdateUserResp\x12\x33\n\nSearchUser\x12\x11.pb.SearchUserReq\x1a\x12.pb.SearchUserRespb\x06proto3'
)




_SIGNINREQ = _descriptor.Descriptor(
  name='SignInReq',
  full_name='pb.SignInReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='phone_number', full_name='pb.SignInReq.phone_number', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='pb.SignInReq.code', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_id', full_name='pb.SignInReq.device_id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=26,
  serialized_end=92,
)


_SIGNINRESP = _descriptor.Descriptor(
  name='SignInResp',
  full_name='pb.SignInResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='is_new', full_name='pb.SignInResp.is_new', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='pb.SignInResp.user_id', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='token', full_name='pb.SignInResp.token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=94,
  serialized_end=154,
)


_USER = _descriptor.Descriptor(
  name='User',
  full_name='pb.User',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='pb.User.user_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='nickname', full_name='pb.User.nickname', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sex', full_name='pb.User.sex', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='avatar_url', full_name='pb.User.avatar_url', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='extra', full_name='pb.User.extra', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='pb.User.create_time', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='update_time', full_name='pb.User.update_time', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=157,
  serialized_end=288,
)


_GETUSERREQ = _descriptor.Descriptor(
  name='GetUserReq',
  full_name='pb.GetUserReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='pb.GetUserReq.user_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=290,
  serialized_end=319,
)


_GETUSERRESP = _descriptor.Descriptor(
  name='GetUserResp',
  full_name='pb.GetUserResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='pb.GetUserResp.user', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=321,
  serialized_end=358,
)


_UPDATEUSERREQ = _descriptor.Descriptor(
  name='UpdateUserReq',
  full_name='pb.UpdateUserReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='nickname', full_name='pb.UpdateUserReq.nickname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sex', full_name='pb.UpdateUserReq.sex', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='avatar_url', full_name='pb.UpdateUserReq.avatar_url', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='extra', full_name='pb.UpdateUserReq.extra', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=360,
  serialized_end=441,
)


_UPDATEUSERRESP = _descriptor.Descriptor(
  name='UpdateUserResp',
  full_name='pb.UpdateUserResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=443,
  serialized_end=459,
)


_SEARCHUSERREQ = _descriptor.Descriptor(
  name='SearchUserReq',
  full_name='pb.SearchUserReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='pb.SearchUserReq.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=461,
  serialized_end=489,
)


_SEARCHUSERRESP = _descriptor.Descriptor(
  name='SearchUserResp',
  full_name='pb.SearchUserResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='users', full_name='pb.SearchUserResp.users', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=491,
  serialized_end=532,
)

_GETUSERRESP.fields_by_name['user'].message_type = _USER
_SEARCHUSERRESP.fields_by_name['users'].message_type = _USER
DESCRIPTOR.message_types_by_name['SignInReq'] = _SIGNINREQ
DESCRIPTOR.message_types_by_name['SignInResp'] = _SIGNINRESP
DESCRIPTOR.message_types_by_name['User'] = _USER
DESCRIPTOR.message_types_by_name['GetUserReq'] = _GETUSERREQ
DESCRIPTOR.message_types_by_name['GetUserResp'] = _GETUSERRESP
DESCRIPTOR.message_types_by_name['UpdateUserReq'] = _UPDATEUSERREQ
DESCRIPTOR.message_types_by_name['UpdateUserResp'] = _UPDATEUSERRESP
DESCRIPTOR.message_types_by_name['SearchUserReq'] = _SEARCHUSERREQ
DESCRIPTOR.message_types_by_name['SearchUserResp'] = _SEARCHUSERRESP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SignInReq = _reflection.GeneratedProtocolMessageType('SignInReq', (_message.Message,), {
  'DESCRIPTOR' : _SIGNINREQ,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.SignInReq)
  })
_sym_db.RegisterMessage(SignInReq)

SignInResp = _reflection.GeneratedProtocolMessageType('SignInResp', (_message.Message,), {
  'DESCRIPTOR' : _SIGNINRESP,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.SignInResp)
  })
_sym_db.RegisterMessage(SignInResp)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), {
  'DESCRIPTOR' : _USER,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.User)
  })
_sym_db.RegisterMessage(User)

GetUserReq = _reflection.GeneratedProtocolMessageType('GetUserReq', (_message.Message,), {
  'DESCRIPTOR' : _GETUSERREQ,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetUserReq)
  })
_sym_db.RegisterMessage(GetUserReq)

GetUserResp = _reflection.GeneratedProtocolMessageType('GetUserResp', (_message.Message,), {
  'DESCRIPTOR' : _GETUSERRESP,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetUserResp)
  })
_sym_db.RegisterMessage(GetUserResp)

UpdateUserReq = _reflection.GeneratedProtocolMessageType('UpdateUserReq', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEUSERREQ,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.UpdateUserReq)
  })
_sym_db.RegisterMessage(UpdateUserReq)

UpdateUserResp = _reflection.GeneratedProtocolMessageType('UpdateUserResp', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEUSERRESP,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.UpdateUserResp)
  })
_sym_db.RegisterMessage(UpdateUserResp)

SearchUserReq = _reflection.GeneratedProtocolMessageType('SearchUserReq', (_message.Message,), {
  'DESCRIPTOR' : _SEARCHUSERREQ,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.SearchUserReq)
  })
_sym_db.RegisterMessage(SearchUserReq)

SearchUserResp = _reflection.GeneratedProtocolMessageType('SearchUserResp', (_message.Message,), {
  'DESCRIPTOR' : _SEARCHUSERRESP,
  '__module__' : 'business.ext_pb2'
  # @@protoc_insertion_point(class_scope:pb.SearchUserResp)
  })
_sym_db.RegisterMessage(SearchUserResp)



_BUSINESSEXT = _descriptor.ServiceDescriptor(
  name='BusinessExt',
  full_name='pb.BusinessExt',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=535,
  serialized_end=739,
  methods=[
  _descriptor.MethodDescriptor(
    name='SignIn',
    full_name='pb.BusinessExt.SignIn',
    index=0,
    containing_service=None,
    input_type=_SIGNINREQ,
    output_type=_SIGNINRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetUser',
    full_name='pb.BusinessExt.GetUser',
    index=1,
    containing_service=None,
    input_type=_GETUSERREQ,
    output_type=_GETUSERRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateUser',
    full_name='pb.BusinessExt.UpdateUser',
    index=2,
    containing_service=None,
    input_type=_UPDATEUSERREQ,
    output_type=_UPDATEUSERRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SearchUser',
    full_name='pb.BusinessExt.SearchUser',
    index=3,
    containing_service=None,
    input_type=_SEARCHUSERREQ,
    output_type=_SEARCHUSERRESP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BUSINESSEXT)

DESCRIPTOR.services_by_name['BusinessExt'] = _BUSINESSEXT

# @@protoc_insertion_point(module_scope)
