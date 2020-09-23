package main

import "fmt"

const con1 = "con1"

func main() {
	//变量声明 没有使用后会报错
	var name string
	name = "dog"
	fmt.Println(name)
	var age = 5 //自动推测类型
	age++
	fmt.Printf("age=%d\n", age)
	fmt.Println("age=", age, "name=", name)
	//批量声明
	var (
		a string  //""
		b bool    //false
		c float32 //0
		//d  //强类型语言必须带上类型
	)
	fmt.Printf("a=%s,b=%,c=%f \n", a, b, c)
	x := "1" //短变量声明 只能在函数内部使用
	fmt.Printf("x=%s\n", x)
	//匿名变量 _用于接收不需要的值,不使用也不会报错,且可重复声明使用
	m, _ := foo()
	fmt.Printf(m + "\n")

	//常量 const不能修改  最好定义在函数外面 也可批量声明
	const pi = 3.14
	const (
		con2 = "22"
		con3 = "33"
		e    = iota //每个const中iota出现时初值为0 const每加一行iota就会停累加一次
		f    = iota
		g
		h
		i = 2
		j //后面不赋值默认等于上一个值
		k
		//空换行iota不会增加
		l = iota
	)
	fmt.Println(con2, con3, e, f, g, h, i, j, k, l)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
	)
	fmt.Println(KB,MB,GB)
	const (
		aa, bb     = iota + 1, iota + 2       //iota=0, aa = 1,bb = 2
		cc, dd, ee = iota, iota + 1, iota + 2 //iota=1,cc =1,dd=2,ee =3
	)
	fmt.Println(aa,bb,cc,dd,ee)

}

//两个返回值的函数
func foo() (string, int) {
	return "哈哈", 22
}
