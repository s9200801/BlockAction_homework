package logger

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

// logger 初始化 可在這調整log的設定
func Init() {
	Logger = logrus.New()
}