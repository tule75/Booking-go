package setting

type Config struct {
	Server      Server       `mapstructure:"server"`
	Mysql       MySQLSetting `mapstructure:"mysql"`
	LogSettings LogSetting   `mapstructure:"log"`
	Redis       RedisSetting `mapstructure:"redis"`
	Kafka       KafkaSetting `mapstructure:"kafka"`
}

type Server struct {
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Addr     string `mapstructure:"addr"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type KafkaSetting struct {
	Addr    string `mapstructure:"addr"`
	Topic   string `mapstructure:"topic"`
	GroupID int    `mapstructure:"groupid"`
}
