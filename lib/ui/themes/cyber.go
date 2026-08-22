package hyprdyn_themes

import (
	"image/color"
)

var Cyber = HyprdynTheme{
	Background:      color.RGBA{R: 8, G: 11, B: 15, A: 255},
	InputBackground: color.RGBA{R: 112, G: 120, B: 137, A: 255},
	InputBorder:     color.RGBA{R: 57, G: 60, B: 71, A: 255},
	Placeholder:     color.RGBA{R: 57, G: 60, B: 71, A: 255},
	ListSeparator:   color.RGBA{R: 57, G: 60, B: 71, A: 255},
	Text:            color.RGBA{R: 86, G: 219, B: 233, A: 255},
	NewText:         color.RGBA{R: 86, G: 231, B: 100, A: 255},
	NewHighLight:    color.RGBA{R: 230, G: 86, B: 146, A: 255},
	Highlight:       color.RGBA{R: 218, G: 231, B: 87, A: 255},
	Suggestion:      color.RGBA{R: 156, G: 156, B: 158, A: 255},
	DisabledText:    color.RGBA{R: 230, G: 100, B: 87, A: 255},
}
