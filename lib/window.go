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
	Address               string          `json:"address"`
	Mapped                bool            `json:"mapped"`
	Hidden                bool            `json:"hidden"`
	Visible               bool            `json:"visible"`
	AcceptsInput          bool            `json:"acceptsInput"`
	At                    [2]int          `json:"at"`
	Size                  [2]int          `json:"size"`
	Workspace             WorkspaceCore   `json:"workspace"`
	Floating              bool            `json:"floating"`
	Monitor               int             `json:"monitor"`
	Class                 string          `json:"class"`
	Title                 string          `json:"title"`
	InitialClass          string          `json:"initialClass"`
	InitialTitle          string          `json:"initialTitle"`
	Pid                   int             `json:"pid"`
	Xwayland              bool            `json:"xwayland"`
	Pinned                bool            `json:"pinned"`
	PinFullscreened       bool            `json:"pinFullscreened"`
	Fullscreen            FullscreenState `json:"fullscreen"`
	FullscreenClient      FullscreenState `json:"fullscreenClient"`
	FullscreenHandler     string          `json:"fullscreenHandler"`
	AllowedOverFullscreen bool            `json:"allowedOverFullscreen"`
	Grouped               []string        `json:"grouped"`
	Tags                  []string        `json:"tags"`
	Swallowing            string          `json:"swallowing"`
	FocusHistoryId        int             `json:"focusHistoryID"`
	InhibitingIdle        bool            `json:"inhibitingIdle"`
	XdgTag                string          `json:"xdgTag"`
	XdgDescription        string          `json:"xdgDescription"`
	ContentType           string          `json:"contentType"`
	TearingHint           bool            `json:"tearingHint"`
	StableId              string          `json:"stableId"`
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
func (w Window) MoveToWorkspaceSilent(target *Workspace, targetName string) {
	var workspaceAddress string

	if target != nil {
		workspaceAddress = target.Address
	} else {
		workspaceAddress = "name:" + targetName
	}

	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"%s\", follow = false, window = \"address:%s\"})", workspaceAddress, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

// MoveToWorkspace moves the window to the named workspace and switches the
// active view to follow it, sending a dispatch command to the Hyprland socket.
func (w Window) MoveToWorkspace(target *Workspace, targetName string) {
	var workspaceAddress string

	if target != nil {
		workspaceAddress = target.Address
	} else {
		workspaceAddress = "name:" + targetName
	}

	arg := fmt.Sprintf("hl.dsp.window.move({ workspace = \"%s\", follow = true, window = \"address:%s\"})", workspaceAddress, w.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}
