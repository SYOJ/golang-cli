package initialize

import (
	"github.com/gin-gonic/gin"
	"golang-cli/web/middleware"
	"golang-cli/web/router"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.New()
	Router.Use(GinLogger(), GinRecovery(true), middleware.Cors())

	ApiRouter := Router.Group("/uapply/v1")
	router.InitUserRouter(ApiRouter)
	router.InitAdminRouter(ApiRouter)
	return Router
}
