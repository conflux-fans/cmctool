package types

import (
	"strconv"
	"time"
)

type MarketPairQueryParams struct {
	Slug     string
	Category string
	Start    int
	Limit    int
}

func (r *MarketPairQueryParams) Build() map[string]string {
	params := map[string]string{
		"centerType":    "all",
		"sort":          "cmc_rank_advanced",
		"direction":     "desc",
		"spotUntracked": "true",
	}
	if r.Slug != "" {
		params["slug"] = r.Slug
	}
	if r.Category != "" {
		params["category"] = r.Category
	}
	if r.Start > 0 {
		params["start"] = strconv.Itoa(r.Start)
	}
	if r.Limit > 0 {
		params["limit"] = strconv.Itoa(r.Limit)
	}
	return params
}

type MarketPairsData struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Symbol         string           `json:"symbol"`
	NumMarketPairs int              `json:"numMarketPairs"`
	MarketPairs    []MarketPairLine `json:"marketPairs"`
	Sort           string           `json:"sort"`
	Direction      string           `json:"direction"`
	// SponsoredExchange []SponsoredExchange `json:"sponsoredExchange"`
}

type Quotes struct {
	ID               string  `json:"id"`
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume24h"`
	DepthPositiveTwo float64 `json:"depthPositiveTwo"`
	DepthNegativeTwo float64 `json:"depthNegativeTwo"`
}

type MarketPairLine struct {
	Rank                int       `json:"rank"`
	ExchangeID          int       `json:"exchangeId"`
	ExchangeName        string    `json:"exchangeName"`
	ExchangeSlug        string    `json:"exchangeSlug"`
	ExchangeNotice      string    `json:"exchangeNotice"`
	OutlierDetected     float64   `json:"outlierDetected"`
	PriceExcluded       float64   `json:"priceExcluded"`
	VolumeExcluded      float64   `json:"volumeExcluded"`
	MarketID            int       `json:"marketId"`
	MarketPair          string    `json:"marketPair"`
	Category            string    `json:"category"`
	MarketURL           string    `json:"marketUrl"`
	MarketScore         string    `json:"marketScore"`
	MarketReputation    float64   `json:"marketReputation"`
	BaseSymbol          string    `json:"baseSymbol"`
	BaseCurrencyID      int       `json:"baseCurrencyId"`
	QuoteSymbol         string    `json:"quoteSymbol"`
	QuoteCurrencyID     int       `json:"quoteCurrencyId"`
	Price               float64   `json:"price"`
	VolumeUsd           float64   `json:"volumeUsd"`
	EffectiveLiquidity  float64   `json:"effectiveLiquidity"`
	LastUpdated         time.Time `json:"lastUpdated"`
	Quote               float64   `json:"quote"`
	VolumeBase          float64   `json:"volumeBase"`
	VolumeQuote         float64   `json:"volumeQuote"`
	FeeType             string    `json:"feeType"`
	DepthUsdNegativeTwo float64   `json:"depthUsdNegativeTwo"`
	DepthUsdPositiveTwo float64   `json:"depthUsdPositiveTwo"`
	ReservesAvailable   int       `json:"reservesAvailable"`
	PorAuditStatus      int       `json:"porAuditStatus"`
	VolumePercent       float64   `json:"volumePercent"`
	IndexPrice          float64   `json:"indexPrice"`
	IsVerified          int       `json:"isVerified"`
	Quotes              []Quotes  `json:"quotes"`
	Type                string    `json:"type"`
	CenterType          string    `json:"centerType"`
	HideStarsMarket     int       `json:"hideStarsMarket"`
}
type SubInfos struct {
	URL              string   `json:"url"`
	ShowMsg          string   `json:"showMsg"`
	SupportCountries []string `json:"supportCountries"`
	ExcludeCountries []string `json:"excludeCountries"`
	CustomOptions    string   `json:"customOptions"`
}
type SponsoredExchange struct {
	ID            int        `json:"id,omitempty"`
	EventID       int        `json:"eventId"`
	Name          string     `json:"name,omitempty"`
	CustomName    string     `json:"customName"`
	SubInfos      []SubInfos `json:"subInfos"`
	CustomOptions string     `json:"customOptions"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      string    `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}
