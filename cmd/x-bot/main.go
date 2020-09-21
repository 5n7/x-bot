package main

import (
	"fmt"
	"os"

	x "github.com/skmatz/x-bot"
)

var (
	port = 8080
)

func main() {
	s := x.NewServer()
	s.Init()
	if err := s.Run(port); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
