package hyprdyn

import (
	"flag"
	"fmt"
	"os"
)

type RuntimeFlags struct {
	IsUiMode bool

	// INFO: Non-Interactive commands
	SetupCmd        *bool
	PrimaryCmd      *bool
	ToggleLayoutCmd *bool

	// INFO: GUI modes
	SelectMode *bool
	SendMode   *bool
	RenameMode *bool
	LayoutMode *bool
}

func CaptureFlags() RuntimeFlags {
	var flags RuntimeFlags

	flags.SetupCmd = flag.Bool("setup", false, "Set configured monitors default workspace names. Useful on startup ie. ('exec-once')")
	flags.PrimaryCmd = flag.Bool("primary", false, "Go to, or spawn your primary workspace. See config:primaryName")
	flags.ToggleLayoutCmd = flag.Bool("layout-toggle", false, "Toggle the tiled layout of the active workspace.")

	flags.SelectMode = flag.Bool("select", false, "Select or create a workspace on current monitor.")
	flags.SendMode = flag.Bool("send", false, "Send the current window to a workspace.")
	flags.RenameMode = flag.Bool("rename", false, "Rename a workspace.")
	flags.LayoutMode = flag.Bool("layout", false, "Set workspace layout.")

	flag.Parse()

	flagCount := 0

	// INFO: Non-Interactive commands
	if *flags.SetupCmd == true {
		flagCount++
	}
	if *flags.PrimaryCmd == true {
		flagCount++
	}
	if *flags.ToggleLayoutCmd == true {
		flagCount++
	}

	// INFO: GUI modes
	if *flags.SelectMode == true {
		flagCount++
	}
	if *flags.SendMode == true {
		flagCount++
	}
	if *flags.RenameMode == true {
		flagCount++
	}
	if *flags.LayoutMode == true {
		flagCount++
	}

	if flagCount > 1 {
		PrintUsage()
		os.Exit(1)
	}

	// if we have a flag assume ui mode
	if flagCount == 1 && (*flags.SetupCmd != true && *flags.PrimaryCmd != true && *flags.ToggleLayoutCmd != true) {
		flags.IsUiMode = true
	}

	return flags
}

func PrintUsage() {
	fmt.Print("Usage: hyprdyn [OPTION]\nOPTIONS:\n")
	flag.PrintDefaults()
}
