package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerInfoConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	ServerInfo ServerInfoConfig `mapstructure:"server_info" json:"server_info"`
	MysqlInfo  MysqlConfig      `mapstructure:"mysql" json:"mysql"`
}
