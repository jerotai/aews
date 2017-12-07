package service

import (
	"aews/core/model/api"
	"log"
)

var snedApi interface {
	SendMessage(title string, mess string)
}

type Alerter struct {
}

func AlerterService (api_type string) *Alerter {
	switch api_type {
		case "slack":
			snedApi = api.SlackApi()
		default:
			log.Fatal("send api undefined")
	}
	return &Alerter{}
}

func (a *Alerter) SnedDayAlerter() {
	snedApi.SendMessage("數據異常警示", "單一遊戲場次累積10萬場，RTP > 100% 現象")
}
