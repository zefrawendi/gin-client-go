package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zefrawendi/gin-client-go/pkg/service"
)

func GetNamespaces(ctx *gin.Context) {
	namespaces, err := service.GetNamespaces()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			http.StatusText(http.StatusInternalServerError): err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, namespaces)
}
