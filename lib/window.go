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
	target := resolveWorkspaceTarget(workspaceName)
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"%s\", follow = false, window = \"address:%s\"})", target, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

// MoveToWorkspace moves the window to the named workspace and switches the
// active view to follow it, sending a dispatch command to the Hyprland socket.
func (w Window) MoveToWorkspace(workspaceName string) {
	target := resolveWorkspaceTarget(workspaceName)
	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"%s\", follow = true, window = \"address:%s\"})", target, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

// resolveWorkspaceTarget maps a user-supplied workspace name to the value that
// should be passed as the `workspace` argument of a dispatch command.
//
// Under Hyprland 0.57 named workspaces are identified by their immutable
// `address`, not by their (possibly renamed) display name. When a workspace
// with the given name or address already exists we target it by address to
// avoid creating duplicates. Otherwise we fall back to the `name:` form so a
// new named workspace is created.
func resolveWorkspaceTarget(name string) string {
	for _, ws := range GetAllWorkspaces(true) {
		if ws.Name == name || ws.Address == name {
			return ws.Address
		}
	}

	return "name:" + name
}
