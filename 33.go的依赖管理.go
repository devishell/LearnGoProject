package main

func main() {
	//依赖管理 -- 方便更新依赖升级后的代码

	//godep依赖管理 如果项目目录有vendor目录会优先从中寻找依赖
	//go get github.com/tools/godep 已经Archived了 终端输入godep save命令生成godeps.json依赖管理文件 update/diff

	//dep官方实验性质依赖管理 已经Archived了 Dep was an official experiment to implement a package manager for Go.
	//As of 2020, Dep is deprecated and archived in favor of Go modules, which have had official support since Go 1.11.
	//For more details, see https://golang.org/ref/mod.
	//Go modules -- 官方依赖
	//终端输入go env/ go mod查看相关命令
//Usage:
//
//	go mod <command> [arguments]
//
//	The commands are:
//
//	download    download modules to local cache
//	edit        edit go.mod from tools or scripts
//	graph       print module requirement graph
//	init        initialize new module in current directory
//	tidy        add missing and remove unused modules
//	vendor      make vendored copy of dependencies
//	verify      verify dependencies have expected content
//	why         explain why packages or modules are needed

	//参考 https://blog.csdn.net/weixin_39003229/article/details/97638573
	//1.go mod init LearnGoProject //需要在GOPATH之外的项目目录下执行 //旧项目会从godeps下拷贝依赖require
	//2.go build、go get自动发现维护依赖
	//偏好设置可以搜索vgo启用go mod integration 并设置https://goproxy.cn代理网址
	//其他命令replace 替换包；download下载到gopath并缓存；edit命令行编辑依赖

	//导包报红问题
	//内部包
	//更换导包路径，原来可能使用的是./这种相对路径，换成{module name}/package1/package2的方式。
	//外部包
	//在root用户下能编译，在其他用户下却没法编译，于是考虑到了可能是权限问题
	//首先要知道gomod包都是下载到了$GOPATH/pkg/mod下面的，使用ls -l查看下报红的包文件夹的权限，发现是root用户的，其他用户没有权限，那就简单了，直接暴力解决：
	//chmod 777 mod -R
	//修改mod文件夹下所有文件夹的权限为所有用户所有组可读写。
}
