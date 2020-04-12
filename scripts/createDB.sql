CREATE DATABASE `DemoApp` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
CREATE TABLE `favorites` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `item_id` int unsigned DEFAULT NULL,
  `user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=407 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `us_item_id` varchar(16) NOT NULL,
  `offer_id` varchar(64) DEFAULT NULL,
  `sku` varchar(16) DEFAULT NULL,
  `sales_unit` varchar(32) DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `name_lc` varchar(256) DEFAULT NULL,
  `thumbnail` varchar(256) DEFAULT NULL,
  `weight_increment` double DEFAULT NULL,
  `average_weight` double DEFAULT NULL,
  `max_allowed` int DEFAULT NULL,
  `product_url` varchar(192) DEFAULT NULL,
  `is_snap_eligible` int DEFAULT NULL,
  `type` varchar(16) DEFAULT NULL,
  `rating` double DEFAULT NULL,
  `reviews_count` int DEFAULT NULL,
  `is_out_of_stock` varchar(8) DEFAULT NULL,
  `list` double DEFAULT NULL,
  `previous_price` double DEFAULT NULL,
  `price_unit_of_measure` varchar(32) DEFAULT NULL,
  `sales_unit_of_measure` varchar(32) DEFAULT NULL,
  `sales_quantity` int DEFAULT NULL,
  `display_condition` varchar(32) DEFAULT NULL,
  `display_price` double DEFAULT NULL,
  `display_unit_price` varchar(32) DEFAULT NULL,
  `is_clearance` varchar(16) DEFAULT NULL,
  `is_rollback` varchar(16) DEFAULT NULL,
  `unit` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `USItemId_UNIQUE` (`us_item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4601 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `items_bak` (
  `id` int NOT NULL AUTO_INCREMENT,
  `us_item_id` varchar(16) NOT NULL,
  `offer_id` varchar(64) DEFAULT NULL,
  `sku` varchar(16) DEFAULT NULL,
  `sales_unit` varchar(32) DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `name_lc` varchar(256) DEFAULT NULL,
  `thumbnail` varchar(256) DEFAULT NULL,
  `weight_increment` double DEFAULT NULL,
  `average_weight` double DEFAULT NULL,
  `max_allowed` int DEFAULT NULL,
  `product_url` varchar(192) DEFAULT NULL,
  `is_snap_eligible` int DEFAULT NULL,
  `type` varchar(16) DEFAULT NULL,
  `rating` double DEFAULT NULL,
  `reviews_count` int DEFAULT NULL,
  `is_out_of_stock` varchar(8) DEFAULT NULL,
  `list` double DEFAULT NULL,
  `previous_price` double DEFAULT NULL,
  `price_unit_of_measure` varchar(32) DEFAULT NULL,
  `sales_unit_of_measure` varchar(32) DEFAULT NULL,
  `sales_quantity` int DEFAULT NULL,
  `display_condition` varchar(32) DEFAULT NULL,
  `display_price` double DEFAULT NULL,
  `display_unit_price` varchar(32) DEFAULT NULL,
  `is_clearance` varchar(16) DEFAULT NULL,
  `is_rollback` varchar(16) DEFAULT NULL,
  `unit` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `USItemId_UNIQUE` (`us_item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `order_items` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int unsigned DEFAULT NULL,
  `item_id` int unsigned DEFAULT NULL,
  `quantity` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=448 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `orders` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
