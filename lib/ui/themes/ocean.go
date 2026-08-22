package hyprdyn_themes

import (
	"image/color"
)

var Ocean = HyprdynTheme{
	Background:      color.RGBA{R: 6, G: 43, B: 58, A: 255},
	InputBackground: color.RGBA{R: 11, G: 64, B: 82, A: 255},
	InputBorder:     color.RGBA{R: 23, G: 98, B: 122, A: 255},
	Placeholder:     color.RGBA{R: 127, G: 175, B: 190, A: 255},
	ListSeparator:   color.RGBA{R: 23, G: 157, B: 158, A: 255},
	Text:            color.RGBA{R: 230, G: 247, B: 250, A: 255},
	NewText:         color.RGBA{R: 255, G: 255, B: 255, A: 255},
	NewHighLight:    color.RGBA{R: 22, G: 135, B: 167, A: 255},
	Highlight:       color.RGBA{R: 14, G: 96, B: 120, A: 255},
	Suggestion:      color.RGBA{R: 169, G: 217, B: 227, A: 255},
	DisabledText:    color.RGBA{R: 104, G: 144, B: 157, A: 255},
}
