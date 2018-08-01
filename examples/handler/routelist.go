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

