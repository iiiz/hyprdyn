package hyprdyn_themes

import (
	"image/color"
)

var Emerald = HyprdynTheme{
	Background:      color.RGBA{R: 6, G: 26, B: 20, A: 255},
	InputBackground: color.RGBA{R: 13, G: 43, B: 33, A: 255},
	InputBorder:     color.RGBA{R: 23, G: 107, B: 77, A: 255},
	Placeholder:     color.RGBA{R: 134, G: 183, B: 165, A: 255},
	ListSeparator:   color.RGBA{R: 23, G: 107, B: 77, A: 255},
	Text:            color.RGBA{R: 229, G: 246, B: 238, A: 255},
	NewText:         color.RGBA{R: 154, G: 230, B: 193, A: 255},
	NewHighLight:    color.RGBA{R: 24, G: 166, B: 106, A: 255},
	Highlight:       color.RGBA{R: 23, G: 107, B: 77, A: 255},
	Suggestion:      color.RGBA{R: 165, G: 214, B: 190, A: 255},
	DisabledText:    color.RGBA{R: 99, G: 140, B: 123, A: 255},
}
