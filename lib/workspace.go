package hyprdyn

import (
	"fmt"
	"regexp"
)

type Workspace struct {
	WorkspaceType
	Monitor         string `json:"monitor"`
	MonitorID       int    `json:"monitorID"`
	Windows         int    `json:"windows"`
	HasFullScreen   bool   `json:"hasfullscreen"`
	LastWindow      string `json:"lastwindow"`
	LastWindowTitle string `json:"lastwindowtitle"`

	// hyprdyn attributes
	Active     bool
	Foreground bool
	Background bool
}

type WorkspaceType struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WorkspaceList []Workspace

var specialWorkspaceRegexp = regexp.MustCompile(`^special(:.*)?$`)
var workspaceNameRegexp = regexp.MustCompile(`^[A-Za-z0-9 ._:@?\/-]{1,255}$`)

func ValidWorkspaceName(name string) bool {
	return workspaceNameRegexp.MatchString(name) && !specialWorkspaceRegexp.MatchString(name)
}

func (ws Workspace) Rename(name string) {
	arg := fmt.Sprintf("hl.dsp.workspace.rename({ workspace = \"%s\", name = \"%s\" })", ws.Address, name)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

func (ws Workspace) FocusOnCurrentMonitor() {
	arg := fmt.Sprintf("hl.dsp.focus({ workspace = \"%s\", on_current_monitor = true }) ", ws.Address)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

func (wss WorkspaceList) FindByName(name string) (ws Workspace, found bool) {
	for _, ws := range wss {
		if ws.Name == name {
			return ws, true
		}
	}

	return Workspace{}, false
}

func (wss WorkspaceList) GetForegroundByMonitor(monitor string) *Workspace {
	for _, ws := range wss {
		if ws.Foreground == true && ws.Monitor == monitor {
			return &ws
		}
	}

	return nil
}

// NOTE: potential foot-gun here unless name is checked for collision prior to calling.
func SpawnWorkspace(name string) {
	arg := fmt.Sprintf("hl.dsp.focus({ workspace = \"name:%s\", on_current_monitor = true }) ", name)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

func GetAllWorkspaces(omitSpecial bool) WorkspaceList {
	var workspaces []Workspace

	response, err := hyprlandClient.sendCommand("workspaces", nil)
	Check(err)

	workspaces, err = UnmarshalHyprlandResponse[[]Workspace](response)
	Check(err)

	monitors := GetMonitors()
	activeWorkspace := GetActiveWorkspace()

	activeAddresses := make(map[string]bool)
	for _, m := range monitors {
		activeAddresses[m.ActiveWorkspace.Address] = true
	}

	var result []Workspace
	for _, ws := range workspaces {
		if ws.Address == activeWorkspace.Address {
			ws.Active = true
		}

		if activeAddresses[ws.Address] {
			ws.Foreground = true
		} else {
			ws.Background = true
		}

		if omitSpecial {
			if !specialWorkspaceRegexp.MatchString(ws.Name) {
				result = append(result, ws)

			}
		} else {
			result = append(result, ws)
		}
	}

	return result
}

func (wss WorkspaceList) Names() []string {
	var names []string

	for _, ws := range wss {
		names = append(names, ws.Name)
	}

	return names
}

func GetActiveWorkspace() Workspace {
	var activeWorkspace Workspace

	response, err := hyprlandClient.sendCommand("activeworkspace", nil)
	Check(err)

	activeWorkspace, err = UnmarshalHyprlandResponse[Workspace](response)
	Check(err)

	return activeWorkspace
}
