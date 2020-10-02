package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func main() {
	//回顾

	//GO写测试，代码检测sonar，如果对代码有要求

	//物理层：传输电信号 数据链路层：以太网协议解析01电信号，网卡mac地址 网络层：网络地址ip，ip数据包 传输层：根据mac和ip链接 应用层：解包数据(封包->解包)

	//粘包 用协议解决？(定义协议 加上包头封包，解包得到长度再解析)
	//conn.Write(Encode("hello client"))
	//Decode(bufio.NewReader(coon))


}
func Encode(msg string)([]byte ,error){
	var length = int32(len(msg))//读取消息长度 转换为int32类型占4个字节
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length) //大端：位数高的在前地址增大方向 小端：一般x86和arm是小端，c51是大端
	if err != nil {
		return pkg.Bytes(),err
	}
	//写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return pkg.Bytes(),err
	}
	return pkg.Bytes(),nil
}
func Decode(reader*bufio.Reader)(string,error){
	//读取消息长度
	lengthBytes,_:=reader.Peek(4)
	lengthBuff:=bytes.NewBuffer(lengthBytes)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//buffered返回缓冲区现有可读的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的消息数据
	pack:=make([]byte,int(length+4))
	_, err = reader.Read(pack)
	if err!=nil {
		return "", err
	}
	return string(pack[4:]), nil
}
