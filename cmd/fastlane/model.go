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
	charactersSetupForm *huh.Form
	characters          map[player]character
	helpMenu            help.Model
	currentState        StateToken
}

func newModel() tea.Model {
	return &model{
		viewArt: artwork{
			art.MainTitle,
		},
		messages: userMessages{
			"Hello", "Goodbye", "Press q to quit OR any other key to start",
		},
		helpMenu:     help.New(),
		characters:   map[player]string{},
		currentState: initializing,
	}
}

func (m model) Init() tea.Cmd {
	return initialized.Cmd
}

func (m model) Update(tMsg tea.Msg) (tea.Model, tea.Cmd) {
	dbg(fmt.Sprintf("update state: %[1]s | msg/type: %[2]v/%[2]T", m.currentState, tMsg))

	switch msg := tMsg.(type) {
	case StateTransition:
		switch msg {
		case initialized:
			m.currentState = confirmingStart
		case startRequested:
			m.currentState = settingPlayerCount
		case playerCountSet:
			m.currentState = settingCharacters
			m.playerSetupForm = nil
		case charactersSet:
			m.currentState = startingGame
			m.charactersSetupForm = nil
		case exitRequested:
			m.currentState = shuttingDown
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, exitRequested.Cmd
		}
	}

	switch m.currentState {
	case confirmingStart:
		return m.updateConfirm(tMsg)
	case settingPlayerCount:
		return m.updateSetPlayerCount(tMsg)
	case settingCharacters:
		return m.updateCharacterSetupForm(tMsg)
	case shuttingDown:
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	dbg("view state: %s", m.currentState)

	switch m.currentState {
	case initializing:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Hello)

	case confirmingStart:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Confirm)

	case settingPlayerCount:
		return m.playerSetupForm.View()

	case settingCharacters:
		return m.charactersSetupForm.View()

	case startingGame:
		return fmt.Sprintln("Lets Play! 🚀\n")

	case shuttingDown:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Goodbye)

	default:
		return fmt.Sprintf("unhandled view state: %s\n", m.currentState)
	}
}
