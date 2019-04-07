package main

import (
	"encoding/json"
	coincheckgo "github.com/Akagi201/coincheckgo"
	"gotrade/app/controllers"
	"gotrade/config"
	"gotrade/utils"
)

type TickerResult struct {
	Last      float64
	Bid       float64
	Ask       float64
	High      float64
	Low       float64
	Volume    float64
	Timestamp int
}

func getTicker(client coincheckgo.CoinCheck) TickerResult {
	resp := client.Ticker.All()
	var tick_result TickerResult
	json_err := json.Unmarshal([]byte(resp), &tick_result)
	if json_err != nil {
		panic(json_err)
	}
	return tick_result
}

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
