package main

import "fmt"

func main() {
	a := 'c'
	fmt.Printf("a=%v\n", a) //%v通用占位符
	fmt.Printf("a=%T\n", a) //%T type类型
	fmt.Printf("%%\n")      //%% 转义%
	b := 66
	fmt.Printf("|%d|\n", b)   //指定整形的长度
	fmt.Printf("|%8d|\n", b)  //左留白
	fmt.Printf("|%-8d|\n", b) //
	fmt.Printf("|%08d|\n", b) //补0
	c := 123.456
	fmt.Printf("%f\n", c)
	fmt.Printf("%g\n", c)   //用最少的数字来表示
	fmt.Printf("%.2f\n", c) //小数部分最多2位数字
	fmt.Printf("%.4g\n", c) //整个长度最多4位数字
	d := false
	fmt.Printf("%t\n", d)
	e := struct {
		m int
		n string
		o bool
	}{}
	fmt.Printf("%v\n", e)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%#v\n", e)

}
