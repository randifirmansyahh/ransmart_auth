package responseModel

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"respon_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
