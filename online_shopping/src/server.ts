import Application from "./application/application";

require('dotenv').config()
const config = {
    inventoryAddr: process.env.INVENTORY_ADDR,
    discountCalculatorAddr: process.env.DISCOUNT_CALCULATOR_ADDR,
    webServerPort: parseInt(process.env.WEB_SERVER_PORT)
};
const app = new Application(config)
app.run()
process.on('SIGTERM', () => {
    console.info('gracefully shutdown');
    app.shutdown();
});
