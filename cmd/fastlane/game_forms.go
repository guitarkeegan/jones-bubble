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
