package db

import (
	"context"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_record"
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"time"
)

type CouponRecord struct {
	CouponNo     *int64 `gorm:"primaryKey"`
	CouponMetaNo *int64 `gorm:"uniqueIndex:user_id_meta_no_index,sort:desc"`
	UserId       *int64 `gorm:"uniqueIndex:user_id_meta_no_index,sort:desc"`
	CouponStatus coupon_record.CouponRecordStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (c *CouponRecord) TableName() string {
	return constants.CouponRecordTableName
}

func (c *CouponRecord) BeforeCreate(tx *gorm.DB) (error error) {
	id := idgen.NextId()
	c.CouponNo = &id
	return
}

func IdempotentCheck(ctx context.Context, check *CouponRecord) bool {
	var count int64
	DB.WithContext(ctx).Model(&CouponRecord{}).Where(check).Count(&count)
	return count > 0
}

func AddCouponToUser(ctx context.Context, data *CouponRecord) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(data).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
