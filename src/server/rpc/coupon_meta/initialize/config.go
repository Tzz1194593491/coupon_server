package initialize

import (
	"github.com/Tzz1194593491/coupon_server/rpc/common"
	"github.com/Tzz1194593491/coupon_server/rpc/coupon_meta/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

// todo 在此处做etcd配置拉取
func initConfig() {
	v := viper.New()
	v.SetConfigFile(common.CouponMetaConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalServerConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}
}
