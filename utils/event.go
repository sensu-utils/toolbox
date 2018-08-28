package utils

import (
	"encoding/json"
	"os"
)

type Event struct {
	ID         string      `json:"id"`
	Timestamp  int         `json:"timestamp"`
	Action     string      `json:"action"`
	Occurences int         `json:"occurrences"`
	Check      EventCheck  `json:"check"`
	Entity     EventEntity `json:"entity"`
}

type EventCheck struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Subscribers []string `json:"subscribers"`
	Interval    int      `json:"interval"`
	Handler     string   `json:"handler"`
	Handlers    []string `json:"handlers"`
	Issued      int      `json:"issued"`
	Output      string   `json:"output"`
	Status      int      `json:"status"`
	History     []int    `json:"history"`
	Source      string   `json:"source"`
	Origin      string   `json:"origin"`
}

type EventEntity struct {
	ID            string   `json:"ID"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Timestamp     int      `json:"timestamp"`
}

func ReadEvent(event *Event) error {
	decoder := json.NewDecoder(os.Stdin)
	return decoder.Decode(event)
}
