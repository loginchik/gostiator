package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"strings"
	"time"
)

type Article struct {
	CitationForm
}

func (form *Article) ExampleForm() {
	form.Title.Example()
	form.YearPublished.Example()
	form.Authors[0].Example()
	form.DOI.Example()
	form.URL.Example()
	form.ParentTitle.Example()
	form.ParentVolume.Example()
	form.ParentNumber.Example()
	form.PageStart.Example()
	form.PageEnd.Example()
}

func (form *CitationForm) ClearForm() {
	form.Title.Clear()
	form.YearPublished.Clear()
	for _, author := range form.Authors {
		author.Clear()
	}
	form.DOI.Clear()
	form.URL.Clear()
	form.ParentTitle.Clear()
	form.ParentVolume.Clear()
	form.ParentNumber.Clear()
	form.PageStart.Clear()
	form.PageEnd.Clear()
	form.Errors = []string{}
}

func (form *Article) ValidateForm() bool {
	form.Errors = []string{}
	var requiredStrings = []*FormEntry{form.Title, form.ParentTitle, form.Authors[0].FirstName, form.Authors[0].LastName}
	for _, rString := range requiredStrings {
		if !(rString.ValidateRequired()) {
			form.Errors = append(form.Errors, rString.ValidationErrors.Empty)
		}
	}
	if !(form.YearPublished.ValidateRequired()) {
		form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Empty)
	} else {
		if !(form.YearPublished.ValidateFormat()) {
			form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Format)
		} else {
			if !(form.YearPublished.ValidateValue()) {
				form.Errors = append(form.Errors, form.YearPublished.ValidationErrors.Invalid)
			}
		}
	}
	if !(form.URL.IsEmpty()) {
		if !(form.URL.ValidateFormat()) {
			form.Errors = append(form.Errors, form.URL.ValidationErrors.Format)
		}
	}
	return len(form.Errors) == 0
}

func (form *Article) Citation() string {
	var citationParts []string

	// author title
	var authors = PeopleFromForm(form.Authors)
	if len(authors) < 4 {
		citationParts = append(citationParts, authors[0].SurnameInitials())
	}
	citationParts = append(citationParts, form.Title.TrimText())

	// / authors — doi part
	citationParts = append(citationParts, "/", ListPeople(authors, 3))
	if !(form.DOI.TrimText() == "") {
		citationParts = append(citationParts, fmt.Sprintf("— DOI: %s.", form.DOI.TrimText()))
	}
	if !(form.URL.TrimText() == "") {
		citationParts = append(citationParts, "— Текст: электронный.")
	}

	// journal info part
	citationParts = append(citationParts,
		"//", form.ParentTitle.TrimText(),
		fmt.Sprintf("— %d.", form.YearPublished.ToNumber()))

	var journalString string
	if (form.ParentVolume.TrimText() != "") && (form.ParentNumber.TrimText() != "") {
		journalString = fmt.Sprintf("%s (%s)", form.ParentVolume.TrimText(), form.ParentNumber.TrimText())
	} else {
		if form.ParentVolume.TrimText() != "" {
			journalString = form.ParentVolume.TrimText()
		} else if form.ParentNumber.TrimText() != "" {
			journalString = form.ParentNumber.TrimText()
		}
	}
	if !(journalString == "") {
		citationParts = append(citationParts, fmt.Sprintf("— %s.", journalString))
	}

	var pagesRange = StringPageRange(form.PageStart.TrimText(), form.PageEnd.TrimText())
	if !(pagesRange == "") {
		citationParts = append(citationParts, fmt.Sprintf("— С. %s.", pagesRange))
	}
	if form.URL.TrimText() != "" {
		var citationDate = time.Time.Format(time.Now(), "02.01.2006")
		var urlString = fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), citationDate)
		citationParts = append(citationParts, urlString)
	}

	// collect parts
	return strings.Join(citationParts, " ")
}

func (form *Article) HistoryRecordType() string {
	return ArticleCT.SystemName
}
func (form *Article) ErrorText() []string {
	return form.Errors
}

func (form *Article) ToCanvasObject() fyne.CanvasObject {
	var journalSecondBlock = container.New(customLayouts.NewRatioLayout(0.25, 0.25, 0.25, 0.25),
		form.ParentVolume, form.ParentNumber, form.PageStart, form.PageEnd)

	var formFields = []fyne.CanvasObject{
		customLayouts.NewFormBlock("Статья", container.NewVBox(form.Title,
			container.New(customLayouts.NewRatioLayout(0.8, 0.2), form.DOI,
				NumberEntryWithButtons(form.YearPublished, 1425, time.Now().Year(), time.Now().Year())))),
		customLayouts.NewFormBlock("Авторы", PeopleContainer(form.Authors)),
		customLayouts.NewFormBlock("Журнал", container.NewVBox(form.ParentTitle, journalSecondBlock)),
		customLayouts.NewFormBlock("Ссылка", container.NewVBox(form.URL)),
	}
	return container.New(customLayouts.NewFormLayout(), formFields...)
}

// NewArticleForm creates new Article object with all the required data to display entries in app
func NewArticleForm(authorsCount uint8) *Article {
	var form = &Article{}
	// title
	form.Title = newFormEntry("Название статьи", true)
	form.Title.Examples = []string{"Очень интересная статья", "Самая увлекательная статья"}
	form.Title.ValidationErrors = newValidationErrors("Название статьи", "строку")
	// authors
	form.authorsFields(authorsCount, true)
	// journal title
	form.ParentTitle = newFormEntry("Название журнала", true)
	form.ParentTitle.Examples = []string{"Веселый журнал", "Журнал о котиках", "Псевдонаучный журнал"}
	form.ParentTitle.ValidationErrors = newValidationErrors("Название журнала", "строку")
	// journal info
	form.ParentVolume = newFormEntry("Том журнала", false)
	form.ParentVolume.Examples = []string{"10", "20", "30"}
	form.ParentNumber = newFormEntry("Номер журнала", false)
	form.ParentNumber.Examples = []string{"1", "2", "3"}
	// year
	form.yearField()
	// pages
	form.pageRangeFields(false)
	// links
	form.doiField()
	form.URL = newURLFormEntry(false)

	return form
}
