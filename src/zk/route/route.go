package route

import (
	"net/http"
)
//Route
type Route struct {
	Path string
	Handler func (w http.ResponseWriter, r *http.Request)
}


func MakeRoute(path string, handler func (w http.ResponseWriter, r *http.Request)) Route{
	return Route{
		Path:path,
		Handler:handler,
	}
}

func RegistRoutes(routes...Route){
	for _,route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}
}
//StaticRoute



type StaticRoute struct {
	Path string
	Dir string
}


func MakeStaticRoute(path string, dir string) StaticRoute{
	return StaticRoute{
		Path:path,
		Dir:dir,
	}
}

func RegistStaticRoutes(staticRoutes...StaticRoute){
	for _,staticRoute := range staticRoutes {
		http.Handle(staticRoute.Path,http.FileServer(http.Dir(staticRoute.Dir)))
	}
}


