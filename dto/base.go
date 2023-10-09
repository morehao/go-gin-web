package dto

type HttpResponse struct {
	ErrNo int                    `json:"errNo"`
	Data  map[string]interface{} `json:"data"`
}
