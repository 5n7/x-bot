package service

import (
	"fmt"
	"strings"

	"github.com/greymd/ojichat/generator"
	"github.com/skmatz/x-bot/etc"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (s *Slack) Ojichat(api *slack.Client, e *slackevents.AppMentionEvent) error {
	messages := strings.Split(etc.RemoveDuplicateSpace(e.Text), " ")
	commands := messages[1:] // first element is the BOT ID (mention)

	var targetName string
	if len(commands) < 2 {
		targetName = fmt.Sprintf("<@%s>", e.User)
	} else {
		targetName = commands[1]
	}

	config := generator.Config{
		TargetName: targetName,
		EmojiNum:   3,
	}
	result, err := generator.Start(config)
	if err != nil {
		return err
	}

	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(result, false)); err != nil {
		return err
	}
	return nil
}
