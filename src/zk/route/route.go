package route

import (
	"net/http"
	"html/template"
)
//Router

type Router struct {
	Routes []Route
}

type Route struct {
	Path string
	Method string //POST.GET.DELETE,PUT etc.
	Router func (w RouteResponse, r RouteRequest)
}
const (
	POST="POST"
	GET="GET"
	PUT="PUT"
	DELETE="DELETE"
)


func MakeGetRoute(path string, router func (response RouteResponse, request RouteRequest)) Route{
	return Route{
		Method:GET,
		Path:path,
		Router:router,
	}
}
func MakePostRoute(path string, router func (response RouteResponse, request RouteRequest)) Route{
	return Route{
		Method:POST,
		Path:path,
		Router:router,
	}
}
func MakeDeleteRoute(path string, router func (response RouteResponse, request RouteRequest)) Route{
	return Route{
		Method:DELETE,
		Path:path,
		Router:router,
	}
}
func MakePutRoute(path string, router func (response RouteResponse, request RouteRequest)) Route{
	return Route{
		Method:PUT,
		Path:path,
		Router:router,
	}
}

func RegistRoutes(rs...Route) Router{
	return Router{rs}
}

//StaticRouter
type StaticRoute struct {
	Path string
	Dir string
}

type StaticRouter struct {
	StaticRoutes []StaticRoute
}

func (staticRoute StaticRoute) Router(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles(staticRoute.Dir+request.URL.Path)
	if err == nil {
		t.Execute(writer,nil)
	}
}

func MakeStaticRoute(path string, dir string) StaticRoute{
	return StaticRoute{
		Path:path,
		Dir:dir,
	}
}

func RegistStaticRoutes(srs...StaticRoute) StaticRouter {
	return StaticRouter{srs}
}


