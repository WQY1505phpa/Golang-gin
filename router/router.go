package router

import (
	"github.com/gin-gonic/gin"
	"wqy/controller"
)

func Router() *gin.Engine {
  r:= gin.Default()
  api := r.Group("/space")
  {
	  api.GET("/getSpaceInfo",controller.GetSpaceInfo)
	  api.GET("/getSpaceList",controller.GetSpaceList)
	  api.POST("/addSpace",controller.AddSpace)
	  api.POST("/updateSpace",controller.UpdateSpace)
	  api.POST("/deleteSpace",controller.DeleteSpace)
  }
  user := r.Group("/user")
  {
	  user.GET("/getUserInfo",controller.GetUserInfo)
	  user.GET("/getSpaceInfo",controller.GetSpaceInfo)
  }
  return r
}
