package hyprdyn

import (
	"fmt"

	"github.com/thiagokokada/hyprland-go"
)

func GetActiveWindow() *hyprland.Window {
	window, err := hyprlandClient.ActiveWindow()
	Check(err)

	return &window
}

func MoveWindowToWorkspaceSilent(window hyprland.Window, workspaceName string) {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("movetoworkspacesilent name:%s,address:%s", workspaceName, window.Address))
	Check(err)
}
