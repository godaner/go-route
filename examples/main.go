package main

import (
	"go-route/examples/handler"
	"go-route/examples/httpSessions"
	"go-route/route"
	"github.com/astaxie/beego/session"
)
const (
	ADDR=":80"
)

func initHttpSessions(){
	httpSessions.HttpSessions , _ = session.NewManager("memory", &session.ManagerConfig{EnableSetCookie: true,CookieName:"gosessionid",Gclifetime:3600})
	go httpSessions.HttpSessions.GC()
}

func init(){
	initHttpSessions()
}

func main() {
	//regist routes first
	handler.RegistRoutes()
	//then , start server
	route.Start(ADDR)
}

