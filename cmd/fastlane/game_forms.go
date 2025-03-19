package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func newChooseDestForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key(destinationKey).
				Options(huh.NewOptions(
					"luxuryApartments",
					"rentOffice",
					"lowCostHousing",
					"pawnShop",
					"zMart",
					"monolithBurgers",
					"qtClothing",
					"socketCity",
					"hiTechU",
					"employmentOffice",
					"factory",
					"bank",
					"blacksMarket",
				)...).
				Title(fmt.Sprintf("Select Destination")).
				Description("Where to?"),
		),
	)
}

// TODO: will need to pass in.. player state, and economy state
func newLuxuryApartmentsForm(isResident bool) *huh.Form {
	if isResident {
		return huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Relax?").
					Description("Take some time for yourself, and kick back!").
					Key(luxuryApartmentsKey).
					Affirmative("Relax").
					Negative("Exit"),
			),
		)
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key(luxuryApartmentsKey).
				Options(huh.NewOptions(
					"Exit",
				)...).
				Title(fmt.Sprintf("No Entry")).
				Description("You will have to visit the rent office in order to enter the building."),
		),
	)
}
