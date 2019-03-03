package toolbox

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/sensu/sensu-go/types"
)

// Event is a wrapper around sensu-go's Event so we can add convenience
// methods.
type Event struct {
	types.Event
	humanStatus          string
	colorizedHumanStatus string
	humanAction          string
}

// ReadEvent performs some sanity checks on the event and reads it from stdin
func ReadEvent(event *Event) error {
	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(event); err != nil {
		return err
	}

	if event.Timestamp <= 0 {
		return errors.New("timestamp is missing or must be greater than zero")
	}

	if event.Entity == nil {
		return errors.New("entity is missing from event")
	}

	if !event.HasCheck() {
		return errors.New("check is missing from event")
	}

	if err := event.Entity.Validate(); err != nil {
		return err
	}

	if err := event.Check.Validate(); err != nil {
		return err
	}

	if err := createHumanStatus(event); err != nil {
		return err
	}

	return nil
}

func createHumanStatus(event *Event) error {
	// If the event has transitioned from incident to resolved, special case the
	// human text to indicate that everything is a'okay now
	if event.IsResolution() {
		event.humanStatus = "RESOLVED"
		event.colorizedHumanStatus = "\x0300,03RESOLVED\x03"
		event.humanAction = "succeeded"
	} else {
		// For each known exit code, configure human understandable status and action messages
		switch event.Check.Status {
		case 0:
			event.humanStatus = "OKAY"
			event.colorizedHumanStatus = "\x0300,03OKAY\x03"
			event.humanAction = "succeeded"
		case 1:
			event.humanStatus = "WARN"
			event.colorizedHumanStatus = "\x0301,08WARN\x03"
			event.humanAction = "warned"

		case 2:
			event.humanStatus = "ALERT"
			event.colorizedHumanStatus = "\x0301,04ALERT\x03"
			event.humanAction = "failed"
		// Technically, codes 3 or higher can be customized, we might want to support that in the future
		// For now, we just mark the event as 'unknown' which is the old Nagios standard.
		default:
			event.humanStatus = "UNKNOWN"
			event.colorizedHumanStatus = "\x0301,14UNKNOWN\x03"
			event.humanAction = "exited"

		}
	}
	return nil
}
