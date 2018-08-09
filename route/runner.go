package route

import (
	"net/http"
	"log"
)

func Start(addr string){
	err := http.ListenAndServe(addr,DispatcherRouter{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
