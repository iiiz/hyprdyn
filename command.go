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
			}
		}

		os.Exit(0)
	} else {
		log.Fatal("Hyprdyn: No config present.")
	}
}

func PrimaryCommand() {
	if config != nil && config.PrimaryName != nil {
		existing, found := workspaces.FindByName(*config.PrimaryName)

		if found {
			existing.FocusOnCurrentMonitor()
		} else {
			hd.SpawnWorkspace(*config.PrimaryName)
		}

		os.Exit(0)

	} else {
		log.Fatal("Hyprdyn: No config or configured primary present.")
	}
}

func ToggleLayoutCommand() {
	if config != nil && len(config.Layouts) > 0 {
		activeWs := hd.GetActiveWorkspace()
		nextIndex := 0

		for i, l := range config.Layouts {
			if l == activeWs.TiledLayout {
				nextIndex = (i + 1) % len(config.Layouts)
				break
			}
		}

		activeWs.SetLayout(config.Layouts[nextIndex])
		os.Exit(0)

	} else {
		log.Fatal("Hyprdyn: No config or configured layouts present.")
	}
}
