package hyprdyn

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/go-playground/validator/v10"

	ht "hyprdyn/lib/ui/themes"
)

type MonitorConfig struct {
	Id          string  `validate:"required" json:"id"`
	DefaultName *string `validate:"required,max=255" json:"defaultName"`
}

type Config struct {
	Monitors     []MonitorConfig          `validate:"omitempty,dive" json:"monitors,omitempty"`
	AutoComplete []string                 `validate:"omitempty,max=100" json:"autoComplete,omitempty"`
	PrimaryName  *string                  `validate:"omitempty,min=1,max=255" json:"primaryName,omitempty"`
	Theme        *string                  `validate:"omitempty,oneof=default emerald cyber nordly ruby snow darksky ocean" json:"theme,omitempty"`
	CustomTheme  *ht.HyprdynHexcolorTheme `validate:"omitempty" json:"customTheme,omitempty"`
}

func ReadConfig() *Config {
	homeDir, err := os.UserHomeDir()
	Check(err)

	jsonFile := filepath.Join(homeDir, ".config/hyprdyn/config.json")

	if _, err := os.Stat(jsonFile); err == nil {
		file, err := os.Open(jsonFile)
		Check(err)
		defer file.Close()

		validate := validator.New()

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
