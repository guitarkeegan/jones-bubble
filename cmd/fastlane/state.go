package main

import tea "github.com/charmbracelet/bubbletea"

type State string

func (t State) String() string {
	return string(t)
}

type StateMsg string

func (t StateMsg) String() string {
	return string(t)
}

func (t StateMsg) Cmd() tea.Msg {
	return t
}

const (
	initializing State    = "initializing"
	initialized  StateMsg = "initialized"

	confirmingStart State    = "confirming start"
	exitRequested   StateMsg = "exit requested"
	startRequested  StateMsg = "start requested"

	settingPlayerCount State    = "setting player count"
	playerCountSet     StateMsg = "player count set"

	settingCharacters State    = "setting characters"
	charactersSet     StateMsg = "characters set"

	startingGame State    = "starting game"
	gameStarted  StateMsg = "game started"

	shuttingDown State = "shutting down"
)
