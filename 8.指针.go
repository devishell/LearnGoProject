package main

import "fmt"

func main() {
	//地址：字节描述的内存地址 指针：带类型的地址
	//&取地址 *取地址的值
	var a int
	fmt.Println(a)
	p:=&a
	fmt.Printf("%T %T %T\n",a,p,*p)
	fmt.Println(a,p,*p)
	//go语言可以读取指针 不能操作指针比较安全
	*p = 10
	fmt.Println(a,p,*p)

	//函数需要传入指针才可修改外部的变量
	b:=[3]int{1,2,3}
	testp1(b);
	fmt.Println(b)
	testp2(&b)
	fmt.Println(b)

	//声明和初始化
	var p1 * int//声明int型指针 nil
	var p2 = new(int)//得到int型的指针初始化
	fmt.Println(p1,p2)
	*p2 = 1
	fmt.Println(*p2)

	//new用来厨师喜欢指针型 make初始化slice map chan
}
func testp1(a [3]int)  {
	a[0] = 10
}
func testp2(a *[3]int)  {
	a[0] = 10
	//*a[1] = 11//go里会自动推测
}
