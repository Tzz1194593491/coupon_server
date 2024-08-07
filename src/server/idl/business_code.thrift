namespace go com.tang.coupon_server.business_code

enum BusinessCode {
    SUCCESS = 200 // 成功
    SYSTEM_ERROR = 500 // 系统错误
    PARAM_MISSING = 501 // 参数缺失
    REPEAT_ERROR = 502 // 重复数据
    SYSTEM_BUSY = 503 // 系统繁忙
}