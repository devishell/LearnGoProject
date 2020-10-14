package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User1 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	//cookie示例
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var user User1
			err := c.ShouldBind(&user)
			if err != nil {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "必须输入用户名和密码！",
				})
				return
			}
			if user.Username == "dwx" && user.Password == "123" {
				//登陆成功 设置cookie 访问的地址必须和cookie设置的domain一致否则不生效，获取不到cookie（神坑）
				c.SetCookie("username", user.Username, 20, "/", "localhost", false, true)
				c.Redirect(http.StatusFound, "/home")
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "用户名或密码错误！",
				})
			}
		} else if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "login.html", nil)
		} else {

		}
	})
	r.GET("/home", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "home.html", nil)
		//判断cookie
		cookie1, err := c.Cookie("username")
		if err == nil {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"username": cookie1,
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
	})
	//没有cokkie时没登录也能看home页面 必须判断是否登陆了 以及页面变化后前面页面记录的数据传递
	//cookie保持状态 浏览器本地的键值对 下次请求会携带cookie信息（不能跨域获取其它网站信息和调用其它网站接口）
	_ = http.Cookie{} //查看结构 一般设置maxage相对时间 而不是expire

	r.Run(":9093")
}
