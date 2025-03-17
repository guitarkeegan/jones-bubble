package main

import "github.com/charmbracelet/lipgloss"

var locationBlock = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("7"))

var titleBlock = lipgloss.NewStyle().
	Background(lipgloss.Color("#EB96A9")).
	Foreground(lipgloss.Color("234"))
