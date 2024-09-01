package handlers

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/api/rpc"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func SendCoupon(c context.Context, ctx *app.RequestContext) {
	var req *coupon_record.CouponRecordSendReq
	if err := ctx.BindJSON(&req); err != nil {
		hlog.Error("参数解析出错", err)
		SendResponse(ctx, business_code.BusinessCode_PARAM_PARSE_FAIL, nil)
		return
	}
	code := rpc.SendCoupon(c, req)
	SendResponse(ctx, code, nil)
}
