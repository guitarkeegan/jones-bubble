package main

// when i click on a store, i see a new menu
// how do i know which one i clicked?
// update the currentLocation
// show the view of the currentLocation
// each location is a custom form

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"maps"
)

type Job struct {
	Title            string
	BaseWage         int // in dollars
	ExperienceReq    int
	DependabilityReq int
	Degrees          []string
	Uniform          string
}

type location struct {
	name              string
	pos               int
	interiorOpenImg   string
	interiorClosedImg string
	isOpen            bool
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

type GameModel struct {
	Board       map[string]*location
	MainMap     string
	GameState   GameState
	ActionsMenu *huh.Form
	CurrentLoc  *location
	GameMsg     string
	Hours       float64
	// really??
	GameMsgCounter int
}

// Returns key: value pairs for openImgs
func loadClosedImgs(filePath string) ([][]string, error) {

	dbg("loadClosedImgs")

	var Closed = "Closed"

	kvOpenImgs := [][]string{}

	err := filepath.WalkDir(filePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(d.Name(), Closed) {
			dbg("  found %s in %s", Closed, d.Name())
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("loadClosedImgs: WalkDir: %v", err)
			}
			nameState := strings.Split(d.Name(), "_")
			dbg("  nameState length: %d", len(nameState))
			kvOpenImgs = append(kvOpenImgs, []string{
				nameState[0],
				string(data),
			})
		}

		return nil
	})
	if err != nil {
		return [][]string{}, fmt.Errorf("loadClosedImgs: %v", err)
	}

	return kvOpenImgs, nil
}

func loadClosedImg(filepath string) (string, error) {

	dbg("loadClosedImg")

	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("loadClosedImg: %w", err)
	}

	dbg("loadClosedImg END")
	return string(data), nil
}

func initializeLocations() map[string]*location {

	dbg("initializeLocations")

	// Map location fields to their data
	locations := map[string]*location{
		"luxuryApartments": {"Luxury Apartments", 0, "", "", false},
		"rentOffice":       {"Rent Office", 1, "", "", false},
		"lowCostHousing":   {"Low Cost Housing", 2, "", "", false},
		"pawnShop":         {"Pawn Shop", 3, "", "", false},
		"zMart":            {"Z-Mart", 4, "", "", false},
		"monolithBurgers":  {"Monolith Burgers", 5, "", "", false},
		"qtClothing":       {"QT Clothing", 6, "", "", false},
		"socketCity":       {"Socket City", 7, "", "", false},
		"hiTechU":          {"Hi-Tech University", 8, "", "", false},
		"employmentOffice": {"Employment Office", 9, "", "", false},
		"factory":          {"Factory", 10, "", "", false},
		"bank":             {"Bank", 11, "", "", false},
		"blacksMarket":     {"Black's Market", 12, "", "", false},
	}

	maps.Copy(locations, locations)

	// return a slice of pairs, key: img data
	closedImgPath := "assets/interiors"
	cImgs, err := loadClosedImgs(closedImgPath)
	if err != nil {
		log.Fatalf("initializeLocations: loadOpenImgs: %v", err)
	}

	dbg("initializeLocations: cImgs count: %d", len(cImgs))
	if len(cImgs) == 0 {
		log.Fatal("loadClosedImgs return 0")
	}

	// TODO: failing here
	for _, kv := range cImgs {
		dbg("  kv length is %d", len(kv))
		dbg("  before kv[0]: %+v", kv[0])
		locations[kv[0]].interiorClosedImg = kv[1]
		dbg("  after kv[0]: %+v", kv[0])
	}

	dbg("initializeLocations end")
	return locations

}

func NewGameModel() *GameModel {
	var closedImgsPath = "assets/interiors/lowCostHousing_Closed"
	closedImgData, err := loadClosedImg(closedImgsPath)
	if err != nil {
		log.Fatalf("NewGameModel: %v", err)
	}
	return &GameModel{
		Board:     initializeLocations(),
		MainMap:   makeMainMap("assets/MainMenu"),
		GameState: initializingMap,
		CurrentLoc: &location{
			name:              "Low Cost Housing",
			pos:               2,
			interiorOpenImg:   "",
			interiorClosedImg: closedImgData,
			isOpen:            false,
		},
	}
}

func (gm GameModel) Init() tea.Cmd {
	return mapInitialized.Cmd
}

