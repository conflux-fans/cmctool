package servicies

import (
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/internal/servicies/volume"
	"github.com/conflux-fans/cmctool/pkg/common"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func Start() {
	c := cron.New()

	var entryId cron.EntryID
	entryId, _ = c.AddFunc(configs.Get().Service.Cron.Schedule, func() {
		common.Retry(3, time.Minute, scheduleDataCollection)
		nextTime := c.Entry(cron.EntryID(entryId)).Next
		logrus.Infof("[Services] Next task execution time: %v\n", nextTime)
	})

	c.Start()
}

func scheduleDataCollection() error {
	cfg := configs.Get()

	volumnFetcher := volume.NewVolumeFetcher(cfg.Server.Cmc)
	logrus.Info("[Services] === Start Collect Volumes ===")
	allTokenMarketPairs, err := volumnFetcher.CollectSpotAndPerpVolumes()
	if err != nil {
		return err
	}

	// pos rewards
	posRewardFetcher := NewPosRewardFetcher(cfg.Server.Scan, cfg.PosAddress)
	posRewards, err := posRewardFetcher.GetPosRewards()
	if err != nil {
		return err
	}

	reporter := NewReporter(allTokenMarketPairs, posRewards)
	excelPath, err := reporter.WriteToExcel()
	if err != nil {
		return err
	}

	return reporter.SendMail(excelPath)
}
