package service

import (
	"errors"
	"fmt"
	"wqy/models"
)
/**
获取单条数据
 */
func GetSpaceInfo(query map[string]interface{}) (*models.Space, error) {
	space, err := models.GetSpaceInfo(query)
	if err != nil {
		return nil, err
	}
	return &space, err
}

/**
添加空间
 */
func AddSpace(space *models.Space) error {
	if space == nil {
		return errors.New("空间不能为空")
	}
	err := models.AddSpace(space)
	if err != nil {
		return fmt.Errorf("添加空间失败 %v", err)
	}
	return nil
}

/**
修改空间
 */
func UpdateSpace(query map[string]interface{},update map[string]interface{}) error{
  return models.UpdateSpace(query,update)
}

/**
删除空间
 */
func DeleteSpace(query map[string]interface{}) error{
	return models.DeleteSpace(query)
}

/**
获取空间列表
 */
func GetSpaceList(query map[string]interface{},p int, ps int) ([]models.Space,error){
	spaces,err := models.GetSpaceList(query,p, ps)
	if err != nil{
		return spaces,err
	}
	return  spaces,err
}
/**
获取空间总数
 */
func GetSpaceCount(query map[string]interface{}) (int, error) {
	res, err := models.GetSpaceCount(query)
	if err != nil {
		return 0, err
	}
	return res, nil
}