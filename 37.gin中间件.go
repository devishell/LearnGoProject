package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//中间件功能 （类似拦截器？）所有接口比如进行验证或者其他中间操作

	r := gin.Default()
	//1.单独使用 放在返回函数之前 可对分组使用
	r.GET("/hello", castTime, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello ok",
		})
	})
	shopping := r.Group("/shopping",castTime)
	//2.对所有的使用 可对分组使用
	r.Use(castTime)
	shopping.Use(castTime)
	//进入gin.Default()内部查看 发现加载了loger中间件 engine := New() engine.Use(Logger(), Recovery())

	//数据绑定
	r.POST("/login", func(c *gin.Context) {
		type User struct {
			Name string `form:"username" binding:"required"`
			Pwd string `form:"password" json:"password" `
		}
		var userinfo User
		err := c.ShouldBind(&userinfo) //根据content-type的类型解析到userinfo中
		//c.ShouldBindJSON();c.ShouldBindXML()
		//c.BindJSON();c.BindXML() must bind
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"msg":"login error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "login ok",
		})
	})

	r.Run(":9092")
}

func castTime(c *gin.Context) {
	start := time.Now()
	c.Next() //运行下一个handlefunc函数
	//统计耗时
	cast := time.Since(start)
	fmt.Println("cast time:", cast)
}
