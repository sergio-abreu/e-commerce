#!/bin/bash

yarn run grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=grpc_js:./src/modules/products_on_sale/infrastructure/discount \
    --js_out=import_style=commonjs,binary:./src/modules/products_on_sale/infrastructure/discount \
    --grpc_out=grpc_js:./src/modules/products_on_sale/infrastructure/discount \
    -I ../proto \
    discount_calculator.proto

yarn run grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=grpc_js:./src/modules/products_on_sale/infrastructure/product \
    --js_out=import_style=commonjs,binary:./src/modules/products_on_sale/infrastructure/product \
    --grpc_out=grpc_js:./src/modules/products_on_sale/infrastructure/product \
    -I ../proto \
    product.proto
