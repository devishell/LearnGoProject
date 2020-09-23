package main

import (
	"fmt"
	p1 "pkg1/pkg2"
	"pkg2"
)

func main() {
	println("123")
	//cd src
	//go build 1.变量常量.go
	//GOOS=windows go build 1.变量常量.go


	//调用其他包的方法 包.方法
	pkg2.Test2()

	//别名 可以解决包名冲突
	p1.NewAbc("JoJo",2)
}

func init(){ //init方法会自动执行，用来初始化操作
	fmt.Println("init")
}
//执行顺序：全局声明->init->main
//最里层的包init先执行 main的init最后执行