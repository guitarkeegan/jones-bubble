package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func initplayerSetupForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Key(playerCountKey).
				Options(huh.NewOptions(1, 2, 3, 4)...).
				Title("Players").
				Description("Choose how many players there will be."),
		),
	)
}

func initCharacterSelectForm(playerCount int) *huh.Form {
	dbg("initCharacterSelectForm start")
	characterSelections := []*huh.Group{}
	for i := 1; i <= playerCount; i++ {
		characterSelections = append(characterSelections, huh.NewGroup(
			huh.NewSelect[string]().
				Key(fmt.Sprintf("player%d", i)).
				Options(huh.NewOptions(john, george, ringo, paul)...).
				Title(fmt.Sprintf("Player %d Character", i)).
				Description(fmt.Sprintf("Choose a character for Player %d.", i)),
		))
	}

	dbg("initCharacterSelectForm end")
	return huh.NewForm(characterSelections...)
}
