package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"strings"
	"time"
)

/* ------------ Book Form Structure  ------------ */

type BookForm struct {
	CitationForm
}

func (form *BookForm) ClearForm() {
	form.Title.Clear()
	form.Description.Clear()
	form.Edition.Clear()
	for _, a := range form.Authors {
		a.Clear()
	}
	for _, e := range form.Editors {
		e.Clear()
	}
	for _, t := range form.Translators {
		t.Clear()
	}
	for _, p := range form.Publishers {
		p.Clear()
	}
	form.YearPublished.Clear()
	form.PagesCount.Clear()
	form.ISBN.Clear()
	form.URL.Clear()
	form.Errors = []string{}
}

func (form *BookForm) ExampleForm() {
	form.Title.Example()
	form.Description.Example()
	form.Edition.Example()
	form.Authors[0].Example()
	form.Editors[0].Example()
	form.Translators[0].Example()
	form.Publishers[0].Example()
	form.YearPublished.Example()
	form.PagesCount.Example()
	form.ISBN.Example()
	form.URL.Example()
}

func (form *BookForm) ValidateForm() bool {
	form.Errors = []string{}
	if !(form.Title.ValidateRequired()) {
		form.Errors = append(form.Errors, form.Title.ValidationErrors.Empty)
	}
	if FieldsEmpty(form.Authors) && FieldsEmpty(form.Editors) {
		form.Errors = append(form.Errors, "Нужно заполнить информацию хотя бы об одном авторе или редакторе")
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
	if !(form.PagesCount.ValidateRequired()) {
		form.Errors = append(form.Errors, form.PagesCount.ValidationErrors.Empty)
	} else {
		if !(form.PagesCount.ValidateFormat()) {
			form.Errors = append(form.Errors, form.PagesCount.ValidationErrors.Format)
		}
	}

	return len(form.Errors) == 0
}

func (form *BookForm) ToCanvasObject() fyne.CanvasObject {
	var bookInfoFirstBlock = container.New(customLayouts.NewRatioLayout(0.7, 0.3),
		form.Title, form.ISBN)
	var bookInfoSecondBlock = container.New(customLayouts.NewRatioLayout(0.25, 0.25, 0.25, 0.25),
		form.Edition, form.Description,
		NumberEntryWithButtons(form.PagesCount, 1, 100000, 100),
		NumberEntryWithButtons(form.YearPublished, 1425, time.Now().Year(), time.Now().Year()))

	var formFields = []fyne.CanvasObject{
		customLayouts.NewFormBlock("Книга", container.NewVBox(bookInfoFirstBlock, bookInfoSecondBlock)),
		customLayouts.NewFormBlock("Авторы", PeopleContainer(form.Authors)),
		customLayouts.NewFormBlock("Редакторы", PeopleContainer(form.Editors)),
		customLayouts.NewFormBlock("Переводчики", PeopleContainer(form.Translators)),
		customLayouts.NewFormBlock("Издатели", OrganizationsContainer(form.Publishers)),
		customLayouts.NewFormBlock("Ссылка", container.NewVBox(form.URL)),
	}
	return container.New(customLayouts.NewFormLayout(), formFields...)
}

func (form *BookForm) Citation() string {
	var authorTitlePart, authorsPart, bookInfoPart string

	// author n. title : description
	var firstPartElements []string
	var authors = PeopleFromForm(form.Authors)
	if (len(authors) > 0) && (len(authors) < 4) {
		firstPartElements = append(firstPartElements, authors[0].SurnameInitials())
	}
	firstPartElements = append(firstPartElements, form.Title.TrimText())
	if form.Description.TrimText() != "" {
		firstPartElements = append(firstPartElements, fmt.Sprintf(": %s", form.Description.TrimText()))
	}
	authorTitlePart = strings.Join(firstPartElements, " ")

	// / authors; ред. editors; пер. translators. - edition. - Текст: электронный.
	var secondPartElements []string
	secondPartElements = append(secondPartElements, "/")
	if len(authors) > 0 {
		secondPartElements = append(secondPartElements, ListPeople(authors, 3))
	}
	var editors = PeopleFromForm(form.Editors)
	if len(editors) > 0 {
		var editorsString = ListPeople(editors, 1)
		if len(secondPartElements) == 2 {
			secondPartElements = append(secondPartElements, fmt.Sprintf("; ред. %s", editorsString))
		}
	}
	var translators = PeopleFromForm(form.Translators)
	if len(translators) > 0 {
		secondPartElements = append(secondPartElements, fmt.Sprintf("; пер. %s", ListPeople(translators, 1)))
	}

	if form.Edition.TrimText() != "" {
		secondPartElements = append(secondPartElements, fmt.Sprintf("— изд. %s.", form.Edition.TrimText()))
	}
	authorsPart = strings.Join(secondPartElements, " ")

	// — address: publisher, address: publisher, year. — pagesCount с. — URL (). — ISBN — Текст: электронный.
	var thirdPartElements []string
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

	thirdPartElements = append(thirdPartElements, fmt.Sprintf("— %d с.", form.PagesCount.ToNumber()))
	if form.URL.TrimText() != "" {
		var referenceDate = time.Time.Format(time.Now(), "02.01.2006")
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— URL: %s (дата обращения: %s).", form.URL.TrimText(), referenceDate))
	}
	if form.ISBN.TrimText() != "" {
		thirdPartElements = append(thirdPartElements, fmt.Sprintf("— ISBN %s.", form.ISBN.TrimText()))
	}
	if form.URL.TrimText() != "" {
		thirdPartElements = append(thirdPartElements, "— Текст: электронный.")
	}
	bookInfoPart = strings.Join(thirdPartElements, " ")

	var referenceParts = []string{authorTitlePart, authorsPart, bookInfoPart}
	return strings.Join(referenceParts, " ")
}

func (form *BookForm) HistoryRecordType() string {
	return BookCT.SystemName
}

func (form *BookForm) ErrorText() []string {
	return form.Errors
}

// NewBookForm creates new BookForm object with all the required data to display entries in app
func NewBookForm(authorsCount uint8, editorsCount uint8, translatorsCount uint8, publishersCount uint8) *BookForm {
	var form = &BookForm{}
	// title
	form.Title = newFormEntry("Название книги", true)
	form.Title.Examples = []string{"Очень интересная книга", "Самая увлекательная книга"}
	form.Title.ValidationErrors = newValidationErrors("Название книги", "строку")
	// description
	form.Description = newFormEntry("Описание", false)
	form.Description.Examples = []string{"неучебное пособие", "Полиграфия"}
	form.Description.ValidationErrors = newValidationErrors("Описание", "строку")
	// edition
	form.Edition = newFormEntry("Издание", false)
	form.Edition.Examples = []string{"1-е, оригинальное", "2-е, неоригинальное"}
	form.Edition.ValidationErrors = newValidationErrors("Издание", "строку")
	// people
	form.authorsFields(authorsCount, true)
	form.editorsFields(editorsCount, false)
	form.translatorsFields(translatorsCount, false)
	// publishers
	form.publishersFields(publishersCount, false)
	// year
	form.yearField()
	// pages count
	form.PagesCount = newNumberFormEntry("Кол-во страниц", true, false)
	form.PagesCount.ValidationErrors = newValidationErrors("Количество страниц", "число")
	// isbn
	form.ISBN = newFormEntry("ISBN", false)
	form.ISBN.Examples = []string{"000-0-00-000000-0"}
	// url
	form.URL = newURLFormEntry(false)

	return form
}
