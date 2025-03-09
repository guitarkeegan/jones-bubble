package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	art "github.com/guitarkeegan/jones_bubble/internal/ascii-art"
)

// TODOS
// Example Weekend Edition
// Show a top info menu, the main map/viewer, and help commands

type artwork struct {
	Title string
}

type userMessages struct {
	Hello   string
	Goodbye string
	Confirm string
}

type model struct {
	viewArt             artwork
	messages            userMessages
	playerSetupForm     *huh.Form
	playerCount         int
	playerGoalsCount    int
	characterGoalsForm  *huh.Form
	charactersSetupForm *huh.Form
	characters          map[player]character
	helpMenu            help.Model
	state               State
}

func newModel() tea.Model {
	return &model{
		viewArt: artwork{
			art.MainTitle,
		},
		messages: userMessages{
			"Hello", "Goodbye", "Press q to quit OR any other key to start",
		},
		helpMenu:   help.New(),
		characters: make(map[player]character),
		state:      initializing,
	}
}

func (m model) Init() tea.Cmd {
	return initialized.Cmd
}

func (m model) Update(tMsg tea.Msg) (tea.Model, tea.Cmd) {
	dbg(fmt.Sprintf("update state: %[1]s | msg/type: %[2]v/%[2]T", m.state, tMsg))

	switch msg := tMsg.(type) {
	case StateMsg:
		switch msg {
		case initialized:
			m.state = confirmingStart
		case startRequested:
			m.state = settingPlayerCount
		case playerCountSet:
			m.state = settingCharacters
			m.playerSetupForm = nil
		case charactersSet:
			m.state = settingGoals
			m.charactersSetupForm = nil
			m.characterGoalsForm = nil
		case goalsSet:
			m.state = startingGame
			m.characterGoalsForm = nil
		case exitRequested:
			m.state = shuttingDown
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, exitRequested.Cmd
		}
	}

	switch m.state {
	case initializing:
		return m, nil
	case confirmingStart:
		return m.updateConfirm(tMsg)
	case settingPlayerCount:
		return m.updateSetPlayerCount(tMsg)
	case settingCharacters:
		return m.updateCharacterSetupForm(tMsg)
	case settingGoals:
		return m.updateCharacterGoalsForm(tMsg)
	case startingGame:
		return m.updateGame(tMsg)
	case shuttingDown:
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m model) View() string {
	dbg("view state: %s", m.state)

	switch m.state {
	case initializing:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Hello)

	case confirmingStart:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Confirm)

	case settingPlayerCount:
		return m.playerSetupForm.View()

	case settingCharacters:
		return m.charactersSetupForm.View()

	case settingGoals:
		return m.characterGoalsForm.View()

	case startingGame:
		var s string
		for k, v := range m.characters {
			s += fmt.Sprintf("%v:\n\n%+v\n\n", k, v)
		}
		return s

	case shuttingDown:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Goodbye)

	default:
		return fmt.Sprintf("unhandled view state: %s\n", m.state)
	}
}
