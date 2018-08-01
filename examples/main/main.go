package main

import (
	"go-route/route"
	"go-route/examples/handler"
	"net/http"
	"log"
	"go-route/examples/httpSessions"
	"github.com/astaxie/beego/session"
)


const(
	ADDR=":80"
)


func StartServer(){

	//routes
	router := handler.Routes()

	//run server
	runServer(router)
}
func runServer(router route.Router) {

	err := http.ListenAndServe(ADDR,route.GetDispatcherRouter(router))
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

