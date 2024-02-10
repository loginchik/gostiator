package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"strings"
	"time"
)

/* ------------ WebText Form Structure  ------------ */

type WebTextForm struct {
	CitationForm
}

func (form *WebTextForm) ExampleForm() {
	form.Title.Example()
	form.Authors[0].Example()
	form.DayPublished.Example()
	form.MonthPublished.Example()
	form.YearPublished.Example()
	form.ParentTitle.Example()
	form.Description.Example()
	form.URL.Example()
	form.DOI.Example()
}

func (form *WebTextForm) Clear() {
	form.Title.Clear()
	for _, a := range form.Authors {
		a.Clear()
	}
	form.DayPublished.Clear()
	form.MonthPublished.Clear()
	form.YearPublished.Clear()
	form.ParentTitle.Clear()
	form.Description.Clear()
	form.URL.Clear()
	form.DOI.Clear()
	form.Errors = []string{}
}

func (form *WebTextForm) ValidateForm() bool {
	form.Errors = []string{}

	var requiredStrings = []*FormEntry{form.Title, form.ParentTitle}
	for _, rString := range requiredStrings {
		if !(rString.ValidateRequired()) {
			form.Errors = append(form.Errors, rString.ValidationErrors.Empty)
		}
	}
	form.ValidateDateFields()

	if !(form.URL.ValidateRequired()) {
		form.Errors = append(form.Errors, form.URL.ValidationErrors.Empty)
	} else if !(form.URL.ValidateFormat()) {
		form.Errors = append(form.Errors, form.URL.ValidationErrors.Format)
	}

	return len(form.Errors) == 0
}

func (form *WebTextForm) Citation() string {
	var authorTitleElements []string
	var authors = PeopleFromForm(form.Authors)
	if (len(authors) > 0) && (len(authors) < 4) {
		authorTitleElements = append(authorTitleElements, authors[0].SurnameInitials())
	}
	authorTitleElements = append(authorTitleElements, form.Title.TrimText())
	var authorTitlePart = strings.Join(authorTitleElements, " ")

	var authorsDOIElements = []string{"/"}
	if len(authors) > 0 {
		authorsDOIElements = append(authorsDOIElements, ListPeople(authors, 3))
	}
	if !(form.DOI.TrimText() == "") {
		var format = "DOI: %s. — Текст: электронный."
		if len(authorsDOIElements) > 1 {
			format = "— " + format
		}
		authorsDOIElements = append(authorsDOIElements, fmt.Sprintf(format, form.DOI.TrimText()))
	} else {
		var t = "Текст: электронный."
		if len(authorsDOIElements) > 1 {
			t = "— " + t
		}
		authorsDOIElements = append(authorsDOIElements, t)
	}
	var authorsDOIPart = strings.Join(authorsDOIElements, " ")

	var websiteElements = []string{"//"}
	websiteElements = append(websiteElements, form.ParentTitle.TrimText())
	if !(form.Description.TrimText() == "") {
		websiteElements = append(websiteElements, fmt.Sprintf(": %s.", form.Description.TrimText()))
	}
	var date = time.Date(
		int(form.YearPublished.ToNumber()),
		time.Month(int(form.MonthPublished.ToNumber())),
		int(form.DayPublished.ToNumber()),
		0, 0, 0, 0, time.UTC,
	)
	websiteElements = append(websiteElements, date.Format("— 02.01.2006."))
	var referenceDate = time.Now().Format("02.01.2006")
	websiteElements = append(websiteElements, fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), referenceDate))
	var websitePart = strings.Join(websiteElements, " ")

	var citationParts = []string{authorTitlePart, authorsDOIPart, websitePart}
	return strings.Join(citationParts, " ")
}

func (form *WebTextForm) HistoryRecordType() string {
	return WebtextCT.SystemName
}

func (form *WebTextForm) ErrorText() []string {
	return form.Errors
}

func (form *WebTextForm) ToCanvasObject() fyne.CanvasObject {
	var authorsContainer = PeopleContainer(form.Authors, "Авторы", 3)
	var webSiteInfoElements = []fyne.CanvasObject{
		form.Description, form.DayPublished, form.MonthPublished, form.YearPublished,
	}
	var webSiteInfoContainer = container.NewAdaptiveGrid(len(webSiteInfoElements), webSiteInfoElements...)
	var linksContainer = container.NewAdaptiveGrid(2, form.URL, form.DOI)
	var formFields = []fyne.CanvasObject{form.Title, authorsContainer, form.ParentTitle, webSiteInfoContainer, linksContainer}
	return container.NewVBox(formFields...)
}

// NewWebtextForm creates new WebTextForm object with all the required data to display entries in app
func NewWebtextForm(authorsCount uint8) *WebTextForm {
	var form = &WebTextForm{}
	// title
	form.Title = newFormEntry("Заголовок", true)
	form.Title.Examples = []string{"Статья о котиках с сайта о собачках", "Запись с женского форума"}
	form.Title.ValidationErrors = newValidationErrors(form.Title.PlaceHolderText, "строку")
	// authors
	form.authorsFields(authorsCount, false)
	// date
	form.DayPublished = newNumberFormEntry("День публикации", true, false)
	form.DayPublished.isDay = true
	form.DayPublished.ValidationErrors = newValidationErrors(form.DayPublished.PlaceHolderText, "число в диапазоне от 1 до 31")
	form.MonthPublished = newNumberFormEntry("Месяц публикации", true, false)
	form.MonthPublished.isMonth = true
	form.MonthPublished.ValidationErrors = newValidationErrors(form.MonthPublished.PlaceHolderText, "номер месяца в диапазоне от 1 до 12")
	form.yearField()
	// site
	form.ParentTitle = newFormEntry("Название сайта", true)
	form.ParentTitle.Examples = []string{"Ëжики рулят", "Модный приговор"}
	form.ParentTitle.ValidationErrors = newValidationErrors(form.ParentTitle.PlaceHolderText, "строку")
	form.Description = newFormEntry("Описание сайта", false)
	form.Description.Examples = []string{"архивная версия сайта", "официальный сайт неофициальной организации"}
	form.URL = newURLFormEntry(true)
	// doi
	form.doiField()
	return form
}
