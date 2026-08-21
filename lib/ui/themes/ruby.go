package hyprdyn_themes

import (
	"image/color"
)

var Ruby = HyprdynTheme{
	Background:      color.RGBA{R: 26, G: 8, B: 11, A: 255},
	InputBackground: color.RGBA{R: 42, G: 13, B: 18, A: 255},
	InputBorder:     color.RGBA{R: 127, G: 29, B: 45, A: 255},
	Placeholder:     color.RGBA{R: 201, G: 139, B: 149, A: 255},
	ListSeparator:   color.RGBA{R: 98, G: 19, B: 35, A: 255},
	Text:            color.RGBA{R: 252, G: 232, B: 236, A: 255},
	NewText:         color.RGBA{R: 255, G: 179, B: 193, A: 255},
	NewHighLight:    color.RGBA{R: 230, G: 57, B: 86, A: 255},
	Highlight:       color.RGBA{R: 143, G: 29, B: 44, A: 255},
	Suggestion:      color.RGBA{R: 232, G: 160, B: 170, A: 255},
	DisabledText:    color.RGBA{R: 143, G: 104, B: 112, A: 255},
}
