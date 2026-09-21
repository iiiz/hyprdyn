package hyprdyn

import (
	"fmt"
)

type FullscreenState int

const (
	FullscreenNone FullscreenState = iota
	FullscreenOn
	FullscreenMax
)

type Window struct {
	Address          string          `json:"address"`
	Mapped           bool            `json:"mapped"`
	Hidden           bool            `json:"hidden"`
	At               [2]int          `json:"at"`
	Size             [2]int          `json:"size"`
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
	InhibitingIdle   bool            `json:"inhibitingIdle"`
}

// GetActiveWindow queries the Hyprland socket for the currently active window
// and unmarshals the response into a Window struct.
func GetActiveWindow() Window {
	var window Window

	res, err := hyprlandClient.sendCommand("activewindow", nil)
	Check(err)

	window, err = UnmarshalHyprlandResponse[Window](res)
	Check(err)

	return window
}

// MoveToWorkspaceSilent moves the window to the named workspace without
// switching the active view, sending a dispatch command to the Hyprland socket.
func (w Window) MoveToWorkspaceSilent(workspaceName string) {
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"name:%s\", follow = false, window = \"address:%s\"})", workspaceName, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

// MoveToWorkspace moves the window to the named workspace and switches the
// active view to follow it, sending a dispatch command to the Hyprland socket.
func (w Window) MoveToWorkspace(workspaceName string) {
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"name:%s\", follow = true, window = \"address:%s\"})", workspaceName, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}
