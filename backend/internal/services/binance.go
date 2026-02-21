package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/ws"
	"github.com/gorilla/websocket"
)

type BinanceTradeResponse struct {
	Stream string `json:"stream"`
	Data   struct {
		Symbol string `json:"s"`
		Price  string `json:"p"`
	} `json:"data"`
}

type BinanceStream struct {
	Hub         *ws.Hub
	SymbolMap   map[string]string // e.g., "btcusdt": "btc"
	prices      map[string]float64
	pricesMu    sync.RWMutex
	StreamNames []string
}

func NewBinanceStream(hub *ws.Hub, symbolMap map[string]string) *BinanceStream {
	var streams []string
	for bSymbol := range symbolMap {
		streams = append(streams, fmt.Sprintf("%s@trade", strings.ToLower(bSymbol)))
	}

	return &BinanceStream{
		Hub:         hub,
		SymbolMap:   symbolMap,
		prices:      make(map[string]float64),
		StreamNames: streams,
	}
}

func (bs *BinanceStream) Start() {
	url := fmt.Sprintf("wss://stream.binance.com:9443/stream?streams=%s", strings.Join(bs.StreamNames, "/"))

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Broadcast loop
	go func() {
		for range ticker.C {
			bs.pricesMu.RLock()
			msg := make(map[string]float64)
			for k, v := range bs.prices {
				msg[k] = v
			}
			bs.pricesMu.RUnlock()

			if len(msg) > 0 {
				bs.Hub.Broadcast <- msg
			}
		}
	}()

	for {
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Printf("dial error for %v: %v", bs.StreamNames, err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Connected to Binance WebSocket for %v", bs.StreamNames)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				c.Close()
				break
			}

			var resp BinanceTradeResponse
			if err := json.Unmarshal(message, &resp); err != nil {
				log.Println("unmarshal:", err)
				continue
			}

			price, err := strconv.ParseFloat(resp.Data.Price, 64)
			if err != nil {
				log.Println("parse price:", err)
				continue
			}

			bSymbol := strings.ToLower(resp.Data.Symbol)
			if key, ok := bs.SymbolMap[bSymbol]; ok {
				bs.pricesMu.Lock()
				bs.prices[key] = price
				bs.pricesMu.Unlock()
			}
		}

		time.Sleep(2 * time.Second)
	}
}
