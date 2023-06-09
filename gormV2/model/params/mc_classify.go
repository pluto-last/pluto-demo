package params

type ReqSaleClass struct {
	ENV          ClassENV `json:"_ENV_"`
	CityID       string   `json:"city_id"`
	AreaID       string   `json:"area_id"`
	MallSaltSign string   `json:"mallSaltSign"`
	ParentID     string   `json:"parent_id"`
	SaltSign     string   `json:"salt_sign"`
}
type ClassENV struct {
	Mno               string `json:"mno"`
	Idfa              string `json:"idfa"`
	AppVersion        string `json:"app_version"`
	Net               string `json:"net"`
	Source            string `json:"source"`
	AppkeyVersion     int    `json:"appkey_version"`
	Lat               string `json:"lat"`
	DistributeChannel string `json:"distribute_channel"`
	IP                string `json:"ip"`
	DeviceName        string `json:"device_name"`
	Bssid             string `json:"bssid"`
	Idfv              string `json:"idfv"`
	DeviceID          string `json:"device_id"`
	OsVersion         string `json:"os_version"`
	Ssid              string `json:"ssid"`
	Lng               string `json:"lng"`
	Location          string `json:"location"`
}

type ClassifyResp struct {
	Ret     int          `json:"ret"`
	Code    int          `json:"code"`
	Data    ClassifyData `json:"data"`
	Msg     string       `json:"msg"`
	Success bool         `json:"success"`
}

type ClassifyData struct {
	List []ClassifyList `json:"list"`
}

type ClassifyList struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
	NameImg  string `json:"nameImg"`
}