func (gm GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	dbg(fmt.Sprintf("update game state: %[1]s | msg/type: %[2]v/%[2]T", gm.GameState, msg))

	var gCmds []tea.Cmd

	switch msg := msg.(type) {

	case GameStateMsg:
		switch msg {
		case destinationSet:
			// set ActionsMenu to nil an replace with location
			// recalculate relative positions in model
			// update time
			gm.ActionsMenu = nil
			gm.GameState = visitingLocation
		case locationVisted:
			gm.ActionsMenu = nil
			gm.GameMsg = ""
			gm.GameState = initializingMap
			gm.GameMsg = ""
			gm.GameMsgCounter = 0
		case rested:
			dbg("  rested")
			gm.GameMsg = GameMsg_LowCostRest
			gm.GameMsgCounter = 2
			gm.ActionsMenu = nil
			gCmds = append(gCmds, Tick())
		case mapInitialized:
		case turnStarted:

		}
	case GameTickMsg:
		dbg("GameMsgCounter: %d", gm.GameMsgCounter)
		gm.GameMsgCounter--
		if gm.GameMsgCounter <= 0 {
			gm.GameMsgCounter = 0
			gm.GameMsg = ""
			// what cmd?
			return gm, nil
		}
		return gm, Tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return gm, exitRequested.Cmd
		}
	}

	switch gm.GameState {
	case initializingMap:
		dbg("initializingMap")
		gm, cmd := gm.updateChooseDestination(msg)
		dbg("cmd: %[1]v, %[1]T", cmd)
		return gm, cmd
	case visitingLocation:
		// switch on current location?
		switch gm.CurrentLoc.name {
		case camelCaseToTitle(luxuryApartments):
			gm, cmd := gm.updateEnterLuxuryApartments(msg)
			return gm, cmd
		case camelCaseToTitle(rentOffice):
			gm, cmd := gm.updateRentOfficeForm(msg)
			gCmds = append(gCmds, cmd)
			return gm, tea.Batch(gCmds...)
		case camelCaseToTitle(lowCostHousing):
			dbg("  visitingLocation: lowCostHousing")
			gm, cmd := gm.updateEnterLowCostOffice(msg)
			gCmds = append(gCmds, cmd)
			return gm, tea.Batch(gCmds...)
		case camelCaseToTitle(employmentOffice):
			gm, cmd := gm.updateEnterEmploymentOffice(msg)
			return gm, cmd
		default:
			dbg("  game_update: on Default case: shouldn't happen")
			return gm, nil
		}

	default:
		return gm, nil
	}

}

func (gm GameModel) View() string {

	dbg("game view state: %s", gm.GameState)

	switch gm.GameState {

	case initializingMap:

		if gm.CurrentLoc == nil || gm.ActionsMenu == nil {
			dbg("  %+v", gm.CurrentLoc)
			return fmt.Sprintln("loading...")
		}

		selectDestForm := gm.ActionsMenu.View()
		f := lipgloss.NewStyle().MaxWidth(20).Render(selectDestForm)
		// TODO render map

		return ClearScreen() + gm.MainMap + "\n" + f

	case visitingLocation:
		switch gm.CurrentLoc.name {
		case camelCaseToTitle(luxuryApartments):
			if gm.CurrentLoc.isOpen {
				return ClearScreen() + "shop is open..."
			}
			return ClearScreen() + gm.CurrentLoc.interiorClosedImg + "\n\n" + gm.ActionsMenu.View()
		case camelCaseToTitle(rentOffice):
			if gm.CurrentLoc.isOpen {
				return "rent office is open"
			}
			return ClearScreen() + gm.CurrentLoc.interiorClosedImg + "\n\n" + gm.ActionsMenu.View()
		case camelCaseToTitle(lowCostHousing):
			if gm.CurrentLoc.isOpen {
				return "lowCostHousing is open"
			}
			return ClearScreen() + gm.CurrentLoc.interiorClosedImg + "\n\n" + gameMsgBlock.Render(gm.GameMsg) +
				"\n" + gm.ActionsMenu.View()
		case camelCaseToTitle(employmentOffice):
			// employment is always open
			return ClearScreen() + gm.CurrentLoc.interiorOpenImg
		}
		return "missed the apartement case"
	case startingTurn:
		return fmt.Sprintln("starting turn...")
	default:
		return fmt.Sprintf("gameState not handled: %s\n", gm.GameState)
	}

}
