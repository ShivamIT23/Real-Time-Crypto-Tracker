package main

import (
	"log"
	"net/http"

	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/handlers"
	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/services"
	"github.com/ShivamIT23/Real-Time-Crypto-Tracker/backend/internal/ws"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// 1. Main Hub (BTC, ETH, SOL + Top 10)
	mainHub := ws.NewHub()
	go mainHub.Run()
	mainStream := services.NewBinanceStream(mainHub, map[string]string{
		"btcusdt":  "btc",
		"ethusdt":  "eth",
		"solusdt":  "sol",
		"bnbusdt":  "bnb",
		"xrpusdt":  "xrp",
		"adausdt":  "ada",
		"dogeusdt": "doge",
		"trxusdt":  "trx",
		"dotusdt":  "dot",
		"ltcusdt":  "ltc",
	})
	go mainStream.Start()

	// 2. ETH Tokens Hub (Ecosystem + DeFi)
	ethHub := ws.NewHub()
	go ethHub.Run()
	ethStream := services.NewBinanceStream(ethHub, map[string]string{
		"linkusdt": "link",
		"uniusdt":  "uni",
		"polusdt":  "pol",
		"pepeusdt": "pepe",
		"shibusdt": "shib",
		"arbusdt":  "arb",
		"opusdt":   "op",
		"ldousdt":  "ldo",
		"aaveusdt": "aave",
		"mkrusdt":  "mkr",
	})
	go ethStream.Start()

	// 3. SOL Tokens Hub (Ecosystem + Memes)
	solHub := ws.NewHub()
	go solHub.Run()
	solStream := services.NewBinanceStream(solHub, map[string]string{
		"jupusdt":    "jup",
		"pythusdt":   "pyth",
		"renderusdt": "render",
		"bonkusdt":   "bonk",
		"wifusdt":    "wif",
		"rayusdt":    "ray",
		"jtousdt":    "jto",
		"popcatusdt": "popcat",
		"tnsrusdt":   "tnsr",
	})
	go solStream.Start()

	// 4. Market Hub (L1s & Alternative Altcoins)
	marketHub := ws.NewHub()
	go marketHub.Run()
	marketStream := services.NewBinanceStream(marketHub, map[string]string{
		"avaxusdt": "avax",
		"nearusdt": "near",
		"atomusdt": "atom",
		"suiusdt":  "sui",
		"aptusdt":  "apt",
		"ftmusdt":  "ftm",
		"algousdt": "algo",
		"hbarusdt": "hbar",
		"icpusdt":  "icp",
		"vetusdt":  "vet",
	})
	go marketStream.Start()

	// 5. Trending Hub (AI & Emerging Tokens)
	trendingHub := ws.NewHub()
	go trendingHub.Run()
	trendingStream := services.NewBinanceStream(trendingHub, map[string]string{
		"fetusdt":   "fet",
		"taousdt":   "tao",
		"flokiusdt": "floki",
		"turbousdt": "turbo",
		"bomeusdt":  "bome",
		"mewusdt":   "mew",
		"brettusdt": "brett",
		"mogusdt":   "mog",
		"neirousdt": "neiro",
	})
	go trendingStream.Start()

	// WebSocket routes
	r.Get("/ws/main", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(mainHub, w, r)
	})

	r.Get("/ws/eth", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(ethHub, w, r)
	})

	r.Get("/ws/sol", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(solHub, w, r)
	})

	r.Get("/ws/market", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(marketHub, w, r)
	})

	r.Get("/ws/trending", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(trendingHub, w, r)
	})

	// Keep old /ws route for backward compatibility if needed, pointing to main
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(mainHub, w, r)
	})

	log.Println("Server running on :8080")
	log.Println("Endpoints: /ws/main, /ws/eth, /ws/sol, /ws/market, /ws/trending")
	http.ListenAndServe(":8080", r)
}
