package db

import (
	"github.com/Tzz1194593491/coupon_server/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	log.Println(constants.MySQLDefaultDSN)
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&CouponMeta{})
	if err != nil {
		klog.Fatal(err)
		return
	}
}
