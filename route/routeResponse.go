package route

import (
	"net/http"
	"encoding/json"
)

type RouteResponse struct {
	ResponseWriter http.ResponseWriter
}
func (response RouteResponse ) WriteJsonStr(obj interface{}){
	json,_ :=json.Marshal(obj)
	jsonStr :=string(json)
	response.ResponseWriter.Write([]byte(jsonStr))
}