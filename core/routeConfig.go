package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	routeConfig()
}

func routeConfig() {
	AddGet("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK, "test.html", nil)
	})

}
