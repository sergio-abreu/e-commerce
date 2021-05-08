import DiscountGRpcRepository from "../modules/products_on_sale/infrastructure/discount/repository";
import InventoryGRpcRepository from "../modules/products_on_sale/infrastructure/inventory/repository";
import ProductsOnSale from "../modules/products_on_sale/application/products_on_sale";
import Config from "./config";
import express = require("express");
import {Server} from "http";
import getProductsWithDiscount from "./controllers/products-with-discounts";

export default class Application {
    protected webServer: Server
    protected webServerPort: number

    constructor(config: Config) {
        this.webServerPort = config.webServerPort
        const discountRepository = new DiscountGRpcRepository(config.discountCalculatorAddr)
        const inventoryRepository = new InventoryGRpcRepository(config.inventoryAddr)
        const productsOnSale = new ProductsOnSale(inventoryRepository, discountRepository)
        const app = express()
        app.get('/products', getProductsWithDiscount(productsOnSale))
        this.webServer = new Server(app)
    }

    run() {
        this.webServer.listen(this.webServerPort)
    }

    shutdown() {
        this.webServer.close((err) => {
            if (err) {
                console.error(err)
            }
        })
    }
}