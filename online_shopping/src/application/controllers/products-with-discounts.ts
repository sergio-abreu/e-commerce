import {Request, Response} from "express";
import ProductsOnSale from "../../modules/products_on_sale/application/products_on_sale";

export default function getProductsWithDiscount(productsOnSale: ProductsOnSale): (request: Request, response: Response) => void {
    return (request: Request, response: Response) => {
        let userId: string
        if ("userId" in request.query) {
            userId = request.query["userId"].toString()
        }
        productsOnSale.getAllProductsWithUserDiscounts(userId).then(products => {
            response.status(200).json(products).end()
        }).catch(err => {
            console.error(err)
            response.status(500).json({"message": "Internal Server Error"}).end()
        })
    }
}