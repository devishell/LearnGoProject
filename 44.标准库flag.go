package main

import (
	"flag"
	"fmt"
)

func main() {
	name:=flag.String("name","张三","姓名")
	var (
		name1 string
		age int
	)
	flag.StringVar(&name1,"name1","李四","姓名")
	flag.IntVar(&age,"age",22,"年龄")
	fmt.Println(*name,name1,age)
	flag.Parse()//必须调用
	fmt.Println(*name,name1,age)


	fmt.Println(flag.Args(),flag.NFlag(),flag.NArg())
	//ting@TingdeMacBook-Pro LearnGoProject % go build 44.标准库flag.go
	//ting@TingdeMacBook-Pro LearnGoProject % 44.标准库flag -name="dwx" -age=34

}
