package markets

type Market struct {
	Market                           string `json:"market"`
	Status                           string `json:"status"`
	BaseAsset                        string `json:"baseAsset"`
	QuoteAsset                       string `json:"quoteAsset"`
	StepSize                         string `json:"stepSize"`
	TickSize                         string `json:"tickSize"`
	IndexPrice                       string `json:"indexPrice"`
	OraclePrice                      string `json:"oraclePrice"`
	PriceChange24H                   string `json:"priceChange24H"`
	NextFundingRate                  string `json:"nextFundingRate"`
	NextFundingAt                    string `json:"nextFundingAt"`
	MinOrderSize                     string `json:"minOrderSize"`
	Type                             string `json:"type"`
	InitialMarginFraction            string `json:"initialMarginFraction"`
	MaintenanceMarginFraction        string `json:"maintenanceMarginFraction"`
	TransferMarginFraction           string `json:"transferMarginFraction"`
	Volume24H                        string `json:"volume24H"`
	Trades24H                        string `json:"trades24H"`
	OpenInterest                     string `json:"openInterest"`
	IncrementalInitialMarginFraction string `json:"incrementalInitialMarginFraction"`
	IncrementalPositionSize          string `json:"incrementalPositionSize"`
	MaxPositionSize                  string `json:"maxPositionSize"`
	BaselinePositionSize             string `json:"baselinePositionSize"`
	AssetResolution                  string `json:"assetResolution"`
	SyntheticAssetId                 string `json:"syntheticAssetId"`
}

type PerpetualMarketsResponse struct {
	Markets map[string]Market `json:"markets"`
}
