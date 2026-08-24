package hyprdyn

type Monitor struct {
	Id               int           `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Make             string        `json:"make"`
	Model            string        `json:"model"`
	Serial           string        `json:"serial"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	RefreshRate      float64       `json:"refreshRate"`
	X                int           `json:"x"`
	Y                int           `json:"y"`
	ActiveWorkspace  WorkspaceType `json:"activeWorkspace"`
	SpecialWorkspace WorkspaceType `json:"specialWorkspace"`
	Reserved         []int         `json:"reserved"`
	Scale            float64       `json:"scale"`
	Transform        int           `json:"transform"`
	Focused          bool          `json:"focused"`
	DpmsStatus       bool          `json:"dpmsStatus"`
	Vrr              bool          `json:"vrr"`
	ActivelyTearing  bool          `json:"activelyTearing"`
	CurrentFormat    string        `json:"currentFormat"`
	AvailableModes   []string      `json:"availableModes"`
}

func GetFocusedMonitor() Monitor {
	var focusedMonitor Monitor

	response, err := hyprlandClient.sendCommand("monitors all", nil)
	Check(err)

	monitors, err := UnmarshalHyprlandResponse[[]Monitor](response)
	Check(err)

	for _, mon := range monitors {
		if mon.Focused {
			focusedMonitor = mon
		}
	}

	return focusedMonitor
}
