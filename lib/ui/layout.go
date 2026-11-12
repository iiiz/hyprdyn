package hyprdyn_ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Layout struct {
	Label     string
	Highlight bool
	Active    bool
	text      *canvas.Text
}

type LayoutList []*Layout

func (l *Layout) textReset() {
	theme := UseTheme()
	if l.Active {
		l.text.Color = theme.DisabledText
	} else {
		l.text.Color = theme.Text
	}
	l.text.TextStyle = fyne.TextStyle{}
	l.text.TextSize = UITextSize
}

func (l *Layout) GetStyledText() fyne.CanvasObject {
	theme := UseTheme()

	if l.text == nil {
		var c color.Color

		if l.Active {
			c = theme.DisabledText
		} else {
			c = theme.Text
		}

		l.text = canvas.NewText(l.Label, c)
	} else {
		l.textReset()
	}

	if l.Highlight {
		l.text.Text = l.Label
		l.text.TextStyle = fyne.TextStyle{Bold: true}
		l.text.Color = theme.Highlight

		return l.text
	}

	return l.text
}
