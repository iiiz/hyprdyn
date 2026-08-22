package hyprdyn_themes

import (
	"image/color"
)

var DefaultTheme = HyprdynTheme{
	Background:      color.RGBA{R: 10, G: 10, B: 10, A: 255},
	InputBackground: color.RGBA{R: 25, G: 31, B: 48, A: 255},
	InputBorder:     color.RGBA{R: 86, G: 118, B: 159, A: 255},
	Placeholder:     color.RGBA{R: 70, G: 79, B: 99, A: 255},
	ListSeparator:   color.RGBA{R: 36, G: 68, B: 99, A: 255},
	Text:            color.RGBA{R: 255, G: 255, B: 255, A: 255},
	NewText:         color.RGBA{R: 0, G: 255, B: 0, A: 255},
	NewHighLight:    color.RGBA{R: 110, G: 190, B: 255, A: 255},
	Highlight:       color.RGBA{R: 90, G: 90, B: 255, A: 255},
	Suggestion:      color.RGBA{R: 180, G: 180, B: 180, A: 255},
	DisabledText:    color.RGBA{R: 255, G: 90, B: 90, A: 255},
}
