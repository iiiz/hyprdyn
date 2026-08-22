package hyprdyn_themes

import (
	"image/color"
)

var Nordly = HyprdynTheme{
	Background:      color.RGBA{R: 42, G: 48, B: 60, A: 255},
	InputBackground: color.RGBA{R: 55, G: 61, B: 78, A: 255},
	InputBorder:     color.RGBA{R: 86, G: 118, B: 159, A: 255},
	Placeholder:     color.RGBA{R: 70, G: 79, B: 99, A: 255},
	ListSeparator:   color.RGBA{R: 76, G: 98, B: 139, A: 255},
	Text:            color.RGBA{R: 195, G: 201, B: 211, A: 255},
	NewText:         color.RGBA{R: 148, G: 172, B: 129, A: 255},
	NewHighLight:    color.RGBA{R: 163, G: 130, B: 159, A: 255},
	Highlight:       color.RGBA{R: 124, G: 174, B: 190, A: 255},
	Suggestion:      color.RGBA{R: 86, G: 118, B: 158, A: 255},
	DisabledText:    color.RGBA{R: 174, G: 89, B: 100, A: 255},
}
