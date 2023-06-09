package params

type McItemReq struct {
	CityID string `json:"city_id"`
	AreaID string `json:"area_id"`
}

type MCDataList struct {
	DataType string `json:"data_type"`
	Data     string `json:"data"`
	CityID   string `json:"city_id"`
}
