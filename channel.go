package youtubehelper

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func getSubscriberCount(channelID string) uint64 {
	service := newYoutubeService(newClient())
	call := service.Channels.List("statistics").Id(channelID).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	item := response.Items[0]
	return item.Statistics.SubscriberCount
}

func mainExample() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	channelID := "UC4YaOt1yT-ZeyB0OmxHgolA"
	subscriberCount := getSubscriberCount(channelID)
	fmt.Printf("channel id: %v\n登録者数: %v\n", channelID, subscriberCount)
}