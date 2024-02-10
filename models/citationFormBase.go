package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"slices"
	"strings"
)

/* ------------ Basic form Structure ------------ */

type CitationInterface interface {
	ClearForm()
	ExampleForm()
	ValidateForm() bool
	Citation() string
	HistoryRecordType() string
	ErrorText() []string
	ToCanvasObject() fyne.CanvasObject
}

type DurationForm struct {
	Seconds *NumberFormEntry
	Minutes *NumberFormEntry
	Hours   *NumberFormEntry
}

func NewDurationForm() *DurationForm {
	var form = DurationForm{}
	form.Seconds = newNumberFormEntry("Секунды", false, false)
	form.Minutes = newNumberFormEntry("Минуты", false, false)
	form.Hours = newNumberFormEntry("Часы", false, false)
	form.Seconds.ValidationErrors = newValidationErrors("Секунды", "число")
	form.Minutes.ValidationErrors = newValidationErrors("Минуты", "число")
	form.Hours.ValidationErrors = newValidationErrors("Часы", "число")
	return &form
}

func (form *DurationForm) Container() *fyne.Container {
	var elements = []fyne.CanvasObject{form.Hours, form.Minutes, form.Seconds}
	return container.NewAdaptiveGrid(len(elements), elements...)
}

func (form *DurationForm) Example() {
	form.Hours.Example()
	form.Minutes.Example()
	form.Seconds.Example()
}

func (form *DurationForm) Clear() {
	form.Hours.Clear()
	form.Minutes.Clear()
	form.Seconds.Clear()
}

func (form *DurationForm) String() string {
	var durationElements []string
	var hours = form.Hours.ToNumber()
	var minutes = form.Minutes.ToNumber()
	var seconds = form.Seconds.ToNumber()

	if hours > 0 {
		durationElements = append(durationElements, fmt.Sprintf("%dч", hours))
	}
	if minutes > 0 {
		durationElements = append(durationElements, fmt.Sprintf("%dмин", minutes))
	}
	if seconds > 0 {
		durationElements = append(durationElements, fmt.Sprintf("%dсек", seconds))
	}
	return strings.Join(durationElements, " ")
}

type CitationForm struct {
	Title          *FormEntry
	Description    *FormEntry
	Edition        *FormEntry
	ContainerBox   *FormEntry
	DayPublished   *NumberFormEntry
	MonthPublished *NumberFormEntry
	YearPublished  *NumberFormEntry

	Authors     []PersonForm
	Editors     []PersonForm
	Translators []PersonForm
	Publishers  []OrganizationForm

	ISBN *FormEntry
	DOI  *FormEntry
	URL  *URLFormEntry

	ParentTitle  *FormEntry
	ParentVolume *FormEntry
	ParentNumber *FormEntry

	PagesCount *NumberFormEntry
	PageStart  *FormEntry
	PageEnd    *FormEntry
	Duration   *DurationForm

	Errors []string
}

func (form *CitationForm) pageRangeFields(required bool) {
	form.PageStart = newFormEntry("Первая страница", required)
	form.PageStart.Examples = []string{"1"}
	form.PageEnd = newFormEntry("Последняя страница", required)
	form.PageEnd.Examples = []string{"100"}
}

func (form *CitationForm) authorsFields(count uint8, requireFirst bool) {
	form.Authors = []PersonForm{}
	for i := 0; uint8(i) < count; i++ {
		var author PersonForm
		author.New((i == 0) && requireFirst, i+1, "Автор")
		form.Authors = append(form.Authors, author)
	}
}

func (form *CitationForm) editorsFields(count uint8, requireFirst bool) {
	form.Editors = []PersonForm{}
	for i := 0; uint8(i) < count; i++ {
		var editor PersonForm
		editor.New((i == 0) && requireFirst, i+1, "Редактор")
		form.Editors = append(form.Editors, editor)
	}
}

func (form *CitationForm) translatorsFields(count uint8, requireFirst bool) {
	form.Translators = []PersonForm{}
	for i := 0; uint8(i) < count; i++ {
		var translator PersonForm
		translator.New((i == 0) && requireFirst, i+1, "Переводчик")
		form.Translators = append(form.Translators, translator)
	}
}

func (form *CitationForm) publishersFields(count uint8, requireFirst bool) {
	form.Publishers = []OrganizationForm{}
	for i := 0; uint8(i) < count; i++ {
		var org OrganizationForm
		org.New((i == 0) && requireFirst, i+1)
		form.Publishers = append(form.Publishers, org)
	}
}

func (form *CitationForm) yearField() {
	form.YearPublished = newNumberFormEntry("Год", true, true)
	form.YearPublished.ValidationErrors = newValidationErrors("Год", "число")
}

func (form *CitationForm) doiField() {
	form.DOI = newFormEntry("DOI", false)
	form.DOI.Examples = []string{"10.0000/n.surname.0000.000000"}
	form.DOI.ValidationErrors = newValidationErrors("DOI", "строку")
}

func (form *CitationForm) checkDay() {
	var yearNum = form.YearPublished.ToNumber()
	var monthNum = form.MonthPublished.ToNumber()
	var dayNum = form.DayPublished.ToNumber()

	if monthNum == 2 {
		if (yearNum%4 == 0) && (dayNum > 29) {
			form.Errors = append(form.Errors, form.DayPublished.ValidationErrors.Invalid)
		} else if (yearNum%4 != 0) && (dayNum > 28) {
			form.Errors = append(form.Errors, form.DayPublished.ValidationErrors.Invalid)
		}
	} else if slices.Contains([]uint16{4, 6, 9, 11}, monthNum) {
		if dayNum > 30 {
			form.Errors = append(form.Errors, form.DayPublished.ValidationErrors.Invalid)
		}
	}
}

func (form *CitationForm) ValidateDateFields() {
	var checkMonthDay = true
	for _, rNumber := range []*NumberFormEntry{form.YearPublished, form.MonthPublished, form.DayPublished} {
		if !(rNumber.ValidateRequired()) {
			form.Errors = append(form.Errors, rNumber.ValidationErrors.Empty)
			checkMonthDay = false
		} else if !(rNumber.ValidateFormat()) {
			form.Errors = append(form.Errors, rNumber.ValidationErrors.Format)
			checkMonthDay = false
		} else if !(rNumber.ValidateValue()) {
			form.Errors = append(form.Errors, rNumber.ValidationErrors.Invalid)
			checkMonthDay = false
		}
	}
	if checkMonthDay {
		form.checkDay()
	}
}
