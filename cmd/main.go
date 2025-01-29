package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gravelstone/gravel"
	"github.com/shayan0v0n/onwallex/internal/common"
	"github.com/shayan0v0n/onwallex/internal/wallex"
)

const (
	configFile = "config/wallex.conf"
)

func main() {
	var channelID, t string
	var timeInterval int
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	channelID = common.Getenv("channelID", scanner)
	t = common.Getenv("TELEGRAM_TOKEN", scanner)
	timeInterval, err = strconv.Atoi(common.Getenv("timeInterval", scanner))
	if err != nil {
		log.Fatal(err)
	}
	client := gravel.NewGravel(t, true)

	ticker := time.NewTicker(time.Duration(timeInterval) * time.Minute)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			message, err := wallex.GetFormattedCryptoPrices()
			if err != nil {
				log.Printf("Error fetching crypto prices: %v", err)
				continue
			}

			err = client.SendMessageToChannel(channelID, message)
			if err != nil {
				log.Printf("Error sending hourly crypto prices: %v", err)
			}
		}
	}()
	for {
		updates, err := client.GetUpdates()
		if err != nil {
			log.Printf("Error fetching updates: %v", err)
			continue
		}

		for _, update := range updates {
			if update.Message.IsCommand() {
				if update.Message != nil {
					if update.Message.Text == "/update" {
						message, err := wallex.GetFormattedCryptoPrices()
						if err != nil {
							log.Printf("Error fetching crypto prices: %v", err)
							sendErr := client.SendMessage(update.Message.Chat.ID, "Error fetching crypto prices.")
							if sendErr != nil {
								log.Printf("Error sending error message: %v", sendErr)
							}
						} else {
							err = client.SendMessage(update.Message.Chat.ID, message)
							if err != nil {
								log.Printf("Error while sending message: %v", err)
							}
						}
					}
				}
			}
		}
	}
}
