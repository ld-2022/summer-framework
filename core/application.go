package core

import (
	"embed"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Application struct {
	Host  string
	Port  int
	ResFs embed.FS
}

func (a *Application) Start() (err error) {
	e := gin.Default()
	for _, route := range routeList {
		e.Handle(route.Method, route.Path, route.Handler...)
	}
	e.LoadHTMLGlob("../resources/template/*")
	err = e.Run(strings.Join([]string{a.Host, strconv.Itoa(a.Port)}, ":"))
	return
}
