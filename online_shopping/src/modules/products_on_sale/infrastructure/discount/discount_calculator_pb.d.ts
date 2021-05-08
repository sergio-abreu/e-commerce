// package: 
// file: discount_calculator.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class RequestForDiscount extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): RequestForDiscount;
    getProductId(): string;
    setProductId(value: string): RequestForDiscount;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestForDiscount.AsObject;
    static toObject(includeInstance: boolean, msg: RequestForDiscount): RequestForDiscount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestForDiscount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestForDiscount;
    static deserializeBinaryFromReader(message: RequestForDiscount, reader: jspb.BinaryReader): RequestForDiscount;
}

export namespace RequestForDiscount {
    export type AsObject = {
        userId: string,
        productId: string,
    }
}

export class Discount extends jspb.Message { 
    getPercentage(): number;
    setPercentage(value: number): Discount;
    getValueInCents(): number;
    setValueInCents(value: number): Discount;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Discount.AsObject;
    static toObject(includeInstance: boolean, msg: Discount): Discount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Discount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Discount;
    static deserializeBinaryFromReader(message: Discount, reader: jspb.BinaryReader): Discount;
}

export namespace Discount {
    export type AsObject = {
        percentage: number,
        valueInCents: number,
    }
}
