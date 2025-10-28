package table

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type UUID struct {
	ID        string     `json:"id"  form:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

func (uuid *UUID) BeforeCreate(scope *gorm.Scope) error {
	var err error
	if uuid.ID == "" {
		err = scope.SetColumn("ID", RandUUID())
	}
	return err
}

func RandUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// XcItem 晓餐商品表
type XcData struct {
	UUID
	ItemData
	ShopID     string `json:"shop_id"` // 唯一ID
	TraceID    string `json:"trace_id"`
	UploadFlag bool   `json:"upload_flag"`
	Expand1    string `json:"expand_1"`
	Expand2    string `json:"expand_2"`
}

func (XcData) TableName() string {
	return "xc_data"
}

type ItemData struct {
	CityCode          string  `json:"city_code" `
	CityName          string  `json:"city_name"`
	FirstCatName      string  `json:"first_cat_name" `  // 一级分类
	SecondCatName     string  `json:"second_cat_name" ` // 二级分类
	ThirdCatName      string  `json:"third_cat_name" `  // 三级分类
	Sort              int     `json:"sort"`             // 分类下索引
	SpecialFlag       string  `json:"special_flag"`     // 是否特价
	Brand             string  `json:"brand"`            // 品牌
	CrawlTime         int64   `json:"crawl_time"`       // 入库时间
	Name              string  `json:"name"`             // 名称
	Desc              string  `json:"desc"`             // 描述
	Origin            string  `json:"origin"`           // 产地
	ShelfLife         string  `json:"shelf_life"`       // 保质期
	PackMethod        string  `json:"pack_method"`      // 包装方式
	Classify          string  `json:"classify"`         // 商品分类
	Specification     string  `json:"specification"`    // 规格
	SaveRequire       string  `json:"save_require"`     // 贮存条件
	OriginalPrice     float64 `json:"original_price"`   // 原价
	SalePrice         float64 `json:"sale_price"`       // 促销价
	SpecificationData string  `json:"specification_data"`
	SkuID             string  `json:"sku_id"` // sku_id
	DataFlag          string  `json:"data_flag"`
}
