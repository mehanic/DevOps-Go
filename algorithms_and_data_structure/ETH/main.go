package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoinGeckoResponse struct {
	Ethereum struct {
		USD float64 `json:"usd"`
	} `json:"ethereum"`
}

type PriceService struct {
	Client  *http.Client
	BaseURL string
}

func NewPriceService() *PriceService {
	return &PriceService{
		Client:  &http.Client{Timeout: 10 * time.Second},
		BaseURL: "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd",
	}
}

func (ps *PriceService) FetchETHPrice() (float64, error) {
	resp, err := ps.Client.Get(ps.BaseURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Ethereum.USD, nil
}

type PriceChecker struct {
	Threshold float64
}

func (pc *PriceChecker) IsAbove(price float64) bool {
	return price > pc.Threshold
}

type App struct {
	PriceService *PriceService
	Checker      *PriceChecker
}

func NewApp() *App {
	return &App{
		PriceService: NewPriceService(),
		Checker:      &PriceChecker{Threshold: 2600.0},
	}
}

func (a *App) Run() {
	price, _ := a.PriceService.FetchETHPrice()
	fmt.Printf("Current ETH price: $%.2f\n", price)
}

func main() {
	app := NewApp()
	app.Run()
}
