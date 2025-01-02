package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gravelstone/gravel"
)

const (
	channelID = "@YOUR_CHANNEL"
)

type CryptoPrice struct {
	USDPrice string `json:"usd_price"`
	IRRPrice string `json:"irr_price"`
}

func main() {
	t := "TELEGRAM_TOKEN"
	client := gravel.NewGravel(t, true)

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			message, err := getFormattedCryptoPrices()
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
						message, err := getFormattedCryptoPrices()
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

func getFormattedCryptoPrices() (string, error) {
	cryptoPrices := map[string]string{
		"Tether":   "https://api.wallex.ir/v1/market?pair=USDT/IRR",
		"Bitcoin":  "https://api.wallex.ir/v1/market?pair=BTC/IRR",
		"Ethereum": "https://api.wallex.ir/v1/market?pair=ETH/IRR",
	}

	results := map[string]CryptoPrice{}

	for name, url := range cryptoPrices {
		price, err := fetchCryptoPrice(url)
		if err != nil {
			return "", fmt.Errorf("error fetching %s price: %v", name, err)
		}
		results[name] = price
	}

	message := fmt.Sprintf(
		`🔴 تتر
%s تومان

🟢 بیت کوین:
%s دلار
%s تومان

🟢 اتریوم:
%s دلار
%s تومان`,
		results["Tether"].IRRPrice,
		results["Bitcoin"].USDPrice, results["Bitcoin"].IRRPrice,
		results["Ethereum"].USDPrice, results["Ethereum"].IRRPrice,
	)

	return message, nil
}

func fetchCryptoPrice(url string) (CryptoPrice, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CryptoPrice{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CryptoPrice{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CryptoPrice{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Stats struct {
			USDPrice string `json:"latest_price_usdt"`
			IRRPrice string `json:"latest_price_irr"`
		} `json:"stats"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return CryptoPrice{}, err
	}

	return CryptoPrice{
		USDPrice: result.Stats.USDPrice,
		IRRPrice: result.Stats.IRRPrice,
	}, nil
}
