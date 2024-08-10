package markets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/veska-io/dydx-v3-client/src/config"
)

func APIRequest(params ...string) (*PerpetualMarketsResponse, error) {
	var marketsData PerpetualMarketsResponse

	url, err := generateUrl(params...)
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
		return nil, err
	}

	err = json.Unmarshal(body, &marketsData)
	if json.Unmarshal(body, &marketsData) != nil {
		return nil, err
	}

	return &marketsData, nil
}

func generateUrl(p ...string) (*url.URL, error) {
	cfg := config.MustNew()
	baseURL := cfg.Url + cfg.MarketsPath

	params := url.Values{}

	if len(p) > 0 && p[0] != "" {
		params.Add("market", p[0])
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	u.RawQuery = params.Encode()
	return u, nil
}
