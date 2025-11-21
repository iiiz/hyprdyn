package hyprdyn

import "github.com/thiagokokada/hyprland-go"

func GetFocusedMonitor() hyprland.Monitor {
	monitors, err := hyprlandClient.Monitors()
	Check(err)

	var focusedMonitor hyprland.Monitor

	for _, mon := range monitors {
		if mon.Focused {
			focusedMonitor = mon
		}
	}

	return focusedMonitor
}
