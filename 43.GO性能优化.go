package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main() {
	//profiling：cpu内存等使用情况的画像， 上线后不要引入
	//内置库runtime/pprof 采集工具型应用数据进行分析 net/http/pprof采集服务型应用数据进行分析
	//使用gin框架则可使用ginprof
	file,err := os.Open("")
	if err == nil {
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
		pprof.WriteHeapProfile(file)
	}
}
