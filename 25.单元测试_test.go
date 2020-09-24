package main

import (
	"fmt"
	"testing"
)

//网上资源 go标准库文档中文 studygo
//后缀_test.go的文件不会被go build编译到可执行文件中
//_test.go文件中有三中类型函数 前缀如下
//1.Test 单元测试
//2.Benchmark 基准测试 //方法左边可以多重模式运行测试性能
//3.Example 示例函数
//go test驱动程序,会遍历所有这些文件函数执行
func TestNewDog(t *testing.T) {
	t.Log("TestNewDog")
	//t.Fail()
	t.Fatalf("出错了原因：%v\n",111)
}
func BenchmarkNewDog(b *testing.B) {
	b.Log("BenchmarkNewDog")
}
func ExampleNewDog() {
	fmt.Println("ExampleNewDog")
}
//ting@TingdeMacBook-Pro LearnGoProject % go test 24.并发的日志.go
//?       command-line-arguments  [no test files]
//ting@TingdeMacBook-Pro LearnGoProject % go test 25.单元测试_test.go
//ok      command-line-arguments  0.424s
//ting@TingdeMacBook-Pro LearnGoProject % go test 25.单元测试_test.go
//--- FAIL: TestNewDog (0.00s)
//25.单元测试_test.go:16: TestNewDog
//25.单元测试_test.go:18: 出错了原因：
//FAIL
//FAIL    command-line-arguments  3.065s
//FAIL
//ting@TingdeMacBook-Pro LearnGoProject %

//整个测试之前做的事 setup teardown
func SetupTetscase(t *testing.T)  {

}
func TestMain(t *testing.M)  {

}

