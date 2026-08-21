package hyprdyn_ui

import (
	"image/color"
	"strconv"

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
		Background:      hexToRGBA(customTheme.Background),
		InputBackground: hexToRGBA(customTheme.InputBackground),
		InputBorder:     hexToRGBA(customTheme.InputBorder),
		Placeholder:     hexToRGBA(customTheme.Placeholder),
		Text:            hexToRGBA(customTheme.Text),
		NewText:         hexToRGBA(customTheme.NewText),
		NewHighLight:    hexToRGBA(customTheme.NewHighLight),
		Highlight:       hexToRGBA(customTheme.Highlight),
		Suggestion:      hexToRGBA(customTheme.Suggestion),
		DisabledText:    hexToRGBA(customTheme.DisabledText),
	}
}

func hexToRGBA(hex string) color.RGBA {
	values, _ := strconv.ParseUint(string(hex[1:]), 16, 32)

	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}
}
