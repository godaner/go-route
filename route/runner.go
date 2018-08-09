package route

import (
	"net/http"
	"log"
)

func Start(){
	err := http.ListenAndServe(conf.Port,DispatcherRouter{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func init(){
	LoadConfiguration()
}