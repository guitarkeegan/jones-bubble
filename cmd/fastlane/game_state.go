package main

import tea "github.com/charmbracelet/bubbletea"

type GameState string

func (gs GameState) String() string {
	return string(gs)
}

type GameStateMsg string

func (gsm GameStateMsg) String() string {
	return string(gsm)
}

func (gsm GameStateMsg) Cmd() tea.Msg {
	return gsm
}

const (
	initializingMap GameState    = "initializing map"
	mapInitialized  GameStateMsg = "map initialized"

	startingTurn GameState    = "starting turn"
	turnStarted  GameStateMsg = "turn started"

	initializingForm GameState    = "initializing form"
	formInitialized  GameStateMsg = "form initialized"

	settingDestination GameState    = "setting destination"
	destinationSet     GameStateMsg = "destination set"
)
