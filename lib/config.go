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
	DefaultName *string `validate:"required,min=1,max=255" json:"defaultName"`
}

type LayoutWorkspaceMapping struct {
	WorkspaceName string `validate:"required,min=1,max=255" json:"workspace"`
	Layout        string `validate:"required,min=1,max=255" json:"layout"`
}

type UseLayoutSetting struct {
	Setup  *bool `validate:"omitempty" json:"setup,omitempty"`
	Spawn  *bool `validate:"omitempty" json:"spawn,omitempty"`
	Rename *bool `validate:"omitempty" json:"rename,omitempty"`
}

type LayoutConfig struct {
	UseLayoutSetting   UseLayoutSetting         `validate:"omitempty" json:"useLayoutSetting,omitempty"`
	LayoutMappings     []LayoutWorkspaceMapping `validate:"omitempty,dive" json:"workspaceMapping,omitempty"`
	QuickToggleOptions []string                 `validate:"omitempty,min=2" json:"quickToggleOptions,omitempty"`
}

type Config struct {
	Monitors     []MonitorConfig          `validate:"omitempty,dive" json:"monitors,omitempty"`
	AutoComplete []string                 `validate:"omitempty,max=100" json:"autoComplete,omitempty"`
	PrimaryName  *string                  `validate:"omitempty,min=1,max=255" json:"primaryName,omitempty"`
	Theme        *string                  `validate:"omitempty,oneof=default emerald cyber nordly ruby snow darksky ocean" json:"theme,omitempty"`
	CustomTheme  *ht.HyprdynHexcolorTheme `validate:"omitempty" json:"customTheme,omitempty"`
	LayoutConfig *LayoutConfig            `validate:"omitempty" json:"layout,omitempty"`

	layoutMap map[string]string
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

		if config.LayoutConfig != nil {
			if len(config.LayoutConfig.LayoutMappings) > 0 {
				config.layoutMap = make(map[string]string)

				for _, lm := range config.LayoutConfig.LayoutMappings {
					config.layoutMap[lm.WorkspaceName] = lm.Layout
				}
			}
		}

		return &config
	}

	return nil
}

func (c *Config) GetMappedLayoutName(workspaceName string) *string {
	if c == nil || c.LayoutConfig == nil || len(c.LayoutConfig.LayoutMappings) == 0 || c.layoutMap == nil {
		return nil
	}

	if layoutName, ok := c.layoutMap[workspaceName]; ok {
		return &layoutName
	}

	return nil
}

// NOTE: default to true
func (u *UseLayoutSetting) SetupEnabled() bool {
	return u.Setup == nil || *u.Setup
}

func (u *UseLayoutSetting) SpawnEnabled() bool {
	return u.Spawn == nil || *u.Spawn
}

func (u *UseLayoutSetting) RenameEnabled() bool {
	return u.Rename == nil || *u.Rename
}
