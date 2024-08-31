package service

import (
	"context"
	"fmt"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/redis"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/pack"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type ManageCouponMeta struct {
	ctx context.Context
}

func NewManageCouponMeta(ctx context.Context) *ManageCouponMeta {
	return &ManageCouponMeta{
		ctx: ctx,
	}
}

func (m *ManageCouponMeta) AddManageCouponMeta(req *coupon_meta.AddCouponMetaReq) (err error) {
	validStartTime, err := utils.StringToTime(req.ValidStartTime)
	validEndTime, err := utils.StringToTime(req.ValidEndTime)
	if validStartTime.After(validEndTime) {
		return fmt.Errorf("validStartTime must be before validEndTime")
	}
	dbCouponMeta := &db.CouponMeta{
		CouponMetaType:  &req.Type,
		ValidStartTime:  validStartTime,
		ValidEndTime:    validEndTime,
		CouponMetaStock: req.Stock,
		CouponMetaStatus: coupon_meta.
			CouponStatusPtr(coupon_meta.CouponStatus_NOT_EXPIRED),
	}
	// 落db
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err := db.AddCouponMeta(m.ctx, dbCouponMeta)
		if err != nil {
			return err
		}
		// 分片逻辑
		redisCouponMeta := redis.WithCouponMeta(dbCouponMeta)
		if req.IsSharding {
			err = redisCouponMeta.ShardingToRedis(m.ctx)
		} else {
			err = redisCouponMeta.SingleToRedis(m.ctx)
		}
		if err != nil {
			klog.Error(err)
			return err
		}
		// 发送mq，进入过期状态机判断
		return nil
	})

	return err
}

func (m *ManageCouponMeta) DeleteManageCouponMeta(req *coupon_meta.DeleteCouponMetaReq) (err error) {
	couponMeta := &db.CouponMeta{CouponMetaNo: &req.CouponMetaNo}
	err = m.checkUpdateValid(couponMeta, err)
	if err != nil {
		return err
	}
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return db.DeleteCouponMeta(m.ctx, couponMeta)
	})
	return err
}

func (m *ManageCouponMeta) UpdateManageCouponMeta(req *coupon_meta.UpdateCouponMetaReq) (err error) {
	err = m.checkUpdateValid(&db.CouponMeta{CouponMetaNo: &req.CouponMetaNo}, err)
	if err != nil {
		return err
	}
	validStartTime, err := utils.StringToTime(req.ValidStartTime)
	validEndTime, err := utils.StringToTime(req.ValidEndTime)
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return db.UpdateCouponMeta(m.ctx, &db.CouponMeta{
			CouponMetaNo:    &req.CouponMetaNo,
			CouponMetaType:  &req.Type,
			ValidStartTime:  validStartTime,
			ValidEndTime:    validEndTime,
			CouponMetaStock: req.Stock,
		})
	})
	return err
}

func (m *ManageCouponMeta) checkUpdateValid(couponMeta *db.CouponMeta, err error) error {
	metaById, err := db.GetCouponMetaById(m.ctx, couponMeta)
	if err != nil {
		return err
	}
	now := time.Now()
	if metaById.ValidStartTime.Before(now) {
		return fmt.Errorf("已经开始发放，不能操作")
	}
	return nil
}

func (m *ManageCouponMeta) GetManageCouponMetaByPage(req *coupon_meta.GetCouponMetaReq) (res []*coupon_meta.CouponMeta, err error) {
	pageInfo := &constants.PageInfo{
		PageSize: int(req.BaseInfo.PageSize),
		PageNum:  int(req.BaseInfo.PageNum),
	}
	byPage, err := db.GetCouponMetaByPage(m.ctx, pageInfo, &db.CouponMeta{
		CouponMetaNo:     req.CouponMetaNo,
		CouponMetaType:   req.Type,
		CouponMetaStatus: req.Status,
	})
	if err != nil {
		return nil, err
	}
	res = pack.CouponMetas(byPage)
	return res, nil
}
