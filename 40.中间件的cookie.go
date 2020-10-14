package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type User2 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
func main() {
	//记录没登录时的页面 登录后跳转回原先访问的页面
	//login？next=/vip 直接访问vipredirect到login页面 登陆成功后跳转回vip页面
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var user User2
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
				//跳转到原页面而不是固定的跳转页面所以要在当前接口（login）跳转的时候提前加上参数
				next := c.DefaultQuery("next", "/home")
				c.Redirect(http.StatusFound, next)
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
			//跳转时加上原页面的参数而不是直接跳转
			c.Redirect(http.StatusMovedPermanently, "/login?next="+c.Request.URL.Path)
			return
		}
	})
	r.GET("/vip", addCookieMiddleware, vipHandler) //访问vip之前先进行cookie判断 且可以保存数据到上下文传递数据
	r.Run(":9094")
}

func addCookieMiddleware(c *gin.Context) {
	cookie1, err := c.Cookie("username")
	if err == nil {
		c.Set("username", cookie1)
		c.Next()
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login?next="+c.Request.URL.Path)
		return
	}
}
func vipHandler(c *gin.Context) {
	temu, exists := c.Get("username")
	if exists {
		s := temu.(string)
		c.HTML(http.StatusOK, "vip.html", gin.H{
			"username": s,
		})
	} else {
		//跳转时加上原页面的参数而不是直接跳转
		c.Redirect(http.StatusMovedPermanently, "/login?next="+c.Request.URL.Path)
	}

}
