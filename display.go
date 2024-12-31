package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

// DisplayRecentTracksTable parses and displays recent tracks in a table
func DisplayRecentTracksTable(response []byte) {
	var data RecentTracksResponse
	if err := json.Unmarshal(response, &data); err != nil {
		log.Fatalf("Error parsing recent tracks: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Track", "Artist"})

	for _, track := range data.RecentTracks.Track {
		table.Append([]string{track.Name, track.Artist.Text})
	}

	fmt.Println("\nHere are your recent tracks:")
	table.Render()
}

// DisplayTopArtistsTable parses and displays top artists in a table
func DisplayTopArtistsTable(response []byte, period string) {
	var data TopArtistsResponse
	if err := json.Unmarshal(response, &data); err != nil {
		log.Fatalf("Error parsing top artists: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Artist", "Play Count"})

	for _, artist := range data.TopArtists.Artist {
		table.Append([]string{artist.Name, artist.PlayCount})
	}

	fmt.Printf("\nYour top artists (%s):\n", period)
	table.Render()
}

// DisplayTopAlbumsTable parses and displays top albums in a table
func DisplayTopAlbumsTable(response []byte, period string) {
	var data TopAlbumsResponse
	if err := json.Unmarshal(response, &data); err != nil {
		log.Fatalf("Error parsing top albums: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Album", "Artist", "Play Count"})

	for _, album := range data.TopAlbums.Album {
		table.Append([]string{album.Name, album.Artist.Name, album.PlayCount})
	}

	fmt.Printf("\nYour top albums (%s):\n", period)
	table.Render()
}

// DisplayTopTracksTable parses and displays top tracks in a table
func DisplayTopTracksTable(response []byte, period string) {
	var data TopTracksResponse
	if err := json.Unmarshal(response, &data); err != nil {
		log.Fatalf("Error parsing top tracks: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Track", "Artist", "Play Count"})

	for _, track := range data.TopTracks.Track {
		table.Append([]string{track.Name, track.Artist.Name, track.PlayCount})
	}

	fmt.Printf("\nYour top tracks (%s):\n", period)
	table.Render()
}
