package hyprdyn_themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type HyprdynHexcolorTheme struct {
	Background      string `validate:"required,hexcolor" json:"background"`
	InputBackground string `validate:"required,hexcolor" json:"inputBackground"`
	InputBorder     string `validate:"required,hexcolor" json:"inputBorder"`
	Placeholder     string `validate:"required,hexcolor" json:"placeholder"`
	ListSeparator   string `validate:"required,hexcolor" json:"listSeparator"`
	Text            string `validate:"required,hexcolor" json:"text"`
	NewText         string `validate:"required,hexcolor" json:"newText"`
	NewHighLight    string `validate:"required,hexcolor" json:"newHighLight"`
	Highlight       string `validate:"required,hexcolor" json:"highlight"`
	Suggestion      string `validate:"required,hexcolor" json:"suggestion"`
	DisabledText    string `validate:"required,hexcolor" json:"disabledText"`
}

type HyprdynTheme struct {
	// fyne theme colors
	Background      color.RGBA
	InputBackground color.RGBA
	InputBorder     color.RGBA
	Placeholder     color.RGBA
	ListSeparator   color.RGBA

	// hyprdyn custom theme attributes
	Text         color.RGBA
	NewText      color.RGBA
	NewHighLight color.RGBA
	Highlight    color.RGBA
	Suggestion   color.RGBA
	DisabledText color.RGBA
}

var _ fyne.Theme = (*HyprdynTheme)(nil)

func (t HyprdynTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary: // used by input border *focused
		return t.InputBorder
	case theme.ColorNameForeground: // input text color
		return t.Text
	case theme.ColorNameBackground:
		return t.Background
	case theme.ColorNameInputBackground:
		return t.InputBackground
	case theme.ColorNameInputBorder: // *no focus
		return t.InputBorder
	case theme.ColorNamePlaceHolder:
		return t.Placeholder
	case theme.ColorNameSeparator:
		return t.ListSeparator
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m HyprdynTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m HyprdynTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m HyprdynTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
