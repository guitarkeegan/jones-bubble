package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateConfirm(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, exitRequested.Cmd

		default:
			return m, startRequested.Cmd
		}

	default:
		return m, nil
	}
}

func (m model) updateSetPlayerCount(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.playerSetupForm == nil {
		m.playerSetupForm = newPlayerSetupForm()
		return m, m.playerSetupForm.Init()
	}

	form, cmd := m.playerSetupForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.playerSetupForm = f
	} else {
		dbg("numOfPlayers form is NOT OK (should never happen)")
	}

	if m.playerSetupForm.State == huh.StateCompleted {
		m.playerCount = m.playerSetupForm.GetInt(playerCountKey)
		return m, playerCountSet.Cmd
	}

	return m, cmd
}

func (m model) updateCharacterSetupForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.charactersSetupForm == nil {
		m.charactersSetupForm = newCharactersSelectForm(m.playerCount)
		return m, m.charactersSetupForm.Init()
	}

	// Process the form
	form, cmd := m.charactersSetupForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.charactersSetupForm = f
	} else {
		dbg("characters form is NOT OK (should never happen)")
	}

	if m.charactersSetupForm.State == huh.StateCompleted {
		for i := 1; i < m.playerCount+1; i++ {
			m.characters[player(i)] = character(m.charactersSetupForm.GetString(fmt.Sprintf("player%d", i)))
		}
		return m, charactersSet.Cmd
	}

	return m, cmd
}
