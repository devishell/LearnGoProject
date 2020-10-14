package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

func main() {
	//cookie可实现简单的保持状态需求 但是保存在浏览器本地端，可能被拦截，私密信息不能存储，且最大支持4k字节
	//session由此出现 存储更多字节 保存在服务器更安全，cookie用来保存sessionid作为服务端识别客户端的中间桥梁
	//1.服务端存储session位置：内存 关系型数据库 redis 文件(少),相对安全
	//2.必须依赖cookie才能使用
	//3.相比cookie需要后端维护session服务

	//gin框架的session扩展
	//流程：无session或已经失效(服务器查不到该sessionid)或session中isLogin为false 跳转登录页面 否则正常响应


}
//每个用户的session数据
type SessionData struct {
	ID string
	Data map[string]interface{}
	rwLock sync.RWMutex//锁Data
	//ExpireTime//过期时间
}
func NewSessionData(id string)*SessionData{
	return &SessionData{
		ID:id,
		Data: make(map[string]interface{},2),
	}
}
//全局的session管理
type SessionMgr struct {
	session map[string]SessionData
	rwLock sync.RWMutex
}
func (s *SessionData)Get(key string)(value interface{},err error){
	//获取读锁
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value,ok := s.Data[key]
	if  !ok{
		err = fmt.Errorf("invalid key")
		return
	}
	return
}
func (s *SessionData)Set(key string,value interface{}){
	//获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data[key] = value
}
func (s *SessionData)Del(key string){
	//获取写锁
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data,key)
}
func (s *SessionMgr)GetSessionData(id string)(data SessionData,err error){
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	data ,ok :=s.session[id]
	if !ok {
		err = fmt.Errorf("session id not exisit")
		return
	}
	return
}
//创建session数据
func (s *SessionMgr)CreateSession()*SessionData{
	id := uuid.NewV4() //防止重复使用uuid go get github.com/satori/go.uuid ,也可以使用时间戳
	return NewSessionData(id.String())
}