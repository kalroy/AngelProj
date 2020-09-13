CREATE TABLE client_cart (
client_id varchar(100),
product_id  varchar(100),
quantity  int(4)
);


CREATE TABLE product_details (
product_id varchar(100),
product_name  varchar(100),
quantity  int(4)
);

CREATE TABLE reserve_purchase (
reserve_id varchar(100),
product_id  varchar(100),
quantity  int(4)
);
