package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/skmatz/x-bot/etc"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func randPick(n int, s []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s[:n]
}

func strJoin(ss []string) string {
	switch len(ss) {
	case 0:
		return ""
	case 1:
		return ss[0]
	default:
		r := ss[0]
		for i, s := range ss[1:] {
			if i != len(ss)-2 {
				r += ", " + s
			} else {
				r += " and " + s
			}
		}
		return r
	}

}

func (s *Slack) Pick(api *slack.Client, e *slackevents.AppMentionEvent) error {
	messages := strings.Split(etc.RemoveDuplicateSpace(e.Text), " ")
	commands := messages[1:] // first element is the BOT ID (mention)

	if len(commands) < 3 {
		return fmt.Errorf("pick command got invalid message: %s", e.Text)
	}

	n, err := strconv.Atoi(commands[1])
	if err != nil {
		return err
	}

	candidates := commands[2:]
	if l := len(candidates); l < n {
		return fmt.Errorf("cannot pick %d out of %d", n, l)
	}

	picked := strJoin(randPick(n, candidates))
	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf("picked: %s", picked), false)); err != nil {
		return err
	}
	return nil
}
