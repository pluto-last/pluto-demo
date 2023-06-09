package table

import "gV2/global"

// McItem 美菜商品表
type McItem struct {
	global.UUID
	ItemData
	UploadFlag bool   `json:"upload_flag"`
	TraceID    string `json:"trace_id"`
}

func (McItem) TableName() string {
	return "mc_item"
}

type ItemData struct {
	CityID        string `json:"city_id" `
	CityName      string `json:"city_name"`
	Keyword       string `json:"keyword"`
	FirstCatID    string `json:"first_cat_id" `
	FirstCatName  string `json:"first_cat_name" `
	SecondCatID   string `json:"second_cat_id" `
	SecondCatName string `json:"second_cat_name" `
	SkuID         int    `json:"sku_id"`
	SpuID         int    `json:"spu_id"`
	Name          string `json:"name"`
	Data          string `json:"data"`       // 原始报文
	CrawlTime     string `json:"crawl_time"` // 上报时间
}
