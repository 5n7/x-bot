package service

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const list = `
- contribute
- list
- pick
- ping
- timer
`

func (s *Slack) List(api *slack.Client, e *slackevents.AppMentionEvent) error {
	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(list, false)); err != nil {
		return err
	}
	return nil
}
