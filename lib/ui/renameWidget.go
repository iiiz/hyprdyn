package hyprdyn_ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type RenameWidget struct {
	widget.Entry

	onSubmit     func(i string)
	onDismiss    func()
	dismissTimer *time.Timer
}

func NewRenameWidget(OnSubmit func(i string), OnDismiss func()) *RenameWidget {
	rename := &RenameWidget{}
	rename.SetPlaceHolder("Rename Workspace")
	rename.onSubmit = OnSubmit
	rename.onDismiss = OnDismiss

	rename.OnSubmitted = func(input string) {
		rename.onSubmit(input)
	}

	rename.ExtendBaseWidget(rename)

	return rename
}

func (rw *RenameWidget) FocusGained() {
	rw.Entry.FocusGained()

	if rw.dismissTimer != nil {
		rw.dismissTimer.Stop()
		rw.dismissTimer = nil
	}
}

func (rw *RenameWidget) FocusLost() {
	rw.Entry.FocusLost()

	if rw.dismissTimer != nil {
		rw.dismissTimer.Stop()
	}

	rw.dismissTimer = time.AfterFunc(dismissGrace, rw.onDismiss)
}

func (rw *RenameWidget) KeyDown(key *fyne.KeyEvent) {
	if key.Name == "Escape" {
		rw.onDismiss()
	}
}
