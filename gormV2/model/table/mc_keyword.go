package table

import "gV2/global"

// McKeyword 搜索关键词表
type McKeyword struct {
	CityName string `json:"cityName" `
	Keyword  string `json:"keyword" `
	Status   string `json:"status"`
}

func (McKeyword) TableName() string {
	return "mc_keyword"
}

type McTask struct {
	global.UUID
	CityID   string `json:"cityID" `
	CityName string `json:"cityName" `
	Location string `json:"location"`
	Keyword  string `json:"keyword"`
	Rounds   string `json:"rounds"`
	Status   string `json:"status"`
	Proxy    string `json:"proxy"`
	Log      string `json:"log"`
}

func (McTask) TableName() string {
	return "mc_task"
}

type McSubTask struct {
	global.UUID
	CityID   string `json:"cityID" `
	CityName string `json:"cityName" `
	Location string `json:"location"`
	Keyword  string `json:"keyword"`
	Rounds   string `json:"rounds"`
	SkuID    string `json:"skuID"`
	Status   string `json:"status"`
	Proxy    string `json:"proxy"`
	Log      string `json:"log"`
}

func (McSubTask) TableName() string {
	return "mc_sub_task"
}
