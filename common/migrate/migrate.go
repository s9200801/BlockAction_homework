package migrate

import (
	"blockaction_homework/app/model"
	"blockaction_homework/common/db"
)

func migrateTable[T interface{}]() {
	sqlDb := db.GetSqlDb()
	t := new(T)
	sqlDb.AutoMigrate(t)
}

// 檢查所有的model是否有建立table
func InitTable() {
	migrateTable[model.Member]()
}
