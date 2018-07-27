package util

import (
	"log"
)
type Response struct {
	Code	int  `json:"code" bson:"code"`
	Msg		string	`json:"msg" bson:"msg"`
	Data	map[string]interface{} `json:"data" bson:"data"`
}

func PrintErr(err error)  {
	if err != nil {
		log.Panic(err)
	}
}