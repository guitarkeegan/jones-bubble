package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

var block1 = lipgloss.NewStyle().
	Background(lipgloss.Color("57")).
	Padding(1)

var block2 = lipgloss.NewStyle().
	Background(lipgloss.Color("213")).
	Padding(1)

var block3 = lipgloss.NewStyle().
	Background(lipgloss.Color("42")).
	Padding(1)

var block4 = lipgloss.NewStyle().
	Background(lipgloss.Color("208")).
	Padding(1)

var block5 = lipgloss.NewStyle().
	Background(lipgloss.Color("171")).
	Padding(1).
	Render("Block 5")

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
	monolith         location
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

func loadLocationFile(name string) string {
	filePath := filepath.Join("assets", fmt.Sprintf("Location_%s", name))
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Log the error but don't crash the application
		log.Printf("WARNING: Failed to load location file %s: %v", filePath, err)
		// Return a placeholder or default content instead
		return fmt.Sprintf("[Missing asset for %s]", name)
	}
	return string(content)
}

func initializeLocations() *GameBoard {

	// Map location fields to their data
	locationMap := map[string]locData{
		"luxuryApartments": {"SecurityApartments", "Luxury Apartments", 0, 0},
		"rentOffice":       {"RentOffice", "Rent Office", 1, 1},
		"lowCostHousing":   {"LowCostHousing", "Low Cost Housing", 2, 2},
		"pawnShop":         {"PawnShop", "Pawn Shop", 3, 3},
		"zMart":            {"ZMart", "Z-Mart", 4, 4},
		"monolith":         {"Monolith", "Monolith", 5, 5},
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
		monolith:         createLocation(locationMap["monolith"]),
		qtClothing:       createLocation(locationMap["qtClothing"]),
		socketCity:       createLocation(locationMap["socketCity"]),
		hiTechU:          createLocation(locationMap["hiTechU"]),
		employmentOffice: createLocation(locationMap["employmentOffice"]),
		factory:          createLocation(locationMap["factory"]),
		bank:             createLocation(locationMap["bank"]),
		blacksMarket:     createLocation(locationMap["blacksMarket"]),
	}
}

func createLocation(data locData) location {
	return location{
		img:              loadLocationFile(data.asset),
		name:             data.name,
		relativeDistance: data.dist,
		pos:              data.pos,
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
		horizontal := lipgloss.JoinHorizontal(
			lipgloss.Top,
			block1.Render(gm.Board.luxuryApartments.img+"\n"+gm.Board.luxuryApartments.name),
			block2.Render(gm.Board.rentOffice.img+"\n"+gm.Board.rentOffice.name),
			block3.Render(gm.Board.lowCostHousing.img+"\n"+gm.Board.lowCostHousing.name),
			block4.Render(gm.Board.pawnShop.img+"\n"+gm.Board.pawnShop.name))
		return horizontal
	case startingTurn:
		return fmt.Sprintln("starting turn...")
	default:
		return fmt.Sprintf("gameState not handled: %s\n", gm.GameState)
	}

}
