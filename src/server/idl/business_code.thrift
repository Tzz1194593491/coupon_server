namespace go com.tang.coupon_server.business_code

enum BusinessCode {
    SUCCESS = 200 // 成功
    SYSTEM_ERROR = 500 // 系统错误
    PARAM_MISSING = 501 // 参数缺失
    REPEAT_ERROR = 502 // 重复数据
    SYSTEM_BUSY = 503 // 系统繁忙
    ADD_FAIL= 504 // 添加失败
    DELETE_FAIL = 505 // 删除失败
    GET_FAIL = 506 // 获取失败
    UPDATE_FAIL = 507 // 更新失败
    PARAM_PARSE_FAIL = 508 // 参数解析失败
}