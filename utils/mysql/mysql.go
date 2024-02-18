package main

import (
	"fmt"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"column:username;size:32;uniqueIndex"`
	Password    string `json:"password" gorm:"column:password"`
	IsAdmin     bool   `json:"is_admin" gorm:"column:is_admin;comment:是否管理员"`
	IsActive    bool   `json:"is_active" gorm:"column:is_active;comment:是否激活"`
	FullName    string `json:"full_name" gorm:"column:full_name;size:128;comment:中文名"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;size:11;comment:手机号码"`
}

var _db *gorm.DB

func GetDB() *gorm.DB {
	return _db
}

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s", config.MysqlUsername, config.MysqlPassword, config.MysqlAddress, config.MysqlPort, config.MysqlDBName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("connect database error")
	}

	if config.LogLevel == "debug" {
		db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		logs.Error(nil, err.Error())
	}
	sqlDb.SetMaxIdleConns(config.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(config.MaxOpenConnection)
	sqlDb.SetConnMaxLifetime(time.Minute)
	status := sqlDb.Stats().OpenConnections
	fmt.Println("打开连接数", status)
	_db = db
}

func MigrateDB() {
	db := GetDB()
	if err := db.AutoMigrate(&User{}); err != nil {
		logs.Error(nil, err.Error())
	}
}

func CreateData(i interface{}) error {
	db := GetDB()
	sqlDB, _ := db.DB()
	fmt.Println(&i)

	defer sqlDB.Close()
	if err := db.Debug().Create(i).Error; err != nil {
		return err
	}
	return nil
}

func GetData(i interface{}, params map[string]interface{}) {
	db := GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	for k, v := range params {
		fmt.Println(k, v)
	}
	//db.First(i, key, value)
}

//func OperateMySQL() {
//	db := GetDB()
//
//	sqlDB, _ := db.DB()
//	defer sqlDB.Close()
//	user := User{Username: "haha", Password: "aabb", IsActive: true, IsAdmin: true}
//	db.Create(&user)
//
//	// Read
//	NewUser := User{}
//	//db.First(&NewUser, 1) 获取主键为1的对象
//	db.First(&NewUser, "username = ?", "bobo") // 获取username为bobo的对象
//	// update
//	db.Model(&NewUser).Update("password", "aaabbb")
//	// multiple fields
//	//db.Model(&NewUser).Updates(User{IsAdmin: false, IsActive: false})
//	db.Model(&NewUser).Updates(map[string]interface{}{
//		"IsAdmin":  false,
//		"IsActive": false,
//	})
//	fmt.Println(NewUser.Username)
//
//	//db.Delete(&NewUser, "username =?", "bobo")
//}

func main() {
	MigrateDB()
	//OperateMySQL()
	user := User{Username: "joewen4", Password: "aabb", IsActive: true, IsAdmin: true, PhoneNumber: "13800138000", FullName: "joewen"}
	logs.Info(nil, "插入数据成功")
	CreateData(&user)
}
