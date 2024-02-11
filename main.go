package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"gostCituations/tabs"
)

func setupUI() {
	var application = app.New()
	var wind = application.NewWindow(application.Metadata().Name)
	wind.SetCloseIntercept(func() {
		if len(application.Driver().AllWindows()) > 1 {
			dialog.ShowConfirm(
				"Закрыть приложение?",
				"Все несохранённые записи будут удалены",
				func(choice bool) {
					if choice {
						application.Quit()
					}
				}, wind)
		} else {
			application.Quit()
		}
	})
	wind.SetContent(tabs.AddWindow(application))
	wind.SetMainMenu(MainMenu(application))
	wind.Show()
	application.Run()
}

func main() {
	setupUI()
}
