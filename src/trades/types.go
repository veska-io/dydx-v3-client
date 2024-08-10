package trades

import (
	types "github.com/veska-io/dydx-v3-client/src/types"
)

type Trade struct {
	Side        string `json:"side"`
	Size        string `json:"size"`
	Price       string `json:"price"`
	CreatedAt   string `json:"createdAt"`
	Liquidation bool   `json:"liquidation"`
}

type TradesResponse struct {
	Trades []Trade           `json:"trades"`
	Errors []types.DYDXError `json:"errors"`
	Market string            `json:"market"`
}
