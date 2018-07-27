package server

import (
	"zk/httpsession"
	"zk/route"
	"log"
	"net/http"
)

const(
	ADDR=":9090"
)

func StartServer(){

	router:=ListAndRegistRoutes()

	staticRouter:=ListAndRegistStaticRoutes()

	httpsession.InitSession()

	runServer(router,staticRouter)
}
func runServer(router route.Router, staticRouter route.StaticRouter) {

	err := http.ListenAndServe(ADDR,route.GetDispatcherRouter(router,staticRouter))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
