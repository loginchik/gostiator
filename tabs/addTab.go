package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gostCituations/models"
	"gostCituations/tabs/components"
)

// TODO: develop file upload

// NewCitationWindow creates form window depending on models.CitationType
// that is passed to the function
func NewCitationWindow(t models.CitationType, application fyne.App) {
	var window = application.NewWindow(t.WindowName)
	var manualContent fyne.CanvasObject
	switch t.SystemName {
	case models.ArticleCT.SystemName:
		var form = models.NewArticleForm(5)
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.BookCT.SystemName:
		var form = models.NewBookForm(5, 3, 3, 3)
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.BookPartCT.SystemName:
		var form = models.NewBookPartForm(5, 3)
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.WebsiteCT.SystemName:
		var form = models.NewWebsiteForm()
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.WebtextCT.SystemName:
		var form = models.NewWebtextForm(5)
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.FilmCT.SystemName:
		var form = models.NewFilmForm()
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	case models.WebvideoCT.SystemName:
		var form = models.NewOnlineVideoForm()
		var formContainer = form.ToCanvasObject()
		manualContent = components.NewFormBox(form, formContainer, window, application)
	default:
		manualContent = widget.NewLabel(t.AppName)
	}
	var formWindowSize = fyne.NewSize(900, manualContent.MinSize().Height)
	window.Resize(formWindowSize)
	window.SetContent(manualContent)
	window.SetCloseIntercept(func() {
		dialog.ShowConfirm("Закрыть окно?", "Все данные будуту потеряны", func(choice bool) {
			if choice {
				window.Close()
			}
		}, window)
	})
	window.Show()
	window.CenterOnScreen()
	window.RequestFocus()
}

// newCitationButton creates a button. Button click toggles NewCitationWindow for
// specified models.CitationType
func newCitationButton(t models.CitationType, application fyne.App) *widget.Button {
	var button = widget.NewButton(t.AppName, func() {
		NewCitationWindow(t, application)
	})
	button.SetIcon(theme.ContentAddIcon())
	return button
}

// AddWindow creates app start window that shows all the available objects from
// models.CitationType. Returns fyne.Container with buttons (each for one citation type)
// as grid elements
func AddWindow(application fyne.App) *fyne.Container {
	var buttonsGrid = container.NewAdaptiveGrid(3)
	for _, t := range models.CitationTypeOptions {
		var button = newCitationButton(t, application)
		var bigLabel = widget.NewLabel(t.Description)
		bigLabel.Wrapping = fyne.TextWrapWord
		var cont = container.NewBorder(button, nil, nil, nil, bigLabel)
		buttonsGrid.Add(cont)
	}
	return buttonsGrid
}
