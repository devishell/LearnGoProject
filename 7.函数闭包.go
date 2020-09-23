package main

import "fmt"

func main() {
	//函数声明
	fmt.Println(add1(1, 2), add2(1, 2), add3(1, 2))

	//可变参数只能有一个且只能放在最后一个参数
	dynmic(1)

	//go里面参数没有默认值
	//支持多返回值

	//defer 逆序执行 延时执行（方便用来处理资源释放问题）
	deferTest()

	//闭包 closure
	i := calc(1, 2, func(i int, i2 int) int {
		return i + i2
	})
	fmt.Println("calc",i,calc(1,2,add3))
	f1()
	fmt.Printf("%T , %T , %v\n",f1,f2(),f2()(10))
	//闭包多返回值
	f1,f2 := f3()
	fmt.Println(f1(2),f2("Joe"))

	//

}
func add1(a int, b int) int {
	return a + b
}
func add2(a, b int) int { //参数合并
	return a + b
}
func add3(a, b int) (ret int, error int) { //省略return后面
	ret = a + b
	error = 0
	return
}

func dynmic(a ...int) {
	fmt.Printf("%v %T\n", a, a)
}

func deferTest() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)
}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
func f1()  {
	num:=10
	func(){
		fmt.Println("num=",num)
	}()
}
func f2() func(int)int {
	ret:= func(a int)int {
		a++
		return a
	}
	return ret
}
func f3() (func(int) int,func(string)string){ //多返回值需要用括号括起来
	ret:= func(a int)int {
		a++
		return a
	}
	ret2:= func(a string)string {
		return "hello,"+a
	}
	return ret,ret2
}
