import Repository from "../../domain/discount/repository";
import Discount from "../../domain/discount/discount";
import {credentials, Metadata} from "grpc";
import {DiscountCalculatorClient} from "./discount_calculator_grpc_pb";
import {RequestForDiscount} from "./discount_calculator_pb";

export default class GRpcRepository implements Repository {
    client: DiscountCalculatorClient;

    constructor(address: string) {
        this.client = new DiscountCalculatorClient(address, credentials.createInsecure())
    }

    getUserDiscountsForProduct(user_id: string, product_id: string): Promise<Discount> {
        return new Promise<Discount>((resolve, reject) => {
            const requestForDiscount = new RequestForDiscount();
            requestForDiscount.setUserId(user_id)
            requestForDiscount.setProductId(product_id)
            this.client.calculateUsersDiscountForProduct(requestForDiscount, new Metadata(), {}, (err, discount) => {
                if (err) {
                    return reject(err)
                }
                return resolve(discount.toObject())
            })
        })
    }
}

