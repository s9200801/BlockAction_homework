package main

import (
	"blockaction_homework/common/db"
	"blockaction_homework/common/migrate"
	"blockaction_homework/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var server *gin.Engine

func init() {
	logger := logrus.New()
	db.Init()
	migrate.InitTable()
	sqlDB := db.GetSqlDb()

	server = gin.Default()

	// 預期外錯誤統一回傳500
	server.Use(gin.CustomRecoveryWithWriter(logger.Out, func(ctx *gin.Context, err any) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}))

	router.Init(server, sqlDB)
}

func main() {
	server.Run()
}
