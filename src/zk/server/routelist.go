package server

import (
	"zk/handler"
	"zk/route"
)

func ListAndRegistRoutes(){
	route.RegistRoutes(route.MakeRoute("/login",userHandler.LoginHandler))
	route.RegistRoutes(route.MakeRoute("/logout",userHandler.Logout))
	route.RegistRoutes(route.MakeRoute("/onlineUser",userHandler.GetOnlineUser))
	route.RegistRoutes(route.MakeRoute("/regist",userHandler.RegistHandler))
}

func ListAndRegistStaticRoutes(){
	route.RegistStaticRoutes(route.MakeStaticRoute("/html/","template"))
	route.RegistStaticRoutes(route.MakeStaticRoute("/css/","template"))
	route.RegistStaticRoutes(route.MakeStaticRoute("/js/","template"))
}
