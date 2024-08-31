namespace go com.tang.coupon_server.base

include "business_code.thrift"

struct BaseResp {
    1: string msg
    2: business_code.BusinessCode code
    3: bool is_error
}

struct BasePageInfo {
    1: i32 page_num (vt.ge = "1")
    2: i32 page_size (vt.ge = "1")
}