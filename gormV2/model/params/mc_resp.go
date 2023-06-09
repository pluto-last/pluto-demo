package params

type CommonResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Ret     int    `json:"ret"`
	Success bool   `json:"success"`
}

type RespSearch struct {
	Code    int        `json:"code"`
	Data    SearchData `json:"data"`
	Msg     string     `json:"msg"`
	Ret     int        `json:"ret"`
	Success bool       `json:"success"`
}

type SearchData struct {
	Rows []Rows `json:"rows"`
}

type Rows struct {
	SkuBase         SkuBase           `json:"skuBase"`
	SuDesc          string            `json:"suDesc"`
	ExtensionSsuIds []ExtensionSsuIds `json:"extensionSsuIds"`
	SsuList         []SsuList         `json:"ssuList"`
}

type SkuBase struct {
	SkuID       int    `json:"skuId"`
	SkuName     string `json:"skuName"`
	SkuUnit     string `json:"skuUnit"`
	BiID        int    `json:"biId"`
	BiName      string `json:"biName"`
	BiAliasName string `json:"biAliasName"`
	SpuID       int    `json:"spuId"`
	SpuName     string `json:"spuName"`
	SaleC1ID    int    `json:"saleC1Id"`
	SaleC2ID    int    `json:"saleC2Id"`
}

type ExtensionSsuIds struct {
	SsuID    int    `json:"ssuId"`
	SkuID    int    `json:"skuId"`
	UniqueID string `json:"uniqueId"`
}

type SsuList struct {
	SsuID        int           `json:"ssuId"`
	Name         string        `json:"name"`
	SuDesc       string        `json:"suDesc"`
	SsuFp        int           `json:"ssuFp"`
	SsuFpText    string        `json:"ssuFpText"`
	PriceSpm     string        `json:"priceSpm"`
	SuitSsuList  []interface{} `json:"suitSsuList"`
	ActivityType int           `json:"activityType"`
	SkuID        int           `json:"sku_id"`
	UniqueID     string        `json:"unique_id"`
	SaleC1ID     int           `json:"sale_c1_id"`
}
