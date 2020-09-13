# AngelProj

# API Spec

## /addToCart

__Method:__ *POST*

__Content Type:__ application/json

__Request Payload__
```javascript
{
  "clientID": "string",
  "productID": "string",
  "quantity": "integer"
}
```

__Response__
```javascript
{
    "status": "string",
    "error": "string"
}
```

## /checkOutCart

__Method:__ *POST*

__Content Type:__ application/json

__Request Payload__
```javascript
{
  "clientID": "string"
}
```

__Response__
```javascript
{
    "status": "string",
    "error": "string"
}
```

# DB Tables

## client_cart
```
+------------+--------------+------+-----+---------+-------+
| Field      | Type         | Null | Key | Default | Extra |
+------------+--------------+------+-----+---------+-------+
| client_id  | varchar(100) | YES  |     | NULL    |       |
| product_id | varchar(100) | YES  |     | NULL    |       |
| quantity   | int          | YES  |     | NULL    |       |
+------------+--------------+------+-----+---------+-------+

```

## product_details
```
+--------------+--------------+------+-----+---------+-------+
| Field        | Type         | Null | Key | Default | Extra |
+--------------+--------------+------+-----+---------+-------+
| product_id   | varchar(100) | YES  |     | NULL    |       |
| product_name | varchar(100) | YES  |     | NULL    |       |
| quantity     | int          | YES  |     | NULL    |       |
+--------------+--------------+------+-----+---------+-------+
```

## reserve_purchase
```
+------------+--------------+------+-----+---------+-------+
| Field      | Type         | Null | Key | Default | Extra |
+------------+--------------+------+-----+---------+-------+
| reserve_id | varchar(100) | YES  |     | NULL    |       |
| product_id | varchar(100) | YES  |     | NULL    |       |
| quantity   | int          | YES  |     | NULL    |       |
+------------+--------------+------+-----+---------+-------+
```

# MicroServices
The project has 4 micro-services:

## API GateWay:
Developed in NodeJS. It is exposed at port __5003__ and has two endpoints:

    /addToCart
    /checkOutCart

The  swagger documentation has the specification. Accessed using http://localhost:5003/documentation (need to start up the containers first)

## Cart Manager: 
Developed in Golang. It is a worker solely responsible to talk to CartDB.

The worker gets job from RMQ channel and performas follwoing categories of task:
    
* Add Item To Cart: Adds an item to cart and responds to the API with the status

* Checkout Cart: Checks out all the items in a cart and moves to payment gateway

## Inventory Manager:
Developed in Golang and exposed as a gRPC service. It is solely responsible to talk to InventoryDB.

It performs a series of important operations:

* Checks Item Quantity in Inventory

* Reserves Items for Check out

* Commits Reserved Items after successful Checkout: This operation depicts the following taks in order - 
    
    1. Update the Inventory Table to reduce the available quantity
    1. Remove the reserved token from the reservation table

* Rollback Reserved Items after unsuccessfuk CheckOut

## CheckOut Manager:
Developed in Golang. It is a worker which gets the job pushed by __*only Cart manager*__.

On getting a job to check out the list of items,it performs the following taks:
1. Asks Invenory Manager to Reseve the items
1. Starts payment processing
1. If successful asks Inventory Manager to commit the reservations
1. If failed asks Inventory Manager to rollback the reservations


# Components

1. Microservices - *Mentoioned above*
2. RabbitMQ - Queuing mechanisms for the workers so that they can scale easily
3. MySQL - Backend DB to hold the data
      
![architecture](https://user-images.githubusercontent.com/7055118/93028369-2d2d5f00-f631-11ea-9a1b-fbb86b85ec10.jpg)

# Check Out Sequence

![sequence_diagram](https://user-images.githubusercontent.com/7055118/93028390-50f0a500-f631-11ea-8458-5e1cd75cb9e1.jpg)

# Setting Up

1. Install Docker and Docker-Compose in the system
1. Make sure you have bash shell 
1. Run setup.sh
1. Certain workers can still fail after starting. Need to get them separately. Pitfalls discussed below.

# What's missing and concerns

* Limited testing done. So bugs are expected.
* No Unit tests available because of limited time.
* Docker containers need to be set properly with some time out to get RMQ up first.
* No Redis setup to improve the performance of the system using caching.
* No Load balancer added to scale up the API or the gRPC service.
* RMQ channels can be better named and typed.
