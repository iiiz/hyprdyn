package main

import (
	"flag"
	"os"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/charmbracelet/log"
	"github.com/thiagokokada/hyprland-go"

	hd "hyprdyn/lib"
	ui "hyprdyn/lib/ui"
)

var isUiMode bool
var selectMode *bool
var sendMode *bool
var renameMode *bool
var activeWindow *hyprland.Window
var workspaces []hd.Workspace

func init() {
	// TODO: check for a config / use config for styling and prefs?
	// TODO: use cobra or pflag for better handling and mutually exclusive flags?
	selectMode = flag.Bool("select", false, "Select or create a workspace on current monitor.")
	sendMode = flag.Bool("send", false, "Send the current window to a workspace.")
	renameMode = flag.Bool("rename", false, "Rename a workspace.")
	flag.Parse()

	flagCount := 0

	if *selectMode == true {
		flagCount++
	}
	if *sendMode == true {
		flagCount++
	}
	if *renameMode == true {
		flagCount++
	}

	if flagCount > 1 {
		log.Fatal("Error: Flags 'select', 'send' and 'rename' cannot be combined.")
	}

	// Running with a
	if flagCount == 1 {
		isUiMode = true
	}

	hd.GetHyprClient()
}

func main() {
	workspaces = hd.GetAllWorkspaces(true)
	activeWindow = hd.GetActiveWindow()

	if isUiMode {
		spawnUi()
	} else {
		focusedMonitor := hd.GetFocusedMonitor()

		var focused string
		var active []string
		var background []string

		for _, ws := range workspaces {
			hd.PrettyPrint(ws)
			log.Info("Ws", "type", ws.WorkspaceType)

			if ws.LastWindow == "0x0" {
				background = append(background, ws.Name)
			} else {

				if ws.MonitorID == focusedMonitor.Id {
					focused = ws.Name
				} else {
					active = append(active, ws.Name)
				}
			}
		}

		log.Info("Workspaces", "focused", focused, "active", active, "background", background)
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
	if *renameMode == true {
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
	if *selectMode == true || *sendMode == true {
		workspaceNames := hd.GetAllWorkspaceNames(true)

		var onResize = func(height float32) {
			window.Resize(fyne.NewSize(300, height))
		}

		var onSubmit = func(input string) {
			// INFO: Refuse to switch/spawn special workspace, looks to be unsupported https://wiki.hypr.land/Configuring/Dispatchers/#workspaces
			if specialRegexp.MatchString(input) {
				return
			}

			if *sendMode == true {
				hd.MoveWindowToWorkspaceSilent(*activeWindow, input)
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

		selector, initialHeight := ui.NewSelectorWidget(workspaceNames, onSubmit, onResize, onDismiss)
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
