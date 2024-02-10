package responses

type Api struct {
	Status  bool `json:"status"`
	Code    int  `json:"code"`
	Message any  `json:"message"`
	Data    any  `json:"data"`
	Error   any  `json:"errors"`
}
