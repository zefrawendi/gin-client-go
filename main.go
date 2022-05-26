package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zefrawendi/gin-client-go/pkg/router"
	"github.com/zefrawendi/gin-client-go/pkg/utils"
	"k8s.io/klog"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	config, err := config.LoadConfig(".")
	if err != nil {
		klog.Fatal("Failed to load config: %v", err)
	}

	router.InitRouter(engine)

	err = engine.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
	if err != nil {
		klog.Errorf("Failed to run gin server: %v", err)
		return
	}
}
