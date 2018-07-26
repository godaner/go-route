package server

import (
	"net/http"
	"log"
	"zk/handler"
)

const(
	ADDR=":9090"
)

func StartServer(){

	handlerMapping :=getHandlerMapping()

	registHandler(handlerMapping)

	fileServerMapping :=getFileServerMappingMapping()

	registFileServer(fileServerMapping)

	runServer()
}

func runServer() {
	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func getFileServerMappingMapping() map[string]http.Handler {

	fileServerMapping := map[string]http.Handler{}

	fileServerMapping["/css/"] = http.FileServer(http.Dir("template"))

	fileServerMapping["/js/"] = http.FileServer(http.Dir("template"))

	fileServerMapping["/html/"] = http.FileServer(http.Dir("template"))

	return fileServerMapping

}
func registFileServer(fileServerMapping map[string]http.Handler) {
	for k,v := range fileServerMapping{
		http.Handle(k, v)
	}
}
func getHandlerMapping() map[string]func(http.ResponseWriter, *http.Request){

	handlerMapping := map[string]func(http.ResponseWriter, *http.Request){}

	handlerMapping["/login"] = handler.LoginHandler

	return handlerMapping
}
func registHandler(handlerMapping map[string]func(http.ResponseWriter, *http.Request)) {
	for k,v := range handlerMapping{
		http.HandleFunc(k, v)
	}

}