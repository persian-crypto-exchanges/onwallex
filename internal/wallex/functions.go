package wallex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CryptoPrice struct {
	USDPrice string `json:"usd_price"`
	IRRPrice string `json:"irr_price"`
}

func GetFormattedCryptoPrices() (string, error) {
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
