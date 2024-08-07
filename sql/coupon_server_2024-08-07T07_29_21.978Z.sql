CREATE TABLE `promotion_coupon_meta`
(
    `coupon_meta_no`   BIGINT NOT NULL AUTO_INCREMENT UNIQUE COMMENT '券模板id',
    `type`             INTEGER COMMENT '券模板类型',
    `valid_start_time` DATETIME COMMENT '券模板过期开始时间',
    `valid_end_time`   DATETIME COMMENT '券模板过期结束时间',
    `status`           INTEGER COMMENT '券模板状态',
    `stock`            INTEGER COMMENT '券的库存',
    `create_time`      DATETIME COMMENT '创建时间',
    `update_time`      DATETIME COMMENT '更新时间',
    `delete_time`      DATETIME COMMENT '删除时间',
    PRIMARY KEY (`coupon_meta_no`)
);


CREATE TABLE `promotion_coupon_record`
(
    `coupon_no`      BIGINT NOT NULL AUTO_INCREMENT UNIQUE COMMENT '券id',
    `coupon_meta_no` BIGINT COMMENT '券模板id',
    `user_id`        BIGINT COMMENT '用户id',
    `status`         INTEGER COMMENT '券状态',
    `create_time`    DATETIME COMMENT '创建时间',
    `update_time`    DATETIME COMMENT '更新时间',
    `delete_time`    DATETIME COMMENT '删除时间',
    PRIMARY KEY (`coupon_no`)
);


CREATE INDEX `promotion_coupon_record_index_0`
    ON `promotion_coupon_record` (`coupon_meta_no`, `user_id`);
