package utils

import (
	"encoding/json"
	"os"
)

type Event struct {
	ID         string
	Action     string
	Timestamp  int
	Occurences int
	Check      EventCheck
	Client     EventClient
}

type EventCheck struct {
	Type             string
	TotalStateChange int `json:"total_state_change"`
	History          []string
	Status           int
	Output           string
	Executed         int
	Issued           int
	Name             string
	Thresholds       struct {
		Critical int
		Warning  int
	}
}

type EventClient struct {
	Timestamp int
	Version   string
	Socket    struct {
		Port int
		Bind string
	}
	Subscriptions []string
	Environment   string
	Address       string
	Name          string
}

func ReadEvent(event *Event) error {
	decoder := json.NewDecoder(os.Stdin)
	return decoder.Decode(event)
}
