package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server      ServerConfig
		Mysql       MysqlConfig
		Redis       RedisConfig
		Cookie      Cookie
		Session     Session
		Metrics     Metrics
		Logger      Logger
		FileStorage FileStorage
		Jaeger      Jaeger
	}

	ServerConfig struct {
		AppVersion        string
		Port              string
		PprofPort         string
		Mode              string
		JwtSecretKey      string
		CookieName        string
		PasswordSalt      string
		ReadTimeout       time.Duration
		WriteTimeout      time.Duration
		SSL               bool
		CtxDefaultTimeout time.Duration
		CSRF              bool
		Debug             bool
	}

	MysqlConfig struct {
		MysqlHost     string
		MysqlPort     string
		MysqlUser     string
		MysqlPassword string
		MysqlDbname   string
		MysqlSSLMode  bool
		MysqlDriver   string
	}

	RedisConfig struct {
		RedisAddr      string
		RedisPassword  string
		RedisDB        string
		RedisDefaultdb string
		MinIdleConns   int
		PoolSize       int
		PoolTimeout    int
		Password       string
		DB             int
	}

	Logger struct {
		Development       bool
		DisableCaller     bool
		DisableStacktrace bool
		Encoding          string
		Level             string
	}

	Cookie struct {
		Name     string
		MaxAge   int
		Secure   bool
		HTTPOnly bool
	}

	Session struct {
		Prefix string
		Name   string
		Expire int
	}

	Metrics struct {
		URL         string
		ServiceName string
	}

	FileStorage struct {
		Endpoint  string
		Bucket    string
		AccessKey string
		SecretKey string
		Secure    bool
	}

	Jaeger struct {
		Host        string
		ServiceName string
		LogSpans    bool
	}
)

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	SetConfig(c)
	return &c, nil
}

var config Config

func GetConfig() Config {
	return config
}

func SetConfig(cfg Config) {
	config = cfg
}
