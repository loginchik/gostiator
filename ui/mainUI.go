package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"gostCituations/ui/menu"
	"gostCituations/ui/services"
)

func SetupUI() {
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
	wind.SetContent(services.AddWindow(application))
	wind.SetMainMenu(menu.MainMenu(application))
	wind.Show()
	application.Run()
}
