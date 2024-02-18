package config

import (
	"fmt"
	"krm-backend/utils/logs"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	once   sync.Once
	Reader = new(Config)
)

type Config struct {
	Server     *Server
	Mysql      *DB
	LocalCache *LocalCache
	Casbin     *Casbin
}

type Server struct {
	Port int64
}

type DB struct {
	Host     string
	Port     int64
	Username string
	Password string
	Dbname   string
	TimeOut  string
}

type LocalCache struct {
	ExpireTime time.Duration
}

type Casbin struct {
	Module string
}

func (config *Config) ReadConfig() *Config {
	once.Do(func() {
		fmt.Println("aaa", Reader)
		viper.SetConfigFile("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		fmt.Println("bbb", Reader)
		var err error
		err = viper.ReadInConfig()
		if err != nil {

			logs.Fatal(nil, fmt.Sprintf("Fatal error config file: %s \n", err))
		}
		err = viper.Unmarshal(config)
		if err != nil {
			logs.Fatal(nil, fmt.Sprintf("Fatal error viper unmarshal config: %s \n", err))
		}
	})
	return Reader
}
