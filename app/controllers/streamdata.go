package controllers

import (
	"gotrade/app/models"
	"gotrade/bitflyer"
	"gotrade/config"
)

func StreamIngestionData(){
	 tickerChannel := make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey,config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode,tickerChannel)
	 go func() {
		 for ticker := range tickerChannel {
			 //log.Printf("action=StreamIngestionData, %v", ticker)
			 for _, duration := range config.Config.Durations {
				 isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				 if isCreated == true && duration == config.Config.TradeDuration {

				 }
			 }
		 }
	 }()
}
