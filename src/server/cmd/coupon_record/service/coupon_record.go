package service

import (
	"context"
	"fmt"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_record/rpc"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"time"
)

type CouponRecordService struct {
	ctx context.Context
}

func NewCouponRecordService(ctx context.Context) *CouponRecordService {
	return &CouponRecordService{
		ctx: ctx,
	}
}

func (c *CouponRecordService) SendCoupon(req *coupon_record.CouponRecordSendReq) (err error) {
	// 获取券模板
	metaResp, err := rpc.GetCouponValidMetaList(c.ctx,
		&coupon_meta.GetCouponValidMetaInfoReq{
			CouponMetaNo: req.GetCouponMetaNo(),
		})
	if err != nil {
		return err
	}
	// 校验券的领取时间
	now := time.Now()
	info := metaResp.CouponMetaInfo
	if info == nil {
		return fmt.Errorf("发券失败")
	}
	startTime, _ := utils.StringToTime(info.ValidStartTime)
	endTime, _ := utils.StringToTime(info.ValidEndTime)
	if startTime.After(now) || endTime.Before(now) {
		return fmt.Errorf("已经过了发券时间")
	}
	// 幂等性判断（用用户id和券模板id判断）- 可用布隆过滤器优化
	check := db.IdempotentCheck(c.ctx, &db.CouponRecord{CouponMetaNo: &req.CouponMetaNo, UserId: &req.UserId})
	if check {
		return fmt.Errorf("用户已经领取过了")
	}
	// 生成订单数据
	order := &db.CouponRecord{
		CouponMetaNo: &req.CouponMetaNo,
		UserId:       &req.UserId,
		CouponStatus: coupon_record.CouponRecordStatus_NOMARL,
	}
	// 库存处理-redis处理-根据券模板信息扣减
	tryIsSuccess := rpc.TryReduceCouponStock(c.ctx, &coupon_meta.TryReduceCouponStockReq{CouponMetaNo: req.CouponMetaNo})
	if !tryIsSuccess {
		return fmt.Errorf("库存不足")
	}
	// 落db
	err = db.AddCouponToUser(c.ctx, order)
	if err != nil {
		return err
	}
	return nil
}
