package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
	"math"
	"strings"
)

/* ------------ Person Structure  ------------ */

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) addPoint() {
	if string(p.FirstName[len(p.FirstName)-1]) != "." {
		p.FirstName += "."
	}
}

func (p *Person) InitialsSurname() string {
	p.addPoint()
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) SurnameInitials() string {
	p.addPoint()
	return fmt.Sprintf("%s %s", p.LastName, p.FirstName)
}

func ListPeople(people []Person, limit int) string {
	var peopleList []string
	for _, person := range people {
		peopleList = append(peopleList, person.InitialsSurname())
	}
	var namesString string
	if len(peopleList) > limit {
		namesString = strings.Join(peopleList[:limit], ", ") + " [и др.]"
	} else {
		namesString = strings.Join(peopleList, ", ")
	}
	return namesString
}

/* ------------ Person Form Structure  ------------ */

type PersonForm struct {
	FirstName *FormEntry
	LastName  *FormEntry
}

func (pf *PersonForm) New(required bool, personI int, fieldTitle string) {
	pf.FirstName = newFormEntry("И.И.", required)
	pf.FirstName.Examples = []string{"П.П.", "С.С.", "Е.Е.", "Г.Г."}
	pf.FirstName.ValidationErrors = newValidationErrors(fmt.Sprintf("%s %d (имя)", fieldTitle, personI), "строку")
	pf.LastName = newFormEntry("Иванов", required)
	pf.LastName.Examples = []string{"Петров", "Сидоров", "Кузнецов", "Крылов"}
	pf.LastName.ValidationErrors = newValidationErrors(fmt.Sprintf("%s %d (фамилия)", fieldTitle, personI), "строку")
}

func (pf *PersonForm) Clear() {
	pf.FirstName.SetText("")
	pf.LastName.SetText("")
}

func (pf *PersonForm) Example() {
	pf.FirstName.Example()
	pf.LastName.Example()
}

func (pf *PersonForm) Empty() bool {
	return pf.FirstName.IsEmpty() && pf.LastName.IsEmpty()
}

func (pf *PersonForm) Full() bool {
	return !(pf.FirstName.IsEmpty()) && !(pf.LastName.IsEmpty())
}

func PeopleContainer(people []PersonForm) *fyne.Container {
	var gridColumns = int(math.Min(float64(len(people)), 3))
	var peopleFields = container.NewAdaptiveGrid(gridColumns)
	for _, person := range people {
		var personContainerLayout = customLayouts.NewRatioLayout(0.3, 0.7)
		var personContainer = container.New(personContainerLayout, person.FirstName, person.LastName)
		peopleFields.Add(personContainer)
	}
	return container.NewVBox(peopleFields)
}

func FieldsEmpty(fields []PersonForm) bool {
	for _, field := range fields {
		if field.Full() {
			return false
		}
	}
	return true
}

func PeopleFromForm(formFields []PersonForm) []Person {
	var people []Person
	for _, field := range formFields {
		if field.Full() {
			var person = Person{
				FirstName: field.FirstName.TrimText(),
				LastName:  field.LastName.TrimText(),
			}
			people = append(people, person)
		}
	}
	return people
}
