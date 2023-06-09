package params

type McItemResp struct {
	Code    int      `json:"code"`
	Data    ItemData `json:"data"`
	Msg     string   `json:"msg"`
	Ret     int      `json:"ret"`
	Success bool     `json:"success"`
}

type ItemData struct {
	Ssu Ssu `json:"ssu"`
	Sku Sku `json:"sku"`
}

type Ssu struct {
	SaleC1ID  int   `json:"sale_c1_id"`
	SaleC2ID  int   `json:"sale_c2_id"`
	SaleC1Ids []int `json:"saleC1Ids"` // 所属一级分类
	SaleC2Ids []int `json:"saleC2Ids"` // 所属二级分类
}

type Sku struct {
	SkuID int    `json:"sku_id"`
	SpuID int    `json:"spu_id"`
	Name  string `json:"name"`
}
