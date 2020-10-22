package main

/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h>
int add(int a,int b){
return a+b;
}
float plus(float a,float b){
printf("a=%.1f,b=%f\n",a,b);
return a-b;
}
*/
import "C"

import "fmt"

func main() {
	//参考 https://blog.csdn.net/zdy0_2004/article/details/79124269
	//go 调用 c/c++ 函数的实现方式有三种,原理相似：
	//1.直接嵌套在go文件中使用，最简单直观的
	fmt.Println(C.add(1,2),C.plus(6.1,6))
	//2.导入动态库 .so 或 dll 的形式，最安全但是很不爽也比较慢的
	//3.直接引用 c/c++ 文件的形式，层次分明，容易随时修改看结果的
}
