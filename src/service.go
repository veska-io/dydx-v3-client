package client

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/veska-io/dydx-v3-client/src/candles"
	"github.com/veska-io/dydx-v3-client/src/funding"
	"github.com/veska-io/dydx-v3-client/src/markets"
	"github.com/veska-io/dydx-v3-client/src/trades"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) GetCandles(
	ticker, resolution string, limit uint8, dateFrom, dateTo time.Time,
) (*candles.CandlesResponse, error) {
	p := struct {
		Ticker     string    `validate:"required"`
		Resolution string    `validate:"required,oneof=1MIN 5MINS 15MINS 30MINS 1HOUR 4HOURS 1DAY"`
		Limit      uint8     `validate:"min=1,max=100"`
		FromISO    time.Time `validate:"ltcsfield=ToISO"`
		ToISO      time.Time `validate:"gtcsfield=FromISO"`
	}{
		Ticker:     ticker,
		Resolution: resolution,
		Limit:      limit,
		FromISO:    dateFrom,
		ToISO:      dateTo,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := candles.APIRequest(ticker, resolution,
		dateFrom.Format("2006-01-02T15:04:05"), dateTo.Format("2006-01-02T15:04:05"), limit)

	if err != nil {
		return nil, fmt.Errorf("failed to get candles: %w", err)
	}

	return resp, nil
}

func (c *Client) GetMarkets(params ...string) (*markets.PerpetualMarketsResponse, error) {
	resp, err := markets.APIRequest(params...)
	if err != nil {
		return nil, fmt.Errorf("failed to get markets: %w", err)
	}

	return resp, nil
}

func (c *Client) GetHistoricalFunding(
	market string, limit uint8, effectiveBeforeOrAt time.Time,
) (*funding.HistoricalFundingResponse, error) {
	p := struct {
		Market string `validate:"required"`
		Limit  uint8  `validate:"min=1,max=100"`
		Before time.Time
	}{
		Market: market,
		Limit:  limit,
		Before: effectiveBeforeOrAt,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := funding.APIRequest(market, limit,
		effectiveBeforeOrAt.Format("2006-01-02T15:04:05"))

	if err != nil {
		return nil, fmt.Errorf("failed to get funding: %w", err)
	}

	return resp, nil
}

func (c *Client) GetTrades(
	market string,
	limit uint8,
	startingBeforeOrAt time.Time,
) (*trades.TradesResponse, error) {
	p := struct {
		Market string `validate:"required"`
		Limit  uint8  `validate:"min=1,max=100"`
		Before time.Time
	}{
		Market: market,
		Limit:  limit,
		Before: startingBeforeOrAt,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.Struct(p)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	resp, err := trades.APIRequest(market, limit, startingBeforeOrAt.Format("2006-01-02T15:04:05"))

	if err != nil {
		return nil, fmt.Errorf("failed to get funding: %w", err)
	}

	return resp, nil
}
