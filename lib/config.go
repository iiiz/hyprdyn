package hyprdyn

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type MonitorConfig struct {
	Id          string  `json:"id"`
	DefaultName *string `json:"defaultName,omitempty"`
}

type Config struct {
	Monitors     []MonitorConfig `json:"monitors,omitempty"`
	AutoComplete []string        `json:"autoComplete,omitempty"`
}

func ReadConfig() *Config {
	homeDir, err := os.UserHomeDir()
	Check(err)

	jsonFile := filepath.Join(homeDir, ".config", "hyprdyn.json")

	if _, err := os.Stat(jsonFile); err == nil {
		file, err := os.Open(jsonFile)
		Check(err)
		defer file.Close()

		var config Config

		err = json.NewDecoder(file).Decode(&config)
		Check(err)

		return &config
	}

	return nil
}
