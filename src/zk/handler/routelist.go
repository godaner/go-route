package handler

import (
	"zk/route"
)

func Routes() route.Router{
	return route.RegistRoutes(
		route.MakePostRoute("/login",LoginHandler),
		route.MakePostRoute("/logout",LogoutHandler),
		route.MakeGetRoute("/onlineUser",GetOnlineUserHandler),
		route.MakePostRoute("/regist",RegistHandler))
}

func StaticRoutes() route.StaticRouter{
	return route.RegistStaticRoutes(
		route.MakeStaticRoute("/html/","template"),
		route.MakeStaticRoute("/css/","template"),
		route.MakeStaticRoute("/js/","template"))
}
