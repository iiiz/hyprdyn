package hyprdyn

// NOTE: Minimal attribute coverage.
type Monitor struct {
	Name            string        `json:"name"`
	ActiveWorkspace WorkspaceCore `json:"activeWorkspace"`
}

func GetMonitors() []Monitor {
	monitorResp, err := hyprlandClient.sendCommand("monitors", nil)
	Check(err)

	monitors, err := UnmarshalHyprlandResponse[[]Monitor](monitorResp)
	Check(err)

	return monitors
}
