package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
	go是解释型语言类似于c语言,编译出的可执行文件不需要安装go环境,可直接运行
	1.编译生成可执行文件
	ting@TingdeMacBook-Pro LearnGoProject % go build main.go
	2.执行可执行文件
	ting@TingdeMacBook-Pro LearnGoProject % ./main
	3.或者直接执行不生成可执行文件 go run main.go
	4.交叉编译生成跨平台可执行文件:
	4-1.edit configurations,enviroment填入GOOS=windows;CGO_ENABLED=0;GOARCH=amd64//生成go_build_LearnGoProject_.exe
	4-2.Mac下编译Linux, Windows平台的64位可执行程序：Terminal下输入
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
	如:@TingdeMacBook-Pro LearnGoProject % GOOS=windows go build main.go//生成main.exe
	*/
	for {
		var reader = bufio.NewReader(os.Stdin)
		fmt.Println("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "quit" {
			return
		} else if text == "foo" {
			fmt.Println("foo")
		} else {
			fmt.Println(text+"[type quit to exit]")
		}
	}
}
