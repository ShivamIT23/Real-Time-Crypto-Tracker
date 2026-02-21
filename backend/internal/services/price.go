package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func FetchPrices(ids []string) map[string]float64 {
	prices := make(map[string]float64)
	if len(ids) == 0 {
		return prices
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + strings.Join(ids, ",") + "&vs_currencies=usd"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("HTTP error:", err)
		return prices
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API error: status %d", resp.StatusCode)
		return prices
	}

	var data map[string]map[string]json.Number
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("JSON error:", err)
		return prices
	}

	for _, id := range ids {
		if coinData, ok := data[id]; ok {
			if val, ok := coinData["usd"]; ok {
				price, err := val.Float64()
				if err == nil {
					prices[id] = price
				}
			}
		}
	}

	return prices
}

func FetchPrice(id string) float64 {
	p := FetchPrices([]string{id})
	return p[id]
}
