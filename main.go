package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	// Load environment variables
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	for {
		// Prompt the user to select an action
		menuSelect := ShowMenu()

		// Handle the user's selection
		switch menuSelect {
		case "exit":
			fmt.Println("Goodbye! Thanks for using the Last.fm CLI.")
			return
		case "clear-terminal":
			ClearTerminal()
			continue
		case "currently-playing":
			HandleCurrentlyPlaying(config.SpotifyDcCookie)
			continue
		}

		// Handle Last.fm API requests
		HandleLastFmRequest(menuSelect, config.LastFmUsername, config.LastFmApiKey)
	}
}

// ShowMenu displays the menu and returns the user's selection
func ShowMenu() string {
	var menuSelect string
	huh.NewSelect[string]().
		Title("What would you like to do?").
		Options(
			huh.NewOption("View my recent tracks", "recent-tracks"),
			huh.NewOption("View my top artists", "top-artists"),
			huh.NewOption("View my top albums", "top-albums"),
			huh.NewOption("View my top tracks", "top-tracks"),
			huh.NewOption("View my currently playing track", "currently-playing"),
			huh.NewOption("Clear Terminal", "clear-terminal"),
			huh.NewOption("Exit", "exit"),
		).
		Value(&menuSelect).
		Run()
	return menuSelect
}
