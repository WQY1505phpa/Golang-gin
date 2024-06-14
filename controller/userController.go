package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(*gin.Context) {
	fmt.Print("获取用户信息")
}

