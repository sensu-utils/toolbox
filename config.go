package toolbox

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadConfig reads in a sensu handler plugin configuration file
func ReadConfig(configName string, target interface{}) error {
	filename := fmt.Sprintf("/etc/sensu/%s.json", configName)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		return err
	}

	return nil
}
