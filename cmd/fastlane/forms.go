package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func newPlayerSetupForm() *huh.Form {
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

func newCharactersSelectForm(playerCount int) *huh.Form {
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

	return huh.NewForm(characterSelections...)
}

func newCharacterSelectGoalsForm(playerNum int) *huh.Form {

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Key(educationGoalKey).
				Options(huh.NewOptions(50, 100, 150, 200, 250)...).
				Title(fmt.Sprintf("Player %d: Education Goal", playerNum)).
				Description("What is your educational goal?"),
			huh.NewSelect[int]().
				Key(happinessGoalKey).
				Options(huh.NewOptions(50, 100, 150, 200, 250)...).
				Title(fmt.Sprintf("Player %d: Happiness Goal", playerNum)).
				Description("What it is your happiness goal?"),
			huh.NewSelect[int]().
				Key(moneyGoalKey).
				Options(huh.NewOptions(50, 100, 150, 200, 250)...).
				Title(fmt.Sprintf("Player %d: Money Goal", playerNum)).
				Description("What it is your money goal?"),
			huh.NewSelect[int]().
				Key(careerGoalKey).
				Options(huh.NewOptions(50, 100, 150, 200, 250)...).
				Title(fmt.Sprintf("Player %d: Career Goal", playerNum)).
				Description("What it is your career goal?"),
		),
	)

}
