package service

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (s *Slack) Ping(api *slack.Client, e *slackevents.AppMentionEvent) error {
	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText("pong", false)); err != nil {
		return err
	}
	return nil
}
