package main

import (
	"net/http"

	//外部包和内部包之间最好空一行
	"github.com/gin-gonic/gin"
)

func main() {
	router:=gin.Default()
	//先加载所有模板文件
	router.LoadHTMLGlob("templates/**")//多及目录这样写templates/**/* //也可以列举router.LoadHTMLFiles("1.html","2.html")
	//加载静态文件 比如引用图片等文件
	router.Static("/sta","./statics") //第一个参数是代码中使用的 第二个参数是真实路径
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":"DogLI",
			"Age":23,
		})
	})
	//多及目录需要在模板里定义如 templates/home/index.html开头定义define "home/index.html" 末尾要加end
	//router.GET("/index/test", func(c *gin.Context) {
	//	c.HTML(http.StatusOK,"home/index.html",gin.H{
	//		"Name":"DogLI",
	//		"Age":23,
	//	})
	//})

	//路由分组
	shoppingGroup := router.Group("/shopping") //以shopping开头的统一管理
	shoppingGroup.GET("/index", func(c *gin.Context) {  //访问路径/shopping/index
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"msg":"shopping/index success",
		})
	})
	shoppingGroup.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"msg":"shopping/home success",
		})
	})
	//继续分层 多层可能效率较低
	sv1 := shoppingGroup.Group("/v1")
	{
		sv1.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"code":0,
				"msg":"shopping/v1/index success!",
			})
		})
		sv1.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"code":0,
				"msg":"shopping/v1/home success",
			})
		})
	}
	//路由原理 httprouter库：构造前缀树而不是反射,相同前缀的接口分到一起比如/search和/set前缀相同检索s更快单独的分为一组

	router.Run(":9091")
}
