package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/charmbracelet/log"

	hd "hyprdyn/lib"
	ui "hyprdyn/lib/ui"
)

func spawnGui() {
	hyprdynApp := app.NewWithID("iiiz.hyprdyn")

	if config != nil {
		if config.Theme != nil {
			ui.SetTheme(*config.Theme)
		} else if config.CustomTheme != nil {
			ui.SetCustomTheme(*config.CustomTheme)
		}
	}

	theme := ui.UseTheme()
	hyprdynApp.Settings().SetTheme(theme)

	window := hyprdynApp.NewWindow("hyprdyn")
	window.SetFixedSize(true)
	window.RequestFocus()

	if *flags.LayoutMode == true {
		LayoutUi(window)
	}

	if *flags.RenameMode == true {
		RenameUi(window)
	}

	if *flags.SelectMode == true || *flags.SendMode == true {
		SelectUi(window)
	}

	window.ShowAndRun()
}

func LayoutUi(window fyne.Window) {
	if config != nil && config.LayoutConfig != nil {
		quickToggleLen := len(config.LayoutConfig.QuickToggleOptions)

		if quickToggleLen > 0 {
			activeWs := hd.GetActiveWorkspace()

			ll := make(ui.LayoutList, quickToggleLen)

			for i, l := range config.LayoutConfig.QuickToggleOptions {
				ll[i] = &ui.Layout{Label: l}

				if l == activeWs.TiledLayout {
					ll[i].Active = true
					ll[i].Highlight = true
				}
			}

			var onDismiss = func() {
				os.Exit(0)
			}

			var onSubmit = func(layout string) {
				activeWs.SetLayout(layout)
				os.Exit(0)
			}

			lw := ui.NewLayoutWidget(ll, onSubmit, onDismiss)

			window.Resize(fyne.NewSize(ui.UIWindowWidth, (ui.UIListItemHeight * float32(quickToggleLen))))
			window.SetContent(container.NewStack(lw))
			window.Canvas().Focus(lw)
		} else {
			log.Fatal("Hyprdyn: No quickToggleOptions present in config.")
		}
	} else {
		log.Fatal("Hyprdyn: No config or configured layout settings present.")
	}
}

func RenameUi(window fyne.Window) {
	var onSubmit = func(input string) {
		if !hd.ValidWorkspaceName(input) {
			return
		}

		if workspaces.FindByName(input) == nil {
			active := hd.GetActiveWorkspace()

			active.Rename(input)

			if config != nil && config.LayoutConfig != nil && config.LayoutConfig.UseLayoutSetting.RenameEnabled() {
				if layoutName := config.GetMappedLayoutName(input); layoutName != nil {
					active.SetLayout(*layoutName)
				}
			}
		}

		os.Exit(0)
	}

	var onDismiss = func() {
		os.Exit(0)
	}

	r := ui.NewRenameWidget(onSubmit, onDismiss)
	window.Resize(fyne.NewSize(ui.UIWindowWidth, ui.UIRowHeight))

	window.SetContent(container.NewStack(r))
	window.Canvas().Focus(r)
}

func SelectUi(window fyne.Window) {
	workspaceNames := workspaces.Names()
	var autoComplete []string

	var onResize = func(height float32) {
		window.Resize(fyne.NewSize(ui.UIWindowWidth, height))
	}

	var onSubmit = func(input string, follow bool) {
		if !hd.ValidWorkspaceName(input) {
			return
		}

		var useLayoutSetting = config != nil && config.LayoutConfig != nil && config.LayoutConfig.UseLayoutSetting.SpawnEnabled()
		var layoutTarget *hd.Workspace
		existing := workspaces.FindByName(input)

		if *flags.SendMode == true {
			if follow {
				activeWindow.MoveToWorkspace(existing, input)

				if useLayoutSetting && existing == nil {
					active := hd.GetActiveWorkspace()
					layoutTarget = &active
				}
			} else {
				activeWindow.MoveToWorkspaceSilent(existing, input)

				if useLayoutSetting && existing == nil {
					layoutTarget = hd.GetAllWorkspaces(true).FindByName(input)

					if layoutTarget == nil {
						log.Warn("Hyprdyn: Workspace not ready, layout configuration may not be respected.")
					}
				}
			}
		} else {
			if existing != nil {
				existing.FocusOnCurrentMonitor()
			} else {
				hd.SpawnWorkspace(input)

				if useLayoutSetting {
					active := hd.GetActiveWorkspace()
					layoutTarget = &active
				}
			}
		}

		if useLayoutSetting && existing == nil && layoutTarget != nil {
			if layoutName := config.GetMappedLayoutName(input); layoutName != nil {
				layoutTarget.SetLayout(*layoutName)
			}
		}

		os.Exit(0)
	}

	var onDismiss = func() {
		os.Exit(0)
	}

	if config != nil && config.AutoComplete != nil {
		autoComplete = append(autoComplete, config.AutoComplete...)
	}

	selector, initialHeight := ui.NewSelectorWidget(workspaceNames, autoComplete, onSubmit, onResize, onDismiss)
	window.Resize(fyne.NewSize(ui.UIWindowWidth, initialHeight))

	window.SetContent(
		container.New(
			layout.NewBorderLayout(selector, nil, nil, nil),
			selector,
			container.NewStack(selector.TabCompleteListWidget),
		),
	)

	window.Canvas().Focus(selector)
}
