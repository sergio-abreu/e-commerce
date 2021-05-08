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

    getUserDiscountsForProduct(userId: string, productId: string): Promise<Discount> {
        return new Promise<Discount>((resolve, reject) => {
            if (userId == undefined || userId == "" || productId == undefined || productId == "") {
                return resolve({valueInCents: 0, percentage: 0})
            }
            const requestForDiscount = new RequestForDiscount();
            requestForDiscount.setUserId(userId)
            requestForDiscount.setProductId(productId)
            this.client.calculateUsersDiscountForProduct(requestForDiscount, new Metadata(), {}, (err, discount) => {
                if (err) {
                    return reject(err)
                }
                return resolve(discount.toObject())
            })
        })
    }
}

