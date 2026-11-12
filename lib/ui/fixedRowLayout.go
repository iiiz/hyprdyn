package hyprdyn_ui

import "fyne.io/fyne/v2"

type fixedRowLayout struct{}

func (l *fixedRowLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, UIRowHeight)
}

func (l *fixedRowLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(fyne.NewSize(size.Width, UIRowHeight))
		o.Move(fyne.NewPos(UIHorizontalPadding, 0))
	}
}
