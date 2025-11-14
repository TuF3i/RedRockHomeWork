package config

import (
	"encoding/json"
	"lesson04/core"
	"os"
)

func (root *Conf) GetConfig(path string) error {
	jsonData := core.Config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	core.Conf = jsonData
	return nil
}
