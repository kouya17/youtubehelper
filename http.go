package youtubehelper

import (
	"os"
	"net/http"
	"log"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func newClient() *http.Client{
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	}
	return client
}

func newYoutubeService(client *http.Client) *youtube.Service {
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	return service
}