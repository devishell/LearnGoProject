package main

import "fmt"

type stu struct {
	int
	name   string
	string //不写变量名 默认以类型作为变量名 但相同类型不能定义多个
	Dog
	*Cat
}
type Dog struct {
	name string
	age  int
}
type Cat struct {
	name string
}
//NewDog的构造函数（如果函数开头大写表示开放）
func NewDog(name string,age int)  *Dog{
	return &Dog{name,age}
}

func main() {
	//嵌套
	stu1 := stu{
		1, "dk", "dwx", Dog{
			"Lili", 5,
		}, &Cat{
			"cc",
		},
	}
	fmt.Println(stu1, stu1.name, stu1.Dog.name, stu1.age) //有相同变量name 必须显式调用 age可以直接隐式调用(可实现继承特性)
}
