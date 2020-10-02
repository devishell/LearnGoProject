package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {
	// cs->bs , http80 https443 超文本协议
	//应用层：http，ftp，smtp 传输层：tcp，udp 网络层：ip，arp，路由器 数据链路层：以太网，网桥 物理层：双绞线，光纤，无线电波

	//http client
	url := "www.eeeoo.cn:80"
	//url:="www.jd.com:80"
	conn, err := net.Dial("tcp", url)
	if err == nil {
		defer conn.Close()
		//发送数据： 请求方法 路径 协议
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		//接收数据
		var buf [1024]byte
		for {
			read, err := conn.Read(buf[:])
			if err == io.EOF {
				fmt.Println(string(buf[:read]))
				break
			} else if err != nil {
				fmt.Println(err)
				break
			} else if err == nil {
				fmt.Println(string(buf[:read]))
			}
		}
	} else {
		fmt.Println(err)
	}

	//http server
	//注册路由
	http.HandleFunc("/", handleFunc) // 访问 http://127.0.0.1:9090
	http.HandleFunc("/login", handleFunc2)// http://127.0.0.1:9090/login
	//监听9090
	err = http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Println(err)
	}

	//http是流式的可能会出现粘包 给数据制造固定的长度来解决

}

func handleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1 style='color:red'>hello index！</h1>"))
}
func handleFunc2(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1 style='color:blue'>登陆成功！</h1>"))
}
