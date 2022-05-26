package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zefrawendi/gin-client-go/pkg/apis"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", apis.Ping)
	r.GET("/namespaces", apis.GetNamespaces)
	r.GET("/namespaces/:namespace/pods", apis.GetPods)
}
