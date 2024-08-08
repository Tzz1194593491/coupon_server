package pack

import (
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/base"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
)

func Success(code business_code.BusinessCode) *base.BaseResp {
	return baseResp(code, false)
}

func Fail(code business_code.BusinessCode) *base.BaseResp {
	return baseResp(code, true)
}

func baseResp(code business_code.BusinessCode, isError bool) *base.BaseResp {
	msg := code.String()
	return &base.BaseResp{
		Msg:     msg,
		Code:    code,
		IsError: isError,
	}
}
