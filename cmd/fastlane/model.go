package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	art "github.com/guitarkeegan/jones_bubble/internal/ascii-art"
)

// TODOS
// Start Screen, press q to quit, or any other to start
// How many players?
// Choose Players
// Example Weekend Edition
// Show a top info menu, the main map/viewer, and help commands

type charDoneMsg string

func charDone() tea.Msg {
	return charDoneMsg("done")
}

type view int

const (
	start view = iota
	numOfPlayers
	chooseCharacter
	game
	end
)

func (v view) String() string {
	lookup := []string{
		"start",
		"numOfPlayers",
		"chooseCharacter",
		"game",
		"end",
	}
	s := "unknown"
	if int(v) < len(lookup) {
		s = lookup[v]
	}
	return s
}

type character = string

const (
	george character = "George"
	john   character = "John"
	paul   character = "Paul"
	ringo  character = "Ringo"
)

type player int

const (
	player1 player = iota + 1
	player2
	player3
	player4
)

const (
	playerCount = "playerCount"
)

type model struct {
	viewArt            []art.Element
	currentState       view
	numPlayers         int
	playerSetupForm    *huh.Form
	characterSetupForm *huh.Form
	playerCharacter    map[player]character
	helpMenu           help.Model
	quitting           bool
}

func newModel() tea.Model {
	return &model{
		viewArt:         []art.Element{art.MainTitle, "", "", "", "endGame"},
		helpMenu:        help.New(),
		currentState:    start,
		playerSetupForm: initplayerSetupForm(),
		playerCharacter: map[player]string{},
	}
}

func (m model) Init() tea.Cmd {
	return m.playerSetupForm.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	dbg(fmt.Sprintf("state: %[1]s | msg/type: %[2]v/%[2]T", m.currentState, msg))
	switch m.currentState {
	case start:
		return m.updateStart(msg)
	case numOfPlayers:
		return m.updateNumOfPlayers(msg)
	case chooseCharacter:
		if m.characterSetupForm == nil {
			m.characterSetupForm = initCharacterSelectForm(m.numPlayers)
			cmd := m.characterSetupForm.Init()
			return m, cmd
		}
		dbg("charDoneMsg")
		// switch msg := msg.(type) {
		// case charDoneMsg:
		dbg("m.characterSetupForm != nil")
		return m.updateCharacterSetupForm(msg)
		//}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return string(m.viewArt[end])
	}

	switch m.currentState {
	case start:
		startMsg := "Press q to quit OR any other key to start\n"
		return string(m.viewArt[start]) + "\n\n" + startMsg
	case numOfPlayers:
		return m.playerSetupForm.View()
	case chooseCharacter:
		if m.characterSetupForm == nil {
			return fmt.Sprintln("loading...")
		}
		return m.characterSetupForm.View()
	case game:
		return fmt.Sprintln("Lets Play! 🚀")
	}

	startMsg := "Press q to quit OR any other key to start\n"
	return string(m.viewArt[start]) + "\n\n" + startMsg
}
