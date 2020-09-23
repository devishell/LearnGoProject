package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Time{})
	fmt.Println(time.Now())

	now:=time.Now()
	fmt.Println(now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second(),now.Nanosecond())
	//精确时间戳unix
	fmt.Println(now.Unix(),now.UnixNano())

	//转换时间戳
	unix:=time.Unix(8322011294,0)
	fmt.Println(unix)

	//格式化 格式化模板是golang诞生日期 20060102150405
	fmt.Println(now.Format("2006-01-02 15:04.000")) //golang不支持静态变量 可以把定义的全局变量认为是静态变量
	fmt.Println(now.Format("2006/1-2 15:04.999"))
}
