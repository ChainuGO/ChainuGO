package api

import (
	"strconv"

	"github.com/ChainuGO/ChainuGO/request_define"
)

// CreateWallet create wallet
// @param ChainTokenId user open id
// @param OrderID chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) CreateOrder(ChainTokenId, OrderID, ReturnUrl string, Amount int64) ([]byte, string, string, string, error) {

 amount := strconv.FormatInt(Amount, 10)
 return s.signPack(
  request_define.RequestUserCreateOrder{
   ChainTokenId: ChainTokenId,
   OrderID:      OrderID,
   Amount:       amount,
   ReturnUrl:    ReturnUrl,
  },
 )
}
