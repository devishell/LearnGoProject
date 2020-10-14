package main

func main() {
	//前后端不分离 需要提前弄好html模板等文件再开发 mvc架构
	//前后端分离之后 前端调用接口处理界面逻辑，各司其职 （感觉前端有点类似于客户端开发了）mvvm架构 可以对接网页app小程序跨终端

	//rest与技术无关 是软件架构风格
	//无restful接口设计：
	//获取书籍 get /get_book
	//创建、更新、删除书籍 post /create_book、/update_book、/delete_book

	//应用restful设计 就是利用http的方法代替这些动作
	//get /book
	//创建post /book 更新put /book 删除delete /book

	//postman调试post接口
}
