package constants

const (
	CPURateLimit  float64 = 80.0
	DefaultLimit          = 10
	MysqlUserName         = "root"
	MysqlPassword         = "123456"
	MysqlDataBase         = "promotion_coupon"
	RedisPassword         = "1234"

	CouponMetaTableName   = "promotion_coupon_meta"
	CouponMetaServiceName = "coupon_meta"
	CouponMetaServicePort = "8889"

	CouponRecordTableName   = "promotion_coupon_record"
	CouponRecordServiceName = "coupon_record"
	CouponRecordServicePort = "8989"

	ApiServicePort = "8080"
	ApiServiceName = "api"
)

var (
	IP, _           = GetOutBoundIP()
	MySQLDefaultDSN = MysqlUserName + ":" + MysqlPassword + "@tcp(" + GetIp("MysqlIp") + ":3306)/" + MysqlDataBase + "?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = GetIp("EtcdIp") + ":2379"
	RedisCluster    = []string{IP + ":6379", IP + ":6380", IP + ":6381", IP + ":6382", IP + ":6383", IP + ":6384"}
)
