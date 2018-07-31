package route

import (
	"net/http"
	"strings"
	"log"
)

type DispatcherRouter struct {
	Router Router
	StaticRouter StaticRouter

}
func GetDispatcherRouter(router Router,staticRouter StaticRouter)DispatcherRouter{
	return DispatcherRouter{
		Router:router,
		StaticRouter:staticRouter,
	}
}

func (dispatcherHandler DispatcherRouter)ServeHTTP(w http.ResponseWriter,r *http.Request){

	prepareRequest(r)

	//start route
	finded := false
	for _,route:= range dispatcherHandler.Router.Routes {
		if matchRoute(route,r) {//match request path and request method:post get delete put etc.
			route.Router(makeRouteResponse(w),makeRouteRequest(r))
			finded = true
			break
		}
	}
	if finded {
		return
	}
	for _,staticRoute:= range dispatcherHandler.StaticRouter.StaticRoutes {

		if matchStaticRoute(staticRoute,r) {//match static source
			staticRoute.Router(w,r)
			break
		}
	}
}
func prepareRequest(r *http.Request) {
	//parse form
	r.ParseForm()
	//log
	log.Println("====> path is : ",r.URL.Path," ====>")
	log.Println(r)
	log.Println(r.Form)
}

func matchStaticRoute(staticRoute StaticRoute, request *http.Request) bool {
	//request /html/login.html
	//staticRoute /html/ template
	b0:=strings.Index(request.URL.Path,staticRoute.Path)==0

	return b0 && request.Method==GET
}
func matchRoute( route Route,r *http.Request) bool{
	return r.URL.Path==route.Path && r.Method == route.Method
}
func makeRouteRequest(request *http.Request) RouteRequest {

	//get form param
	params:=map[string]interface{}{}

	for k,v :=range request.Form {
		if len(v)==1 {
			params[k] = v[0]
		}else{
			params[k] = v
		}
	}
	return RouteRequest{
		Request:request,
		Params:params,
	}
}
func makeRouteResponse(writer http.ResponseWriter) RouteResponse {
	return RouteResponse{writer}
}