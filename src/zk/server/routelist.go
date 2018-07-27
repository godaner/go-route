package server

import (
	"zk/router"
	"zk/route"
)

func ListAndRegistRoutes() route.Router{
	return route.RegistRoutes(
		route.MakePostRoute("/login",userHandler.LoginRouter),
		//route.MakeRoute("/logout",userHandler.Logout),
		//route.MakeRoute("/onlineUser",userHandler.GetOnlineUser),
		//route.MakeRoute("/regist",userHandler.RegistHandler))
	)
}

func ListAndRegistStaticRoutes() route.StaticRouter{
	return route.RegistStaticRoutes(
		route.MakeStaticRoute("/html/","template"),
		route.MakeStaticRoute("/css/","template"),
		route.MakeStaticRoute("/js/","template"))
}
