package smmgo

import "github.com/valyala/fasthttp"

type Session struct {
	Client  *fasthttp.Client
	BaseURL string
	ApiKey  string
}

type ServiceList []struct {
	Service  int    `json:"service"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Rate     string `json:"rate"`
	Min      string `json:"min"`
	Max      string `json:"max"`
	Refill   bool   `json:"refill"`
	Cancel   bool   `json:"cancel"`
}

type BalanceResponse struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
}

type OrderResponse struct {
	OrderId string `json:"order"`
}
