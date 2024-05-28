package models

import (
	"wqy/config"
)

const (
	SpaceTable string = "tb_space"
)

type Space struct {
	Name  string
	Attr  int
	State int
}
/**
获取单条空间数据
 */
func GetSpaceInfo(query map[string]interface{}) (Space, error) {
	space := new(Space)
	DB := config.DbHelper
	result := DB.Table(SpaceTable).Where(query).First(space)
	if result.Error != nil {
		return *space, result.Error
	}
	return *space, nil
}
/**
添加空间
 */
func AddSpace(space *Space) error {
	DB := config.DbHelper
	result := DB.Table(SpaceTable).Create(space)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
修改空间
 */
func UpdateSpace(query, set map[string]interface{}) error {
	result := config.DbHelper.Table(SpaceTable).Where(query).Updates(set)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
/**
删除空间
 */
func DeleteSpace(query map[string]interface{}) error{
	result := config.DbHelper.Table(SpaceTable).Where(query).Delete(query)
	if result.Error != nil{
		return result.Error
	}
	return nil
}