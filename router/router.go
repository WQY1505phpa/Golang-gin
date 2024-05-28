package router

import (
	"github.com/gin-gonic/gin"
	"wqy/controller"
)

func Router() *gin.Engine {
  r:= gin.Default()
  api := r.Group("/space")
  {
	  api.GET("/getSpaceList",controller.GetSpaceList)
	  api.POST("/addSpace",controller.AddSpace)
	  api.POST("/updateSpace",controller.UpdateSpace)
	  api.POST("/deleteSpace",controller.DeleteSpace)
  }
  return r
}
