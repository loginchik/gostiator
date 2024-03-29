package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"strings"
	"time"
)

type Film struct {
	CitationForm
}

func (form *Film) ClearForm() {
	form.Title.Clear()
	form.Description.Clear()
	form.ContainerBox.Clear()
	form.Authors[0].Clear()
	form.Publishers[0].Name.Clear()
	form.Publishers[1].Clear()
	form.Duration.Clear()
	form.YearPublished.Clear()
}

func (form *Film) ExampleForm() {
	form.Title.Example()
	form.Description.Example()
	form.ContainerBox.Example()
	form.Authors[0].Example()
	form.Publishers[0].Name.Example()
	form.Publishers[1].Example()
	form.Duration.Example()
	form.YearPublished.Example()
}

func (form *Film) ValidateForm() bool {
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

	return len(form.Errors) == 0
}

func (form *Film) Citation() string {
	var titleElements = []string{form.Title.TrimText()}
	if form.Description.TrimText() != "" {
		titleElements = append(titleElements, fmt.Sprintf(": %s", form.Description.TrimText()))
	}
	var titlePart = strings.Join(titleElements, " ")

	var creatorsElements []string
	var director = PeopleFromForm(form.Authors)[0]
	creatorsElements = append(creatorsElements, fmt.Sprintf("/ реж. %s", director.InitialsSurname()))
	var creationStudio = form.Publishers[0].Name.TrimText()
	if creationStudio != "" {
		creatorsElements = append(creatorsElements, creationStudio)
	}
	var creatorsPart = strings.Join(creatorsElements, "; ") + "."

	var metaElements []string
	var firstMetaString string
	var duration = form.Duration.String()

	if form.Publishers[1].Full() {
		var distributor = OrganizationsFromForm(form.Publishers[1:])[0]
		firstMetaString = fmt.Sprintf("— %s, %d.", distributor.String(), form.YearPublished.ToNumber())
	} else {
		firstMetaString = fmt.Sprintf("— [не ук.], %d.", form.YearPublished.ToNumber())
	}
	metaElements = append(metaElements, firstMetaString)

	var secondMetaString string
	if form.Publishers[1].Full() {
		if len(duration) > 0 {
			secondMetaString = fmt.Sprintf("— %s (%s).", form.ContainerBox.TrimText(), duration)
		} else {
			secondMetaString = fmt.Sprintf("— %s.", form.ContainerBox.TrimText())
		}
	} else {
		if len(duration) > 0 {
			secondMetaString = fmt.Sprintf("— [не ук.] (%s)", duration)
		} else {
			secondMetaString = "— [не ук.]."
		}
	}
	metaElements = append(metaElements, secondMetaString)

	metaElements = append(metaElements, "— Изображение (двухмерное, движущееся): видео.")
	var metaPart = strings.Join(metaElements, " ")

	var citationParts = []string{titlePart, creatorsPart, metaPart}
	return strings.Join(citationParts, " ")
}

func (form *Film) HistoryRecordType() string {
	return FilmCT.SystemName
}

func (form *Film) ErrorText() []string {
	return form.Errors
}

func (form *Film) ToCanvasObject() fyne.CanvasObject {
	var filmInfoFirstBlock = container.New(customLayouts.NewRatioLayout(0.7, 0.3),
		form.Title, form.Description)
	var filmInfoSecondBlock = container.New(customLayouts.NewRatioLayout(0.3, 0.5, 0.2),
		form.ContainerBox, form.Duration.Container(),
		NumberEntryWithButtons(form.YearPublished, 1425, time.Now().Year(), time.Now().Year()))
	var creatorsBlock = container.New(customLayouts.NewRatioLayout(0.5, 0.5),
		PeopleContainer(form.Authors), form.Publishers[0].Name)
	var distributorsBlock = OrganizationsContainer([]OrganizationForm{form.Publishers[1]})

	var formFields = []fyne.CanvasObject{
		customLayouts.NewFormBlock("Фильм", container.NewVBox(filmInfoFirstBlock, filmInfoSecondBlock)),
		customLayouts.NewFormBlock("Режиссёр и киностудия", creatorsBlock),
		customLayouts.NewFormBlock("Дистрибутор", distributorsBlock),
	}
	return container.New(customLayouts.NewFormLayout(), formFields...)
}

func NewFilmForm() *Film {
	var form = &Film{}
	form.Title = newFormEntry("Название", true)
	form.Title.Examples = []string{"Фильм о котиках", "Фильм о ежиках"}
	form.Title.ValidationErrors = newValidationErrors("Название фильма", "строку")
	form.Description = newFormEntry("Описание", false)
	form.Description.Examples = []string{"", "документальный фильм", "художественный фильм"}
	form.Description.ValidationErrors = newValidationErrors("Описание", "строку")
	form.ContainerBox = newFormEntry("Контейнер", false)
	form.ContainerBox.Examples = []string{"2 DVD", "1 HD DVD"}
	form.ContainerBox.ValidationErrors = newValidationErrors("Контейнер", "строку")
	// creators and distributor
	form.authorsFields(1, true)
	form.publishersFields(2, false)
	var creationStudio = OrganizationForm{
		newFormEntry("", false),
		newFormEntry("Название киностудии", false),
	}
	creationStudio.Name.Examples = []string{"НеМосфильм", "НеКиностудия"}
	creationStudio.Name.ValidationErrors = newValidationErrors("Название киностудии", "строку")
	creationStudio.Name.Refresh()
	form.Publishers[0] = creationStudio
	// year
	form.yearField()
	form.Duration = NewDurationForm()

	form.Errors = []string{}
	return form
}
