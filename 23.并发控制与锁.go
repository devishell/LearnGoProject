package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64 //全局变量x
var wg sync.WaitGroup
var lock sync.Mutex     //互斥锁 避免对公共资源x的修改不一致
var lockRW sync.RWMutex //读写互斥锁 主要为了提高"读比写多"的时候提高性能 读锁可以同时获取 但写锁互斥且写的时候读锁获取不了
func add() {
	for i := 100; i > 0; i-- {
		lock.Lock() //其他goroutine等待释放锁
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	//互斥锁
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	lockRW.RLock() //加读锁
	lockRW.RUnlock()//释放读锁

	//懒加载 延迟初始化 并发时保证只执行一次
	var loadOnce sync.Once
	var xx int
	f:=func() {
		xx = 2
		fmt.Println("xx=",xx)
	}
	for i:=3;i>0;i-- {
		loadOnce.Do(f) //xx=只会打印一次
	}

	//并发读写数据结构map,sync.Map可直接使用且封装了很多数据操作api
	map1:= sync.Map{}
	for i:=20;i>0;i-- {
		wg.Add(1)
		go func() {
			map1.Store("1",1)
			map1.Store(2,2)
			fmt.Println(map1.Load(2))
			fmt.Println(map1.Load("1"))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(map1.Load(2))
	fmt.Println(map1.Load(1))
	fmt.Println(map1.Load("1"))

	//定时器
	tick:= time.Tick(time.Second) //每秒都会执行的定时器
	for i := range tick { //相当于会定时往tick里面存值
		fmt.Println("time is ",i)
		//time.Sleep(time.Second*5)// 取值不一定立即取
	}

}
