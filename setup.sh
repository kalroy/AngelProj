#!/bin/bash

# building up APIGateway
cd ./apiGateway
docker build -t apigateway .
cd ..

# building cartManager
cd ./cartManager
docker build -t cartmanager .
cd ..

# building checkManager
cd ./checkOutManager
docker build -t checkoutmanager .
cd ..

cd ./inventoryManager
docker build -t inventorymanager .
cd ..

docker-compose up -d
