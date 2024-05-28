package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wqy/models"
	"wqy/service"
)
type Space struct {
	Name  string
	Attr  int
}
/**
获取单条空间信息
 */
func GetSpaceList(c *gin.Context) {
	SpaceId := c.DefaultQuery("spaceId","0")
	if SpaceId == "0" {
		c.JSON(http.StatusOK,gin.H{"error":"spaceId not is empty"})
		return
	}
	query := map[string]interface{}{
		"id" : SpaceId,
		"state" : 1,
	}
	//获取空间信息
    space,error := service.GetSpaceInfo(query)
	if error != nil{
		c.JSON(http.StatusOK,gin.H{"error":"Failed to get space information"})
		return
	}
	c.JSON(http.StatusOK,space)
}
/**
添加空间
 */
func AddSpace(c *gin.Context){
	spaceName := c.DefaultPostForm("name","0")
	spaceAttr := c.DefaultPostForm("attr","1")
	if spaceName == "0"{
		c.JSON(http.StatusOK,gin.H{"error":"spaceName not is empty"})
		return
	}
	attr,err := strconv.Atoi(spaceAttr)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"error":"属性值必须是整数"})
		return
	}
	space := &models.Space{
		Name: spaceName,
		Attr: attr,
	}
	if err := service.AddSpace(space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加空间失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "空间添加成功"})
}
/**
编辑空间
 */
func UpdateSpace(c *gin.Context){
	id := c.DefaultPostForm("id","0")
	name := c.DefaultPostForm("name","0")
	if id == "0" || name == "0" {
		c.JSON(http.StatusOK,gin.H{"error":"id或者空间名称不能为空"})
	}
	spaceId,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":"id必须是整形"})
		return
	}
	query := map[string]interface{}{
		"id":spaceId,
	}
	update := map[string]interface{}{}
	update["name"] = name
    if err :=service.UpdateSpace(query,update);err !=nil{
		c.JSON(http.StatusOK,gin.H{"error":"修改失败"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"修改成功"})
}
/**
删除空间
 */
func DeleteSpace(c *gin.Context){
	id := c.DefaultPostForm("id","0")
	if id == "0" {
		c.JSON(http.StatusOK,gin.H{"error":"空间id不能为空"})
		return
	}
	//条件
	spaceId,err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":"空间id不能为空"})
	}
	cond := map[string]interface{}{}
	cond["id"] = spaceId
	if err := service.DeleteSpace(cond); err !=nil {
		c.JSON(http.StatusOK,gin.H{"error":"删除失败"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"删除成功"})
}