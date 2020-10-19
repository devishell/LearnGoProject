package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func worker(ctx context.Context) {
	defer wg1.Done()
	go worker2(ctx)
LABEL:
	for {
		fmt.Println("working ...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}
}
func worker2(ctx context.Context) {
	defer wg1.Done()
LABEL:
	for {
		fmt.Println("working2 ...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}
}
func main() {
	//退出http包中请求使用的goroutine，释放goroutine占用的资源-->select chan或者全局变量判断（有缺点所以用context改进）
	cancelctx, cancelFunc := context.WithCancel(context.Background())
	wg1.Add(2)
	go worker(cancelctx)
	time.Sleep(time.Second * 3)
	cancelFunc() //调用cancel函数通知子goroutine退出
	wg1.Wait()
	fmt.Println("over")
	time.Sleep(time.Second * 2)
	fmt.Println("make sure over")
	//context可以控制多个goroutine退出

	//context接口
	//context.WithCancel()//用的最多
	//context.WithDeadline()//超时了也要调用cancel
	//context.WithTimeout()//超时了也要调用cancel
	//context.WithValue()//课要使用自定义类型
	//两个根函数
	context.Background()
	context.TODO()
	//不能把context放入结构体中，要以参数形式显示传递，作为第一个参数
	//传入context参数的时候不知道传什么就传TODO不要传nil
	//context线程安全可以在多个groutine中传递

}