package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//互联网基础是TCP、UDP --> socket抽象层 --> 应用层

	//服务端监听端口
	go func() {
		listen, err2 := net.Listen("tcp", "127.0.0.1:20000")
		//客户端连接服务器
		if err2 != nil {
			fmt.Println("监听失败", err2)
		} else {
			defer listen.Close()
			for {
				con, err3 := listen.Accept() //没有链接会阻塞等待连接
				if err3 == nil {
					go process(con) //处理这个连接后重新等待另外的连接
					con.Write([]byte("hello client！"))
				} else {
					fmt.Println(err3)
				}
			}
		}
	}()

	//客户端连接服务器
	go func() {
		time.Sleep(time.Second/10)
		conn, err := net.Dial("tcp", "127.0.0.1:20000")
		if err == nil {
			fmt.Println("客户端连接服务器成功了")
			for i := 3; i > 0; i-- {
				write, err1 := conn.Write([]byte("hello server!\n")) //可以换成控制台输入
				if err1 != nil {
					fmt.Println(write, err1)
				}
			}
			go func() {	//也可以监听服务器消息
				defer conn.Close()
				var buf [1024]byte
				read, err := conn.Read(buf[:]) //这里实际应该循环读取判断EOF等
				if err == nil {
					fmt.Println("服务器的消息：", string(buf[:read]))
				} else {
					fmt.Println("服务器的消息接收失败了", err)
				}
			}()
		}else {
			fmt.Println("客户端连接服务器失败了",err)
		}
	}()

	//UDP 服务端
	go func() {
		listenUDP, err5 := net.ListenUDP("udp", &net.UDPAddr{
			IP:net.ParseIP("127.0.0.1"),
			Port: 21000,
		})
		if err5==nil {
			go func() {
				defer listenUDP.Close()
				//监听接收消息
				for {
					var buf [1024]byte
					read,addr, err := listenUDP.ReadFromUDP(buf[:]) //这里实际应该循环读取判断EOF等
					if err == nil {
						fmt.Println("客户端的消息：", string(buf[:read]))
						listenUDP.WriteToUDP([]byte("hello udp client"),addr)//发送
					} else {
						fmt.Println("接收失败了", err)
					}
				}
			}()
		}else {
			fmt.Println("启动udp server出错了",err5)
		}
	}()
	//udp客户端
	go func() {
		time.Sleep(time.Second/10)
		conn, err := net.Dial("udp", "127.0.0.1:21000")
		if err == nil {
			fmt.Println("udp客户端连接服务器成功了")
			for i := 3; i > 0; i-- {
				write, err1 := conn.Write([]byte("hello udp server!\n")) //可以换成控制台输入
				if err1 != nil {
					fmt.Println(write, err1)
				}
			}
			go func() {	//也可以监听服务器消息
				defer conn.Close()
				var buf [1024]byte
				read, err := conn.Read(buf[:]) //这里实际应该循环读取判断EOF等
				if err == nil {
					fmt.Println("服务器的消息：", string(buf[:read]))
				} else {
					fmt.Println("服务器的消息接收失败了", err)
				}
			}()
		}else {
			fmt.Println("客户端连接服务器失败了",err)
		}
	}()
	time.Sleep(time.Second)
}
func process(con net.Conn) {
	defer con.Close() //最好在go里面关闭
	//接收连接的数据
	fmt.Println("正在监听")
	var buf [1024]byte
	read, err := con.Read(buf[:]) //这里实际应该循环读取判断EOF等
	if err == nil {
		fmt.Println("客户端的消息：", string(buf[:read]))
	} else {
		fmt.Println("接收失败了", err)
	}
}
