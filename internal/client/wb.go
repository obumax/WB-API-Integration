package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"wb-api-integration/pkg/logger"
)

type WBClient struct {
	standardToken string
	advToken      string
	statToken     string
	client        *http.Client
}

func NewWBClient(standardToken, advToken, statToken string) *WBClient {
	return &WBClient{
		standardToken: standardToken,
		advToken:      advToken,
		statToken:     statToken,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *WBClient) makeRequest(url, token string, method string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	logger.Info(fmt.Sprintf("Making %s request to %s", method, url))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func (c *WBClient) GetCards() ([]byte, error) {
	url := "https://suppliers-api.wildberries.ru/content/v1/cards/cursor/list"
	body := map[string]interface{}{
		"sort": map[string]string{
			"cursor": "",
			"limit":  "100",
		},
	}
	return c.makeRequest(url, c.standardToken, "POST", body)
}

// Aliases for compatibility with handler
type WbClient = WBClient

func NewWbClient(standardToken, advToken, statToken string) *WbClient {
	return NewWBClient(standardToken, advToken, statToken)
}
func (c *WBClient) GetOrders() ([]byte, error) {
	url := "https://statistics-api.wildberries.ru/api/v1/supplier/orders"
	dateFrom := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	fullURL := fmt.Sprintf("%s?dateFrom=%s", url, dateFrom)
	return c.makeRequest(fullURL, c.statToken, "GET", nil)
}

func (c *WBClient) UpdatePrices(prices interface{}) ([]byte, error) {
	url := "https://suppliers-api.wildberries.ru/public/api/v1/prices"
	return c.makeRequest(url, c.standardToken, "POST", prices)
}

func (c *WBClient) UpdateStocks(stocks interface{}) ([]byte, error) {
	url := "https://suppliers-api.wildberries.ru/api/v3/stocks"
	return c.makeRequest(url, c.standardToken, "POST", stocks)
}

func (c *WBClient) GetProducts() ([]byte, error) {
	return c.GetCards()
}

func (c *WBClient) GetAnalytics() ([]byte, error) {
	url := "https://statistics-api.wildberries.ru/api/v1/supplier/sales"
	dateFrom := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	fullURL := fmt.Sprintf("%s?dateFrom=%s", url, dateFrom)
	return c.makeRequest(fullURL, c.statToken, "GET", nil)
}

func (c *WBClient) GetSales() ([]byte, error) {
	url := "https://statistics-api.wildberries.ru/api/v1/supplier/sales"
	dateFrom := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	fullURL := fmt.Sprintf("%s?dateFrom=%s", url, dateFrom)
	return c.makeRequest(fullURL, c.statToken, "GET", nil)
}
