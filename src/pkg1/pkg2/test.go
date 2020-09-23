package pkg2

import "fmt"

func test(){

}

func init(){
	fmt.Println("test init")
}

type Abc struct {
	Name string;age int
}
//方法名大写才能被调用
func NewAbc(name string,age int) Abc {
	return Abc{name,age}
}
