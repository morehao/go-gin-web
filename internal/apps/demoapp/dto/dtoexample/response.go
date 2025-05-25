package dtoexample

import "fmt"

type FormatResResp struct {
	Items []FormatDataItem `json:"items"`
	FormatDataItem
	ItemMap   map[string]FormatDataItem `json:"itemMap"`
	PriceList []float64                 `json:"priceList" precision:"2"`
	NameList  []string                  `json:"nameList"`
	NameMap   map[string][]string       `json:"nameMap"`
	PriceMap  map[string]float64        `json:"priceMap" precision:"2"`
}

type FormatDataItem struct {
	Price     float64          `json:"price" precision:"2"`
	PriceList []float64        `json:"priceList" precision:"2"`
	DescList  []string         `json:"descList"`
	Children  []FormatDataItem `json:"children"`
}

type SSEMessage struct {
	ID    string `json:"id"`
	Event string `json:"event"`
	Data  string `json:"data"`
}

func (msg SSEMessage) Format() string {
	result := ""
	if msg.ID != "" {
		result += fmt.Sprintf("id: %s\n", msg.ID)
	}
	if msg.Event != "" {
		result += fmt.Sprintf("event: %s\n", msg.Event)
	}
	result += fmt.Sprintf("data: %s\n\n", msg.Data)
	return result
}
