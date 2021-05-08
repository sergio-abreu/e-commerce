// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var discount_calculator_pb = require('./discount_calculator_pb.js');

function serialize_Discount(arg) {
  if (!(arg instanceof discount_calculator_pb.Discount)) {
    throw new Error('Expected argument of type Discount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Discount(buffer_arg) {
  return discount_calculator_pb.Discount.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RequestForDiscount(arg) {
  if (!(arg instanceof discount_calculator_pb.RequestForDiscount)) {
    throw new Error('Expected argument of type RequestForDiscount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RequestForDiscount(buffer_arg) {
  return discount_calculator_pb.RequestForDiscount.deserializeBinary(new Uint8Array(buffer_arg));
}


var DiscountCalculatorService = exports.DiscountCalculatorService = {
  calculateUsersDiscountForProduct: {
    path: '/DiscountCalculator/CalculateUsersDiscountForProduct',
    requestStream: false,
    responseStream: false,
    requestType: discount_calculator_pb.RequestForDiscount,
    responseType: discount_calculator_pb.Discount,
    requestSerialize: serialize_RequestForDiscount,
    requestDeserialize: deserialize_RequestForDiscount,
    responseSerialize: serialize_Discount,
    responseDeserialize: deserialize_Discount,
  },
};

exports.DiscountCalculatorClient = grpc.makeGenericClientConstructor(DiscountCalculatorService);
