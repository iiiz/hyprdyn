package hyprdyn

import (
	"fmt"
)

type FullscreenState int

type Window struct {
	Address          string          `json:"address"`
	Mapped           bool            `json:"mapped"`
	Hidden           bool            `json:"hidden"`
	At               []int           `json:"at"`
	Size             []int           `json:"size"`
	Workspace        WorkspaceType   `json:"workspace"`
	Floating         bool            `json:"floating"`
	Pseudo           bool            `json:"pseudo"`
	Monitor          int             `json:"monitor"`
	Class            string          `json:"class"`
	Title            string          `json:"title"`
	InitialClass     string          `json:"initialClass"`
	InitialTitle     string          `json:"initialTitle"`
	Pid              int             `json:"pid"`
	Xwayland         bool            `json:"xwayland"`
	Pinned           bool            `json:"pinned"`
	Fullscreen       FullscreenState `json:"fullscreen"`
	FullscreenClient FullscreenState `json:"fullscreenClient"`
	Grouped          []string        `json:"grouped"`
	Tags             []string        `json:"tags"`
	Swallowing       string          `json:"swallowing"`
	FocusHistoryId   int             `json:"focusHistoryID"`
}

func GetActiveWindow() Window {
	var window Window

	res, err := hyprlandClient.sendCommand("activewindow", nil)
	Check(err)

	window, err = UnmarshalHyprlandResponse[Window](res)
	Check(err)

	return window
}

func (w Window) MoveToWorkspaceSilent(workspaceName string) {
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"name:%s\", follow = false, window = \"address:%s\"})", workspaceName, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

func (w Window) MoveToWorkspace(workspaceName string) {
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"name:%s\", follow = true, window = \"address:%s\"})", workspaceName, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}
