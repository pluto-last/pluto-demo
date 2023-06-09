package table

type SysProxy struct {
	Date     string `json:"date"`
	UseCount int    `json:"useCount"`
}

func (SysProxy) TableName() string {
	return "sys_proxy"
}
