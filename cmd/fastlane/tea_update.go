package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateStart(msg tea.Msg) (tea.Model, tea.Cmd) {
	// start game
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			m.currentState = numOfPlayers
		}
	}

	return m, nil
}

func (m model) updateNumOfPlayers(msg tea.Msg) (tea.Model, tea.Cmd) {
	dbg("updateNumOfPlayers start")
	// Process the form
	form, cmd := m.playerSetupForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		dbg("  ok")
		m.playerSetupForm = f
	} else {
		dbg("  !ok")
	}

	if m.playerSetupForm.State == huh.StateCompleted {
		// set player count and update view
		dbg("  huh.StateCompleted")
		m.numPlayers = m.playerSetupForm.GetInt(playerCount)
		m.currentState = chooseCharacter
	}

	dbg("  current state: %v", m.currentState)
	dbg("  numberOfPlayers: %d", m.numPlayers)
	dbg("updateNumOfPlayers end")
	return m, cmd
}

func (m model) updateCharacterSetupForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	dbg("updateCharacterSetupForm start")
	// Process the form
	form, cmd := m.characterSetupForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		dbg("  ok")
		m.characterSetupForm = f
	} else {
		dbg("  !ok")
	}

	if m.characterSetupForm.State == huh.StateCompleted {
		dbg("  stateCompleted start")
		// map players to characters
		for i := 1; i < m.numPlayers+1; i++ {
			m.playerCharacter[player(i)] = character(m.characterSetupForm.GetString(fmt.Sprintf("player%d", i)))
		}
		dbg("  stateCompleted end")
		m.currentState = game
	}

	dbg("updateCharacterSetupForm end")
	return m, cmd
}
