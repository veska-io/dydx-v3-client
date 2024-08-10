package candles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veska-io/dydx-v3-client/src/config"
)

/*
	ratelimit-remaining: 99
	ratelimit-reset: 1723115940868
*/

func APIRequest(market, resolution string, fromISO, toISO string, limit uint8) (*CandlesResponse, error) {
	var candlesResponse CandlesResponse

	url, err := generateUrl(market, resolution, fromISO, toISO, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate URL: %w", err)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get response from external API: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &candlesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(candlesResponse.Errors) > 0 {
		return nil, fmt.Errorf("external API returned errors: %s",
			candlesResponse.Errors[0].Msg)
	}

	return &candlesResponse, nil
}

func generateUrl(market, resolution string, fromISO, toISO string, limit uint8) (*url.URL, error) {
	cfg := config.MustNew()
	baseURL := cfg.Url + cfg.CandlesPath + "/" + market

	params := url.Values{}
	params.Add("resolution", resolution)

	if limit > 0 {
		params.Add("limit", fmt.Sprint(limit))
	}

	if fromISO != "" {
		params.Add("fromISO", fromISO)
	}

	if toISO != "" {
		params.Add("toISO", toISO)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	u.RawQuery = params.Encode()
	return u, nil
}
