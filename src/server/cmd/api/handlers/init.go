package handlers

import (
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/business_code"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    business_code.BusinessCode `json:"code"`
	Message string                     `json:"message"`
	Data    interface{}                `json:"data"`
}

// SendResponse pack response
func SendResponse(ctx *app.RequestContext, code business_code.BusinessCode, data interface{}) {
	ctx.JSON(consts.StatusOK, Response{
		Code:    code,
		Message: code.String(),
		Data:    data,
	})
}
