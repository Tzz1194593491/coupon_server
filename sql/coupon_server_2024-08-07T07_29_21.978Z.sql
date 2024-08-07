CREATE TABLE `promotion_coupon_meta` (
	`coupon_meta_no` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`type` INTEGER,
	`valid_start_time` DATETIME,
	`valid_end_time` DATETIME,
	`status` INTEGER,
	`create_time` DATETIME,
	`update_time` DATETIME,
	`delete_time` DATETIME,
	PRIMARY KEY(`coupon_meta_no`)
);


CREATE TABLE `promotion_coupon_record` (
	`coupon_no` BIGINT NOT NULL AUTO_INCREMENT UNIQUE,
	`coupon_meta_no` BIGINT,
	`user_id` BIGINT,
	`status` INTEGER,
	`create_time` DATETIME,
	`update_time` DATETIME,
	`delete_time` DATETIME,
	PRIMARY KEY(`coupon_no`)
);


CREATE INDEX `promotion_coupon_record_index_0`
ON `promotion_coupon_record` (`coupon_meta_no`, `user_id`);
