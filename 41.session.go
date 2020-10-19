package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"
)

type User3 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var mgr *SessionMgrMem
var mgr2 *SessionMgrRedis

func init() {
	mgr = &SessionMgrMem{
		Session: make(map[string]*SessionData, 1024),
	}
	mgr2 = &SessionMgrRedis{
		Session: make(map[string]*SessionData,1024),client: &redis.Pool{
			MaxIdle: 16,
			MaxActive: 0,
			IdleTimeout: 300,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp","127.0.0.1:6379")
			},
		},
	}
}
func main() {
	//cookie可实现简单的保持状态需求 但是保存在浏览器本地端，可能被拦截，私密信息不能存储，且最大支持4k字节
	//session由此出现 存储更多字节 保存在服务器更安全，cookie用来保存sessionid作为服务端识别客户端的中间桥梁
	//1.服务端存储session位置：内存 关系型数据库 redis 文件(少),相对安全
	//2.必须依赖cookie才能使用
	//3.相比cookie需要后端维护session服务

	//gin框架的session扩展
	//流程：无session或已经失效(服务器查不到该sessionid)或session中isLogin为false 跳转登录页面 否则正常响应

	//go语言开发相比于Java、Python等更注重底层而不是框架 所以熟悉cookie和session的原理实现很重要
	//memory版session服务 仅供参考不能在生产使用断电就没了
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(SessionMiddleware(mgr))
	r.Any("/login", func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var user User3
			err := c.ShouldBind(&user)
			if err != nil {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "必须输入用户名和密码！",
				})
				return
			}
			if user.Username == "dwx" && user.Password == "123" {
				session_v, ok := c.Get(SessionContextName)
				if ok {
					session := session_v.(*SessionData)
					session.Set("isLogin", true)
				}
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
	r.GET("/home", authMiddleWare, func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	r.GET("/vip", authMiddleWare, func(c *gin.Context) {
		c.HTML(http.StatusOK, "vip.html", nil)
	}) //访问vip之前先进行cookie判断 且可以保存数据到上下文传递数据
	r.Run(":9094")
}

//每个用户的session数据
type SessionData struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex //锁Data
	//ExpireTime//过期时间
}

func NewSessionData(id string) *SessionData {
	return &SessionData{
		ID:   id,
		Data: make(map[string]interface{}, 2),
	}
}

//全局的session管理
type SessionMgrMem struct {
	Session map[string]*SessionData
	rwLock  sync.RWMutex
}

//redis版SessionMgr
type SessionMgrRedis struct {
	Session map[string]*SessionData
	rwLock  sync.RWMutex
	client *redis.Pool
}
type SessionDataRedis struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex //锁Data
	expire int//过期时间
	client * redis.Pool
}

func (s *SessionData) Get(key string) (value interface{}, err error) {
	//获取读锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}
	return
}
func (s *SessionData) Set(key string, value interface{}) {
	//获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data[key] = value
}
func (s *SessionData) Del(key string) {
	//获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data, key)
}
func (s *SessionMgrMem) GetSessionData(id string) (data *SessionData, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	data, ok := s.Session[id]
	if !ok {
		err = fmt.Errorf("session id not exisit")
		return
	}
	return
}

//创建session数据
func (s *SessionMgrMem) CreateSession() *SessionData {
	id := uuid.NewV4() //防止重复使用uuid go get github.com/satori/go.uuid/go get github.com/gofrs/uuid ,也可以使用时间戳
	sessionData := NewSessionData(id.String())
	s.Session[sessionData.ID] = sessionData
	return sessionData
}

const (
	SessionCookieName  = "session_id" //cookie中的key
	SessionContextName = "session"    //context中的key
)

//全局session中间件
func SessionMiddleware(s *SessionMgrMem) gin.HandlerFunc {
	if s == nil {
		panic("should init SessionMgrMem first!")
	}
	//取cookie-->找session->取sessiondata->传递sessiondata(利用gin的context保存而不需要重复这一长串解析过程)
	return func(c *gin.Context) {
		var session *SessionData
		sessionid, err := c.Cookie(SessionCookieName)
		if err != nil {
			session = s.CreateSession() //Cookie中不存在sessionid创建新的
			sessionid = session.ID
		} else {
			session, err = s.GetSessionData(sessionid)
			if err != nil {
				session = s.CreateSession() //SessionMgr中不存在对应id的session创建新的
				sessionid = session.ID
			}
		}
		//传递给上下文
		c.Set(SessionContextName, session)
		//设置cookie 必须在下个函数返回之前设置
		c.SetCookie(SessionCookieName, sessionid, 3600, "/", "127.0.0.1", false, true)
		c.Next()
	}
}
func authMiddleWare(c *gin.Context) {
	session_v, _ := c.Get(SessionContextName)
	sd := session_v.(*SessionData)
	isLogin_v, err := sd.Get("isLogin")
	if err != nil {
		c.Redirect(http.StatusFound, "/login?next="+c.Request.URL.Path)
		return
	}
	isLogin, ok := isLogin_v.(bool)
	if !ok {
		c.Redirect(http.StatusFound, "/login?next="+c.Request.URL.Path)
		return
	}
	if !isLogin {
		c.Redirect(http.StatusFound, "/login?next="+c.Request.URL.Path)
		return
	}
	c.Next()
}
//redis实现
func (s *SessionMgrRedis) GetSessionData(id string) (data *SessionData, err error) {
	if s.Session == nil {
		conn := s.client.Get()
		defer conn.Close()
		value, err := redis.String(conn.Do("Get", id)) //转化为string否则无法将interface{}强转为[]byte
		if err == nil {
			err := json.Unmarshal([]byte(value), &s.Session)
			if err != nil {
				return nil, err
			}
		}else {
			return nil, err
		}
	}
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	data, ok := s.Session[id]
	if !ok {
		err = fmt.Errorf("session id not exisit")
		return
	}
	return
}
func (s *SessionDataRedis)save(){
	bytes, err := json.Marshal(s.Data)
	if err != nil {
		fmt.Println("narshal session data failed,err:",err)
		return
	}
	conn := s.client.Get()
	defer conn.Close()
	conn.Do("Set",s.ID,bytes,time.Second*time.Duration(s.expire))
}

//分布式全局id生成 go get github.com/sony/sonyflake

//注意问题 go中json序列化空接口存储的数字类型(整形、浮点型等)全部为float64类型