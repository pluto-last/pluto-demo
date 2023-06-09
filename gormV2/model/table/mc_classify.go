package table

import "gV2/global"

// McClassify 美菜类目表
type McClassify struct {
	global.UUID
	CityName string `json:"city_name"`
	CityID   string `json:"city_id"`
	ClassID  string `json:"class_id"`
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
}

func (McClassify) TableName() string {
	return "mc_classify"
}
