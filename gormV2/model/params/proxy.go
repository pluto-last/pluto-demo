package params

type RespZhiMa struct {
	Code int `json:"code"`
	Data []struct {
		IP         string `json:"ip"`
		Port       int    `json:"port"`
		ExpireTime string `json:"expire_time"`
		City       string `json:"city"`
		Isp        string `json:"isp"`
	} `json:"data"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type RespKuai struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Count          int      `json:"count"`
		ProxyList      []string `json:"proxy_list"`
		OrderLeftCount int      `json:"order_left_count"`
		DedupCount     int      `json:"dedup_count"`
	} `json:"data"`
}
