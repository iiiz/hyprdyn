package hyprdyn_ui

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/charmbracelet/log"

	ht "hyprdyn/lib/ui/themes"
)

var theme = "default"
var useCustom = false
var userTheme ht.HyprdynTheme

func UseTheme() ht.HyprdynTheme {
	if !useCustom {
		switch theme {
		case "default":
			return ht.DefaultTheme
		case "emerald":
			return ht.Emerald
		case "cyber":
			return ht.Cyber
		case "nordly":
			return ht.Nordly
		case "ruby":
			return ht.Ruby
		case "snow":
			return ht.Snow
		case "darksky":
			return ht.DarkSky
		case "ocean":
			return ht.Ocean
		}
	} else {
		return userTheme
	}

	return ht.DefaultTheme
}

func SetTheme(name string) {
	theme = name
}

func SetCustomTheme(customTheme ht.HyprdynHexcolorTheme) {
	useCustom = true
	userTheme = ht.HyprdynTheme{
		Background:      mustHexToRGBA(customTheme.Background),
		InputBackground: mustHexToRGBA(customTheme.InputBackground),
		InputBorder:     mustHexToRGBA(customTheme.InputBorder),
		Placeholder:     mustHexToRGBA(customTheme.Placeholder),
		ListSeparator:   mustHexToRGBA(customTheme.ListSeparator),
		Text:            mustHexToRGBA(customTheme.Text),
		NewText:         mustHexToRGBA(customTheme.NewText),
		NewHighLight:    mustHexToRGBA(customTheme.NewHighLight),
		Highlight:       mustHexToRGBA(customTheme.Highlight),
		Suggestion:      mustHexToRGBA(customTheme.Suggestion),
		DisabledText:    mustHexToRGBA(customTheme.DisabledText),
	}
}

func mustHexToRGBA(hex string) color.RGBA {
	colorValue, err := hexToRGBA(hex)
	if err != nil {
		log.Fatal(err)
	}

	return colorValue
}

func hexToRGBA(hex string) (color.RGBA, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return color.RGBA{}, fmt.Errorf("invalid hex color %q, expected format #FFFFFF", hex)
	}

	values, err := strconv.ParseUint(hex[1:], 16, 32)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color %q: %w", hex, err)
	}

	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}, nil
}
