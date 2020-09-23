package main

import "fmt"

func main() {
	if x := 69; x > 80 {
		fmt.Println("x>80 优秀")
	} else if x < 60 {
		fmt.Println("x<60 不及格")
	} else {
		fmt.Println("")
	}

	age := 12
	for age > 0 { //循环
		age--
		fmt.Println(age)
		if age < 3 {
			break
			//return
			//goto
			//panic()
		}
	}

	age = 8
	switch {
	case age > 3:
		fmt.Println("age >3")
		fallthrough //会无条件执行下一个语句
	case age > 15:
		fmt.Println("age >15")
		//fallthrough
	case age > 7:
		fmt.Println("age >7")
	}

	for i := 1; i < 5; i++ {
		for j:=1;j<2;j++ {
			if i*j==8 {
				goto label//主要方便跳出多层循环
				//continue label
			}
		}
	}
	label: //定义一个标签

}
