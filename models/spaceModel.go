package models

import (
	"wqy/config"
)

const (
	SpaceTable string = "tb_space"
)

type Space struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	Attr  int `json:"attr"`
	State int `json:"state"`
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
获取空间列表
 */
func GetSpaceList(query map[string]interface{}, p int, ps int) ([]Space, error) {
	var spaces []Space
	offset := (p - 1) * ps
	DB := config.DbHelper
	result := DB.Table(SpaceTable).Where(query).Offset(offset).Limit(ps).Find(&spaces)
	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
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
/**
获取空间总数
 */
func GetSpaceCount(query map[string]interface{}) (int, error) {
	var count int
	result := config.DbHelper.Table(SpaceTable).Where(query).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}