package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	art "github.com/guitarkeegan/jones_bubble/internal/ascii-art"
)

func main() {
	_, err := tea.NewProgram(newModel()).Run()
	if err != nil {
		log.Fatalf("whoopsies: %s", err)
	}
	fmt.Println(art.MainTitle)
}
