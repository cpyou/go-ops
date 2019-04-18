package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	MySQL      MySQLConfig
	PgSQL      PgSQLConfig
	Prometheus PrometheusConfig
	Ops        OpsConfig
	MongoDB    MongoConfig
	Cache      CacheConfig
	Jwt        JwtConfig
	Cookie     CookieConfig
	Common     CommonLog
	Kafka      KafkaConfig
}

type MySQLConfig struct {
	HOST     string
	Username string
	Password string
	Port     string
	Dbname   string
	Dbtype   string
	Prefix   string
}

type PgSQLConfig struct {
	HOST     string
	Username string
	Password string
	Port     string
	Dbname   string
	Dbtype   string
}

// prometheus
type PrometheusConfig struct {
	Url string
}

// ops
type OpsConfig struct {
	Url      string
	Token    string
	AlarmUrl string
}

// mongodb
type MongoConfig struct {
	Uri     string
}

type CacheConfig struct {
	Host        string
	Password    string
	Dbname      int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var config Config

func GetConfig() *Config {
	return &config
}

type CommonLog struct {
	DEBUG           bool
	HTTP_PORT       int
	APP_SECRET      string
	ACCESS_LOG      bool
	ACCESS_LOG_PATH string
	ERROR_LOG       bool
	ERROR_LOG_PATH  string
	INFO_LOG        bool
	INFO_LOG_PATH   string
	SQL_LOG         bool
	TEMPLATE_PATH   string // 静态文件相对路径
	READ_TIMEOUT    int
	WRITE_TIMEOUT   int
	PAGE_SIZE       int
}

// jwt
type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

// cookie
type CookieConfig struct {
	NAME string
}

type KafkaConfig struct {
	Host string
	Port string
}

func init() {
	ginEnv := os.Getenv("gin_env")
	if ginEnv == "" {
		ginEnv = "local"
	}
	viper.SetConfigName(ginEnv)    // 设置配置文件名 (不带后缀)
	viper.AddConfigPath("config/") // 第一个搜索路径
	viper.WatchConfig()            // 监控配置文件热重载
	err := viper.ReadInConfig()    // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&config) // 将配置信息绑定到结构体上
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(config)
}
