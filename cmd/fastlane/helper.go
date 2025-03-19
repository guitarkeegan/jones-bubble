package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/charmbracelet/lipgloss"
)

func loadLocationFile(name string) string {
	dbg("loadLocationFile")
	filePath := filepath.Join("assets", fmt.Sprintf("Location_%s", name))
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Log the error but don't crash the application
		log.Printf("WARNING: Failed to load location file %s: %v", filePath, err)
		// Return a placeholder or default content instead
		return fmt.Sprintf("[Missing asset for %s]", name)
	}
	dbg("loadLocationFile end")
	return string(content)
}

func (gm GameModel) getLocationBlock(l string) lipgloss.Style {
	if gm.CurrentLoc.name == l {
		return currentLocationBlock
	}
	return locationBlock
}

func camelCaseToTitle(input string) string {
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	formatted := re.ReplaceAllString(input, `$1 $2`)

	t := cases.Title(language.English)
	return t.String(formatted)
}
