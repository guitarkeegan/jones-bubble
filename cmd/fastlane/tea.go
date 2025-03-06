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

func (m model) Init() tea.Cmd {
	return m.playerSetupForm.Init()
}

func InitModel() tea.Model {
	return model{
		viewArt:      []art.Element{art.MainTitle, "", "", "", "endGame"},
		helpMenu:     help.New(),
		currentState: start,
		playerSetupForm: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Key(playerCount).
					Options(huh.NewOptions(1, 2, 3, 4)...).
					Title("Players").
					Description("Choose how many players there will be."),
			),
		),
	}
}

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

func initCharacterSelectForm(playerCount int) *huh.Form {
	dbg("initCharacterSelectForm start")
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

	dbg("initCharacterSelectForm end")
	return huh.NewForm(characterSelections...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	dbg("Update: %v", m.currentState)

	switch m.currentState {
	case start:
		return m.updateStart(msg)
	case numOfPlayers:
		return m.updateNumOfPlayers(msg)
	case chooseCharacter:
		if m.characterSetupForm == nil {
			dbg("characterSetupForm == nil")
			// TODO: intialize character selecter
			m.characterSetupForm = initCharacterSelectForm(m.numPlayers)
			dbg("m.characterSetupForm.Init()")
			cmd := m.characterSetupForm.Init()
			dbg("returning")
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
