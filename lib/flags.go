package hyprdyn

import (
	"flag"

	"github.com/charmbracelet/log"
)

type RuntimeFlags struct {
	IsUiMode bool

	// setup mode
	SetupMode *bool

	// normal modes
	SelectMode *bool
	SendMode   *bool
	RenameMode *bool
	PrimaryCmd *bool
}

func CaptureFlags() RuntimeFlags {
	var flags RuntimeFlags

	flags.SetupMode = flag.Bool("setup", false, "Set configured monitors default workspace names. Useful on startup ie. ('exec-once')")
	flags.SelectMode = flag.Bool("select", false, "Select or create a workspace on current monitor.")
	flags.SendMode = flag.Bool("send", false, "Send the current window to a workspace.")
	flags.RenameMode = flag.Bool("rename", false, "Rename a workspace.")
	flags.PrimaryCmd = flag.Bool("primary", false, "Go to, or spawn your primary workspace. See config:primaryName")

	flag.Parse()

	flagCount := 0

	if *flags.SetupMode == true {
		flagCount++
	}
	if *flags.SelectMode == true {
		flagCount++
	}
	if *flags.SendMode == true {
		flagCount++
	}
	if *flags.RenameMode == true {
		flagCount++
	}
	if *flags.PrimaryCmd == true {
		flagCount++
	}

	if flagCount > 1 {
		log.Fatal("Error: Flags 'select', 'send', 'rename', ... cannot be combined.")
	} else if flagCount == 0 {
		log.Fatal("Error: No flags specified.")
	}
	// if we have a flag assume ui mode
	if flagCount == 1 && *flags.SetupMode != true || *flags.PrimaryCmd != true {
		flags.IsUiMode = true
	}

	return flags
}
