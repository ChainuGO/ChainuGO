package response_define

type ResponseUserCreateOrder struct {
	Sign      string                      `json:"sign"`
	Timestamp string                      `json:"timestamp"`
	Data      ResponseUserCreateOrderData `json:"data"`
	Msg       string                      `json:"msg"`
	Code      int                         `json:"code"`
}

type ResponseUserCreateOrderData struct {
	PayUrl      string  `json:"payUrl"`
	PlatOrderId string  `json:"platOrderId"`
	OrderId     string  `json:"orderId"`
	Amount      int     `json:"amount"`
	PlatAmount  float64 `json:"platAmount"`
	Addr        string  `json:"addr"`
	Token       string  `json:"token"`
	Status      int32   `json:"status"`
	EndTime     int     `json:"endTime"`
	ReturnUrl   string  json:"returnUrl"
}
