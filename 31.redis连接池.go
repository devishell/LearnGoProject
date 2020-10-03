package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//安装redis: %brew install redis 启动redis服务端% redis-server 服务端口6379
	//ting@TingdeMacBook-Pro LearnGoProject % go get github.com/gomodule/redigo/redis
	pool:=&redis.Pool{
		MaxIdle: 16,
		MaxActive: 0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}
	rc:=pool.Get()
	defer rc.Close()
	_, err := rc.Do("Set", "a", 100)
	if err!=nil{
		fmt.Println(err)
		return
	}
	a, err := redis.Int(rc.Do("Get", "a"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("从redis获取到 a =",a)

}
