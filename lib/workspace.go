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
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type WorkspaceList []Workspace

func (ws Workspace) Rename(name string) {
	arg := fmt.Sprintf("hl.dsp.workspace.rename({ workspace = \"%d\", name = \"%s\" })", ws.Id, name)

	_, err := hyprlandClient.sendCommand("dispatch", &arg)
	Check(err)
}

func (ws Workspace) FocusOnCurrentMonitor() {
	arg := fmt.Sprintf("hl.dsp.focus({ workspace = \"name:%s\", on_current_monitor = true }) ", ws.Name)

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

	activeWorkspace := GetActiveWorkspace()
	specialRegexp := regexp.MustCompile("^special(?:[:]{1}.*)*$")

	var result []Workspace
	for _, ws := range workspaces {
		if ws.Id == activeWorkspace.Id {
			ws.Active = true
		}

		if ws.LastWindow == "0x0" && ws.Windows > 0 {
			ws.Background = true
		} else {
			ws.Foreground = true
		}

		if omitSpecial {
			if !specialRegexp.MatchString(ws.Name) {
				result = append(result, ws)

			}
		} else {
			result = append(result, ws)
		}
	}

	return result
}

func GetAllWorkspaceNames(omitSpecial bool) []string {
	var workspaces []Workspace

	response, err := hyprlandClient.sendCommand("workspaces", nil)
	Check(err)

	workspaces, err = UnmarshalHyprlandResponse[[]Workspace](response)
	Check(err)

	var workspaceNames []string
	specialRegexp := regexp.MustCompile("^special(?:[:]{1}.*)*$")

	for _, ws := range workspaces {
		if omitSpecial {
			if !specialRegexp.MatchString(ws.Name) {
				workspaceNames = append(workspaceNames, ws.Name)
			}
		} else {
			workspaceNames = append(workspaceNames, ws.Name)
		}
	}

	return workspaceNames
}

func GetActiveWorkspace() Workspace {
	var activeWorkspace Workspace

	response, err := hyprlandClient.sendCommand("activeworkspace", nil)
	Check(err)

	activeWorkspace, err = UnmarshalHyprlandResponse[Workspace](response)
	Check(err)

	return activeWorkspace
}
