package handler

import (
	"net/http"
	"strings"
	"log"
	"zk/util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	log.Println("====> path is : %s ====>",r.URL.Path)
	log.Println(r.Form)
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}
	mString := util.Interface2JsonString(map[string]string{"success":"ok"})
	w.Write([]byte(mString))

}