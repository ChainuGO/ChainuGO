package request_define

type RequestUserCreateOrder struct {
	ChainTokenId string `json:"ChainTokenId"`
	OrderID      string `json:"OrderID"`
	Amount       string `json:"Amount"`
	ReturnUrl    string json:"ReturnUrl"
}
