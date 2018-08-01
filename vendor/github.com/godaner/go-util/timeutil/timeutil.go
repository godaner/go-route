package timeutil

import (
	"strconv"
	"time"
)

var  t = time.Now()
func UnixStr()(string){
	return strconv.FormatInt(t.UTC().UnixNano(), 10)[:10]
}
func Unix() (int64){
	return t.Unix()
}
