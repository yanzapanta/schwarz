CREATE DATABASE IF NOT EXISTS `coupons`;

DROP TABLE IF EXISTS `coupon`;

CREATE TABLE `coupon` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `discount` int NOT NULL,
  `min_basket_value` int NOT NULL,
  `valid_from` datetime DEFAULT NULL,
  `valid_to` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` int DEFAULT '1',
  `updated_at` datetime DEFAULT NULL,
  `is_active` int DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
