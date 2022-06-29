package smmgo

import (
	"encoding/json"
	"net/url"

	"github.com/valyala/fasthttp"
)

func New(baseURL string, ApiKey string) *Session {
	session := &Session{
		Client:  &fasthttp.Client{},
		BaseURL: baseURL,
		ApiKey:  ApiKey,
	}

	return session
}

func (session *Session) Request(query map[string]string) (*fasthttp.Response, error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	base, _ := url.Parse(session.BaseURL)
	params := url.Values{}

	for key, value := range query {
		params.Add(key, value)
	}

	base.RawQuery = params.Encode()

	request.Header.SetMethod("GET")
	request.SetRequestURI(base.String())

	err := session.Client.Do(request, response)
	fasthttp.ReleaseRequest(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (session *Session) GetBalance() *BalanceResponse {
	payload := map[string]string{
		"key":    session.ApiKey,
		"action": "balance",
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	balance := &BalanceResponse{}
	json.Unmarshal(response.Body(), &balance)
	return balance
}

func (session *Session) GetServices() *ServiceList {
	payload := map[string]string{
		"key":    session.ApiKey,
		"action": "services",
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	serviceList := &ServiceList{}
	json.Unmarshal(response.Body(), &serviceList)
	return serviceList
}

func (session *Session) AddDefaultOrder(service, link, quantity string) *OrderResponse {
	payload := map[string]string{
		"key":      session.ApiKey,
		"action":   "add",
		"service":  service,
		"link":     link,
		"quantity": quantity,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}

func (session *Session) AddPackageOrder(service, link string) *OrderResponse {
	payload := map[string]string{
		"key":     session.ApiKey,
		"action":  "add",
		"service": service,
		"link":    link,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}

func (session *Session) AddCustomCommentsOrder(service, link, comments string) *OrderResponse {
	payload := map[string]string{
		"key":      session.ApiKey,
		"action":   "add",
		"service":  service,
		"link":     link,
		"comments": comments,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}

func (session *Session) AddCommentLikesOrder(service, link, username, quantity string) *OrderResponse {
	payload := map[string]string{
		"key":      session.ApiKey,
		"action":   "add",
		"service":  service,
		"link":     link,
		"quantity": quantity,
		"username": username,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}

func (session *Session) AddPollOrder(service, link, quantity, answer_number string) *OrderResponse {
	payload := map[string]string{
		"key":           session.ApiKey,
		"action":        "add",
		"service":       service,
		"link":          link,
		"quantity":      quantity,
		"answer_number": answer_number,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}

func (session *Session) AddSubscriptionsOrder(service, link, username, min, max, posts, old_posts, delay, expiry string) *OrderResponse {
	payload := map[string]string{
		"key":       session.ApiKey,
		"action":    "add",
		"service":   service,
		"link":      link,
		"username":  username,
		"min":       min,
		"max":       max,
		"posts":     posts,
		"old_posts": old_posts,
		"delay":     delay,
		"expiry":    expiry,
	}

	response, err := session.Request(payload)
	if err != nil {
		panic(err.Error())
	}

	order := &OrderResponse{}
	json.Unmarshal(response.Body(), &order)
	return order
}
