package main

import "github.com/charmbracelet/lipgloss"

var titleBlock = lipgloss.NewStyle().
	Background(lipgloss.Color("#EB96A9")).
	Foreground(lipgloss.Color("234"))

var gameMsgBlock = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f0b4dc")).
	Italic(true)
