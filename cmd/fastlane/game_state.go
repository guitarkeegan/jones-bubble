package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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

type GameTickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return GameTickMsg(t)
	})
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

	visitingLocation GameState    = "visiting location"
	locationVisted   GameStateMsg = "location visited"

	rested GameStateMsg = "rested"

	removeGameMsg GameStateMsg = "remove game msg"
)
