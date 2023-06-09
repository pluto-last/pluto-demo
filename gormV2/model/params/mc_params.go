package params

import (
	"encoding/json"
	"gV2/utils"
	"time"
)

type SessionReq struct {
	ClientPubKey     string `json:"clientPubKey"`
	Timestamp        int64  `json:"timestamp"`
	AppKey           string `json:"appKey"`
	NonceStr         string `json:"nonceStr"`
	DeviceSdkVersion string `json:"deviceSdkVersion"`
	Platform         string `json:"platform"`
}

type RespSession struct {
	Data struct {
		ServerPubKey string `json:"serverPubKey"`
		SessionID    string `json:"sessionId"`
	} `json:"data"`
	Ret int `json:"ret"`
}

type DeviceTokenData struct {
	Language           string   `json:"language"`
	PhysicalCPU        int      `json:"physicalCpu"`
	Hostname           string   `json:"hostname"`
	Apputm             string   `json:"apputm"`
	KernOSVersion      string   `json:"kernOSVersion"`
	FreeSpace          int64    `json:"freeSpace"`
	Battery            int      `json:"battery"`
	TotalSpace         int64    `json:"totalSpace"`
	Cost               int      `json:"cost"`
	Timezone           string   `json:"timezone"`
	PackageName        string   `json:"packageName"`
	Brightness         float64  `json:"brightness"`
	Width              int      `json:"width"`
	Screen             string   `json:"screen"`
	SystemUptime       int      `json:"systemUptime"`
	AppBuildNumber     string   `json:"appBuildNumber"`
	Model              string   `json:"model"`
	KernOSRev          int      `json:"kernOSRev"`
	Memory             int64    `json:"memory"`
	KernOSType         string   `json:"kernOSType"`
	BatteryState       int      `json:"batteryState"`
	SimOperator        string   `json:"simOperator"`
	LastTokenGetTime   int      `json:"lastTokenGetTime"`
	SdkVersion         string   `json:"sdkVersion"`
	CPUType            string   `json:"cpuType"`
	Track              bool     `json:"track"`
	Platform           string   `json:"platform"`
	DNS                []string `json:"dns"`
	DeviceAlias        string   `json:"deviceAlias"`
	Idfv               string   `json:"idfv"`
	KernVersion        string   `json:"kernVersion"`
	Idfa               string   `json:"idfa"`
	CountryIso         string   `json:"countryIso"`
	T                  int64    `json:"t"`
	DeviceToken        string   `json:"deviceToken"`
	Common             Common   `json:"common"`
	DeviceOrientation  string   `json:"deviceOrientation"`
	NetworkType        string   `json:"networkType"`
	UserInterfaceIdiom int      `json:"userInterfaceIdiom"`
	AppVersion         string   `json:"appVersion"`
	Debugging          bool     `json:"debugging"`
	Boot               int      `json:"boot"`
	Height             int      `json:"height"`
	Brand              string   `json:"brand"`
	KernOSRelease      string   `json:"kernOSRelease"`
	CPUFreq            int      `json:"cpuFreq"`
	AvailableMemory    int      `json:"availableMemory"`
	OsVersion          string   `json:"osVersion"`
	Os                 string   `json:"os"`
}

type TokenReq struct {
	Data             string `json:"data"`
	Timestamp        int64  `json:"timestamp"`
	AppKey           string `json:"appKey"`
	NonceStr         string `json:"nonceStr"`
	DeviceSdkVersion string `json:"deviceSdkVersion"`
	Platform         string `json:"platform"`
	SessionId        string `json:"sessionId"`
}

type Common struct {
	NonceStr         string `json:"nonceStr"`
	AppKey           string `json:"appKey"`
	Platform         string `json:"platform"`
	Timestamp        int64  `json:"timestamp"`
	DeviceSdkVersion string `json:"deviceSdkVersion"`
}

type RespDeviceToken struct {
	Data struct {
		DeviceAlias  string `json:"deviceAlias"`
		DeviceToken  string `json:"deviceToken"`
		TokenGetTime int64  `json:"tokenGetTime"`
	} `json:"data"`
	Ret int `json:"ret"`
}

type RespCommon struct {
	Ret     int         `json:"ret"`
	Code    int         `json:"code"`
	Data    string      `json:"data"`
	Msg     string      `json:"msg"`
	ErrMsg  interface{} `json:"errMsg"`
	Error   interface{} `json:"error"`
	Success bool        `json:"success"`
}

