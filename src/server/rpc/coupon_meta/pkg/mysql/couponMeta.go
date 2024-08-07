package mysql

import (
	"errors"
	"github.com/Tzz1194593491/coupon_server/kitex_gen/com/tang/coupon_server/coupon_meta"
	"github.com/Tzz1194593491/coupon_server/rpc/common"
	"gorm.io/gorm"
	"log"
	"time"
)

type CouponMeta struct {
	CouponMetaNo     int64
	CouponMetaType   coupon_meta.CouponMetaType
	ValidStartTime   time.Time
	ValidEndTime     time.Time
	CouponMetaStatus coupon_meta.CouponStatus
	CouponMetaStock  int
	CreateTime       time.Time
	UpdateTime       time.Time
	Deleted          gorm.DeletedAt
}

type CouponMetaManager struct {
	db *gorm.DB
}

// NewCouponMetaManager creates a mysql manager.
func NewCouponMetaManager(db *gorm.DB) *CouponMetaManager {
	err := db.AutoMigrate(&CouponMeta{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &CouponMetaManager{
		db: db,
	}
}

func (m *CouponMetaManager) GetCouponMetaByPage(pageInfo *common.PageInfo, couponMeta *CouponMeta) (res []*CouponMeta, err error) {
	offset := (pageInfo.PageNum - 1) * pageInfo.PageSize
	tx := m.db.Where(couponMeta).Limit(pageInfo.PageNum).Offset(offset).Find(&res)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, tx.Error
	}
	return res, nil
}

func (m *CouponMetaManager) GetCouponMetaIsValid(couponMeta *CouponMeta) (res *CouponMeta, err error) {
	tx := m.db.Where("validStartTime < ? AND validEndTime > ?", time.Now(), time.Now()).
		First(&res, couponMeta.CouponMetaNo)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, tx.Error
	}
	return res, nil
}

func (m *CouponMetaManager) AddCouponMeta(couponMeta *CouponMeta) (err error) {
	tx := m.db.Create(couponMeta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *CouponMetaManager) UpdateCouponMeta(couponMeta *CouponMeta) (err error) {
	tx := m.db.Model(&CouponMeta{}).Updates(&couponMeta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *CouponMetaManager) DeleteCouponMeta(couponMeta *CouponMeta) (err error) {
	tx := m.db.Delete(couponMeta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
