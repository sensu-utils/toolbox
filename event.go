package toolbox

import (
	"encoding/json"
	"os"

	"github.com/sensu/sensu-go/types"
)

// Event is a wrapper around sensu-go's Event so we can add convenience
// methods.
type Event struct {
	types.Event
}

func ReadEvent(event *Event) error {
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(event); err != nil {
		return err
	}

	if err := event.Check.Validate(); err != nil {
		return err
	}

	return nil
}
