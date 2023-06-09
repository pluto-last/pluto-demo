package table

import "gV2/global"

// McCity mc城市表
type McCity struct {
	global.UUID
	CityID   string  `json:"cityID" `
	CityName string  `json:"cityName" `
	Lng      float64 `json:"lng"`
	Lat      float64 `json:"lat"`
	Address  string  `json:"address"`
}

func (McCity) TableName() string {
	return "mc_city"
}
