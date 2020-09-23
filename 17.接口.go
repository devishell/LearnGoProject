package main

import (
	"fmt"
	"strings"
)

type CCat struct {
	name string
}

type DDog struct {
	n string
	a int
}

//接口后缀一般加er 如writer stringer
type Sayer interface {
	Say() string
}

func (c CCat) Say() string {
	return "喵喵喵"
}
func (d DDog) Say() string {
	return "汪汪"
}

func main() {
	//接口(协议)是一种抽象的类型，区别于具体的类型比如int string会存储数据，接口定义一套协议、规范。
	var s Sayer
	fmt.Printf("%T %v\n", s, s)
	var s1 = CCat{"Jerry"}
	s2 := DDog{"Tom", 5}
	fmt.Printf("%v %v\n", s1, s2)
	s = s1 //把结构体赋值给接口
	fmt.Printf("%T %v\n", s, s)
	s = s2
	fmt.Printf("%T %v\n", s, s)
	//接口运用
	var sayers []Sayer
	sayers = append(sayers, s1)
	sayers = append(sayers, s2)
	fmt.Println(sayers)
	for _, sayer := range sayers {
		println(sayer.Say())
	}
	//方法返回接口 而不是具体的结构体
	ss := getSayer("汪")
	fmt.Println(ss)
	//多个协议的接口必须同时满足才能算
	var animals []Animal
	animals = append(animals, s2)
	fmt.Println(animals)
	//空接口 可用来接收任意类型,相当于Object或泛型
	showType(123)
	showType(animals)
	var info = make(map[interface{}]interface{}, 3) //键值对为空接口的map
	info["丁三"] = 100
	info["李思"] = "晕"
	info[1] = 1
	info["2da"] = false
	fmt.Println(info)
	//判断空接口传入的类型 .(type)
	v, ok := info[1].(int)
	if ok {
		fmt.Println(v,info[1])
	} else {
		fmt.Println(v,info[1])
	}
	showType2(123)
	showType2(false)
	showType2(animals)
	sp:="123"
	showType2(&sp)
	showType2(&s1)
	//如果用指针实现接收者的接口没问题，语法糖 （*s1）.say,但是反过来本身是指针式接收者的接口 不能传普通接收者 大部分几乎都是指针式接收者
	p :=&s1
	p.Say()
	s = p
	fmt.Println(s)
}
func getSayer(s string) Sayer {
	if s == "喵喵喵" {
		return CCat{"Je"}
	} else if strings.Contains(s, "汪") {
		return DDog{"To", 2}
	}
	return nil
}

//定义多个接口
type Animal interface {
	Say() string
	Move()
}

func (d DDog) Move() {

}

func showType(a interface{}) {
	fmt.Printf("%T\n", a)
}
func showType2(a interface{}){
	switch v := a.(type) {
	case int:
		fmt.Printf("a is a int value is %d\n",v)
	case string:
		fmt.Printf("a is a string value is %v\n",v)
	case *string:
		fmt.Printf("a is a string pointer,value is %v\n",*v)
	default:
		fmt.Printf("a is a %T value is %v\n",v,v)
	}
}
