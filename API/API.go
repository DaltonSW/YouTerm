package API

import (
	"errors"

	// youtube "google.golang.org/api/youtube/v3"
	"os"
)

func GetKeyFromEnv() (string, error) {
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		return api, nil
	}
	//Prompt user for API key

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}
