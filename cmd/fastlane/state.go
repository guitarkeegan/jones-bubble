package main

import tea "github.com/charmbracelet/bubbletea"

type StateToken string

var (
	initializing       StateToken = "initializing"
	confirmingStart    StateToken = "confirming start"
	settingPlayerCount StateToken = "setting player count"
	settingCharacters  StateToken = "setting characters"
	startingGame       StateToken = "starting game"
	shuttingDown       StateToken = "shutting down"
)

func (t StateToken) String() string {
	return string(t)
}

type StateTransition string

func (t StateTransition) String() string {
	return string(t)
}

func (t StateTransition) Cmd() tea.Msg {
	return t
}

const (
	initialized    StateTransition = "initialized"
	startRequested StateTransition = "start requested"
	playerCountSet StateTransition = "player count set"
	charactersSet  StateTransition = "characters set"
	gameStarted    StateTransition = "game started"
	exitRequested  StateTransition = "exit requested"
)
