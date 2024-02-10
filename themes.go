package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.RGBA{R: 200, G: 200, B: 150, A: 255}
	}

	if name == theme.ColorNameButton {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.RGBA{R: 100, G: 200, B: 150, A: 255}
	}

	if name == theme.ColorNameForeground {
		if variant == theme.VariantLight {
			return color.Black
		}
		return color.RGBA{R: 120, G: 0, B: 225, A: 255}
	}

	if name == theme.ColorNameInputBackground {
		if variant == theme.VariantLight {
			return color.Black
		}
		return color.RGBA{R: 255, G: 0, B: 225, A: 255}
	}

	if name == theme.ColorNameInputBorder {
		if variant == theme.VariantLight {
			return color.Black
		}
		return color.RGBA{R: 255, G: 0, B: 225, A: 255}
	}

	// placeholder hint
	if name == theme.ColorNamePlaceHolder {
		if variant == theme.VariantLight {
			return color.Black
		}
		return color.RGBA{R: 0, G: 0, B: 225, A: 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
