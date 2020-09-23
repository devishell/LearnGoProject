package main

import (
	"fmt"
)

func main() {

	var a [5]int
	a = [5]int{1, 2}
	fmt.Println(a)
	var b = [8]int{1, 2, 3, 4: 8, 5: 5}
	fmt.Println(b)
	for index, value := range b {
		fmt.Println(index, b[index], value)
	}
	var c = [...]int{1, 2, 3, 4, 5} //[...]会自动推测长度 不可变
	var d = [5]int{2, 2, 3, 3, 3}   //不可变
	var dd = []int{2, 2, 3, 3, 3}   //[]长度可变 切片
	fmt.Println(c, d)
	e := c //新数组声明赋值 数组拷贝,新赋值不会影响原数组
	e[1] = 99
	d = c //旧数组赋值 数组拷贝,不会影响原数组
	d[0] = 66
	//dd = c//类型不同 不能赋值
	ee := dd
	ee[1] = 88
	fmt.Println("值类型复制", c, d, e) //值类型复制的时候都是拷贝 不会影响原数组
	fmt.Println("引用类型复制", dd, ee) //引用类型复制的时候都是指向同一个块内存 会同时改变
	//多维数组
	m := [5][3]int{
		{1, 3, 2},
		{3},
		{2: 45},
		3: {22, 2, 3},
		4: {23},
	}
	fmt.Println(m)

	//切片slice 基于数组的一层封装，自动扩容，是引用类型，数组是值类型容量不可变
	//丛数组得到切片
	dd = c[:] //全部
	fmt.Println(dd, len(dd), cap(dd))
	dd = c[0:2] //0到2但不包含2  //dd = c[:2]
	fmt.Println(dd, len(dd), cap(dd))
	//cap容量：切片底层的数组容量从开头算到数组末尾 且切片还可以被切限制按底层数组的长度不能超过范围
	ddd := dd[1:5]
	fmt.Println(ddd, len(ddd), cap(ddd))
	d4 := ddd[:2]
	fmt.Printf("c=%p dd=%p ddd=%p d4=%p\n", &c, dd, ddd, d4) //切的位置不同 对应切片的指针指向的位置也不同
	//扩容
	fmt.Printf("%v %p %d %d\n", d4, d4, len(d4), cap(d4))
	d4 = append(d4, 22, 33)
	fmt.Printf("%v %p %d %d\n", d4, d4, len(d4), cap(d4)) //容量够时指针不变
	d4 = append(d4, 55)
	fmt.Printf("%v %p %d %d\n", d4, d4, len(d4), cap(d4)) //容量不够时会翻倍且指针变了(换了块新内存)
	//指针变了 改变值不会影响原来的底层数组
	d4[0] = 88
	fmt.Println(c, d, dd, ddd, d4)

	//切片必须初始化或append后才能使用 make用于给引用类型初始化申请内存
	var slice1 []int
	//slice1 = []int{1,2,3}
	//slice1 = append(slice1,1)
	slice1 = make([]int, 0, 100) //切片初始化时尽量容量大一点免得动态扩容
	fmt.Println("slice1=", slice1, len(slice1), cap(slice1), cap(make([]string, 6)))

	//copy delete
	slice2 := []byte{1, 2, 3}
	slice3 := make([]byte, len(slice2))
	copy(slice3, slice2)
	fmt.Println(slice2, slice3)

}
