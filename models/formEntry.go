package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/* ------------ Form Entry Structures  ------------ */

// validationErrors objects contain strings that can be displayed in app in case
// something is wrong with user input data. Empty - if field is required but not filled,
// Format - if field can't be converted to target data type, Invalid - if field data
// doesn't match expected values (non-existent date, for example).
type validationErrors struct {
	Empty   string
	Format  string
	Invalid string
}

// newValidationErrors creates validationErrors object with some predefined values.
//
// Empty pattern: "Нужно заполнить поле fieldTitle"
//
// Format pattern: "В поле fieldTitle введите formatTitle",
//
// Invalid pattern: "fieldTitle кажется подозрительным".
func newValidationErrors(fieldTitle string, formatTitle string) validationErrors {
	return validationErrors{
		Empty:   fmt.Sprintf("Нужно заполнить поле %s", fieldTitle),
		Format:  fmt.Sprintf("В поле %s введите %s", fieldTitle, formatTitle),
		Invalid: fmt.Sprintf("%s кажется подозрительным", fieldTitle),
	}
}

type FormEntry struct {
	widget.Entry
	PlaceHolderText  string
	Required         bool
	Examples         []string
	ValidationErrors validationErrors
}

func (formEntry *FormEntry) CreateRenderer() fyne.WidgetRenderer {
	return formEntry.Entry.CreateRenderer()
}

func (formEntry *FormEntry) Clear() {
	formEntry.SetText("")
}

func (formEntry *FormEntry) Example() {
	var examplesCount = len(formEntry.Examples)
	var exampleText = formEntry.Examples[rand.Intn(examplesCount)]
	formEntry.SetText(exampleText)
}

func (formEntry *FormEntry) IsEmpty() bool {
	return formEntry.Text == ""
}

func (formEntry *FormEntry) TrimText() string {
	return strings.Trim(formEntry.Text, " ")
}

func (formEntry *FormEntry) ValidateRequired() bool {
	if formEntry.Required {
		return !(formEntry.Text == "")
	} else {
		return true
	}
}

type NumberFormEntry struct {
	FormEntry
	isYear  bool
	isMonth bool
	isDay   bool
}

func (numberEntry *NumberFormEntry) ValidateFormat() bool {
	_, err := strconv.ParseInt(numberEntry.Text, 10, 0)
	return err == nil
}

func (numberEntry *NumberFormEntry) ValidateValue() bool {
	intValue, _ := strconv.ParseInt(numberEntry.Text, 10, 0)
	if numberEntry.isYear {
		return intValue >= 1425
	} else if numberEntry.isMonth {
		return (intValue >= 1) && (intValue <= 12)
	} else if numberEntry.isDay {
		return (intValue >= 1) && (intValue <= 31)
	} else {
		return true
	}
}

func (numberEntry *NumberFormEntry) ToNumber() uint16 {
	valueInt, parseError := strconv.ParseUint(numberEntry.Text, 10, 16)
	if parseError != nil {
		return 0
	} else {
		return uint16(valueInt)
	}
}

type URLFormEntry struct {
	FormEntry
}

func (urlEntry *URLFormEntry) ValidateFormat() bool {
	parsedUrl, err := url.Parse(urlEntry.Text)
	parsed := err == nil
	hostFull := parsedUrl.Host != ""
	schemeFull := parsedUrl.Scheme != ""
	return parsed && hostFull && schemeFull
}

func newFormEntry(placeholder string, required bool) *FormEntry {
	var entry = &FormEntry{}
	if required {
		entry.PlaceHolderText = fmt.Sprintf("%s *", placeholder)
	} else {
		entry.PlaceHolderText = placeholder
	}
	entry.Required = required
	entry.SetPlaceHolder(entry.PlaceHolderText)
	entry.ExtendBaseWidget(entry)
	return entry
}

func newNumberFormEntry(placeholder string, required bool, isYear bool) *NumberFormEntry {
	var entry = &NumberFormEntry{}
	if required {
		entry.PlaceHolderText = fmt.Sprintf("%s *", placeholder)
	} else {
		entry.PlaceHolderText = placeholder
	}
	entry.Required = required
	entry.isYear = isYear
	if entry.isYear {
		entry.Examples = []string{time.Now().Format("2006")}
	} else {
		entry.Examples = []string{"1", "2", "3", "4", "5", "10"}
	}
	entry.SetPlaceHolder(entry.PlaceHolderText)
	entry.ExtendBaseWidget(entry)
	return entry
}

func updateValue(lowerLimit int, upperLimit int, newValue int) int {
	if newValue <= lowerLimit {
		return lowerLimit
	}
	if newValue >= upperLimit {
		return upperLimit
	}
	return newValue
}

func NumberEntryWithButtons(numberField *NumberFormEntry, lowerLimit int, upperLimit int, startingValue int) *fyne.Container {
	var increaseButton = widget.NewButton("+", func() {
		if numberField.Text == "" {
			numberField.SetText(strconv.Itoa(startingValue))
			return
		}

		currentValue, err := strconv.ParseInt(numberField.Text, 10, 64)
		if err != nil {
			numberField.SetText(strconv.Itoa(startingValue))
			return
		} else {
			var newValue = int(currentValue) + 1
			var trueNewValue = updateValue(lowerLimit, upperLimit, newValue)
			if trueNewValue != int(currentValue) {
				numberField.SetText(strconv.Itoa(trueNewValue))
				return
			}
		}

	})
	var decreaseButton = widget.NewButton("-", func() {
		if numberField.Text == "" {
			numberField.SetText(strconv.Itoa(startingValue))
			return
		}

		currentValue, err := strconv.ParseInt(numberField.Text, 10, 64)
		if err != nil {
			numberField.SetText(strconv.Itoa(startingValue))
			return
		} else {
			var newValue = int(currentValue) - 1
			var trueNewValue = updateValue(lowerLimit, upperLimit, newValue)
			if trueNewValue != int(currentValue) {
				numberField.SetText(strconv.Itoa(trueNewValue))
				return
			}
		}
	})
	return container.NewBorder(nil, nil, decreaseButton, increaseButton, numberField)
}

func newURLFormEntry(required bool) *URLFormEntry {
	var entry = &URLFormEntry{}
	if required {
		entry.PlaceHolderText = "URL *"
	} else {
		entry.PlaceHolderText = "URL"
	}
	entry.Examples = []string{"https://google.com"}
	entry.ValidationErrors = newValidationErrors("URL", "ссылку, начинающуюся с https://")
	entry.Required = required
	entry.SetPlaceHolder(entry.PlaceHolderText)
	entry.ExtendBaseWidget(entry)
	return entry
}
