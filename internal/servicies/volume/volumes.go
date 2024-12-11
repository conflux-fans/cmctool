package volume

import (
	"time"

	"github.com/conflux-fans/cmctool/pkg/cmcsdk"
	"github.com/conflux-fans/cmctool/pkg/cmcsdk/types"
	"github.com/sirupsen/logrus"
)

type VolumeFetcher struct {
	cmcClient *cmcsdk.Client
}

func NewVolumeFetcher(cmcBaseUrl string) *VolumeFetcher {
	return &VolumeFetcher{
		cmcClient: cmcsdk.NewClient(cmcBaseUrl),
	}
}

func (v *VolumeFetcher) CollectSpotAndPerpVolumes() (map[string]*types.MarketPairsResp, error) {
	spots, err := v.CollectSpotVloumes()
	if err != nil {
		return nil, err
	}
	logrus.Info("Get spots done")

	perps, err := v.CollectPerpVolumes()
	if err != nil {
		return nil, err
	}
	logrus.Info("Get perps done")

	allTokenMarketPairs := spots
	for k, v := range perps {
		allTokenMarketPairs[k] = v
	}

	return allTokenMarketPairs, nil
}

func (v *VolumeFetcher) CollectSpotVloumes() (spots map[string]*types.MarketPairsResp, err error) {
	// time.Sleep(time.Second)
	btcSpot, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"bitcoin", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get btc spot done")

	time.Sleep(time.Second)
	ethSpot, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"ethereum", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get eth spot done")

	time.Sleep(time.Second)
	cfxSpot, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"conflux-network", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get cfx spot done")

	spots = map[string]*types.MarketPairsResp{
		"BTC-Spot": btcSpot,
		"ETH-Spot": ethSpot,
		"CFX-Spot": cfxSpot,
	}
	return
}

func (v *VolumeFetcher) CollectPerpVolumes() (perps map[string]*types.MarketPairsResp, err error) {
	time.Sleep(time.Second)
	btcPerp, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"bitcoin", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get btc perp done")

	time.Sleep(time.Second)
	ethPerp, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"ethereum", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get eth perp done")

	time.Sleep(time.Second)
	cfxPerp, err := v.cmcClient.MarketPairLatest(types.MarketPairQueryParams{"conflux-network", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	logrus.Infoln("get cfx perp done")

	perps = map[string]*types.MarketPairsResp{
		"BTC-Perp": btcPerp,
		"ETH-Perp": ethPerp,
		"CFX-Perp": cfxPerp,
	}
	return
}
