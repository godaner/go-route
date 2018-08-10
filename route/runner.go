package route

import (
	"net/http"
	"go-util/errorutil"
	"log"
)

func Start(addr string){
	log.Println("Server api is : ", Routes)
	log.Println("Server will start at : [ ", addr," ]")
	err := http.ListenAndServe(addr,DispatcherRouter{})
	errorutil.CheckErr(err)
}
