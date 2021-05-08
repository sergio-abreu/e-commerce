// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var inventory_pb = require('./inventory_pb.js');

function serialize_Products(arg) {
  if (!(arg instanceof inventory_pb.Products)) {
    throw new Error('Expected argument of type Products');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Products(buffer_arg) {
  return inventory_pb.Products.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_Void(arg) {
  if (!(arg instanceof inventory_pb.Void)) {
    throw new Error('Expected argument of type Void');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Void(buffer_arg) {
  return inventory_pb.Void.deserializeBinary(new Uint8Array(buffer_arg));
}


var InventoryService = exports.InventoryService = {
  getAllProducts: {
    path: '/Inventory/GetAllProducts',
    requestStream: false,
    responseStream: false,
    requestType: inventory_pb.Void,
    responseType: inventory_pb.Products,
    requestSerialize: serialize_Void,
    requestDeserialize: deserialize_Void,
    responseSerialize: serialize_Products,
    responseDeserialize: deserialize_Products,
  },
};

exports.InventoryClient = grpc.makeGenericClientConstructor(InventoryService);
