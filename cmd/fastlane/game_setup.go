package main

import (
	"fmt"
	"io/fs"
	"log"
	"maps"
	"os"
	"path/filepath"
	"strings"
)

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

	for _, kv := range cImgs {
		dbg("  kv length is %d", len(kv))
		dbg("  before kv[0]: %+v", kv[0])
		locations[kv[0]].interiorClosedImg = kv[1]
		dbg("  after kv[0]: %+v", kv[0])
	}

	dbg("initializeLocations end")
	return locations

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
