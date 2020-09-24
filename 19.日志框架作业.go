package main

import (
	"fmt"
	os "os"
	"path"
	"runtime"
	"time"
)

func main() {
	//1.日志级别 日志管理类
	const (
		INFO = iota
		DEBUG
		WARN
		ERROR
	)
	type FileLogger struct {
		level       int
		logFilePath string //初始化类需要的参数
		logFileName string
		f           *os.File
	}

	//2.日志按格式写入文件 时间+ level + 哪个类哪行 + desc
	//初始化
	log := &FileLogger{DEBUG, "", "debug.log", nil}
	f, _ := os.OpenFile(log.logFilePath+log.logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	log.f = f //赋值给log.f
	defer log.f.Close()
	//打印+写日志 方法封装
	//log.f.WriteString()
	now := time.Now().Format("2006-01-02 15:04:05.000")
	funcName, fileName, line := getCallerInfo()
	desc := "desc"
	fmt.Printf("[%s] [%s] [%s] [%s] [%d] %s\n", now, getLevelStr(log.level), fileName, funcName, line, desc)
	fmt.Fprintf(log.f, "[%s] [%s] [%s] [%s] [%d行] %s\n", now, getLevelStr(log.level), fileName, funcName, line, desc)

	//文件状态api
	finfo,_:= log.f.Stat()
	//名称 大小 权限 修改时间 是否是目录
	fmt.Println(finfo.Name(),finfo.Size(),finfo.Mode(),finfo.ModTime(),finfo.IsDir())

	//os api
	os.Stat("debug.log")
	//os.Chdir()
	//os.Chmod()
	//os.Create()
	//os.Stdout
	//os.Rename() //重命名
	//os.SameFile() //比较
}

//runtime
func getCallerInfo() (funcName string, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		return
	}
	//根据pc拿到当前执行的函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(file)
	fmt.Println(funcName, fileName, line) //方法名 文件名 行数
	return funcName, fileName, line
}

func getLevelStr(level int) string {
	switch level {
	case 0:
		return "INFO"
	case 1:
		return "DEBUG"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	default:
		return "INFO"
	}
}
