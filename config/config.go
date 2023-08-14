package config

type ServerConfig struct {
	AppName   string      `mapstructure:"appName" json:"appName"`
	Port      int         `mapstructure:"port" json:"port"`
	LogInfo   LogConfig   `mapstructure:"log" json:"log"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	IMInfo    ImConfig    `mapstructure:"im" json:"im"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
}

type LogConfig struct {
	LogPath       string `mapstructure:"logPath" json:"logPath"`
	LogFormat     string `mapstructure:"logFormat" json:"logFormat"`
	LogFileExt    string `mapstructure:"logFileExt" json:"logFileExt"`
	LogMaxSize    int    `mapstructure:"logMaxSize" json:"logMaxSize"`
	LogMaxBackups int    `mapstructure:"logMaxBackups" json:"logMaxBackups"`
	LogMaxAge     int    `mapstructure:"logMaxAge" json:"logMaxAge"`
	LogCompress   bool   `mapstructure:"logCompress" json:"logCompress"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ImConfig struct {
	TokenUrl     string `mapstructure:"tokenUrl" json:"tokenUrl"`
	GrantType    string `mapstructure:"grantType" json:"grantType"`
	ClientId     string `mapstructure:"clientId" json:"clientId"`
	ClientSecret string `mapstructure:"clientSecret" json:"clientSecret"`
}

type MysqlConfig struct {
	DbType   string `mapstructure:"dbType" json:"dbType"`
	DbName   string `mapstructure:"dbName" json:"dbName"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"Username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type RedisConfig struct {
	Addr        string `mapstructure:"addr" json:"addr"`
	Port        int    `mapstructure:"port" json:"port"`
	Password    string `mapstructure:"password" json:"password"`
	Db          int    `mapstructure:"db" json:"db"`
	DialTimeout int    `mapstructure:"dialTimeout" json:"dialTimeout"`
}
