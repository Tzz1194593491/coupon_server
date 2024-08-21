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
		hlog.Error("参数解析出错", err)
		SendResponse(ctx, business_code.BusinessCode_PARAM_PARSE_FAIL, nil)
		return
	}
	businessCode := rpc.AddCouponMeta(c, addParam)
	SendResponse(ctx, businessCode, nil)
}

func DeleteCouponMeta(c context.Context, ctx *app.RequestContext) {
	var deleteParam *coupon_meta.DeleteCouponMetaReq
	if err := ctx.BindJSON(&deleteParam); err != nil {
		hlog.Error("参数解析出错", err)
		SendResponse(ctx, business_code.BusinessCode_PARAM_PARSE_FAIL, nil)
		return
	}
	businessCode := rpc.DeleteCouponMeta(c, deleteParam)
	SendResponse(ctx, businessCode, nil)
}

func UpdateCouponMeta(c context.Context, ctx *app.RequestContext) {
	var updateParam *coupon_meta.UpdateCouponMetaReq
	if err := ctx.BindJSON(&updateParam); err != nil {
		hlog.Error("参数解析出错", err)
		SendResponse(ctx, business_code.BusinessCode_PARAM_PARSE_FAIL, nil)
		return
	}
	businessCode := rpc.UpdateCouponMeta(c, updateParam)
	SendResponse(ctx, businessCode, nil)
}

func GetCouponMetaByPage(c context.Context, ctx *app.RequestContext) {
	var getParam *coupon_meta.GetCouponMetaReq
	if err := ctx.BindJSON(&getParam); err != nil {
		hlog.Error("参数解析出错", err)
		SendResponse(ctx, business_code.BusinessCode_PARAM_PARSE_FAIL, nil)
		return
	}
	businessCode, data := rpc.GetCouponMeta(c, getParam)
	SendResponse(ctx, businessCode, data)
}
