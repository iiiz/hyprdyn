package hyprdyn_themes

import (
	"image/color"
)

var DarkSky = HyprdynTheme{
	Background:      color.RGBA{R: 8, G: 10, B: 13, A: 255},
	InputBackground: color.RGBA{R: 17, G: 20, B: 25, A: 255},
	InputBorder:     color.RGBA{R: 42, G: 48, B: 56, A: 255},
	Placeholder:     color.RGBA{R: 115, G: 123, B: 135, A: 255},
	ListSeparator:   color.RGBA{R: 0, G: 0, B: 0, A: 0},
	Text:            color.RGBA{R: 242, G: 244, B: 247, A: 255},
	NewText:         color.RGBA{R: 255, G: 255, B: 255, A: 255},
	NewHighLight:    color.RGBA{R: 59, G: 70, B: 84, A: 255},
	Highlight:       color.RGBA{R: 36, G: 44, B: 54, A: 255},
	Suggestion:      color.RGBA{R: 174, G: 183, B: 195, A: 255},
	DisabledText:    color.RGBA{R: 92, G: 101, B: 113, A: 255},
}
