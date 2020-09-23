package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//int 自动匹配平台的长度32或64
	//int8相当于byte int16==short int64==long
	//对应无符号uint uint8 uint16等
	var (
		a int = 100
		b int = 077
		c int = 0xff
	)
	//%b表示用二进制显示 %o 八进制 %x 16进制 %p 变量内存地址(16进制表示)
	fmt.Printf("%b %o %x %X %p\n", a, b, c, c, &a)

	//float32 float64
	var f1 float32 = math.MaxFloat32
	var f2 float64 = math.MaxFloat64
	fmt.Println(f1, f2)

	//bool无法与其他类型强制转换 默认false

	//string 编码UTF-8
	s2 := fmt.Sprintf("%s***---%s ii", "s1", "s2") //返回拼接字符串
	fmt.Println(s2)

	//转义符号 \n \' \t \\
	var s1 = `daf"f,afmad736\daf\"
da;fa;\dasf[]apfa[f\n\t
f` //多行文本 不转义 使用键盘左上角~`符号
	fmt.Println(s1, len(s1))

	//strings常用函数
	//ting@TingdeMacBook-Pro LearnGoProject % go doc strings //查看文档
	//ting@TingdeMacBook-Pro LearnGoProject % go doc strings.split //查看文档
	var s3 = "string,test,str "
	fmt.Println(strings.Split(s3, ","))
	//UTF8下一个中文3个字节 ascii占1个字节
	var c2 rune = '丁' //int32
	var c3 byte = 'y' //int8
	s4 := "丁wx123"
	fmt.Println(len(string(c2)),len(string(c3)),len(s4),s4[0])
	//遍历字符串 for range是按rune解决痛点 普通for是按位
	for i,v:= range s4{
		fmt.Printf("%d %c\n",i,v)
	}
	//str转byte数组 类似于string(c2)的强制类型转换
	bytearray := []byte(s4)
	for i := range bytearray {
		fmt.Println(i)
	}

}
