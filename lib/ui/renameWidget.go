package hyprdyn_ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type RenameWidget struct {
	widget.Entry

	onSubmit  func(i string)
	onDismiss func()
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

func (rw *RenameWidget) FocusLost() {
	rw.onDismiss()
}

func (rw *RenameWidget) KeyDown(key *fyne.KeyEvent) {
	if key.Name == "Escape" {
		rw.onDismiss()
	}
}
