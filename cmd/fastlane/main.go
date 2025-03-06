package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(newModel()).Run()
	if err != nil {
		log.Fatalf("whoopsies: %s", err)
		os.Exit(1)
	}
}
