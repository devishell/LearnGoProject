package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// https://github.com/sirupsen/logrus
	// Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.

	//安装go get github.com/sirupsen/logrus
	logrus.Info("test info")
	logrus.Debug("debug test")
	logrus.WithFields(logrus.Fields{
		"name":"CatD","age":26,
	}).Warn("test warn level")
	//设置级别
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("test info")
	logrus.Debug("debug test")
	logrus.WithFields(logrus.Fields{
		"name":"CatD","age":26,
	}).Warn("test warn level")
	//全局附带字段
	loggerFiled1 := logrus.WithFields(logrus.Fields{
		"field1":"11",
	})
	loggerFiled1.Info("loggerFiled1 test1")
	loggerFiled1.Info("loggerFiled1 test2")
	//钩子 hook
	//logrus.AddHook()
	//格式化
	//logrus.SetFormatter(logrus.JSONFormatter{})
	//记录函数名 会增加性能开销
	logrus.SetReportCaller(true)

	//gin中使用logrus框架
	log_global.Info("test log")
	file, err := os.OpenFile("global.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err!=nil {
		fmt.Println("open global.log failed，err：",err)
		return
	}
	log_global.Out = file //记录到文件中
	log_global.Level = logrus.DebugLevel
	//gin设置模式
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = log_global.Out //让gin把日志也记录到log_global的输出中
	log_global.Info("test log")
}
var log_global = logrus.New() //定义全局可使用的logrus
