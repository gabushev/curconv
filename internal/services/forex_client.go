package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const fetchPricesUrlPattern = "/crypto/fetch-prices?pairs=%s&api_key=%s"

type ForexClient struct {
	client           *http.Client
	forexApiBasename string
	apiKey           string
}

func NewAPIClient(forexApiBasename string, apiKey string) ForexClient {
	c := ForexClient{
		forexApiBasename: forexApiBasename,
		apiKey:           apiKey,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.client = &http.Client{Transport: tr}
	return c
}

func (c ForexClient) GetPrices(pairs []string) (map[string]float64, error) {
	url := fmt.Sprintf(c.forexApiBasename+fetchPricesUrlPattern, strings.Join(pairs, ","), c.apiKey)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rate struct {
		Prices map[string]float64 `json:"prices"`
	}
	err = json.NewDecoder(resp.Body).Decode(&rate)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the response, err: %v", err)
	}

	return rate.Prices, nil
}
