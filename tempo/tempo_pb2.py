# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tempo.proto
# Protobuf Python Version: 4.25.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0btempo.proto\x12\ncbiomodels\"\x8c\x01\n\x05\x45vent\x12\x12\n\nchromosome\x18\x01 \x01(\t\x12\x15\n\rstartPosition\x18\x02 \x01(\x05\x12\x13\n\x0b\x65ndPosition\x18\x03 \x01(\x05\x12\x11\n\trefAllele\x18\x04 \x01(\t\x12\x17\n\x0ftumorSeqAllele1\x18\x05 \x01(\t\x12\x17\n\x0ftumorSeqAllele2\x18\x06 \x01(\t\"z\n\x0cTempoMessage\x12\x13\n\x0b\x63moSampleId\x18\x01 \x01(\t\x12\x19\n\x11normalCmoSampleId\x18\x02 \x01(\t\x12\x17\n\x0fpipelineVersion\x18\x03 \x01(\t\x12!\n\x06\x65vents\x18\x04 \x03(\x0b\x32\x11.cbiomodels.Eventb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tempo_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_EVENT']._serialized_start=28
  _globals['_EVENT']._serialized_end=168
  _globals['_TEMPOMESSAGE']._serialized_start=170
  _globals['_TEMPOMESSAGE']._serialized_end=292
# @@protoc_insertion_point(module_scope)