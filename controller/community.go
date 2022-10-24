package controller

import (
	"bullbeckt/logic"
	"github.com/gin-gonic/gin"
)

func CommunityHandler(c *gin.Context) {
	//查询
	data, err := logic.CommunityList()
	if err != nil {
		return
	}
	ResponseSuccess(c, data)
}
