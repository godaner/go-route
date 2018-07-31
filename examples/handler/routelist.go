package handler

import (
	"go-route/route"
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
		route.MakeStaticRoute("/html/","examples/template"),
		route.MakeStaticRoute("/css/","examples/template"),
		route.MakeStaticRoute("/js/","examples/template"))
}
