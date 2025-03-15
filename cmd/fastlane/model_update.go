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
			// TODO: need to update this to work with the new
			// character
			m.characters[player(i)] = character{name: name(m.charactersSetupForm.GetString(fmt.Sprintf("player%d", i)))}
		}
		return m, charactersSet.Cmd
	}

	return m, cmd
}

func (m model) updateCharacterGoalsForm(msg tea.Msg) (tea.Model, tea.Cmd) {

	dbg(fmt.Sprintf("updateCharacterGoalsForm: %[1]s | msg/type: %[2]v/%[2]T", m.state, msg))

	if m.characterGoalsForm == nil {
		dbg("  cgf is nil")
		m.characterGoalsForm = newCharacterSelectGoalsForm(m.playerGoalsCount + 1)
		return m, m.characterGoalsForm.Init()
	}

	var cmds []tea.Cmd

	form, cmd := m.characterGoalsForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		dbg("  update cgf")
		m.characterGoalsForm = f
		cmds = append(cmds, cmd)
	} else {
		dbg("  characters goals form is NOT OK (should never happen)")
	}

	if m.characterGoalsForm.State == huh.StateCompleted {
		dbg("  completed state")
		// TODO refactor this
		m.playerGoalsCount++
		if m.playerGoalsCount < m.playerCount {
			career := m.characterGoalsForm.GetInt(careerGoalKey)
			money := m.characterGoalsForm.GetInt(moneyGoalKey)
			edu := m.characterGoalsForm.GetInt(educationGoalKey)
			happiness := m.characterGoalsForm.GetInt(happinessGoalKey)
			if p, ok := m.characters[player(m.playerGoalsCount)]; ok {
				p.careerGoal = career
				p.moneyGoal = money
				p.educationGoal = edu
				p.happinessGoal = happiness
				m.characters[player(m.playerGoalsCount)] = p
			}
			return m, charactersSet.Cmd
		}
		cmds = append(cmds, goalsSet.Cmd)

	}
	dbg("  outer return")
	dbg("updateCharacterGoalsForm End")
	return m, tea.Batch(cmds...)
}

func (m model) updateGame(msg tea.Msg) (tea.Model, tea.Cmd) {

	// TODO:
	// show map and current character
	// select location

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, exitRequested.Cmd
		}
	}

	return m, nil
}
