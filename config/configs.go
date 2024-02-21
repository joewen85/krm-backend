package config

import (
	"krm-backend/utils/logs"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000"
)

var (
	LogLevel          string
	Port              string
	SigningKey        string
	ExpireTime        uint64
	Username          string
	Password          string
	MysqlAddress      string
	MysqlPort         uint
	MysqlUsername     string
	MysqlPassword     string
	MysqlDBName       string
	MaxIdleConnection int
	MaxOpenConnection int
	MetaDataNamespace string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
func NewReturnData() ReturnData {
	return ReturnData{
		Status:  200,
		Message: "",
		Data:    make(map[string]interface{}),
	}
}

func initLogSetting(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
		_ = os.Setenv("SECRET_KEY", "joe123456")
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	defer logs.Info(nil, "初始化成功")
	// environment
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8888")
	viper.SetDefault("METADATA_NAMESPACE", "mgmt")
	viper.AutomaticEnv()

	// log setting
	LogLevel = viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	initLogSetting(LogLevel)
	MetaDataNamespace = viper.GetString("METADATA_NAMESPACE")

	//	jwt
	SigningKey = viper.GetString("SECRET_KEY")
	if SigningKey == "" {
		logs.Error(nil, "environment not set SECRET_KEY")
		os.Exit(1)
	}
	viper.SetDefault("JWT_EXPIRE_TIME", 86400)
	ExpireTime = viper.GetUint64("JWT_EXPIRE_TIME")

	// mysql
	viper.SetDefault("MYSQL_ADDRESS", "127.0.0.1")
	viper.SetDefault("MYSQL_PORT", 3306)
	viper.SetDefault("MYSQL_USERNAME", "root")
	viper.SetDefault("MAX_IDLE_CONNECTION", 25)
	viper.SetDefault("MAX_OPEN_CONNECTION", 25)

	//MysqlAddress = viper.GetString("MYSQL_ADDRESS")
	MysqlPort = viper.GetUint("MYSQL_PORT")
	MysqlUsername = viper.GetString("MYSQL_USERNAME")
	//MysqlPassword = viper.GetString("MYSQL_PASSWORD")
	//MysqlDBName = viper.GetString("MYSQL_DBNAME")
	MaxIdleConnection = viper.GetInt("MAX_IDLE_CONNECTION")
	MaxOpenConnection = viper.GetInt("MAX_OPEN_CONNECTION")

	MysqlAddress = "127.0.0.1"
	MysqlPassword = "mismis"
	MysqlDBName = "devops_go"
	MysqlPort = 3306
}
