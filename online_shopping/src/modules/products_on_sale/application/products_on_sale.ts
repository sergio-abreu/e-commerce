import inventoryRepository from "../domain/inventory/repository";
import discountRepository from "../domain/discount/repository";
import ProductWithUserDiscounts from "./products_with_user_discount";
import Discount from "../domain/discount/discount";

export default class ProductsOnSale {
    private inventoryRepository: inventoryRepository;
    private discountRepository: discountRepository;

    constructor(inventoryRepository: inventoryRepository, discountRepository: discountRepository) {
        this.inventoryRepository = inventoryRepository;
        this.discountRepository = discountRepository;
    }

    getAllProductsWithUserDiscounts(userId: string): Promise<Array<ProductWithUserDiscounts>> {
        return new Promise<Array<ProductWithUserDiscounts>>((resolve, reject) => {
            this.inventoryRepository.getAll().then(async products => {
                let productsWithUserDiscounts: Array<ProductWithUserDiscounts> = []
                for (const product of products) {
                    let productWithUserDiscount: ProductWithUserDiscounts = {
                        id: product.id,
                        priceInCents: product.priceInCents,
                        title: product.title,
                        description: product.description,
                        discount: {
                            percentage: 0,
                            valueInCents: 0,
                        }
                    };
                    await this.discountRepository.getUserDiscountsForProduct(userId, product.id).then((discount: Discount) => {
                        productWithUserDiscount.priceInCents = discount.valueInCents > 0 ? discount.valueInCents : product.priceInCents;
                        productWithUserDiscount.discount.percentage = discount.percentage;
                        productWithUserDiscount.discount.valueInCents = discount.valueInCents > 0 ? product.priceInCents - discount.valueInCents : discount.valueInCents;
                    }).catch(err => console.warn(err));
                    productsWithUserDiscounts = productsWithUserDiscounts.concat(productWithUserDiscount);
                }
                return resolve(productsWithUserDiscounts)
            }).catch(err => reject(err))
        })
    }
}

