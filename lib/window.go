package hyprdyn

import (
	"fmt"

	"github.com/thiagokokada/hyprland-go"
)

type Window struct {
	hyprland.Window
}

func GetActiveWindow() Window {
	window, err := hyprlandClient.ActiveWindow()
	Check(err)

	return Window{Window: window}
}

func (w Window) MoveToWorkspaceSilent(workspaceName string) {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("movetoworkspacesilent name:%s,address:%s", workspaceName, w.Address))
	Check(err)
}

func (w Window) MoveToWorkspace(workspaceName string) {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("movetoworkspace name:%s,address:%s", workspaceName, w.Address))
	Check(err)
}
