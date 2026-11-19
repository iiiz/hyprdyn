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

// MoveToWorkspace moves the window to the named workspace. When silent is true
// the active view is not switched; otherwise Hyprland follows the move.
func (w Window) MoveToWorkspace(target *Workspace, targetName string, silent bool) {
	var workspaceAddress string
	created := false

	if target != nil {
		workspaceAddress = target.Address
	} else {
		if existing := GetAllWorkspaces(true).FindByName(targetName); existing != nil {
			workspaceAddress = existing.Address
		} else {
			workspaceAddress = newWorkspaceAddress()
			created = true
		}
	}

	follow := "true"
	if silent {
		follow = "false"
	}

	arg := fmt.Sprintf(
		"hl.dsp.window.move({ workspace = \"%s\", follow = %s, window = \"address:%s\" })",
		workspaceAddress,
		follow,
		w.Address,
	)
	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)

	if created {
		renameArg := fmt.Sprintf(
			"hl.dsp.workspace.rename({ workspace = \"%s\", name = \"%s\" })",
			workspaceAddress,
			targetName,
		)
		_, err := hyprlandClient.sendCommand("dispatch", &renameArg)
		Check(err)
	}
}
