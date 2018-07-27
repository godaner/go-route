package httpsession

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)
var HttpSessions *session.Manager

//init session
func InitSession() {
	HttpSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go HttpSessions.GC()
}