package utils

import (
	"encoding/json"
	"os"
)

type Event struct {
	Timestamp  int         `json:"timestamp"`
	Check      EventCheck  `json:"check"`
	Entity     EventEntity `json:"entity"`
}

type EventCheck struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Output      string   `json:"output"`
	State       string   `json:"state"`
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
