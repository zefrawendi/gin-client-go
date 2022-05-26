package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zefrawendi/gin-client-go/pkg/service"
)

func GetPods(ctx *gin.Context) {
	pods, err := service.GetPods(ctx.Param("namespace"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			http.StatusText(http.StatusInternalServerError): err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, pods)
}
