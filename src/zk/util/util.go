package util

import "encoding/json"

func Interface2JsonString(m interface{})string{
	json,_ :=json.Marshal(m)
	mString :=string(json)
	return mString
}
