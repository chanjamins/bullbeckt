package router

import "github.com/gin-gonic/gin"

func SetRouter(mode string) *gin.Engine {
	//设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
    r.Use(logger.)
}
