package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//c/s 可定制化高，用户体验好 开发成本高，添加新功能需要升级
	//b/s web开发成本低 无法做复杂功能 常用到net/http包

	//html基础 -> hello.html
	http.HandleFunc("/", helloHtml) //按住command点击访问 http://127.0.0.1:9091
	http.HandleFunc("/login", helloHtml)
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:9091", nil)
}

func helloHtml(writer http.ResponseWriter, request *http.Request) {
	file, err := ioutil.ReadFile("hello.html")
	if err == nil {
		writer.Write(file)
	}else {
		writer.Write([]byte(err.Error()))
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.Host, request)
	request.ParseForm()
	fmt.Println(request.Form)
	writer.Write([]byte("index\n"))
	writer.Write([]byte(request.Form.Encode() + "\n"))
	writer.Write([]byte("获取到密码 "+request.Form.Get("pwd") + "\n"))
}

func search(writer http.ResponseWriter, request *http.Request) {

}
