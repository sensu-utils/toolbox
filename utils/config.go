package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadConfig(configName string, target interface{}) error {
	//filename := fmt.Sprintf("/etc/sensu/%s.json", configName)
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
