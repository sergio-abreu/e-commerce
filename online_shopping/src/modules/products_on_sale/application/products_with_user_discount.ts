import Product from "../domain/inventory/product";
import Discount from "../domain/discount/discount";

export default interface ProductWithUserDiscounts extends Product {
    discount: Discount
}