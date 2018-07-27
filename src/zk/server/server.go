package server

import (
	"net/http"
	"log"
	"zk/httpsession"
)

const(
	ADDR=":9090"
)

func StartServer(){

	ListAndRegistRoutes()

	ListAndRegistStaticRoutes()

	httpsession.InitSession()

	runServer()
}

func runServer() {
	err := http.ListenAndServe(ADDR,nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}