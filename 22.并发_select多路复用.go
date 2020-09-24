package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//select 提高同时接收发送多个通道值的效率，且容易关闭chan
	//例子1
	var ch1 = make(chan string, 1)
	var ch2 = make(chan string, 1)
	exitChan := make(chan bool, 1)
	go fu1(ch1)
	go fu12(ch2)
	//监听键盘输入,输入q直接触发exitchan
	go func() {
		for {
			tmp := make([]byte, 1)
			os.Stdin.Read(tmp) //输入后必须按enter才能退出
			//fmt.Scan(&tmp) //fmt.Scanln(&tmp) //必须直接按enter
			if tmp[0] == 'q' {
				exitChan <- true
			}
			fmt.Println("input is ", string(tmp))
		}
	}()
	ret := " "
	for ret != "" {
		select { //select本身也是阻塞的
		case ret = <-ch1:
			fmt.Println(ret)
		case ret = <-ch2:
			fmt.Println(ret)
		case ret1 := <-exitChan: //检测到退出chan 此时可以提前退出并关闭通道
			fmt.Println("ret1=", ret1)
			close(exitChan)
			close(ch1)
			close(ch2)
			break
		default:
			fmt.Println("暂时取不到值 ret=", ret)
			time.Sleep(time.Second / 10)
		}
	}
	fmt.Println("退出了ret=", ret) //关闭通道后 ret是对应类型的零值

	//单向通道 只能发送或只能接收 一般用于在设计函数的参数里面定义单向通道
	var c1 chan<- int //只能发送
	var c2 <-chan int //只能接收
}
func fu1(ch chan string) {
	for i := 100; i > 0; i-- {
		ch <- fmt.Sprintf("fu1:%d", i)
		time.Sleep(time.Second / 20)
	}
	close(ch)
}
func fu12(ch chan string) {
	for i := 100; i > 0; i-- {
		ch <- fmt.Sprintf("fu2:%d", i)
		time.Sleep(time.Second / 20)
	}
	close(ch)
}
