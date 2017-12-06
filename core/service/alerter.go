package service

import (
	"aews/core/model/api"
)

type Alerter struct {
}

func AlerterService () *Alerter {
	return &Alerter{}
}

func (a *Alerter) Check() {
	return
}

func (a *Alerter) SnedDayAlerter() {
	api.SlackApi().SendMessage("數據異常警示", "單一遊戲場次累積10萬場，RTP > 100%現象")
}
