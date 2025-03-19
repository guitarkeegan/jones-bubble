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
		gm.CurrentLoc = gm.Board[gm.ActionsMenu.GetString(destinationKey)]
		// update the game model
		return gm, destinationSet.Cmd
	}

	dbg("updateChooseDestinationEnd")
	return gm, cmd
}

func (gm GameModel) updateEnterLocation(msg tea.Msg) (tea.Model, tea.Cmd) {

	dbg("updateEnterLocation")
	if gm.ActionsMenu == nil {
		dbg("  is nil")
		// TODO: this won't work because the field expects
		// a *huh.Form
		// gm.ActionsMenu = newEnterLocationForm()
		return gm, gm.ActionsMenu.Init()
	}

	dbg("%v", msg)
	dbg("  not nil")

	dbg("end")
	return gm, nil
}
