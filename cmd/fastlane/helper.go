package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

func createLocation(data locData) location {
	return location{
		img:              loadLocationFile(data.asset),
		name:             data.name,
		relativeDistance: data.dist,
		pos:              data.pos,
	}
}
