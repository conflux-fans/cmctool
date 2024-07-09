package types

type Response[T any] struct {
	Data   T      `json:"data"`
	Status Status `json:"status"`
}

type MarketPairsResp = Response[*MarketPairsData]
