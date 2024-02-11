package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"strings"
	"time"
)

/* ------------ Website Form Structure  ------------ */

type WebsiteForm struct {
	CitationForm
}

func (form *WebsiteForm) ExampleForm() {
	form.Title.Example()
	form.Description.Example()
	form.URL.Example()
}

func (form *WebsiteForm) ClearForm() {
	form.Title.Clear()
	form.Description.Clear()
	form.URL.Clear()
	form.Errors = []string{}
}

func (form *WebsiteForm) ValidateForm() bool {
	form.Errors = []string{}

	if !(form.Title.ValidateRequired()) {
		form.Errors = append(form.Errors, form.Title.ValidationErrors.Empty)
	}
	if !(form.URL.ValidateRequired()) {
		form.Errors = append(form.Errors, form.URL.ValidationErrors.Empty)
	}

	return len(form.Errors) == 0
}

func (form *WebsiteForm) Citation() string {
	var referenceFirstPart = form.Title.TrimText()
	if form.Description.TrimText() != "" {
		referenceFirstPart += fmt.Sprintf(": %s.", form.Description.TrimText())
	}
	var referenceDate = time.Now().Format("02.01.2006")
	var urlString = fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), referenceDate)
	var citationParts = []string{referenceFirstPart, urlString, "— Текст: электронный."}
	return strings.Join(citationParts, " ")
}

func (form *WebsiteForm) HistoryRecordType() string {
	return WebsiteCT.SystemName
}

func (form *WebsiteForm) ErrorText() []string {
	return form.Errors
}

func (form *WebsiteForm) ToCanvasObject() fyne.CanvasObject {
	var formFields = []fyne.CanvasObject{
		customLayouts.NewFormBlock("Сайт", container.New(customLayouts.NewRatioLayout(0.7, 0.3),
			form.Title, form.Description)),
		customLayouts.NewFormBlock("Ссылка", container.NewVBox(form.URL)),
	}
	return container.New(customLayouts.NewFormLayout(), formFields...)
}

// NewWebsiteForm creates new WebsiteForm object with all the required data to display entries in app
func NewWebsiteForm() *WebsiteForm {
	var form = &WebsiteForm{}
	form.Title = newFormEntry("Название сайта", true)
	form.Title.Examples = []string{"Официальный сайт неофициальной организации"}
	form.Title.ValidationErrors = newValidationErrors("Название сайта", "строку")
	form.Description = newFormEntry("Описание сайта", false)
	form.Description.Examples = []string{"официальный сайт", "архивный сайт"}
	form.Description.ValidationErrors = newValidationErrors("Описание сайта", "строку")
	form.URL = newURLFormEntry(true)
	return form
}
