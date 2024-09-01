namespace go com.tang.coupon_server.coupon_record

include "base.thrift"

enum CouponRecordStatus {
    NOMARL = 0 // 正常
    EXPIRED = 1 // 已过期
    USERD = 100 // 已使用
}

struct CouponRecord {
    1: i64 coupon_no
    2: i64 coupon_meta_no
    3: i64 user_id
    4: CouponRecordStatus coupon_status
}


struct CouponRecordSendReq {
    1: i64 coupon_meta_no
    2: i64 user_id
}

struct CouponRecordSendResp {
    1: i64 coupon_meta_no
    2: i64 user_id
    255: base.BaseResp base_resp
}

struct CouponRecordListReq {
    1: i64 user_id
}

struct CouponRecordListResp {
    1: i64 user_id
    2: list<CouponRecord> records
    3: base.BasePageInfo base_info
    255: base.BaseResp base_resp
}

struct CouponRecordUsedReq {
    1: i64 coupon_meta_no
    2: i64 user_id
}

struct CouponRecordUsedResp {
    1: i64 coupon_meta_no
    2: i64 user_id
    255: base.BaseResp base_resp
}

service CouponRecordService {
    CouponRecordSendResp SendCoupon(1: CouponRecordSendReq req) // 给指定用户发券
    CouponRecordListResp GetCouponRecordList(1: CouponRecordListReq req) // 给指定用户发券
    CouponRecordUsedResp UseCoupon(1: CouponRecordUsedReq req) // 给指定用户发券
}