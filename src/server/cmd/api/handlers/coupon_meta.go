package handlers

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/api/rpc"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func AddCouponMeta(c context.Context, ctx *app.RequestContext) {
	var addParam *coupon_meta.AddCouponMetaReq
	if err := ctx.BindJSON(&addParam); err != nil {
		hlog.Error("参数解析出错")
		SendResponse(ctx, business_code.BusinessCode_ADD_FAIL, nil)
		return
	}
	businessCode := rpc.AddCouponMeta(c, addParam)
	SendResponse(ctx, businessCode, nil)
}
