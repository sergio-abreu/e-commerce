import Discount from "./discount";

export default  interface Repository {
    getUserDiscountsForProduct(user_id: string, product_id: string): Promise<Discount>
}