CREATE TABLE `client_cart` (
  `client_id` varchar(100) NOT NULL,
  `product_id` varchar(100) NOT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`client_id`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `product_details`;
CREATE TABLE `product_details` (
  `product_id` varchar(100) NOT NULL,
  `product_name` varchar(100) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `reserve_purchase` (
  `reserve_id` varchar(100) NOT NULL,
  `product_id` varchar(100) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`reserve_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
