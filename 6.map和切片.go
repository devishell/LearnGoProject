package main

import "fmt"

func main() {
	//map即映射 无序
	var map1 map[string]int
	map1 = make(map[string]int,10)
	map1["a"] = 1
	map1["b"] = 1
	map1["c"] = 2
	map1["c"] = 3
	fmt.Println(map1)
	//声明时初始化
	map2:= map[string]bool{
		"ok":true,"no":false,
	}
	fmt.Println(map2)

	//判断数组是否存在某个值 ok返回bool
	v,ok := map1["a"]
	fmt.Println(v,ok)

	//遍历
	for k, v := range map1 {
		fmt.Println(k,v)
	}
	for  i := range map1 {
		fmt.Println(i)
	}

	//删除
	delete(map1,"b")
	fmt.Println(map1)

	//map切片（数组）切片里的值是map
	mapSlice :=make([]map[int]string,4,8)
	fmt.Println(mapSlice)
	mapSlice = append(mapSlice, map[int]string{1:"11"})//添加的顺序不固定
	mapSlice = append(mapSlice, map[int]string{2:"22"})
	mapSlice[2] = map[int]string{3:"33"}
	mapSlice[2][7] = "77"
	for i, m := range mapSlice {
		fmt.Println(i,m)
	}
	//slicemap map里的值是切片
	sliceMap :=make(map[int][]string,10)
	fmt.Println(sliceMap)
	sliceMap[0] = make([]string,2)
	sliceMap[0][0] = "str1"
	sliceMap[0][1] = "str2"
	fmt.Println(sliceMap)

}
