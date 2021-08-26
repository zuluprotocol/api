// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var coreapi_v1_coreapi_pb = require('../../coreapi/v1/coreapi_pb.js');
var assets_pb = require('../../assets_pb.js');
var vega_pb = require('../../vega_pb.js');
var events_v1_events_pb = require('../../events/v1/events_pb.js');

function serialize_vega_coreapi_v1_ListAccountsRequest(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListAccountsRequest)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListAccountsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListAccountsRequest(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListAccountsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListAccountsResponse(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListAccountsResponse)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListAccountsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListAccountsResponse(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListAccountsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListAssetsRequest(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListAssetsRequest)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListAssetsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListAssetsRequest(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListAssetsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListAssetsResponse(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListAssetsResponse)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListAssetsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListAssetsResponse(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListAssetsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListNetworkParametersRequest(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListNetworkParametersRequest)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListNetworkParametersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListNetworkParametersRequest(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListNetworkParametersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListNetworkParametersResponse(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListNetworkParametersResponse)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListNetworkParametersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListNetworkParametersResponse(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListNetworkParametersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListPartiesRequest(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListPartiesRequest)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListPartiesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListPartiesRequest(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListPartiesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListPartiesResponse(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListPartiesResponse)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListPartiesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListPartiesResponse(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListPartiesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListValidatorsRequest(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListValidatorsRequest)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListValidatorsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListValidatorsRequest(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListValidatorsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_vega_coreapi_v1_ListValidatorsResponse(arg) {
  if (!(arg instanceof coreapi_v1_coreapi_pb.ListValidatorsResponse)) {
    throw new Error('Expected argument of type vega.coreapi.v1.ListValidatorsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_vega_coreapi_v1_ListValidatorsResponse(buffer_arg) {
  return coreapi_v1_coreapi_pb.ListValidatorsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var CoreApiServiceService = exports.CoreApiServiceService = {
  listAccounts: {
    path: '/vega.coreapi.v1.CoreApiService/ListAccounts',
    requestStream: false,
    responseStream: false,
    requestType: coreapi_v1_coreapi_pb.ListAccountsRequest,
    responseType: coreapi_v1_coreapi_pb.ListAccountsResponse,
    requestSerialize: serialize_vega_coreapi_v1_ListAccountsRequest,
    requestDeserialize: deserialize_vega_coreapi_v1_ListAccountsRequest,
    responseSerialize: serialize_vega_coreapi_v1_ListAccountsResponse,
    responseDeserialize: deserialize_vega_coreapi_v1_ListAccountsResponse,
  },
  listAssets: {
    path: '/vega.coreapi.v1.CoreApiService/ListAssets',
    requestStream: false,
    responseStream: false,
    requestType: coreapi_v1_coreapi_pb.ListAssetsRequest,
    responseType: coreapi_v1_coreapi_pb.ListAssetsResponse,
    requestSerialize: serialize_vega_coreapi_v1_ListAssetsRequest,
    requestDeserialize: deserialize_vega_coreapi_v1_ListAssetsRequest,
    responseSerialize: serialize_vega_coreapi_v1_ListAssetsResponse,
    responseDeserialize: deserialize_vega_coreapi_v1_ListAssetsResponse,
  },
  listNetworkParameters: {
    path: '/vega.coreapi.v1.CoreApiService/ListNetworkParameters',
    requestStream: false,
    responseStream: false,
    requestType: coreapi_v1_coreapi_pb.ListNetworkParametersRequest,
    responseType: coreapi_v1_coreapi_pb.ListNetworkParametersResponse,
    requestSerialize: serialize_vega_coreapi_v1_ListNetworkParametersRequest,
    requestDeserialize: deserialize_vega_coreapi_v1_ListNetworkParametersRequest,
    responseSerialize: serialize_vega_coreapi_v1_ListNetworkParametersResponse,
    responseDeserialize: deserialize_vega_coreapi_v1_ListNetworkParametersResponse,
  },
  listParties: {
    path: '/vega.coreapi.v1.CoreApiService/ListParties',
    requestStream: false,
    responseStream: false,
    requestType: coreapi_v1_coreapi_pb.ListPartiesRequest,
    responseType: coreapi_v1_coreapi_pb.ListPartiesResponse,
    requestSerialize: serialize_vega_coreapi_v1_ListPartiesRequest,
    requestDeserialize: deserialize_vega_coreapi_v1_ListPartiesRequest,
    responseSerialize: serialize_vega_coreapi_v1_ListPartiesResponse,
    responseDeserialize: deserialize_vega_coreapi_v1_ListPartiesResponse,
  },
  listValidators: {
    path: '/vega.coreapi.v1.CoreApiService/ListValidators',
    requestStream: false,
    responseStream: false,
    requestType: coreapi_v1_coreapi_pb.ListValidatorsRequest,
    responseType: coreapi_v1_coreapi_pb.ListValidatorsResponse,
    requestSerialize: serialize_vega_coreapi_v1_ListValidatorsRequest,
    requestDeserialize: deserialize_vega_coreapi_v1_ListValidatorsRequest,
    responseSerialize: serialize_vega_coreapi_v1_ListValidatorsResponse,
    responseDeserialize: deserialize_vega_coreapi_v1_ListValidatorsResponse,
  },
};

exports.CoreApiServiceClient = grpc.makeGenericClientConstructor(CoreApiServiceService);