package admintypes

import (
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types/constants/enums"
	"github.com/shopspring/decimal"
)

type FinanceExportUrlResponse struct {
	Url string `json:"url"`
}

type FinanceExportUrlRequest struct {
	Year  int    `form:"year" binding:"required"`
	Month int    `form:"month" binding:"required"`
	Type  string `form:"type" binding:"required"`
}

type FinanceListRequest struct {
	StartAt    int64    `form:"startAt" binding:"required"`
	EndAt      int64    `form:"endAt" binding:"required"`
	Page       int      `form:"page" binding:"required"`
	Limit      int      `form:"limit" binding:"required"`
	Categories []string `form:"categories"`
}

type FinanceChannelListRequest struct {
	StartAt    int64    `form:"startAt" binding:"required"`
	EndAt      int64    `form:"endAt" binding:"required"`
	Page       int      `form:"page" binding:"required"`
	Limit      int      `form:"limit" binding:"required"`
	Categories []string `form:"categories"`
	Channels   []string `form:"channels"`
}

type FinanceSummaryRequest struct {
	StartAt int64 `form:"startAt" binding:"required"`
	EndAt   int64 `form:"endAt" binding:"required"`
}

type FinanceUserWeb3Act struct {
	RecordId         int64           `json:"recordId"`
	Category         string          `json:"category"`
	FromAddress      string          `json:"fromAddress"`
	ToAddress        string          `json:"toAddress"`
	Token            string          `json:"token"`
	Chain            enums.ChainType `json:"chain"`
	Amount           string          `json:"amount"`
	TxHash           string          `json:"txHash"`
	Timestamp        int64           `json:"timestamp"`
	CollectHotTxHash *string         `json:"collectHotTxHash"`
}

type FinanceExchange struct {
	RecordId         int64           `json:"recordId"`
	Category         string          `json:"category"`
	TxHash           string          `json:"txHash"`
	Chain            enums.ChainType `json:"chain"`
	Timestamp        int64           `json:"timestamp"`
	FromTokenOrFiat  string          `json:"fromTokenOrFiat"`
	FromAmount       string          `json:"fromAmount"`
	FromAddress      string          `json:"fromAddress"`
	ToTokenOrFiat    string          `json:"toTokenOrFiat"`
	ToAmount         string          `json:"toAmount"`
	ToAddress        string          `json:"toAddress"`
	FeeTokenOrFiat   string          `json:"feeTokenOrFiat"`
	FeeAmount        string          `json:"feeAmount"`
	Price            string          `json:"price"`
	ChannelFeeAmount decimal.Decimal `json:"channelFeeAmount"`
	Channel          string          `json:"channel"`
}

type FinanceOpenListItem struct {
	RecordId         int64           `json:"recordId"`
	TxHash           string          `json:"txHash"`
	Chain            enums.ChainType `json:"chain"`
	Timestamp        int64           `json:"timestamp"`
	FromTokenOrFiat  string          `json:"fromTokenOrFiat"`
	FromAmount       string          `json:"fromAmount"`
	FromAddress      string          `json:"fromAddress"`
	ToAddress        string          `json:"toAddress"`
	FeeTokenOrFiat   string          `json:"feeTokenOrFiat"`
	ChannelFeeAmount decimal.Decimal `json:"channelFeeAmount"`
	Channel          string          `json:"channel"`
}

type FinanceChannelListItem struct {
	RecordId         int64           `json:"recordId"`
	Category         string          `json:"category"`
	TxHash           string          `json:"txHash"`
	Chain            enums.ChainType `json:"chain"`
	Timestamp        int64           `json:"timestamp"`
	FeeTokenOrFiat   string          `json:"feeTokenOrFiat"`
	ChannelFeeAmount decimal.Decimal `json:"channelFeeAmount"`
	Channel          string          `json:"channel"`
	UserAddress      string          `json:"userAddress"`
}

type FinanceUserWeb3ActSummaryItem struct {
	Token            string          `json:"token"`
	DepositAmount    decimal.Decimal `json:"depositAmount"`
	WithdrawAmount   decimal.Decimal `json:"withdrawAmount"`
	CollectHotAmount decimal.Decimal `json:"collectHotAmount"`
}

type FinanceExchangeSummary struct {
	ExchangeToken []FinanceTokenOrFiatSummaryItem `json:"exchangeToken"`
	ExchangeFiat  []FinanceTokenOrFiatSummaryItem `json:"exchangeFiat"`
}

type FinanceTokenOrFiatSummaryItem struct {
	Token  string `json:"token"`
	Fiat   string `json:"fiat"`
	Amount string `json:"amount"`
}

func (f *FinanceTokenOrFiatSummaryItem) TokenOrFiat() string {
	if f.Fiat != "" {
		return f.Fiat
	}
	return f.Token
}

type FinanceOpenSummary struct {
	Sum      []FinanceOpenSummaryItem `json:"sum"`
	MakeCard []FinanceOpenSummaryItem `json:"makeCard"`
	Post     []FinanceOpenSummaryItem `json:"post"`
}

type FinanceOpenSummaryItem struct {
	Token  string `json:"token"`
	Amount string `json:"amount"`
}

type ChannelSummaryDetailItem struct {
	Channel string                          `json:"channel"`
	Amount  []FinanceTokenOrFiatSummaryItem `json:"amount"`
}

type ChannelSummary struct {
	Details []ChannelSummaryDetailItem      `json:"details"`
	Total   []FinanceTokenOrFiatSummaryItem `json:"total"`
}

type FinanceVaultsBalance struct {
	Hot  map[string]decimal.Decimal `json:"hot"`
	Cold map[string]decimal.Decimal `json:"cold"`
}

type FinanceBankBalance struct {
	TopupBalance string `json:"topupBalance"`
	CardBalance  string `json:"cardBalance"`
}

type FinanceBalanceResponse struct {
	Web3     FinanceVaultsBalance `json:"web3"`
	Bank     FinanceBankBalance   `json:"bank"`
	Internal FinanceBankBalance   `json:"internal"`
}

type InsertAdminBalanceRecordRequest struct {
	TopupBalance decimal.Decimal `json:"topup_balance"`
	CardBalance  int64           `json:"card_balance"`
	Reason       string          `json:"reason"`
}
