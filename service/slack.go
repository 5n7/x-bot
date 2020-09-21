package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/skmatz/x-bot/etc"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	slackToken = os.Getenv("SLACK_TOKEN")
)

type Slack struct{}

func NewSlack() *Slack {
	return &Slack{}
}

func (s *Slack) Callback(event slackevents.EventsAPIEvent) error {
	api := slack.New(slackToken)

	switch e := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		messages := strings.Split(etc.RemoveDuplicateSpace(e.Text), " ")
		if len(messages) < 2 {
			return fmt.Errorf("invalid message: %s", e.Text)
		}

		switch messages[1] {
		case "contribute":
			if err := s.Contribute(api, e); err != nil {
				return err
			}
		case "pick":
			if err := s.Pick(api, e); err != nil {
				return err
			}
		case "ping":
			if err := s.Ping(api, e); err != nil {
				return err
			}
		case "timer":
			if err := s.Timer(api, e); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Slack) Challenge(b []byte) (string, error) {
	var r *slackevents.ChallengeResponse
	if err := json.Unmarshal(b, &r); err != nil {
		return "", err
	}
	return r.Challenge, nil
}
