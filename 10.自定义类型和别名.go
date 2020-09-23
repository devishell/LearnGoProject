package main

import "fmt"

//自定义类型 最好加上注释
type NewInt int

//类型别名 增加可读性 编译后不存在
type INT = int

func main() {
	var a NewInt = 1 //鼠标放NewInt/INT上注释会显示
	var b INT = 1
	fmt.Printf("%T %T\n",a,b)
	var c uint8
	var d byte
	fmt.Printf("%T %T\n",c,d)

}
