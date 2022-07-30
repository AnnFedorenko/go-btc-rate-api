package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Rate struct {
	Price string
}

const BinanceURL = "https://api.binance.com/api/v3/ticker/price?symbol=BTCUAH"

func GetRateFromBinance() (Rate, error) {
	var newRate Rate
	resp, err := http.Get(BinanceURL)
	if err != nil {
		return newRate, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &newRate); err != nil {
		return newRate, err
	}

	return newRate, nil
}
