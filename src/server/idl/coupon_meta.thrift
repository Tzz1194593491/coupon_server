namespace go com.tang.coupon_server.coupon_meta

include "base.thrift"

enum CouponMetaType {
    E_COMMERCE = 0 // 电商
    TO_SHOP = 1 // 到店
}

enum CouponStatus {
    EXPIRED = 0 // 已过期
    NOT_EXPIRED = 1 // 未过期
}

struct CouponMeta {
    1: i64 coupon_meta_no
    2: CouponMetaType type
    3: i64 valid_start_time
    4: i64 valid_end_time
    5: CouponStatus status
    6: i32 stock
    7: i64 create_time
    8: i64 update_time
    9: i64 delete_time
}

// 查询券模板功能

struct GetCouponMetaReq {
    1: required i64 coupon_meta_no
    2: optional CouponMetaType type
    3: optional CouponStatus status
    255: required base.BasePageInfo base_info
}

struct GetCouponMetaResp {
    1: list<CouponMeta> coupon_meta
    255: base.BaseResp baseResp
}

// 增加券模板功能

struct AddCouponMetaReq {
    1: required CouponMetaType type
    2: required i64 valid_start_time
    3: required i64 valid_end_time
    4: required i32 stock
}

struct AddCouponMetaResp {
    255: base.BaseResp baseResp
}

// 删除券模板功能

struct DeleteCouponMetaReq {
    1: required i64 coupon_meta_no // 券模板id
}

struct DeleteCouponMetaResp {
    255: base.BaseResp baseResp
}

// 修改券模板功能

struct UpdateCouponMetaReq {
    1: required i64 coupon_meta_no // 券模板id
    2: required CouponMetaType type
    3: required i64 valid_start_time
    4: required i64 valid_end_time
    5: required i32 stock
}

struct UpdateCouponMetaResp {
    255: base.BaseResp baseResp
}

// 获取券模板有效期

struct GetCouponMetaIsValidReq {
    1: required i64 coupon_meta_no // 券模板id
}

struct GetCouponMetaIsValidResp {
    1: bool is_valid // 是否有效
}

// 获取券模板库存

struct GetCouponMetaStockReq {
    1: required i64 coupon_meta_no // 券模板id
}

struct GetCouponMetaStockResp {
    1: i64 coupon_meta_no // 券模板id
    2: i32 stock // 库存信息
}

service CouponMetaService {
    GetCouponMetaResp GetCouponMeta(1: GetCouponMetaReq req) // 分页查询券模板
    AddCouponMetaResp AddCouponMeta(1: AddCouponMetaReq req) // 增加券模板
    DeleteCouponMetaResp deleteCouponMeta(1: DeleteCouponMetaReq req) // 删除券模板
    UpdateCouponMetaResp updateCouponMeta(1: UpdateCouponMetaReq req) // 更新券模板
    GetCouponMetaIsValidReq getCouponMetaIsValid(1: GetCouponMetaIsValidResp res) // 判断券模板是否过期
    GetCouponMetaStockResp getCouponMetaStock(1: GetCouponMetaStockReq res) // 获取券模板库存
}