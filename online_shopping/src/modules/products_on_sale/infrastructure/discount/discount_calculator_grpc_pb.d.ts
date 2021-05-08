// package: 
// file: discount_calculator.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as discount_calculator_pb from "./discount_calculator_pb";

interface IDiscountCalculatorService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    calculateUsersDiscountForProduct: IDiscountCalculatorService_ICalculateUsersDiscountForProduct;
}

interface IDiscountCalculatorService_ICalculateUsersDiscountForProduct extends grpc.MethodDefinition<discount_calculator_pb.RequestForDiscount, discount_calculator_pb.Discount> {
    path: "/DiscountCalculator/CalculateUsersDiscountForProduct";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<discount_calculator_pb.RequestForDiscount>;
    requestDeserialize: grpc.deserialize<discount_calculator_pb.RequestForDiscount>;
    responseSerialize: grpc.serialize<discount_calculator_pb.Discount>;
    responseDeserialize: grpc.deserialize<discount_calculator_pb.Discount>;
}

export const DiscountCalculatorService: IDiscountCalculatorService;

export interface IDiscountCalculatorServer {
    calculateUsersDiscountForProduct: grpc.handleUnaryCall<discount_calculator_pb.RequestForDiscount, discount_calculator_pb.Discount>;
}

export interface IDiscountCalculatorClient {
    calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
}

export class DiscountCalculatorClient extends grpc.Client implements IDiscountCalculatorClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    public calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    public calculateUsersDiscountForProduct(request: discount_calculator_pb.RequestForDiscount, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_calculator_pb.Discount) => void): grpc.ClientUnaryCall;
}
