package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"strings"
	"time"
)

/* ------------ Book Part Form Structure  ------------ */

type BookPartForm struct {
	CitationForm
}

func (form *BookPartForm) ExampleForm() {
	form.Title.Example()
	form.Authors[0].Example()
	form.Publishers[0].Example()
	form.ParentTitle.Example()
	form.YearPublished.Example()
	form.PageStart.Example()
	form.PageEnd.Example()
	form.DOI.Example()
	form.URL.Example()
}

func (form *BookPartForm) ClearForm() {
	form.Title.Clear()
	for _, a := range form.Authors {
		a.Clear()
	}
	for _, p := range form.Publishers {
		p.Clear()
	}
	form.ParentTitle.Clear()
	form.YearPublished.Clear()
	form.PageStart.Clear()
	form.PageEnd.Clear()
	form.DOI.Clear()
	form.URL.Clear()
	form.Errors = []string{}
}

func (form *BookPartForm) ValidateForm() bool {
	form.Errors = []string{}
	var requiredStrings = []*FormEntry{form.Title, form.ParentTitle}
	for _, rString := range requiredStrings {
		if !(rString.ValidateRequired()) {
			form.Errors = append(form.Errors, rString.ValidationErrors.Empty)
		}
	}
	if FieldsEmpty(form.Authors) {
		form.Errors = append(form.Errors, "Нужно заполнить инфморацию о хотя бы одном авторе")
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

	return len(form.Errors) == 0
}

func (form *BookPartForm) Citation() string {
	var authorTitlePart, authorsPart, bookInfoPart string

	// author. title
	var firstPartElements []string
	var authors = PeopleFromForm(form.Authors)
	if len(authors) < 4 {
		firstPartElements = append(firstPartElements, authors[0].SurnameInitials())
	}
	firstPartElements = append(firstPartElements, form.Title.TrimText())
	authorTitlePart = strings.Join(firstPartElements, " ")

	// / authors. — DOI. — Текст: электронный
	var secondPartElements = []string{"/", ListPeople(authors, 3)}
	if !(form.DOI.TrimText() == "") {
		secondPartElements = append(secondPartElements, fmt.Sprintf("— DOI: %s.", form.DOI.TrimText()))
	}
	if !(form.URL.TrimText() == "") {
		secondPartElements = append(secondPartElements, "— Текст: электронный.")
	}
	authorsPart = strings.Join(secondPartElements, " ")

	// // book. — address: name, year. — pages.
	var thirdPartElements = []string{"//", form.ParentTitle.TrimText()}
	var publishers = OrganizationsFromForm(form.Publishers)
	if len(publishers) == 0 {
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— [не ук.]: [не ук.], %d.", form.YearPublished.ToNumber()))
	} else {
		var pubStrings []string
		if len(publishers) > 2 {
			pubStrings = append(pubStrings, publishers[0].String()+" [и др.]")
		} else {
			for _, org := range publishers {
				pubStrings = append(pubStrings, org.String())
			}
		}
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— %s, %d.", strings.Join(pubStrings, ", "), form.YearPublished.ToNumber()))
	}

	var pagesString = StringPageRange(form.PageStart.TrimText(), form.PageEnd.TrimText())
	if !(pagesString == "") {
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— С. %s.", pagesString))
	}

	// — URL ()
	if !(form.URL.TrimText() == "") {
		var referenceDate = time.Now().Format("02.01.2006")
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), referenceDate))
	}
	bookInfoPart = strings.Join(thirdPartElements, " ")

	var citationParts = []string{authorTitlePart, authorsPart, bookInfoPart}
	return strings.Join(citationParts, " ")
}

func (form *BookPartForm) HistoryRecordType() string {
	return BookPartCT.SystemName
}

func (form *BookPartForm) ErrorText() []string {
	return form.Errors
}

func (form *BookPartForm) ToCanvasObject() fyne.CanvasObject {
	var partInfo = container.New(customLayouts.NewRatioLayout(0.7, 0.3),
		form.DOI, NumberEntryWithButtons(form.YearPublished, 1425, time.Now().Year(), time.Now().Year()))
	var bookInfoBlock = container.New(customLayouts.NewRatioLayout(0.6, 0.2, 0.2),
		form.ParentTitle, form.PageStart, form.PageEnd)
	var formFields = []fyne.CanvasObject{
		customLayouts.NewFormBlock("Часть книги", container.NewVBox(form.Title, partInfo)),
		customLayouts.NewFormBlock("Авторы", PeopleContainer(form.Authors)),
		customLayouts.NewFormBlock("Книга", bookInfoBlock),
		customLayouts.NewFormBlock("Издатели", OrganizationsContainer(form.Publishers)),
		customLayouts.NewFormBlock("Ссылка", container.NewVBox(form.URL)),
	}
	return container.New(customLayouts.NewFormLayout(), formFields...)
}

// NewBookPartForm creates new BookPartForm object with all the required data to display entries in app
func NewBookPartForm(authorsCount uint8, publishersCount uint8) *BookPartForm {
	var form = &BookPartForm{}
	// title
	form.Title = newFormEntry("Название части книги", true)
	form.Title.Examples = []string{"Очень интересная глава", "Самая увлекательная часть"}
	form.Title.ValidationErrors = newValidationErrors("Название части книги", "строку")
	// authors
	form.authorsFields(authorsCount, true)
	// publishers
	form.publishersFields(publishersCount, false)
	// book title
	form.ParentTitle = newFormEntry("Название книги", true)
	form.ParentTitle.Examples = []string{"Интересная книга", "Увлекательная книга"}
	form.ParentTitle.ValidationErrors = newValidationErrors("Название книги", "строку")
	// year
	form.yearField()
	// pages
	form.pageRangeFields(false)
	// doi
	form.doiField()
	// url
	form.URL = newURLFormEntry(false)

	return form
}
