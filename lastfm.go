package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/go-resty/resty/v2"
)

type RecentTracksResponse struct {
	RecentTracks struct {
		Track []struct {
			Name   string `json:"name"`
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type TopArtistsResponse struct {
	TopArtists struct {
		Artist []struct {
			Name      string `json:"name"`
			PlayCount string `json:"playcount"`
		} `json:"artist"`
	} `json:"topartists"`
}

type TopAlbumsResponse struct {
	TopAlbums struct {
		Album []struct {
			Name   string `json:"name"`
			Artist struct {
				Name string `json:"name"`
			} `json:"artist"`
			PlayCount string `json:"playcount"`
		} `json:"album"`
	} `json:"topalbums"`
}

type TopTracksResponse struct {
	TopTracks struct {
		Track []struct {
			Name   string `json:"name"`
			Artist struct {
				Name string `json:"name"`
			} `json:"artist"`
			PlayCount string `json:"playcount"`
		} `json:"track"`
	} `json:"toptracks"`
}

// MakeLastFmRequest makes a request to the Last.fm API
func MakeLastFmRequest(username, apiKey, method, period string) ([]byte, error) {
	client := resty.New()
	req := client.R().
		SetQueryParams(map[string]string{
			"method":  method,
			"user":    username,
			"api_key": apiKey,
			"format":  "json",
		})

	// Add the period parameter if applicable
	if method == "user.gettopartists" || method == "user.gettopalbums" || method == "user.gettoptracks" {
		req.SetQueryParam("period", period)
	}

	resp, err := req.Get("http://ws.audioscrobbler.com/2.0/")

	if err != nil {
		return nil, fmt.Errorf("Oops! Something went wrong. Please try again later.")
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("The API returned an error. Please check your inputs and try again.")
	}

	return resp.Body(), nil
}

// HandleLastFmRequest handles Last.fm API requests based on the user's selection
func HandleLastFmRequest(menuSelect, username, apiKey string) {
	var method string
	var period string

	switch menuSelect {
	case "recent-tracks":
		method = "user.getrecenttracks"
	case "top-artists":
		method = "user.gettopartists"
		period = PromptForPeriod()
	case "top-albums":
		method = "user.gettopalbums"
		period = PromptForPeriod()
	case "top-tracks":
		method = "user.gettoptracks"
		period = PromptForPeriod()
	default:
		log.Fatalf("Invalid selection: %s", menuSelect)
	}

	// Show a spinner while making the API request
	var response []byte
	err := spinner.New().
		Title(fmt.Sprintf("Fetching your %s...", menuSelect)).
		Action(func() {
			var err error
			response, err = MakeLastFmRequest(username, apiKey, method, period)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}).
		Run()

	if err != nil {
		log.Fatalf("Error showing spinner: %v", err)
	}

	// Parse and display the response based on the selected method
	switch menuSelect {
	case "recent-tracks":
		DisplayRecentTracksTable(response)
	case "top-artists":
		DisplayTopArtistsTable(response, period)
	case "top-albums":
		DisplayTopAlbumsTable(response, period)
	case "top-tracks":
		DisplayTopTracksTable(response, period)
	}

	// Add a separator for better readability
	fmt.Print("\n----------------------------------------\n")
}

// PromptForPeriod prompts the user to select a time period
func PromptForPeriod() string {
	var period string
	huh.NewSelect[string]().
		Title("Select a time period:").
		Options(
			huh.NewOption("Last 7 days", "7day"),
			huh.NewOption("Last 1 month", "1month"),
			huh.NewOption("Last 3 months", "3month"),
			huh.NewOption("Last 6 months", "6month"),
			huh.NewOption("Last 12 months", "12month"),
			huh.NewOption("Overall", "overall"),
		).
		Value(&period).
		Run()
	return period
}
