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
	if config != nil && config.Layouts != nil {
		activeWs := hd.GetActiveWorkspace()

		ll := make(ui.LayoutList, len(config.Layouts))

		for i, l := range config.Layouts {
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

		window.Resize(fyne.NewSize(ui.UIWindowWidth, (ui.UIListItemHeight * float32(len(config.Layouts)))))
		window.SetContent(container.NewStack(lw))
		window.Canvas().Focus(lw)
	} else {
		log.Fatal("Hyprdyn: No config or configured layouts present.")
	}
}

func RenameUi(window fyne.Window) {
	var onSubmit = func(input string) {
		if !hd.ValidWorkspaceName(input) {
			return
		}

		_, found := workspaces.FindByName(input)

		if !found {
			active := hd.GetActiveWorkspace()

			active.Rename(input)
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

		if *flags.SendMode == true {
			if follow {
				activeWindow.MoveToWorkspace(input)
			} else {
				activeWindow.MoveToWorkspaceSilent(input)
			}
		} else {
			existing, found := workspaces.FindByName(input)

			if found {
				existing.FocusOnCurrentMonitor()
			} else {
				hd.SpawnWorkspace(input)
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
