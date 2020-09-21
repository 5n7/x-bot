package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/skmatz/x-bot/etc"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	timeFormat     = "15:04:05"
	timeZoneName   = "Asia/Tokyo"
	timeZoneOffset = 9 * 60 * 60
)

func now() string {
	utc := time.Now().UTC()
	jst := time.FixedZone(timeZoneName, timeZoneOffset)
	return utc.In(jst).Format(timeFormat)
}

func (s *Slack) Timer(api *slack.Client, e *slackevents.AppMentionEvent) error {
	messages := strings.Split(etc.RemoveDuplicateSpace(e.Text), " ")
	commands := messages[1:] // first element is the BOT ID (mention)

	if len(commands) < 2 {
		return fmt.Errorf("timer command got invalid message: %s", e.Text)
	}

	dur, err := strconv.Atoi(commands[1])
	if err != nil {
		return err
	}

	if len(commands) > 2 {
		switch commands[2] {
		case "sec":
		case "min":
			dur *= 60
		}
	}

	var memo string
	if len(commands) > 3 {
		memo = "; " + strings.Join(commands[3:], " ")
	}

	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf("timer started at %s%s", now(), memo), false)); err != nil {
		return err
	}

	timer := time.NewTimer(time.Second * time.Duration(dur))
	defer timer.Stop()
	select {
	case <-timer.C:
		if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf("timer finished at %s%s", now(), memo), false)); err != nil {
			return err
		}
	}
	return nil
}
