package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wqy/common"
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
func GetSpaceInfo(c *gin.Context) {
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
	res := common.JsonRes("200", "空间添加成功", map[string]interface{}{"result": 1})
	c.JSON(http.StatusOK, res)
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
	res := common.JsonRes("200","修改成功", map[string]interface{}{"result":1})
	c.JSON(http.StatusOK,res)
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
	res := common.JsonRes("200","删除成功", map[string]interface{}{"result":1})
	c.JSON(http.StatusOK,res)
}
/**
获取空间列表
 */
func GetSpaceList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	query := map[string]interface{}{
		"state": 1,
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数 page 解析错误"})
		return
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数 pageSize 解析错误"})
		return
	}
	result, err := service.GetSpaceList(query, p, ps)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "数据获取失败"})
		return
	}
	//获取空间总数
    count,_ := service.GetSpaceCount(query)
	data := map[string]interface{}{
		"count":count,
		"list": result,
	}
	res := common.JsonRes("200", "获取成功", data)
	c.JSON(http.StatusOK, res)
}