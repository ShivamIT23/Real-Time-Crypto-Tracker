package services

import (
	"encoding/json"
	"net/http"
)

func FetchBTC() float64 {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	var data map[string]map[string]float64
	json.NewDecoder(resp.Body).Decode(&data)

	return data["bitcoin"]["usd"]
}

