package constants

const (
	CouponMetaTableName           = "promotion_coupon_meta"
	CouponMetaServiceName         = "coupon_meta"
	CPURateLimit          float64 = 80.0
	DefaultLimit                  = 10
	MysqlUserName                 = "root"
	MysqlPassword                 = "123456"
	MysqlDataBase                 = "promotion_coupon"
)

var (
	MySQLDefaultDSN = MysqlUserName + ":" + MysqlPassword + "@tcp(" + GetIp("MysqlIp") + ":3306)/" + MysqlDataBase + "?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = GetIp("EtcdIp") + ":2379"
)
