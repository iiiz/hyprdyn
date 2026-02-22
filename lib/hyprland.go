package hyprdyn

import "github.com/thiagokokada/hyprland-go"

var hyprlandClient *hyprland.RequestClient

func GetHyprClient() {
	hyprlandClient = hyprland.MustClient()
}
