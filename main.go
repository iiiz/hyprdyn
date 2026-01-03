package main

import (
	"os"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	hd "hyprdyn/lib"
	ui "hyprdyn/lib/ui"
)

var config hd.Config
var flags hd.RuntimeFlags
var activeWindow hd.Window
var workspaces hd.WorkspaceList

func init() {
	if c := hd.ReadConfig(); c != nil {
		config = *c
	}

	flags = hd.CaptureFlags()

	hd.GetHyprClient()
}

func main() {
	workspaces = hd.GetAllWorkspaces(true)
	activeWindow = hd.GetActiveWindow()

	if *flags.SetupMode == true {
		for _, monitorConfig := range config.Monitors {
			ws := workspaces.GetForegroundByMonitor(monitorConfig.Id)

			if ws != nil && monitorConfig.DefaultName != nil {
				ws.Rename(*monitorConfig.DefaultName)
			}
		}

		os.Exit(0)
	}

	if *flags.PrimaryCmd == true && config.PrimaryName != nil {
		var existingWorkspace *hd.Workspace

		for _, ws := range workspaces {
			if ws.Name == *config.PrimaryName {
				existingWorkspace = &ws
			}
		}

		if existingWorkspace != nil {
			existingWorkspace.FocusOnCurrentMonitor()
		} else {
			hd.SpawnWorkspace(*config.PrimaryName)
		}

		os.Exit(0)
	}

	if flags.IsUiMode {
		spawnUi()
	} else {
		// TODO: re-implement with better methods
		// focusedMonitor := hd.GetFocusedMonitor()
		//
		// var focused string
		// var active []string
		// var background []string
		//
		// for _, ws := range workspaces {
		// 	if ws.LastWindow == "0x0" {
		// 		background = append(background, ws.Name)
		// 	} else {
		//
		// 		if ws.MonitorID == focusedMonitor.Id {
		// 			focused = ws.Name
		// 		} else {
		// 			active = append(active, ws.Name)
		// 		}
		// 	}
		// }
		//
		// log.Info("Workspaces:", "focused", focused, "active", active, "background", background)
	}
}

func spawnUi() {
	hyprdynApp := app.New()
	window := hyprdynApp.NewWindow("hyprdyn")
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.RequestFocus()

	specialRegexp := regexp.MustCompile("^special(?:[:]{1}.*)*$")

	var onDismiss = func() {
		os.Exit(0)
	}

	/**
	* Rename Mode
	**/
	if *flags.RenameMode == true {
		var onSubmit = func(input string) {
			var existingWorkspace *hd.Workspace

			// INFO: Refuse to switch/spawn special workspace, looks to be unsupported https://wiki.hypr.land/Configuring/Dispatchers/#workspaces
			if specialRegexp.MatchString(input) {
				return
			}

			for _, ws := range workspaces {
				if ws.Name == input {
					existingWorkspace = &ws
				}
			}

			if existingWorkspace == nil {
				active := hd.GetActiveWorkspace()

				active.Rename(input)
			}

			os.Exit(0)
		}

		r := ui.NewRenameWidget(onSubmit, onDismiss)
		window.Resize(fyne.NewSize(300, 42))

		window.SetContent(container.NewStack(r))
		window.Canvas().Focus(r)
	}

	/**
	* Select Mode
	**/
	if *flags.SelectMode == true || *flags.SendMode == true {
		workspaceNames := hd.GetAllWorkspaceNames(true)

		var onResize = func(height float32) {
			window.Resize(fyne.NewSize(300, height))
		}

		var onSubmit = func(input string) {
			// INFO: Refuse to switch/spawn special workspace, looks to be unsupported https://wiki.hypr.land/Configuring/Dispatchers/#workspaces
			if specialRegexp.MatchString(input) {
				return
			}

			if *flags.SendMode == true {
				activeWindow.MoveToWorkspaceSilent(input)
			} else {
				var existingWorkspace *hd.Workspace

				for _, ws := range workspaces {
					if ws.Name == input {
						existingWorkspace = &ws
					}
				}

				if existingWorkspace != nil {
					existingWorkspace.FocusOnCurrentMonitor()
				} else {
					hd.SpawnWorkspace(input)
				}
			}

			os.Exit(0)
		}

		selector, initialHeight := ui.NewSelectorWidget(workspaceNames, config.AutoComplete, onSubmit, onResize, onDismiss)
		window.Resize(fyne.NewSize(300, initialHeight))

		window.SetContent(
			container.New(
				layout.NewBorderLayout(selector, nil, nil, nil),
				selector,
				container.NewStack(selector.TabCompleteListWidget),
			),
		)

		window.Canvas().Focus(selector)
	}

	window.ShowAndRun()
}
