package servicies

import (
	"github.com/conflux-fans/go-scan-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PosRewardFetcher struct {
	posAddress []ethcommon.Hash
	scanClient *client.Client
}

func NewPosRewardFetcher(scanUrl string, posAddress []ethcommon.Hash) *PosRewardFetcher {
	scanClient := client.NewClient(scanUrl)
	return &PosRewardFetcher{
		posAddress: posAddress,
		scanClient: scanClient,
	}
}

func (p *PosRewardFetcher) GetPosRewards() (map[common.Hash]decimal.Decimal, error) {
	logrus.Info("[PosRewardFetcher] get pos rewards")
	rewards := make(map[common.Hash]decimal.Decimal)
	for _, posAddress := range p.posAddress {
		posOverView, err := p.scanClient.GetPosAccountOverview(posAddress)
		if err != nil {
			return nil, err
		}
		rewardInWei, err := decimal.NewFromString(posOverView.TotalReward)
		if err != nil {
			return nil, err
		}

		rewardInEth := rewardInWei.Div(decimal.New(1, 18))
		rewards[posAddress] = rewardInEth
	}
	return rewards, nil
}
