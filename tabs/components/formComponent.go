package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gostCituations/models"
	"strings"
	"time"
)

func ClearButton(form models.CitationInterface, formContainer fyne.CanvasObject) *widget.Button {
	return widget.NewButtonWithIcon("Очистить", theme.ContentClearIcon(), func() {
		form.ClearForm()
		formContainer.Refresh()
	})
}

func ExampleButton(form models.CitationInterface, formContainer fyne.CanvasObject) *widget.Button {
	return widget.NewButtonWithIcon("Пример заполнения", theme.InfoIcon(), func() {
		form.ExampleForm()
		formContainer.Refresh()
	})
}

func HelpButton() *widget.Button {
	return widget.NewButtonWithIcon("Помощь", theme.HelpIcon(), func() {

	})
}

func TopButtons(form models.CitationInterface, formContainer fyne.CanvasObject) *fyne.Container {
	var buttons = []fyne.CanvasObject{
		ExampleButton(form, formContainer),
		ClearButton(form, formContainer),
		//HelpButton(),
	}
	return container.NewAdaptiveGrid(len(buttons), buttons...)
}

func handleGenerationRequest(form models.CitationInterface, parentWindow fyne.Window, application fyne.App) {
	if form.ValidateForm() {
		var citation = []string{form.Citation()}
		var resultWindow = ResultWindow(citation, application)
		resultWindow.Show()
		var historyRecord = models.HistoryRecord{
			DateStamp:  time.Now(),
			Content:    citation[0],
			RecordType: form.HistoryRecordType(),
		}
		historyRecord.Save()
		parentWindow.Close()
	} else {
		var messageText = strings.Join(form.ErrorText(), "\n")
		var dialogBlock = dialog.NewInformation("Недостаточно данных для генерации цитаты", messageText, parentWindow)
		dialogBlock.Show()
	}
}

func NewFormBox(form models.CitationInterface, formContainer fyne.CanvasObject, parentWindow fyne.Window, application fyne.App) fyne.CanvasObject {
	var topButtonsContainer = TopButtons(form, formContainer)
	var generateButton = widget.NewButton("", func() {
		handleGenerationRequest(form, parentWindow, application)
	})
	generateButton.SetIcon(theme.DocumentSaveIcon())

	var generateShortcut = &desktop.CustomShortcut{Modifier: fyne.KeyModifierControl, KeyName: fyne.KeyG}
	var exampleShortcut = &desktop.CustomShortcut{Modifier: fyne.KeyModifierControl, KeyName: fyne.KeyE}
	parentWindow.Canvas().AddShortcut(generateShortcut, func(sh fyne.Shortcut) {
		handleGenerationRequest(form, parentWindow, application)
	})
	parentWindow.Canvas().AddShortcut(exampleShortcut, func(sh fyne.Shortcut) {
		form.ExampleForm()
		formContainer.Refresh()
	})

	return container.NewVBox(topButtonsContainer, formContainer, generateButton)
}
