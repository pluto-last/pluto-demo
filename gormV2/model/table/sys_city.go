package table

type City struct {
	Code     string
	Name     string
	Province string
	City     string
}

func (City) TableName() string {
	return "sys_city"
}

type Province struct {
	Code     string
	Name     string
	Province string
}

func (Province) TableName() string {
	return "sys_province"
}

type Area struct {
	Code     string
	Name     string
	Province string
	City     string
}

func (Area) TableName() string {
	return "sys_area"
}

type Street struct {
	Code     string
	Name     string
	Province string
	City     string
	Area     string
}

func (Street) TableName() string {
	return "sys_street"
}
