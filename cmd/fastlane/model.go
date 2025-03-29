package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	art "github.com/guitarkeegan/jones_bubble/internal/ascii-art"
)

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
	GameModel           tea.Model
	state               State
	currentTurn         player
}

func newModel() tea.Model {
	return &model{
		viewArt: artwork{
			art.MainTitle,
		},
		messages: userMessages{
			"Hello", "Goodbye", "Press q to quit OR any other key to start",
		},
		helpMenu:    help.New(),
		characters:  make(map[player]character),
		state:       initializing,
		currentTurn: player(1),
		GameModel:   NewGameModel(),
	}
}

func (m model) Init() tea.Cmd {
	return initialized.Cmd
}

func (m model) Update(tMsg tea.Msg) (tea.Model, tea.Cmd) {
	dbg(fmt.Sprintf("update state: %[1]s | msg/type: %[2]v/%[2]T", m.state, tMsg))

	var gCmd tea.Cmd

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
		case exitRequested:
			m.state = shuttingDown
		}
		// mixing the models a bit here...
	case GameStateMsg:
		switch msg {
		case rested:
			// better way to do this?
			// TODO: not updating properly
			dbg("m.currentTurn: %d", m.currentTurn)
			currentRest := m.characters[m.currentTurn].relaxation + increaseRest
			dbg("  currentRest: %d", currentRest)
			currentCharacter := m.characters[m.currentTurn]
			currentCharacter.relaxation = currentRest
			m.characters[m.currentTurn] = currentCharacter
			m.GameModel, gCmd = m.GameModel.Update(tMsg)
			dbg("Characters: %+v", m.characters)
			return m, gCmd
		}
	case GameTickMsg:
		m.GameModel, gCmd = m.GameModel.Update(tMsg)
		return m, gCmd
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
		var gCmd tea.Cmd
		m.GameModel, gCmd = m.GameModel.Update(tMsg)
		return m, gCmd
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
		if m.characterGoalsForm == nil {
			return "loading..."
		}
		return m.characterGoalsForm.View()

	case startingGame:
		return m.GameModel.View()

	case shuttingDown:
		return fmt.Sprintf("%s\n%s\n", m.viewArt.Title, m.messages.Goodbye)

	default:
		return fmt.Sprintf("unhandled view state: %s\n", m.state)
	}
}
