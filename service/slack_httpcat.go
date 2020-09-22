package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/skmatz/x-bot/etc"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const httpCatURL = "https://http.cat"

func (s *Slack) HTTPCat(api *slack.Client, e *slackevents.AppMentionEvent) error {
	messages := strings.Split(etc.RemoveDuplicateSpace(e.Text), " ")
	commands := messages[1:] // first element is the BOT ID (mention)

	if len(commands) < 2 {
		return fmt.Errorf("httpcat command got invalid message: %s", e.Text)
	}

	status := commands[1]
	r, err := http.Get(fmt.Sprintf("%s/%s", httpCatURL, status))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	_, err = api.UploadFile(slack.FileUploadParameters{
		File:            "",
		Content:         "",
		Reader:          r.Body,
		Filetype:        "",
		Filename:        "httpcat",
		Title:           "",
		InitialComment:  "",
		Channels:        []string{e.Channel},
		ThreadTimestamp: "",
	})
	if err != nil {
		return err
	}
	return nil
}
