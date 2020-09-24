package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup

func hello() {
	fmt.Println("Hello !")
	wait.Done() //计数器减一
}
func main() { //main也是goroutine在调度 函数结束了goroutine也就结束了
	//go天生支持并发，很容易写出效率不错的并发代码

	//并发和并行区别
	//并发：同时和两个女朋友聊天 同一时间点只能和一个聊天；同一时间段同时在做多个事情，充分利用cpu时间
	//并行：两个人都在和女朋友聊天 ；统一时刻在做多个事情

	//进程 线程 协程
	//进程：启动一个程序软件就是一个进程
	//线程：操作系统多任务 ，听歌上网；操作系统调度的最小单位
	//协程：用户态的线程,pythin,php,kotlin,很多语言没有协程，不阻塞等待

	//goroutine 由运行时runtime调用 无序
	//一个函数可以被创建多个groutine 一个goroutine对应一个函数
	wait.Add(2)
	go hello()
	go hello()
	fmt.Println("go hello")
	//time.Sleep(time.Second) //创建goroutine需要一段时间 所以需要等待一会
	wait.Wait() //阻塞等待函数执行完成
	//goroutine的栈可按需增大缩小最大可到1Gb，初始一般为2kb 而线程一般是2Mb多了会崩溃
	//且线程调用成本比较高，goroutine为m:n调度（m个goroutine调度到n个线程）
	//runtime.GOMAXPROCS(1)//设置go程序调用的逻辑核心 m:n中n为1

	//channel goroutine之间的连接通道 便于数据交换 属于引用类型
	var ch1 chan int //定义一鞥channel类型传递存储的int类型
	var ch2 chan string
	ch3 := make(chan int, 1)
	fmt.Println(ch1, ch2, ch3)
	//通道的操作：发送 、接收 、关闭
	ch3 <- 10    //10发送到ch3中
	ret := <-ch3 //ch3的值保存到ret中
	fmt.Println(ret)
	close(ch3) //关闭通道后还可以取到对应类型的零值(如果多个goroutine在赋值 关闭操作后还是可以拿到值的) 但不能发送值了且不能重复关闭
	ret = <-ch3
	fmt.Println(ret)

	ch4 := make(chan bool) //不初始化的话取值就会阻塞等待
	fmt.Println(len(ch4), cap(ch4))
	go rcv(ch4)
	ch4 <- true
	//close(ch4) //可以不关闭 可以被垃圾回收机制回收
	//无缓冲通道 存了值如果没取的话再次存值会阻塞
	//缓冲通道 异步操作
	//在不停的接收值时怎样判断通道没关闭来取值？
	go func() {
		ret1, ok := <-ch4 //也可以用for ret:= range ch4{} 循环从ch4中取值 不需要判断是否关闭
		if ok {
			fmt.Println(ret1)
		}
	}()
	ch4 <- false

	//通道练习题 生产者消费者
	item := make(chan *item, 100)
	result := make(chan *result, 100)
	for n := 3; n > 0; n-- { //多个生产者 每个生产者都在不停生产
		go producer(item)
	}
	for n := 2; n > 0; n-- { //多个消费者 每个消费者都在不停消费
		go consumer(item, result)
	}
	printResult(result) //只要result里面有值就不会退出 一直取值,不加go防止主线程退出

	rand.Seed(time.Now().UnixNano()) //加随机数种子后每次执行产生的随机数都会不一样 否则多次执行都相同
	ret2 := rand.Int63()             //int64正整数
	fmt.Println(ret2)
	//

	//并发控制与锁

	fmt.Println("main 函数结束了")
}

func rcv(ch chan bool) {
	ret := <-ch //ch不初始化 取值这里会阻塞直到ch有值
	fmt.Println(ret)
}

type item struct {
	id  int64
	num int64
}
type result struct {
	*item
	sum int64
}

func producer(ch chan *item) {
	var id int64
	for {
		id++
		number := rand.Int63()
		tmp := &item{id, number}
		ch <- tmp               //把随机数发送到通道中
		time.Sleep(time.Second) //免得生产太快阻塞太多
	}
}

//计算数字每个位的和
func calc1(num int64) int64 {
	//123%10=12..3 sum=0+3
	//12%10=1..2 sum=0+3+2
	//1%10=0..1 sum=0+3+2+1
	var sum int64
	for num > 0 {
		sum = sum + num%10 //取余
		num = num / 10     //取整
	}
	return sum
}

//消费者
func consumer(ch chan *item, res chan *result) {
	for tmp := range ch { //取值一般用for range代替
		sum := calc1(tmp.num)
		retObj := &result{
			tmp, sum,
		}
		res <- retObj
		//time.Sleep(time.Second) //全力消费
	}
}
func printResult(res chan *result) {
	for ret := range res {
		fmt.Println(ret.id, ret.num, ret.sum)
	}
}
