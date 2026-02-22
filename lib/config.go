package hyprdyn

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"gopkg.in/go-playground/validator.v8"
)

type MonitorConfig struct {
	Id          string  `validate:"required" json:"id"`
	DefaultName *string `validate:"required,max=255" json:"defaultName"`
}

type Config struct {
	Monitors     []MonitorConfig `validate:"omitempty,dive" json:"monitors,omitempty"`
	AutoComplete []string        `validate:"omitempty,max=100" json:"autoComplete,omitempty"`
	PrimaryName  *string         `validate:"omitempty,min=1,max=255" json:"primaryName,omitempty"`
}

var validate *validator.Validate

func ReadConfig() *Config {
	homeDir, err := os.UserHomeDir()
	Check(err)

	jsonFile := filepath.Join(homeDir, ".config", "hyprdyn.json")

	if _, err := os.Stat(jsonFile); err == nil {
		file, err := os.Open(jsonFile)
		Check(err)
		defer file.Close()

		validatiorConfig := &validator.Config{TagName: "validate"}
		validate = validator.New(validatiorConfig)

		var config Config

		err = json.NewDecoder(file).Decode(&config)
		Check(err)

		if err := validate.Struct(&config); err != nil {
			log.Fatal(err)
		}

		return &config
	}

	return nil
}
