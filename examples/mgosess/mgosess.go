package mgosess
import (

	"gopkg.in/mgo.v2"
)

const(
	URL="192.168.2.184:27017"
	DB="gowebdemo"
)
func OpenSession() *mgo.Session{
	session,err:=mgo.Dial(URL)
	if err != nil { panic(err) }
	session.SetMode(mgo.Monotonic, true)
	return session
}