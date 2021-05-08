import Repository from "../../domain/inventory/repository";
import {credentials, Metadata} from "grpc";
import {InventoryClient} from "./inventory_grpc_pb";
import {Void} from "./inventory_pb";
import Product from "../../domain/inventory/product";

export default class GRpcRepository implements Repository {
    client: InventoryClient;

    constructor(address: string) {
        this.client = new InventoryClient(address, credentials.createInsecure())
    }

    getAll(): Promise<Array<Product>> {
        return new Promise<Array<Product>>((resolve, reject) => {
            this.client.getAllProducts(new Void(), new Metadata(), {}, (err, products) => {
                if (err) {
                    return reject(err)
                }
                return resolve(products.toObject().productsList)
            })
        })
    }
}

