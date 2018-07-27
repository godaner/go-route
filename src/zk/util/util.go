package util

import (
	"encoding/json"
	"net/http"
	"log"
)
type Response struct {
	Code	int  `json:"code" bson:"code"`
	Msg		string	`json:"msg" bson:"msg"`
	Data	map[string]interface{} `json:"data" bson:"data"`
}
func (r Response) Response2(w http.ResponseWriter){
	WriteJsonStr(w,r)
}

func Interface2JsonString(obj interface{})string{
	json,_ :=json.Marshal(obj)
	mString :=string(json)
	return mString
}

func WriteJsonStr(w http.ResponseWriter,obj interface{}){
	s := Interface2JsonString(obj)
	w.Write([]byte(s))
}

func PrintErr(err error)  {
	if err != nil {
		log.Panic(err)
	}
}