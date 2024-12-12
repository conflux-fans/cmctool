package servicies

import (
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"

	"github.com/conflux-fans/cmctool/pkg/common"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func Start() {
	c := cron.New()

	volumeEntryId := runVolumeTask(c)
	posRewardEntryId := runPosRewardTask(c)

	defer func() {
		logNextRunTime(c, volumeEntryId)
		logNextRunTime(c, posRewardEntryId)
	}()

	c.Start()
}

func logNextRunTime(c *cron.Cron, entryId cron.EntryID) {
	nextTime := c.Entry(cron.EntryID(entryId)).Next
	logrus.Infof("[Services] Next task execution time: %v\n", nextTime)
}

func runVolumeTask(c *cron.Cron) cron.EntryID {
	entryId, err := c.AddFunc(configs.Get().Service.Cron.Volume, func() {
		common.Retry(3, time.Minute, NewVolumeFetcher().FetchAndMail)
	})
	if err != nil {
		panic(err)
	}
	return entryId
}

func runPosRewardTask(c *cron.Cron) cron.EntryID {
	cfg := configs.Get()
	entryId, err := c.AddFunc(configs.Get().Service.Cron.PosReward, func() {
		common.Retry(3, time.Minute, NewPosRewardFetcher(cfg.PosValidatorsByScan, cfg.PosValidatorsByContract).FetchAndMail)
	})
	if err != nil {
		panic(err)
	}
	return entryId
}

// func runScheduleTask() {
// 	cfg := configs.Get()

// 	var w sync.WaitGroup

// 	// w.Add(1)
// 	// go func() {
// 	// 	common.Retry(3, time.Minute, NewVolumeFetcher().FetchAndMail)
// 	// 	w.Done()
// 	// }()

// 	w.Add(1)
// 	go func() {
// 		common.Retry(3, time.Minute, NewPosRewardFetcher(cfg.PosValidatorsByScan, cfg.PosValidatorsByContract).FetchAndMail)
// 		w.Done()
// 	}()

// 	w.Wait()
// }

// func scheduleDataCollection() error {
// 	cfg := configs.Get()

// 	volumnFetcher := volume.NewVolumeFetcher(cfg.Server.Cmc)
// 	logrus.Info("[Services] === Start Collect Volumes ===")
// 	allTokenMarketPairs, err := volumnFetcher.CollectSpotAndPerpVolumes()
// 	if err != nil {
// 		return err
// 	}

// 	// pos rewards
// 	posRewardFetcher := NewPosRewardFetcher(cfg.PosValidatorsByScan)
// 	posRewards, err := posRewardFetcher.GetPosRewards()
// 	if err != nil {
// 		return err
// 	}

// 	reporter := NewReporter(allTokenMarketPairs, posRewards)
// 	excelPath, err := reporter.WriteAllToExcel()
// 	if err != nil {
// 		return err
// 	}

// 	return reporter.SendMail(excelPath)
// }
