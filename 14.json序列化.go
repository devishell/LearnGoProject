package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name string //结N构体字段开头大写表示公开 小写表示私有
	age  int    //小写私有无法反序列化
	Len  int    `json:"len"` //这种方式可以使json里面字段小写 解决前端需求
}
func main() {
	man:=Man{
		"Ma",12,22,
	}
	fmt.Println(man)
	//序列化
	bytes, err := json.Marshal(man)
	if err!=nil {
		fmt.Printf("%#v",string(bytes)) //bytes转换为string
	}
	//反序列化
	man2 := &Man{}
	str:="{\"Name\":\"Joe\", \"age\":32 ,\"len\":23}" //此时age无法被反序列化
	json.Unmarshal([]byte(str), man2)
	fmt.Printf("%#v",*man2)

}
