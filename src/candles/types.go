package candles

import (
	types "github.com/veska-io/dydx-v3-client/src/types"
)

type Candle struct {
	StartedAt            string `json:"startedAt"`
	Market               string `json:"market"`
	Resolution           string `json:"resolution"`
	Low                  string `json:"low"`
	High                 string `json:"high"`
	Open                 string `json:"open"`
	Close                string `json:"close"`
	BaseTokenVolume      string `json:"baseTokenVolume"`
	UsdVolume            string `json:"usdVolume"`
	Trades               string `json:"trades"`
	StartingOpenInterest string `json:"startingOpenInterest"`
}

type CandlesResponse struct {
	Candles []Candle          `json:"candles"`
	Errors  []types.DYDXError `json:"errors"`
}
