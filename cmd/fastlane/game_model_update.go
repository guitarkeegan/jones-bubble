package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (gm GameModel) updateChooseDestination(msg tea.Msg) (tea.Model, tea.Cmd) {
	dbg("updateChooseDestination")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		gm.ActionsMenu = newChooseDestForm()
		return gm, gm.ActionsMenu.Init()
	}
	dbg("  not nil")

	form, cmd := gm.ActionsMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		gm.ActionsMenu = f
	} else {
		dbg("actions form is NOT OK (should never happen)")
	}

	if gm.ActionsMenu.State == huh.StateCompleted {
		dbg("  completed")
		// update the game model
		gm.CurrentLoc = gm.Board[gm.ActionsMenu.GetString(destinationKey)]
		dbg("  gm.currentLoc: %s", gm.CurrentLoc.name)
		return gm, destinationSet.Cmd
	}

	dbg("updateChooseDestinationEnd")
	return gm, cmd
}

func (gm GameModel) updateEnterLuxuryApartments(msg tea.Msg) (tea.Model, tea.Cmd) {

	const (
		EXIT = "Exit"
		REST = "Rest"
	)

	dbg("updateEnterLocation")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		// TODO: check game model to see where player lives
		gm.ActionsMenu = newLuxuryApartmentsForm(false)
		return gm, gm.ActionsMenu.Init()
	}

	dbg("  not nil")

	form, cmd := gm.ActionsMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		gm.ActionsMenu = f
	} else {
		dbg("  actions form is NOT OK (should never happen)")
	}

	if gm.ActionsMenu.State == huh.StateCompleted {
		dbg("  completed")
		action := gm.ActionsMenu.GetString(luxuryApartmentsKey)
		if action == EXIT {
			return gm, locationVisted.Cmd
		}

		dbg("end")
		// TODO: send a Msg that will update the character's hapiness
		return gm, destinationSet.Cmd
	}

	dbg("end")
	return gm, cmd
}

func (gm GameModel) updateEnterLowCostOffice(msg tea.Msg) (tea.Model, tea.Cmd) {

	const (
		EXIT = "Exit"
		REST = "Rest"
	)

	dbg("updateEnterLowCostOffice")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		// TODO: check game model to see where player lives
		gm.ActionsMenu = newLowCostHousingForm()
		return gm, gm.ActionsMenu.Init()
	}

	dbg("  not nil")

	form, cmd := gm.ActionsMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		gm.ActionsMenu = f
	} else {
		dbg("  actions form is NOT OK (should never happen)")
	}

	if gm.ActionsMenu.State == huh.StateCompleted {
		dbg("  completed")
		action := gm.ActionsMenu.GetString(lowCostHousingKey)
		switch action {
		case EXIT:
			return gm, locationVisted.Cmd
		case REST:
			// TODO: rested.Cmd
			return gm, rested.Cmd
		}
		dbg("end")
		// TODO: implement menu
		return gm, destinationSet.Cmd
	}

	dbg("end")
	return gm, cmd
}

func (gm GameModel) updateEnterEmploymentOffice(msg tea.Msg) (tea.Model, tea.Cmd) {

	const EXIT = "Exit"

	dbg("updateEnterEmploymentOffice")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		// TODO: check game model to see where player lives
		gm.ActionsMenu = newEmploymentOfficeForm()
		return gm, gm.ActionsMenu.Init()
	}

	dbg("  not nil")

	form, cmd := gm.ActionsMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		gm.ActionsMenu = f
	} else {
		dbg("  actions form is NOT OK (should never happen)")
	}

	if gm.ActionsMenu.State == huh.StateCompleted {
		dbg("  completed")
		action := gm.ActionsMenu.GetString(luxuryApartmentsKey)
		if action == EXIT {
			return gm, locationVisted.Cmd
		}

		dbg("end")
		// TODO: send a Msg that will update the character's hapiness
		// if they chose rest
		return gm, destinationSet.Cmd
	}

	dbg("end")
	return gm, cmd
}

func (gm GameModel) updateRentOfficeForm(msg tea.Msg) (tea.Model, tea.Cmd) {

	const EXIT = "Exit"

	dbg("updateRentOfficeForm")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		// TODO: check game model to see where player lives
		gm.ActionsMenu = newRentOfficeForm()
		return gm, gm.ActionsMenu.Init()
	}

	dbg("  not nil")

	form, cmd := gm.ActionsMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		gm.ActionsMenu = f
	} else {
		dbg("  actions form is NOT OK (should never happen)")
	}

	if gm.ActionsMenu.State == huh.StateCompleted {
		dbg("  completed")
		action := gm.ActionsMenu.GetString(rentOfficeKey)
		if action == EXIT {
			return gm, locationVisted.Cmd
		}

		dbg("end")
		// TODO: send a Msg that will update the character's hapiness
		// if they chose rest
		return gm, destinationSet.Cmd
	}

	dbg("end")
	return gm, cmd
}
