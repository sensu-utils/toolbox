package toolbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

func ReadConfig(configName string, target interface{}) error {
	filename := fmt.Sprintf("/etc/sensu/%s.json", configName)

	// First attempt to open the file
	file, err := os.Open(filename)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(target)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		// We don't need the error if it was a does not exist because it's still
		// possible to configure via envars.
		return err
	}

	// Override any json value with envars
	err = env.Parse(target)
	if err != nil {
		return err
	}

	return nil
}
