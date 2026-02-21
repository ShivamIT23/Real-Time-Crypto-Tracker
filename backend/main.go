package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/handlers"
	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/services"
	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/ws"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	hub := ws.NewHub()
	go hub.Run()

	// WebSocket route
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(hub, w, r)
	})

	// Start price broadcaster
	go func() {
		for {
			price := services.FetchBTC()

			hub.Broadcast <- map[string]float64{
				"btc": price,
			}

			time.Sleep(1 * time.Second)
		}
	}()

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}

