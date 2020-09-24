package main

func main() {
	//github搜索 ini go/ mysql go类似的方式来查找轮子
	//发送日志到通道暂存 然后go读通道写入文件或打印到控制台
	type LogData struct {
		//定义日志信息结构体
		Message string
		Level int
		Line int
		TimeStr string
		FuncName string
		FileName string
	}
	//var logDataChan * LogData
	//logDataChan <-&LogData{}  //发送
	//LogData<-logDataChan  //初始化时开启go 循环读取通道存储日志

	//如果通道满了(比如消费者不干了)阻塞可以取出来丢弃
	//select {
	//case logDataChan <-&LogData{}:
	//default:
	//	<-logDataChan //丢弃最早一条
	//	logDataChan <-&LogData{}
	//}
}
