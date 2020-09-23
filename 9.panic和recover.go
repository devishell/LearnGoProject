package main

import "fmt"

func main() {
	a := recov()
	fmt.Println(a)
	fmt.Println("正常运行")
}

func recov() []int {
	defer func() {
		//recover尝试恢复
		err := recover()
		fmt.Println("recover捕获到异常", err)
	}()
	//错误异常panic
	var a []int //只声明
	//a[0]=100 //panic 编译不出错 但运行时出错
	if true {
		panic("手动抛出的异常panic1")
	}else {
		panic("手动抛出的异常panic2")
	}
	return a
}
