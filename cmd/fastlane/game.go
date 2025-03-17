package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define location data (asset name, display name, distance, position)
type locData struct {
	asset string
	name  string
	dist  int
	pos   int
}

type location struct {
	img              string
	name             string
	relativeDistance int
	pos              int
}

type Accommodator interface {
	HowMuch() int
	PayRent(payment int) int
	Rest(time int) int
	Talk(event string) string
	Evict(p player) int
}

type Employer interface {
	Apply(edu, exp int) bool
	Work(time int) int
	Promote(p player) int
	Quit()
	Talk(event string) string
}

type Seller[T any] interface {
	Buy(money int) T
	Talk(event string) string
}

type Buyer[T any] interface {
	Sell(item T) int
	Talk(event string) string
}

type GameBoard struct {
	luxuryApartments location
	rentOffice       location
	lowCostHousing   location
	pawnShop         location
	zMart            location
	monolithBurgers  location
	qtClothing       location
	socketCity       location
	hiTechU          location
	employmentOffice location
	factory          location
	bank             location
	blacksMarket     location
}

type GameModel struct {
	Board     *GameBoard
	GameState GameState
}

func initializeLocations() *GameBoard {

	// Map location fields to their data
	locationMap := map[string]locData{
		"luxuryApartments": {"SecurityApartments", "Luxury Apartments", 0, 0},
		"rentOffice":       {"RentOffice", "Rent Office", 1, 1},
		"lowCostHousing":   {"LowCostHousing", "Low Cost Housing", 2, 2},
		"pawnShop":         {"PawnShop", "Pawn Shop", 3, 3},
		"zMart":            {"ZMart", "Z-Mart", 4, 4},
		"monolithBurgers":  {"MonolithBurgers", "Monolith Burgers", 5, 5},
		"qtClothing":       {"QTClothing", "QT Clothing", 6, 6},
		"socketCity":       {"SocketCity", "Socket City", 7, 7},
		"hiTechU":          {"HiTechU", "Hi-Tech University", 8, 8},
		"employmentOffice": {"EmploymentOffice", "Employment Office", 9, 9},
		"factory":          {"Factory", "Factory", 10, 10},
		"bank":             {"Bank", "Bank", 11, 11},
		"blacksMarket":     {"BlacksMarket", "Black's Market", 12, 12},
	}

	// Create board using the location map
	return &GameBoard{
		luxuryApartments: createLocation(locationMap["luxuryApartments"]),
		rentOffice:       createLocation(locationMap["rentOffice"]),
		lowCostHousing:   createLocation(locationMap["lowCostHousing"]),
		pawnShop:         createLocation(locationMap["pawnShop"]),
		zMart:            createLocation(locationMap["zMart"]),
		monolithBurgers:  createLocation(locationMap["monolithBurgers"]),
		qtClothing:       createLocation(locationMap["qtClothing"]),
		socketCity:       createLocation(locationMap["socketCity"]),
		hiTechU:          createLocation(locationMap["hiTechU"]),
		employmentOffice: createLocation(locationMap["employmentOffice"]),
		factory:          createLocation(locationMap["factory"]),
		bank:             createLocation(locationMap["bank"]),
		blacksMarket:     createLocation(locationMap["blacksMarket"]),
	}
}

func NewGameModel() *GameModel {
	return &GameModel{
		Board:     initializeLocations(),
		GameState: initializingMap,
	}
}

func (gm GameModel) Init() tea.Cmd {
	return mapInitialized.Cmd
}

func (gm GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	dbg(fmt.Sprintf("update state: %[1]s | msg/type: %[2]v/%[2]T", gm.GameState, msg))

	switch msg := msg.(type) {

	case GameStateMsg:
		switch msg {
		case mapInitialized:
		case turnStarted:

		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return gm, exitRequested.Cmd
		}
	}

	switch gm.GameState {
	case initializingMap:
		return gm, nil
	default:
		return gm, nil
	}

}

func (gm GameModel) View() string {

	dbg("gm: View")

	switch gm.GameState {

	case initializingMap:
		row1 := lipgloss.JoinHorizontal(
			lipgloss.Top,
			locationBlock.Render(gm.Board.luxuryApartments.img+"\n"+titleBlock.Render(gm.Board.luxuryApartments.name)),
			locationBlock.Render(gm.Board.rentOffice.img+"\n"+titleBlock.Render(gm.Board.rentOffice.name)),
			locationBlock.Render(gm.Board.lowCostHousing.img+"\n"+titleBlock.Render(gm.Board.lowCostHousing.name)),
			locationBlock.Render(gm.Board.pawnShop.img+"\n"+titleBlock.Render(gm.Board.pawnShop.name)),
			locationBlock.Render(gm.Board.zMart.img+"\n"+titleBlock.Render(gm.Board.zMart.name)),
		)
		row2 := lipgloss.JoinHorizontal(
			lipgloss.Top,
			locationBlock.Render(gm.Board.monolithBurgers.img+"\n"+titleBlock.Render(gm.Board.monolithBurgers.name)),
			locationBlock.Render(gm.Board.qtClothing.img+"\n"+titleBlock.Render(gm.Board.qtClothing.name)),
			locationBlock.Render(gm.Board.socketCity.img+"\n"+titleBlock.Render(gm.Board.socketCity.name)),
			locationBlock.Render(gm.Board.hiTechU.img+"\n"+titleBlock.Render(gm.Board.hiTechU.name)),
			locationBlock.Render(gm.Board.employmentOffice.img+"\n"+titleBlock.Render(gm.Board.employmentOffice.name)),
		)
		row3 := lipgloss.JoinHorizontal(
			lipgloss.Top,
			locationBlock.Render(gm.Board.factory.img+"\n"+titleBlock.Render(gm.Board.factory.name)),
			locationBlock.Render(gm.Board.bank.img+"\n"+titleBlock.Render(gm.Board.bank.name)),
			locationBlock.Render(gm.Board.blacksMarket.img+"\n"+titleBlock.Render(gm.Board.blacksMarket.name)),
		)

		return row1 + "\n" + row2 + "\n" + row3

	case startingTurn:
		return fmt.Sprintln("starting turn...")
	default:
		return fmt.Sprintf("gameState not handled: %s\n", gm.GameState)
	}

}
