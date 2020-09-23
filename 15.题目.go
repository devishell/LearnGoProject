package main

import "fmt"

func t1() int {
	x:=5
	defer func() {
		x++
	}()
	return x //1.返回值=5 2.5++ 3.返回5
}
func t2() (x int) {
	defer func() {
		x++
	}()
	return 5//1.返回值=5,x=5 2.x++,x=6 3.返回x=6
}
func t3() (y int) {
	x:=5
	defer func() {
		x++
	}()
	return x//1.返回值=x=5,y=x=5 2.x++ 3.返回y=5
}
func t4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5//1.返回值=5，x=5 2.值拷贝不影响返回值x 3.返回x=5
}


func main() {
	//函数return语句底层实现 不是原子的分为几步
	//1.返回值赋值 2.ret指令
	//defer执行的时机在return之间，即要返回之前执行
	fmt.Println(t1(),t2(),t3(),t4())

}
