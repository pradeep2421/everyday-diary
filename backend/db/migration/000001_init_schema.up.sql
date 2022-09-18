
CREATE TABLE `product_tags` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `simple_table` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `merchant_periods` (
  `id` int PRIMARY KEY, 
  `merchant_id` int,
  `country_code` int,
  `start_date` datetime,
  `end_date` datetime
);





