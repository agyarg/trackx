package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type SpotifyAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type SpotifyCurrentlyPlayingResponse struct {
	Item struct {
		Name    string `json:"name"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
	} `json:"item"`
}

// GetSpotifyAccessToken fetches the Spotify access token using the SPOTIFY_DC_COOKIE
func GetSpotifyAccessToken(spotifyDcCookie string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Cookie", fmt.Sprintf("sp_dc=%s", spotifyDcCookie)).
		Get("https://open.spotify.com/get_access_token")

	if err != nil {
		return "", fmt.Errorf("failed to fetch Spotify access token: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("Spotify API returned non-200 status: %s", resp.Status())
	}

	var tokenResponse SpotifyAccessTokenResponse
	if err := json.Unmarshal(resp.Body(), &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to parse Spotify access token response: %v", err)
	}

	return tokenResponse.AccessToken, nil
}

// DisplayCurrentlyPlayingTrack fetches and displays the user's currently playing track
func DisplayCurrentlyPlayingTrack(accessToken string) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get("https://api.spotify.com/v1/me/player/currently-playing")

	if err != nil {
		log.Fatalf("Error fetching currently playing track: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Spotify API returned non-200 status: %s", resp.Status())
	}

	var currentlyPlaying SpotifyCurrentlyPlayingResponse
	if err := json.Unmarshal(resp.Body(), &currentlyPlaying); err != nil {
		log.Fatalf("Error parsing currently playing track: %v", err)
	}

	if currentlyPlaying.Item.Name == "" {
		fmt.Println("Nothing is currently playing.")
		return
	}

	fmt.Printf("\nNow Playing: %s by %s\n", currentlyPlaying.Item.Name, currentlyPlaying.Item.Artists[0].Name)
}

// HandleCurrentlyPlaying handles the "currently playing" menu option
func HandleCurrentlyPlaying(spotifyDcCookie string) {
	accessToken, err := GetSpotifyAccessToken(spotifyDcCookie)
	if err != nil {
		log.Fatalf("Error fetching Spotify access token: %v", err)
	}
	DisplayCurrentlyPlayingTrack(accessToken)
}
