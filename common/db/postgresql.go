package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDb *gorm.DB

// DB初始化
func Init() {
	dns := "host=host.docker.internal user=postgres password=123 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	sqlDb, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// 取得db物件
func GetSqlDb() *gorm.DB {
	return sqlDb
}
