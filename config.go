package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LastFmUsername  string
	LastFmApiKey    string
	SpotifyDcCookie string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		LastFmUsername:  os.Getenv("LASTFM_USERNAME"),
		LastFmApiKey:    os.Getenv("LASTFM_API_KEY"),
		SpotifyDcCookie: os.Getenv("SPOTIFY_DC_COOKIE"),
	}, nil
}
