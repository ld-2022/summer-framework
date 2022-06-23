package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	routeList []Route
)

type Route struct {
	Method, Path string
	Handler      []gin.HandlerFunc
}

func AddRoute(route Route) {
	routeList = append(routeList, route)
}

func Add(method, path string, handler []gin.HandlerFunc) {
	AddRoute(Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func AddGet(path string, handler ...gin.HandlerFunc) {
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    path,
		Handler: handler,
	})
}

func AddPost(path string, handler []gin.HandlerFunc) {
	AddRoute(Route{
		Method:  http.MethodPost,
		Path:    path,
		Handler: handler,
	})
}
