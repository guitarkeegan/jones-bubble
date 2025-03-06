package main

import (
	"fmt"
	"log"
	"os"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	file, err := os.Create("logs")
	if err != nil {
		log.Fatal("nooooo!!!")
	}
	// truncate = delete the rest
	return func(format string, as ...any) {
		fmt.Fprintf(file, format+"\n", as...)
	}
}()
