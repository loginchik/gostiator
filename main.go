package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func setupUI() {
	var application = app.New()
	var wind = application.NewWindow("Hello!")
	wind.Resize(fyne.NewSize(500, 400))
	wind.ShowAndRun()
}

func main() {
	setupUI()
}
