package youtubehelper

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

func GetMovieTitles(channelID string) []string {
	service := newYoutubeService(newClient())
	call := service.Search.List("snippet").ChannelId(channelID)
	response, err := call.Do()
	var titleList []string
	//for i := 0; i < 5; i++ {
	for true {
		if err != nil {
			log.Printf("%v", err)
			return titleList
		}
		for _, item := range response.Items {
			title := item.Snippet.Title
			titleList = append(titleList, title)
		}
		if response.NextPageToken == "" {
			log.Println("end.")
			break
		}
		call = service.Search.List("snippet").ChannelId(channelID).PageToken(response.NextPageToken)
		response, err = call.Do()
	}
	for _, title := range titleList {
		log.Printf("title: %v", title)
	}
	return titleList
}

func searchExample() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	channelID := "UCSFCh5NL4qXrAy9u-u2lX3g"
	titleList := GetMovieTitles(channelID)
	fp, err := os.Create("title_" + channelID + ".txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()
	for _, title := range titleList {
		_, err := fp.WriteString(title + "\n")
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}