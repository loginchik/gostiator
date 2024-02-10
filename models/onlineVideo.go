package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"net/url"
	"strings"
	"time"
)

type OnlineVideo struct {
	CitationForm
}

func (form *OnlineVideo) ExampleForm() {
	form.Title.Example()
	form.Description.Example()
	form.Duration.Example()
	form.Authors[0].Example()
	form.URL.Example()
	form.YearPublished.Example()
}

func (form *OnlineVideo) ClearForm() {
	form.Title.Clear()
	form.Description.Clear()
	form.Authors[0].Clear()
	form.URL.Clear()
	form.Duration.Clear()
	form.YearPublished.Example()
	form.Errors = []string{}
}

func (form *OnlineVideo) Citation() string {
	var titleElements []string
	var title = fmt.Sprintf("%s (%d)", form.Title.TrimText(), form.YearPublished.ToNumber())
	titleElements = append(titleElements, title)
	if !(form.Description.TrimText() == "") {
		titleElements = append(titleElements, form.Description.TrimText())
	}
	var titlePart = strings.Join(titleElements, ": ")

	var secondPartElements = []string{"/"}
	var author = PeopleFromForm(form.Authors)[0]
	secondPartElements = append(secondPartElements, author.InitialsSurname()+".")
	if !(form.Duration.String() == "") {
		var durationString = fmt.Sprintf("— %s (время воспроизведения).", form.Duration.String())
		secondPartElements = append(secondPartElements, durationString)
	}
	var referenceDate = time.Now().Format("02.01.2006")
	secondPartElements = append(secondPartElements, fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), referenceDate))
	var urlHost, hostErr = url.Parse(form.URL.TrimText())
	if hostErr == nil {
		secondPartElements = append(secondPartElements, fmt.Sprintf("Доступно на %s.", urlHost))
	}
	var secondPart = strings.Join(secondPartElements, " ")

	var citationParts = []string{titlePart, secondPart}
	return strings.Join(citationParts, " ")
}

func (form *OnlineVideo) ValidateForm() bool {
	form.Errors = []string{}

	if !(form.Title.ValidateRequired()) {
		form.Errors = append(form.Errors, form.Title.ValidationErrors.Empty)
	}
	if FieldsEmpty(form.Authors) {
		form.Errors = append(form.Errors, "Нужно заполнить инфморацию о хотя бы одном авторе")
	}

	for _, f := range []*NumberFormEntry{form.Duration.Hours, form.Duration.Minutes, form.Duration.Seconds} {
		if !(f.IsEmpty()) {
			if !(f.ValidateFormat()) {
				form.Errors = append(form.Errors, f.ValidationErrors.Format)
			}
		}
	}

	if !(form.YearPublished.ValidateRequired()) {
		form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Empty)
	} else if !(form.YearPublished.ValidateFormat()) {
		form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Format)
	} else if !(form.YearPublished.ValidateValue()) {
		form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Invalid)
	}

	if !(form.URL.ValidateRequired()) {
		form.Errors = append(form.Errors, form.URL.ValidationErrors.Empty)
	} else if !(form.URL.ValidateFormat()) {
		form.Errors = append(form.Errors, form.URL.ValidationErrors.Format)
	}

	return len(form.Errors) == 0
}

func (form *OnlineVideo) HistoryRecordType() string {
	return WebvideoCT.SystemName
}

func (form *OnlineVideo) ErrorText() []string {
	return form.Errors
}

func (form *OnlineVideo) ToCanvasObject() fyne.CanvasObject {
	var basicInfoElements = []fyne.CanvasObject{form.Description, form.YearPublished, form.Duration.Container()}
	var basicInfoBlock = container.NewAdaptiveGrid(len(basicInfoElements), basicInfoElements...)
	var authorBlock = PeopleContainer(form.Authors, "Автор", 1)

	var formFields = []fyne.CanvasObject{form.Title, basicInfoBlock, authorBlock, form.URL}
	return container.NewVBox(formFields...)
}

func NewOnlineVideoForm() *OnlineVideo {
	var form = &OnlineVideo{}
	form.Title = newFormEntry("Название видео", true)
	form.Title.Examples = []string{"Видео о котиках", "Ëжики завоёвыват мир"}
	form.Title.ValidationErrors = newValidationErrors("Название видео", "строку")

	form.Description = newFormEntry("Описание видео", false)
	form.Description.Examples = []string{""}
	form.Description.ValidationErrors = newValidationErrors("Описание видео", "строку")

	form.authorsFields(1, false)
	form.Duration = NewDurationForm()
	form.yearField()
	form.URL = newURLFormEntry(true)

	form.Errors = []string{}
	return form
}
