package hyprdyn_themes

import (
	"image/color"
)

var Snow = HyprdynTheme{
	Background:      color.RGBA{R: 248, G: 250, B: 252, A: 255},
	InputBackground: color.RGBA{R: 255, G: 255, B: 255, A: 255},
	InputBorder:     color.RGBA{R: 203, G: 213, B: 225, A: 255},
	Placeholder:     color.RGBA{R: 100, G: 116, B: 139, A: 255},
	ListSeparator:   color.RGBA{R: 203, G: 213, B: 225, A: 255},
	Text:            color.RGBA{R: 30, G: 41, B: 59, A: 255},
	NewText:         color.RGBA{R: 37, G: 99, B: 235, A: 255},
	NewHighLight:    color.RGBA{R: 147, G: 197, B: 253, A: 255},
	Highlight:       color.RGBA{R: 148, G: 163, B: 184, A: 255},
	Suggestion:      color.RGBA{R: 71, G: 85, B: 105, A: 255},
	DisabledText:    color.RGBA{R: 100, G: 116, B: 139, A: 255},
}
