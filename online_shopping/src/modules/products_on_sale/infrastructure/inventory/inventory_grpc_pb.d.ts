// package: 
// file: inventory.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as inventory_pb from "./inventory_pb";

interface IInventoryService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getAllProducts: IInventoryService_IGetAllProducts;
}

interface IInventoryService_IGetAllProducts extends grpc.MethodDefinition<inventory_pb.Void, inventory_pb.Products> {
    path: "/Inventory/GetAllProducts";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<inventory_pb.Void>;
    requestDeserialize: grpc.deserialize<inventory_pb.Void>;
    responseSerialize: grpc.serialize<inventory_pb.Products>;
    responseDeserialize: grpc.deserialize<inventory_pb.Products>;
}

export const InventoryService: IInventoryService;

export interface IInventoryServer {
    getAllProducts: grpc.handleUnaryCall<inventory_pb.Void, inventory_pb.Products>;
}

export interface IInventoryClient {
    getAllProducts(request: inventory_pb.Void, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
    getAllProducts(request: inventory_pb.Void, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
    getAllProducts(request: inventory_pb.Void, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
}

export class InventoryClient extends grpc.Client implements IInventoryClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getAllProducts(request: inventory_pb.Void, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
    public getAllProducts(request: inventory_pb.Void, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
    public getAllProducts(request: inventory_pb.Void, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: inventory_pb.Products) => void): grpc.ClientUnaryCall;
}
