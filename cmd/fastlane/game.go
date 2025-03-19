package main

// when i click on a store, i see a new menu
// how do i know which one i clicked?
// update the currentLocation
// show the view of the currentLocation
// each location is a custom form

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type location struct {
	img              string
	name             string
	relativeDistance int
	pos              int
}

var (
	luxuryApartments = "luxuryApartments"
	rentOffice       = "rentOffice"
	lowCostHousing   = "lowCostHousing"
	pawnShop         = "pawnShop"
	zMart            = "zMart"
	monolithBurgers  = "monolithBurgers"
	qtClothing       = "qtClothing"
	socketCity       = "socketCity"
	hiTechU          = "hiTechU"
	employmentOffice = "employmentOffice"
	factory          = "factory"
	bank             = "bank"
	blacksMarket     = "blacksMarket"
)

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

type GameModel struct {
	Board       map[string]*location
	GameState   GameState
	ActionsMenu *huh.Form
	CurrentLoc  *location
}

func initializeLocations() map[string]*location {

	dbg("initializeLocations")

	// Map location fields to their data
	locations := map[string]*location{
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

	for key, loc := range locations {
		dbg("  loc: %s", loc.img)
		img := loadLocationFile(loc.img)
		loc.img = img
		locations[key] = loc
	}

	dbg("initializeLocations end")
	return locations

}

func NewGameModel() *GameModel {
	return &GameModel{
		Board:     initializeLocations(),
		GameState: initializingMap,
		CurrentLoc: &location{
			img:              loadLocationFile("LowCostHousing"),
			name:             "Low Cost Housing",
			relativeDistance: 2,
			pos:              2,
		},
	}
}

func (gm GameModel) Init() tea.Cmd {
	return mapInitialized.Cmd
}

func (gm GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	dbg(fmt.Sprintf("update game state: %[1]s | msg/type: %[2]v/%[2]T", gm.GameState, msg))

	// var cmd tea.Cmd

	switch msg := msg.(type) {

	case GameStateMsg:
		switch msg {
		case destinationSet:
			// set ActionsMenu to nil an replace with location
			gm.ActionsMenu = nil
			gm.GameState = visitingLocation
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
		// TODO: do this
		dbg("initializingMap")
		gm, cmd := gm.updateChooseDestination(msg)
		dbg("cmd: %[1]v, %[1]T", cmd)
		return gm, cmd
	case visitingLocation:
		gm, cmd := gm.updateEnterLocation(msg)
		return gm, cmd

	default:
		return gm, nil
	}
}

func (gm GameModel) View() string {

	dbg("game view state: %s", gm.GameState)

	switch gm.GameState {

	case initializingMap:

		if gm.CurrentLoc == nil || gm.ActionsMenu == nil {
			dbg("%+v", gm.CurrentLoc)
			return fmt.Sprintln("loading...")
		}

		selectDestForm := gm.ActionsMenu.View()
		row1 := lipgloss.JoinHorizontal(
			lipgloss.Top,
			gm.getLocationBlock(gm.Board[luxuryApartments].name).Render(gm.Board[luxuryApartments].img+"\n"+titleBlock.Render(gm.Board[luxuryApartments].name)),
			gm.getLocationBlock(gm.Board[rentOffice].name).Render(gm.Board[rentOffice].img+"\n"+titleBlock.Render(gm.Board[rentOffice].name)),
			gm.getLocationBlock(gm.Board[lowCostHousing].name).Render(gm.Board[lowCostHousing].img+"\n"+titleBlock.Render(gm.Board[lowCostHousing].name)),
			gm.getLocationBlock(gm.Board[pawnShop].name).Render(gm.Board[pawnShop].img+"\n"+titleBlock.Render(gm.Board[pawnShop].name)),
			gm.getLocationBlock(gm.Board[zMart].name).Render(gm.Board[zMart].img+"\n"+titleBlock.Render(gm.Board[zMart].name)),
		)
		row2 := lipgloss.JoinHorizontal(
			lipgloss.Center,
			gm.getLocationBlock(gm.Board[monolithBurgers].name).Render(gm.Board[monolithBurgers].img+"\n"+titleBlock.Render(gm.Board[monolithBurgers].name)),
			gm.getLocationBlock(gm.Board[qtClothing].name).Render(gm.Board[qtClothing].img+"\n"+titleBlock.Render(gm.Board[qtClothing].name)),
			gm.getLocationBlock(gm.Board[socketCity].name).Render(gm.Board[socketCity].img+"\n"+titleBlock.Render(gm.Board[socketCity].name)),
			gm.getLocationBlock(gm.Board[hiTechU].name).Render(gm.Board[hiTechU].img+"\n"+titleBlock.Render(gm.Board[hiTechU].name)),
			gm.getLocationBlock(gm.Board[employmentOffice].name).Render(gm.Board[employmentOffice].img+"\n"+titleBlock.Render(gm.Board[employmentOffice].name)),
		)
		row3 := lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			gm.getLocationBlock(gm.Board[factory].name).Render(gm.Board[factory].img+"\n"+titleBlock.Render(gm.Board[factory].name)),
			gm.getLocationBlock(gm.Board[bank].name).Render(gm.Board[bank].img+"\n"+titleBlock.Render(gm.Board[bank].name)),
			gm.getLocationBlock(gm.Board[blacksMarket].name).Render(gm.Board[blacksMarket].img+"\n"+titleBlock.Render(gm.Board[blacksMarket].name)),
			lipgloss.NewStyle().MaxWidth(20).Render(selectDestForm),
		)

		return row1 + "\n" + row2 + "\n" + row3

	case startingTurn:
		return fmt.Sprintln("starting turn...")
	default:
		return fmt.Sprintf("gameState not handled: %s\n", gm.GameState)
	}

}
