package service

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const (
	repositoryURL = "https://github.com/skmatz/x-bot"
)

func (s *Slack) Contribute(api *slack.Client, e *slackevents.AppMentionEvent) error {
	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf(":tada: %s", repositoryURL), false)); err != nil {
		return err
	}
	return nil
}
