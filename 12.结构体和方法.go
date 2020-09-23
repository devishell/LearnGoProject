package main

import (
	"fmt"
)

type People struct {
	name   string
	gender string
}

func main() {
	//学完了但是不会动手运用 需要进行需求分析 把问题分解，一步步理清思路

	//方法不同于函数 函数指定接收者后就是方法
	p := People{
		"Tomm", "man",
	}
	p.walk()
	fmt.Println(p)
	p.dream()
	fmt.Println(p)

	jerry:=NewDog("Jerry",3)
	fmt.Println(jerry)
}

//p后面类型的首字母小写，相当于this，self  （有点像扩展函数） 不能在该类型的其他包扩展但可以扩展自定义类型
func (p People) walk() {
	p.name = "无法修改p" //因为这里是拷贝了一份 要想修改必须传入指针
	fmt.Printf("%s can walk\n", p.name)
}
func (p *People) dream() {
	p.name = "dreamman"
	fmt.Println(p)
}

type MyP People //主要用来扩展系统的类型
func (m MyP) test() {
	println(m.name)
}
