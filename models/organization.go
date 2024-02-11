package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gostCituations/ui/customLayouts"
)

/* ------------ Organization Structure  ------------ */

type Organization struct {
	Address string
	Name    string
}

func (org *Organization) String() string {
	if !(org.Address == "") && !(org.Name == "") {
		return fmt.Sprintf("%s: %s", org.Address, org.Name)
	} else {
		if org.Address == "" {
			return fmt.Sprintf("[не ук.]: %s", org.Name)
		} else if org.Name == "" {
			return fmt.Sprintf("%s: [не ук.]", org.Address)
		} else {
			return "[не ук.]: [не ук.]"
		}
	}
}

/* ------------ Organization Form Structure  ------------ */

type OrganizationForm struct {
	Address *FormEntry
	Name    *FormEntry
}

func (of *OrganizationForm) New(required bool, orgI int) {
	of.Name = newFormEntry("Издатель", required)
	of.Name.Examples = []string{"Хорошее издательство"}
	of.Name.ValidationErrors = newValidationErrors(fmt.Sprintf("Издатель %d", orgI), "строку")
	of.Address = newFormEntry("Город", required)
	of.Address.Examples = []string{"Город Мечты"}
	of.Address.ValidationErrors = newValidationErrors(fmt.Sprintf("Издатель %d", orgI), "строку")
}

func (of *OrganizationForm) Clear() {
	of.Name.Clear()
	of.Address.Clear()
}

func (of *OrganizationForm) Example() {
	of.Name.SetText(of.Name.Examples[0])
	of.Address.SetText(of.Address.Examples[0])
}

func (of *OrganizationForm) Full() bool {
	return !((of.Name.TrimText() == "") && (of.Address.TrimText() == ""))
}

func OrganizationsContainer(organizations []OrganizationForm) *fyne.Container {
	var organizationsFields = container.NewAdaptiveGrid(len(organizations))
	for _, org := range organizations {
		var orgLayout = customLayouts.NewRatioLayout(0.6, 0.4)
		var orgContainer = container.New(orgLayout, org.Name, org.Address)
		organizationsFields.Add(orgContainer)
	}
	return container.NewVBox(organizationsFields)
}

func OrganizationsFromForm(fields []OrganizationForm) []Organization {
	var orgs []Organization
	for _, field := range fields {
		var address = field.Address.TrimText()
		var name = field.Name.TrimText()
		if !((address == "") && (name == "")) {
			var org = Organization{Name: name, Address: address}
			orgs = append(orgs, org)
		}
	}
	return orgs
}
