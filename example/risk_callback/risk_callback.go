package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ChainuGO/ChainuGO/api"
	"github.com/ChainuGO/ChainuGO/response_define"
	"github.com/ChainuGO/ChainuGO/rsa_utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	r.POST("/risk_cb", func(c *gin.Context) {
		req := response_define.RequestWithdrawCb{}

		body, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read request body",
			})
			return
		}

		fmt.Println("Raw JSON:", string(body))
		err = json.Unmarshal(body, &req)
		if err != nil {
			logrus.Warnln("Unmarshal fail")
			return
		}

		mapData, err := rsa_utils.StructToMap(req)
		if err != nil {
			logrus.Warnln("StructToMap fail, err", err.Error())
			response_define.FailWithMessage("StructToMap fail "+err.Error(), c)
			return
		}

		logrus.Infoln("mapData", mapData)
		logrus.Infoln("Sign", req.Sign)

		err = apiObj.VerifyRiskRSAsignature(mapData, req.Sign)
		if err != nil {
			logrus.Warnln("VerifyRiskRSAsignature fail, err", err.Error(), "end")
			response_define.FailWithMessage("verify RSA signature fail "+err.Error(), c)
			return
		}

		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		rsp := response_define.ResponseWithdraw{
			Code:      "0",
			Timestamp: timestamp,
			Message:   "",
			Sign:      "",
		}

		jStr, err := json.Marshal(&req)
		if err != nil {
			logrus.Warnln("json.Marshal fail, err", err.Error())
			return
		}

		reqMapObj := rsa_utils.ToStringMap(jStr)
		clientSign, err := apiObj.GenerateRSASignature(reqMapObj)
		if err != nil {
			logrus.Warnln("apiObj.GenerateRSASignature fail, err", err.Error())
			return
		}
		rsp.Sign = clientSign

		logrus.Infoln("VerifyRiskRSAsignature success.")
		c.JSON(http.StatusOK, rsp)
	})
	r.Run(":9003")
}
