package service

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/dal/db"
	"github.com/Tzz1194593491/coupon_server/cmd/coupon_meta/pack"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/Tzz1194593491/coupon_server/pkg/utils"
	"gorm.io/gorm"
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
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return db.AddCouponMeta(m.ctx, &db.CouponMeta{
			CouponMetaType:  req.Type,
			ValidStartTime:  utils.UnixSecondToTime(req.ValidStartTime),
			ValidEndTime:    utils.UnixSecondToTime(req.ValidEndTime),
			CouponMetaStock: req.Stock,
		})
	})
	return err
}

func (m *ManageCouponMeta) DeleteManageCouponMeta(req *coupon_meta.DeleteCouponMetaReq) (err error) {
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return db.DeleteCouponMeta(m.ctx, &db.CouponMeta{
			CouponMetaNo: req.CouponMetaNo,
		})
	})
	return err
}

func (m *ManageCouponMeta) UpdateManageCouponMeta(req *coupon_meta.UpdateCouponMetaReq) (err error) {
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		return db.DeleteCouponMeta(m.ctx, &db.CouponMeta{
			CouponMetaNo:    req.CouponMetaNo,
			CouponMetaType:  req.Type,
			ValidStartTime:  utils.UnixSecondToTime(req.ValidStartTime),
			ValidEndTime:    utils.UnixSecondToTime(req.ValidEndTime),
			CouponMetaStock: req.Stock,
		})
	})
	return err
}

func (m *ManageCouponMeta) GetManageCouponMetaByPage(req *coupon_meta.GetCouponMetaReq) (res []*coupon_meta.CouponMeta, err error) {
	pageInfo := &constants.PageInfo{
		PageSize: int(req.BaseInfo.PageSize),
		PageNum:  int(req.BaseInfo.PageNum),
	}
	byPage, err := db.GetCouponMetaByPage(m.ctx, pageInfo, &db.CouponMeta{
		CouponMetaNo:     req.CouponMetaNo,
		CouponMetaType:   *req.Type,
		CouponMetaStatus: *req.Status,
	})
	if err != nil {
		return nil, err
	}
	res = pack.CouponMetas(byPage)
	return res, nil
}
