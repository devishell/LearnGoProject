package main

import "fmt"

func main() {
	//go通过struct自定义类型实现面向对象
	type Person struct {
		name   string
		age    int
		gender string
		hobby  []string
	}
	//实例化1
	person1 := Person{
		name:  "Tom",
		age:   45,
		hobby: []string{"足球", "篮球", "唱歌"},
	}
	fmt.Println(person1, person1.age)
	//实例化2
	var person2 = new(Person)
	person2.age = 25
	person2.name = "Well"
	fmt.Println(person2, *person2)
	//实例化3
	var person3 = &Person{
		age: 23,
	}
	person3.name = "Bill"
	fmt.Println(person3, *person3)
	//实例化4 但是必须全部初始化
	person4 := Person{
		"Joe", 24, "", nil,
	}
	fmt.Println(person4)

	//匿名结构体 通常用来只使用一次的场合
	var user struct {
		name string
		age  int
	}
	fmt.Println(user)

	//指针
	p5 := &person4
	p5.name = "Jhon" //相当于(*p5)
	fmt.Println(p5)

	//读取控制台输入
	var str string
	var i int
	fmt.Scan(&str, &i)
	fmt.Println(str, i)

}
