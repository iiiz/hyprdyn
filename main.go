package main

import (
	"os"

	hd "hyprdyn/lib"
)

var config *hd.Config
var flags hd.RuntimeFlags
var activeWindow hd.Window
var workspaces hd.WorkspaceList

func init() {
	config = hd.ReadConfig()
	flags = hd.CaptureFlags()
	hd.InitHyprlandClient()
}

func main() {
	workspaces = hd.GetAllWorkspaces(true)
	activeWindow = hd.GetActiveWindow()

	if *flags.SetupCmd == true {
		SetupCommand()
	}

	if *flags.PrimaryCmd == true {
		PrimaryCommand()
	}

	if *flags.ToggleLayoutCmd == true {
		ToggleLayoutCommand()
	}

	if flags.IsUiMode {
		spawnGui()
	} else {

		// TODO: Maybe string format output of active, foreground and background WS, not in scope for now.
		hd.PrintUsage()
		os.Exit(1)
	}
}