type ReqLocation struct {
	Platform           int         `json:"platform"`
	Pticket            interface{} `json:"pticket"`
	PageEnv            string      `json:"page_env"`
	Source             string      `json:"source"`
	ThirdSource        string      `json:"third_source"`
	ThirdSourceContent string      `json:"third_source_content"`
	Lat                float64     `json:"lat"`
	Lng                float64     `json:"lng"`
	ReceiveID          string      `json:"receive_id"`
	PoiAddress         string      `json:"poi_address"`
	AddressDetail      string      `json:"address_detail"`
}

type RespLocation struct {
	Code int `json:"code"`
	Data struct {
		Location string `json:"location"`
	} `json:"data"`
}

type ReqSearch struct {
	Keyword         string    `json:"keyword"`
	SaltSign        string    `json:"salt_sign"`
	ENV             SearchENV `json:"_ENV_"`
	SortType        int       `json:"sort_type"`
	Source          string    `json:"source"`
	Size            int       `json:"size"`
	CityID          string    `json:"city_id"`
	AreaID          string    `json:"area_id"`
	AttrFlag        bool      `json:"attrFlag"`
	FilterInvoice   int       `json:"filterInvoice"`
	AttrFilterLabel []interface {
	} `json:"attr_filter_label"`
	MallSaltSign string `json:"mallSaltSign"`
	NewSearch    int    `json:"newSearch"`
	Page         int    `json:"page"`
}
type SearchENV struct {
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

type ReqDetail struct {
	SsuID        string    `json:"ssu_id"`
	ENV          DetailENV `json:"_ENV_"`
	AreaID       string    `json:"area_id"`
	SkuID        string    `json:"sku_id"`
	CityID       string    `json:"city_id"`
	MallSaltSign string    `json:"mallSaltSign"`
	SaltSign     string    `json:"salt_sign"`
}
type DetailENV struct {
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

// idfa idfv 随机生成
func GetDeviceTokenData() ([]byte, error) {

	data := DeviceTokenData{
		Language:           "zh-Hans",
		PhysicalCPU:        6,
		Hostname:           "iPhone",
		Apputm:             "App Store",
		KernOSVersion:      "18D70",
		FreeSpace:          53356213423,
		Battery:            100,
		TotalSpace:         63894048768,
		Cost:               986,
		Timezone:           "Asia/Shanghai",
		PackageName:        "com.meicai.MeicaiStoreP",
		Brightness:         35.15685498714447,
		Width:              375,
		Screen:             "1125.0,2436.0",
		SystemUptime:       1662921169,
		AppBuildNumber:     "908",
		Model:              "iPhone10,3",
		KernOSRev:          0,
		Memory:             2964504576,
		KernOSType:         "Darwin",
		BatteryState:       0,
		SimOperator:        "",
		LastTokenGetTime:   0,
		SdkVersion:         "1.1.4",
		CPUType:            "arm64",
		Track:              false,
		Platform:           "IOS",
		DeviceAlias:        "",
		Idfv:               utils.RandIDFA(),
		KernVersion:        "Darwin Kernel Version 20.3.0: Tue Jan  5 18:34:47 PST 2021; root:xnu-7195.80.35~2/RELEASE_ARM64_T8015",
		Idfa:               utils.RandIDFA(),
		CountryIso:         "CN",
		T:                  time.Now().UnixNano() / 1e6,
		DeviceToken:        "",
		DeviceOrientation:  "portrait",
		NetworkType:        "WIFI",
		UserInterfaceIdiom: 0,
		AppVersion:         "5.4.0",
		Debugging:          false,
		Boot:               3952204,
		Height:             812,
		Brand:              "Apple",
		KernOSRelease:      "20.3.0",
		CPUFreq:            0,
		AvailableMemory:    698073088,
		OsVersion:          "14.4.2",
		Os:                 "iOS",
	}

	data.DNS = append(data.DNS, "180.101.49.13")
	data.DNS = append(data.DNS, "180.101.49.14")

	data.Common = Common{
		NonceStr:         utils.RandIDFA(),
		AppKey:           "mall",
		Platform:         "IOS",
		Timestamp:        time.Now().Unix(),
		DeviceSdkVersion: "1.1.4",
	}

	return json.Marshal(&data)
}
