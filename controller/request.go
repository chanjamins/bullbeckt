package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxUserIDKey = "userID"

var ErrWithouLogin = errors.New("用户未登录..")

func getCurrentUerID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrWithouLogin
		return
	}
	//强制转换
	userID, ok = uid.(int64)
	if !ok {
		err = ErrWithouLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	// 在路径中获取 Get在请求头
	pageIndex := c.Query("page")

	pageSize := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageIndex, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
