package server

import (
	"zk/router"
	"zk/route"
)

func ListAndRegistRoutes() route.Router{
	return route.RegistRoutes(
		route.MakePostRoute("/login",userHandler.LoginRouter),
		route.MakePostRoute("/logout",userHandler.LogoutRouter),
		route.MakeGetRoute("/onlineUser",userHandler.GetOnlineUserRouter),
		route.MakePostRoute("/regist",userHandler.RegistRouter))
}

func ListAndRegistStaticRoutes() route.StaticRouter{
	return route.RegistStaticRoutes(
		route.MakeStaticRoute("/html/","template"),
		route.MakeStaticRoute("/css/","template"),
		route.MakeStaticRoute("/js/","template"))
}
