package main

import (
	"zk/route"
	"zk/handler"
	"net/http"
	"log"
	"github.com/astaxie/beego/session"
	"zk/httpSessions"
)


const(
	ADDR=":9090"
)


func StartServer(){

	//static source routes
	staticRouter := handler.StaticRoutes()

	//dynamic routes
	router := handler.Routes()

	//run server
	runServer(router,staticRouter)
}
func runServer(router route.Router, staticRouter route.StaticRouter) {

	err := http.ListenAndServe(ADDR,route.GetDispatcherRouter(router,staticRouter))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initHttpSessions(){
	httpSessions.HttpSessions , _ = session.NewManager("memory", &session.ManagerConfig{EnableSetCookie: true,CookieName:"gosessionid",Gclifetime:3600})
	go httpSessions.HttpSessions.GC()
}

func init(){
	initHttpSessions()
}


func main() {
	StartServer()
}

