package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const jokeURL = "https://official-joke-api.appspot.com/jokes/random"

type jokeResponse struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func (s *Slack) Joke(api *slack.Client, e *slackevents.AppMentionEvent) error {
	r, err := http.Get(jokeURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var j jokeResponse
	json.NewDecoder(r.Body).Decode(&j)

	msg := fmt.Sprintf(`
> %s
> %s
`, j.Setup, j.Punchline)
	if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(msg, false)); err != nil {
		return err
	}
	return nil
}
