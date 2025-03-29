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

func newEmploymentOfficeForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key(luxuryApartmentsKey).
				Options(huh.NewOptions(
					"Exit",
				)...).
				Title(fmt.Sprintf("No Entry")).
				Description("Get ya job here!"),
		),
	)
}

func newRentOfficeForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key(rentOfficeKey).
				Options(huh.NewOptions(
					"Exit",
				)...).
				Title(fmt.Sprintf("We are Closed")).
				Description("Closed for now"),
		),
	)
}

func newLowCostHousingForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key(lowCostHousingKey).
				Options(huh.NewOptions(
					// TODO: update rest stat and +6 hours
					"Rest",
					"Exit",
				)...).
				Title(fmt.Sprintf("We are Closed")).
				Description("Closed for now"),
		),
	)
}
