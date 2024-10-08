package db

import (
	"context"
	"errors"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"time"
)

type CouponMeta struct {
	CouponMetaNo     *int64 `gorm:"primaryKey"`
	CouponMetaType   *coupon_meta.CouponMetaType
	ValidStartTime   time.Time
	ValidEndTime     time.Time
	CouponMetaStatus *coupon_meta.CouponStatus
	CouponMetaStock  int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (c *CouponMeta) TableName() string {
	return constants.CouponMetaTableName
}

func (c *CouponMeta) BeforeCreate(tx *gorm.DB) (error error) {
	id := idgen.NextId()
	c.CouponMetaNo = &id
	return
}

func GetCouponMetaByPage(ctx context.Context, pageInfo *constants.PageInfo, couponMeta *CouponMeta) (res []*CouponMeta, err error) {
	offset := (pageInfo.PageNum - 1) * pageInfo.PageSize
	tx := DB.WithContext(ctx).Where(couponMeta).Limit(pageInfo.PageSize).Offset(offset).Find(&res)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, tx.Error
	}
	return res, nil
}

func GetCouponMetaById(ctx context.Context, couponMeta *CouponMeta) (res *CouponMeta, err error) {
	tx := DB.WithContext(ctx).Where(couponMeta).First(&res)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, tx.Error
	}
	return res, nil
}

func AddCouponMeta(ctx context.Context, tx *gorm.DB, couponMeta *CouponMeta) (err error) {
	res := tx.WithContext(ctx).Create(couponMeta)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func UpdateCouponMeta(ctx context.Context, couponMeta *CouponMeta) (err error) {
	tx := DB.WithContext(ctx).Model(&CouponMeta{}).
		Where("coupon_meta_no = ?", couponMeta.CouponMetaNo).
		Omit("coupon_meta_no").
		Updates(&couponMeta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteCouponMeta(ctx context.Context, couponMeta *CouponMeta) (err error) {
	tx := DB.WithContext(ctx).Delete(couponMeta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
