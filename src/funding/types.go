package funding

import (
	types "github.com/veska-io/dydx-v3-client/src/types"
)

type HistoricalFunding struct {
	Market      string `json:"market"`
	Rate        string `json:"rate"`
	Price       string `json:"price"`
	EffectiveAt string `json:"effectiveAt"`
}

type HistoricalFundingResponse struct {
	HistoricalFunding []HistoricalFunding `json:"historicalFunding"`
	Errors            []types.DYDXError   `json:"errors"`
}
