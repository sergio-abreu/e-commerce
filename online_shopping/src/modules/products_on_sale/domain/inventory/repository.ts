import Product from "./product";

export default interface Repository {
    getAll(): Promise<Array<Product>>;
}