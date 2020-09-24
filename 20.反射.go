package main

import (
	"fmt"
	"reflect"
)

func main() {
	//话外：一般是Web框架O、RM框架、配置文件解析裤用到 工作用的少，学反射可以更好理解框架

	//原理 程序运行的时候在判断更灵活 但是效率低、代码难以理解、出错发现晚
	//相关：空接口 类型断言x.(type) -> 使用反射可以直接拿到动态类型和动态值

	//反射的typeOf
	var a interface{}
	a = false
	t := reflect.TypeOf(a)
	fmt.Println(t,t.Name(),t.String(),t.Size())
	a = []int{1,2,3}
	t = reflect.TypeOf(a)
	fmt.Println(t,t.Name(),t.String(),t.Size()) //name可能为空
	//代码补全也是用的反射

	//反射的valueOf
	v := reflect.ValueOf(a)
	fmt.Println(v,v.String(),v.Kind())
	fmt.Printf("%T\n",v)

	//kind 类型
	kind:=v.Kind() //根据kind进行修改值
	kind = t.Kind()//也可以通过type获取
	if kind == reflect.Ptr { //指针
		//v.Elem().SetFloat()
		//v.Elem().SetInt()
	}else if kind == reflect.Int {
		//v.SetInt()
	} else if kind == reflect.Struct {  //结构体反射 获取里面的字段 结构体和字段需要大写
		//v.Field()
		//v.FieldByName()
		//v.FieldByIndex()
		fmt.Println(t.NumMethod(),t.NumField(),v.NumMethod(),v.NumField())
	}

	//作业？ 实现配置文件一行行解析
	type Config struct {
		FileName string
		MaxSize int
	}





}
