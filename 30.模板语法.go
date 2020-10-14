package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	//html/template html文本模板->类似jsp
	//文档查看 https://studygolang.com/pkgdoc net->template

	//
	http.HandleFunc("/index.html", template1) // http://127.0.0.1:9092/template1.html
	http.ListenAndServe("127.0.0.1:9092", nil)

}

func template1(writer http.ResponseWriter, request *http.Request) {
	template1, err2 := ioutil.ReadFile("index.html")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	//添加自定义函数需要在parse模板之前
	kuaFunc := func(arg string)(string,error) {
		return arg+"真帅",nil
	}
	files, err :=template.New("template1").Funcs(template.FuncMap{"kua":kuaFunc}).Parse(string(template1))
	//files, err := template.ParseFiles("index.html")	//直接打开模板 不能在添加函数
	if err != nil {
		fmt.Println(err)
		return
	}
	//用数据渲染模板 本质上是字符串替换
	data := User{ //数据也可以从数据库查询获得
		"dwx", 27,
	}
	err = files.Execute(writer, data)
	if err!=nil {
		fmt.Println(err)
	}
}

type User struct {
	Name string
	Age  int
}
