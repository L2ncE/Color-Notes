package model

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	APPID       string        `mapstructure:"appID"`
	APPSecret   string        `mapstructure:"appSecret"`
	GormInfo    GormConfig    `mapstructure:"gorm"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	MongoDBInfo MongoDBConfig `mapstructure:"mongodb"`
}

type GormConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type MongoDBConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DBName string `mapstructure:"dbName"`
}
