package hyprdyn

import (
	"fmt"
	"regexp"

	"github.com/thiagokokada/hyprland-go"
)

// shallow embed to ease separation of concerns and extend funcs
type Workspace struct {
	hyprland.Workspace
}

func (ws Workspace) Rename(name string) {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("renameworkspace %d %s", ws.Id, name))
	Check(err)
}

func (ws Workspace) FocusOnCurrentMonitor() {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("focusworkspaceoncurrentmonitor name:%s", ws.Name))
	Check(err)
}

// NOTE: potential foot-gun here unless name is checked for collision prior to calling.
func SpawnWorkspace(name string) {
	_, err := hyprlandClient.Dispatch(fmt.Sprintf("workspace name:%s", name))
	Check(err)
}

func GetAllWorkspaces(omitSpecial bool) []Workspace {
	hyprWorkspaces, err := hyprlandClient.Workspaces()
	Check(err)

	var workspaces []Workspace
	specialRegexp := regexp.MustCompile("^special(?:[:]{1}.*)*$")

	for _, ws := range hyprWorkspaces {
		if omitSpecial {
			if !specialRegexp.MatchString(ws.Name) {
				workspaces = append(workspaces, Workspace{Workspace: ws})
			}
		} else {
			workspaces = append(workspaces, Workspace{Workspace: ws})
		}
	}

	return workspaces
}

func GetAllWorkspaceNames(omitSpecial bool) []string {
	workspaces, err := hyprlandClient.Workspaces()
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
	hyprWorkspace, err := hyprlandClient.ActiveWorkspace()
	Check(err)

	return Workspace{Workspace: hyprWorkspace}
}
