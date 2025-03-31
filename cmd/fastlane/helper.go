package main

import (
	"log"
	"os"
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func camelCaseToTitle(input string) string {
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	formatted := re.ReplaceAllString(input, `$1 $2`)

	t := cases.Title(language.English)
	return t.String(formatted)
}

func makeMainMap(path string) string {
	img, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("loadMainMap: %v", err)
	}

	return string(img)
}

func ClearScreen() string {
	return "\033[H\033[2J"
}
