// package: 
// file: inventory.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Void extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Void.AsObject;
    static toObject(includeInstance: boolean, msg: Void): Void.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Void, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Void;
    static deserializeBinaryFromReader(message: Void, reader: jspb.BinaryReader): Void;
}

export namespace Void {
    export type AsObject = {
    }
}

export class Product extends jspb.Message { 
    getId(): string;
    setId(value: string): Product;
    getPriceInCents(): number;
    setPriceInCents(value: number): Product;
    getTitle(): string;
    setTitle(value: string): Product;
    getDescription(): string;
    setDescription(value: string): Product;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Product.AsObject;
    static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Product;
    static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
    export type AsObject = {
        id: string,
        priceInCents: number,
        title: string,
        description: string,
    }
}

export class Products extends jspb.Message { 
    clearProductsList(): void;
    getProductsList(): Array<Product>;
    setProductsList(value: Array<Product>): Products;
    addProducts(value?: Product, index?: number): Product;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Products.AsObject;
    static toObject(includeInstance: boolean, msg: Products): Products.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Products, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Products;
    static deserializeBinaryFromReader(message: Products, reader: jspb.BinaryReader): Products;
}

export namespace Products {
    export type AsObject = {
        productsList: Array<Product.AsObject>,
    }
}
