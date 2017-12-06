package api

import (
	"aews/config"
	"github.com/nlopes/slack"

	"log"
)

type Slack struct {
}

func SlackApi() *Slack {
	return &Slack{}
}

// slack send message
// input title , mess
// 使用方式 SendMessage(title, message)
func (s *Slack) SendMessage(title string, mess string) {
	c := config.Instance("slack")
	api := slack.New(c.GetString("earlyWarningSystem.token"))
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "",
		Text:    mess,
	}

	params.Attachments = []slack.Attachment{attachment}
	_, _, err := api.PostMessage(c.GetString("earlyWarningSystem.channel"), title, params)

	if err != nil {
		log.Fatal(err)
		return
	}
}
