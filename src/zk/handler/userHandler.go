package handler

import (
	"net/http"
	"log"
	"zk/util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("====> path is : ",r.URL.Path," ====>")
	log.Println(r.Form)
	mString := util.Interface2JsonString(map[string]string{"success":"ok"})
	w.Write([]byte(mString))

}