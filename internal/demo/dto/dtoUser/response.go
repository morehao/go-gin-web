package dtoUser

type GetUserRes struct {
	ID   uint64 `json:"id"`   // 用户ID
	Name string `json:"name"` // 用户名
}

type FormatDataRes struct {
	Items []Item `json:"items"`
	Item
	ItemMap   map[string]Item     `json:"itemMap"`
	PriceList []float64           `json:"priceList" precision:"2"`
	NameList  []string            `json:"nameList"`
	NameMap   map[string][]string `json:"nameMap"`
	PriceMap  map[string]float64  `json:"priceMap" precision:"2"`
}

type Item struct {
	Price     float64   `json:"price" precision:"2"`
	PriceList []float64 `json:"priceList" precision:"2"`
	DescList  []string  `json:"descList"`
	Children  []Item    `json:"children"`
}
