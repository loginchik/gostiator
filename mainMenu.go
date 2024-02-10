package main

import (
	"fyne.io/fyne/v2"
	"gostCituations/models"
	"gostCituations/tabs"
)

// historyMenu "Архив" fyne.Menu element.
func historyMenu(application fyne.App) *fyne.Menu {
	var menu = fyne.NewMenu("Архив", fyne.NewMenuItem("Показать", func() {
		var historyWindow = application.NewWindow("Архив")
		historyWindow.SetContent(tabs.HistoryWindow())
		historyWindow.Resize(fyne.NewSize(600, 400))
		historyWindow.Show()
		historyWindow.CenterOnScreen()
	}))
	return menu
}

// newCitationMenu creates "Создать" fyne.Menu element with fyne.MenuItem for
// each models.CitationType available at the moment. In fact,
// it repeats the options from app starting window, though it can be useful
func newCitationMenu(application fyne.App) *fyne.Menu {
	var article = fyne.NewMenuItem(models.ArticleCT.AppName, func() {
		tabs.NewCitationWindow(models.ArticleCT, application)
	})
	var conference = fyne.NewMenuItem(models.ConferenceCT.AppName, func() {
		tabs.NewCitationWindow(models.ConferenceCT, application)
	})
	var book = fyne.NewMenuItem(models.BookCT.AppName, func() {
		tabs.NewCitationWindow(models.BookCT, application)
	})
	var bookpart = fyne.NewMenuItem(models.BookPartCT.AppName, func() {
		tabs.NewCitationWindow(models.BookPartCT, application)
	})
	var website = fyne.NewMenuItem(models.WebsiteCT.AppName, func() {
		tabs.NewCitationWindow(models.WebsiteCT, application)
	})
	var webtext = fyne.NewMenuItem(models.WebtextCT.AppName, func() {
		tabs.NewCitationWindow(models.WebtextCT, application)
	})
	var webvideo = fyne.NewMenuItem(models.WebvideoCT.AppName, func() {
		tabs.NewCitationWindow(models.WebvideoCT, application)
	})
	var film = fyne.NewMenuItem(models.FilmCT.AppName, func() {
		tabs.NewCitationWindow(models.FilmCT, application)
	})

	var menuOptions = []*fyne.MenuItem{
		article, conference, book, bookpart,
		website, webtext, webvideo, film,
	}
	var menu = fyne.NewMenu("Создать", menuOptions...)
	return menu
}

// mainMenu creates fyne.MainMenu from the output of newCitationMenu and historyMenu.
func mainMenu(application fyne.App) *fyne.MainMenu {
	var menu = fyne.NewMainMenu(
		newCitationMenu(application),
		historyMenu(application),
	)
	return menu
}
