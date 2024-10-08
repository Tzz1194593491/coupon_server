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
    3: string valid_start_time
    4: string valid_end_time
    5: CouponStatus status
    6: i32 stock
    7: string create_time
    8: string update_time
    9: string delete_time
}

// 查询券模板功能

struct GetCouponMetaReq {
    1: optional i64 coupon_meta_no
    2: optional CouponMetaType type
    3: optional CouponStatus status
    255: required base.BasePageInfo base_info
}

struct GetCouponMetaResp {
    1: list<CouponMeta> coupon_meta
    2: base.BasePageInfo base_info
    255: base.BaseResp baseResp
}

// 增加券模板功能

struct AddCouponMetaReq {
    1: required CouponMetaType type (vt.defined_only = "true")
    2: required string valid_start_time
    3: required string valid_end_time
    4: required i32 stock (vt.gt = "0")
    5: required bool is_sharding
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
    2: required CouponMetaType type (vt.in = "0",vt.in = "1")
    3: required string valid_start_time
    4: required string valid_end_time
    5: required i32 stock (vt.gt = "0")
}

struct UpdateCouponMetaResp {
    255: base.BaseResp baseResp
}

// 获取有效券模板（服务之间调用）

struct GetCouponValidMetaInfoReq {
    1: required i64 coupon_meta_no // 券模板id
}

struct GetCouponValidMetaInfoResp {
    1: map<string,CouponMeta> coupon_meta_map // 废弃
    2: CouponMeta coupon_meta_info
    255: base.BaseResp baseResp
}

// 尝试扣减库存（服务之间调用）
struct TryReduceCouponStockReq {
    1: required i64 coupon_meta_no // 券模板id
}

struct TryReduceCouponStockResp {
    1: bool isSuccess
}

service CouponMetaService {
    GetCouponMetaResp GetCouponMeta(1: GetCouponMetaReq req) // 分页查询券模板
    AddCouponMetaResp AddCouponMeta(1: AddCouponMetaReq req) // 增加券模板
    DeleteCouponMetaResp deleteCouponMeta(1: DeleteCouponMetaReq req) // 删除券模板
    UpdateCouponMetaResp updateCouponMeta(1: UpdateCouponMetaReq req) // 更新券模板
    GetCouponValidMetaInfoResp getCouponValidMetaInfo(1: GetCouponValidMetaInfoReq req) // 获取有效券模板
    TryReduceCouponStockResp TryReduceCouponStock(1: TryReduceCouponStockReq req) // 尝试扣减库存
}