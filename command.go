package main

import (
	"os"

	"github.com/charmbracelet/log"

	hd "hyprdyn/lib"
)

func SetupCommand() {
	if config != nil {
		for _, monitorConfig := range config.Monitors {
			ws := workspaces.GetForegroundByMonitor(monitorConfig.Id)

			if ws != nil && monitorConfig.DefaultName != nil {
				ws.Rename(*monitorConfig.DefaultName)

				if config.LayoutConfig != nil && config.LayoutConfig.UseLayoutSetting.SetupEnabled() {
					if layoutName := config.GetMappedLayoutName(*monitorConfig.DefaultName); layoutName != nil {
						ws.SetLayout(*layoutName)
					}
				}
			}
		}

		os.Exit(0)
	} else {
		log.Fatal("Hyprdyn: No config present.")
	}
}

func PrimaryCommand() {
	if config != nil && config.PrimaryName != nil {
		existing := workspaces.FindByName(*config.PrimaryName)

		if existing != nil {
			existing.FocusOnCurrentMonitor()
		} else {
			hd.SpawnWorkspace(*config.PrimaryName)
			active := hd.GetActiveWorkspace()

			if config.LayoutConfig != nil && config.LayoutConfig.UseLayoutSetting.SpawnEnabled() {
				if layoutName := config.GetMappedLayoutName(*config.PrimaryName); layoutName != nil {
					active.SetLayout(*layoutName)
				}
			}
		}

		os.Exit(0)

	} else {
		log.Fatal("Hyprdyn: No config or configured primary present.")
	}
}

func ToggleLayoutCommand() {
	if config != nil && config.LayoutConfig != nil {
		quickToggleLen := len(config.LayoutConfig.QuickToggleOptions)

		if quickToggleLen > 0 {
			activeWs := hd.GetActiveWorkspace()
			nextIndex := 0

			for i, l := range config.LayoutConfig.QuickToggleOptions {
				if l == activeWs.TiledLayout {
					nextIndex = (i + 1) % quickToggleLen
					break
				}
			}

			activeWs.SetLayout(config.LayoutConfig.QuickToggleOptions[nextIndex])
			os.Exit(0)

		} else {
			log.Fatal("Hyprdyn: No config or configured layouts present.")
		}
	}
}
