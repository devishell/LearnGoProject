package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件 默认只读
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("打开文件出错：", err)
	} else {
		defer file.Close() //关闭文件
		//读文件
		//方法1 不推荐
		//var tmp [128]byte
		//for  {
		//	n, err2 := file.Read(tmp[:])
		//	if err2 == io.EOF { //end of file
		//		fmt.Println("读完了")
		//		break
		//	}
		//	if err2!=nil {
		//		fmt.Println("读取出错：",err2)
		//		break
		//	}
		//	fmt.Printf("读取了%d字节,内容：\n%v\n",n,string(tmp[:]))
		//}
		//方法2 bufio读文件 先读到缓存里
		reader:= bufio.NewReader(file)
		for {
			str,err3 := reader.ReadString('\n') //分行读取
			if err3 == io.EOF {
				fmt.Print(str,"\n")
				break
			}
			if err3!=nil {
				fmt.Print("读取出错：",err3)
				break
			}
			fmt.Print(str)
		}
	}
	//读文件方法3 调用ioutil
	fmt.Println("---------调用ioutil读取---------")
	bytes,err4:=ioutil.ReadFile("1.txt")
	if err4 != nil {
		fmt.Println("读取出错:",err4)
	}else {
		fmt.Println(string(bytes))
	}

	//写文件
	//以指定模式打开文件 返回操作的文件指针 perm文件权限,八进制数 r读04 w写02 x执行01
	file,err = os.OpenFile("1.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	if  err!=nil{
		fmt.Println("打开文件出错：",err)
	}else {
		defer file.Close()
		//写文件方法1
		file.Write([]byte("我是append []byte\n"))
		file.WriteString("我是append str\n") //执行后查看一下文件
		//写文件方法2 bufio
		writer:=bufio.NewWriter(file)
		writer.Write([]byte("我是 bufio append []byte\n"))
		writer.WriteString("我是 bufio append str\n")
		writer.Flush()//必须执行才能写入文件
	}
	//写文件方法3
	//ioutil.WriteFile("2.txt",[]byte("我是append []byte\n"),0755)

	//拷贝文件 1.读打开src 2.写打开des 3.io.copy(des,src)
	//io.Copy()






}
