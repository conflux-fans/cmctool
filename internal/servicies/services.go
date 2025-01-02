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
	var entryId cron.EntryID
	_entryId, err := c.AddFunc(configs.Get().Service.Cron.Volume, func() {
		err := common.Retry(3, time.Minute, NewVolumeFetcher().FetchAndMail)
		if err != nil {
			logrus.Errorf("[Services] Run volume task failed: %v", err)
		}
		logNextRunTime(c, entryId)
	})
	if err != nil {
		panic(err)
	}
	entryId = _entryId
	return entryId
}

func runPosRewardTask(c *cron.Cron) cron.EntryID {
	var entryId cron.EntryID
	cfg := configs.Get()
	_entryId, err := c.AddFunc(cfg.Service.Cron.PosReward, func() {
		err := common.Retry(3, time.Minute, NewPosRewardFetcher(cfg.PosValidatorsByScan, cfg.PosValidatorsByContract).FetchAndMail)
		if err != nil {
			logrus.Errorf("[Services] Run pos reward task failed: %v", err)
		}
		logNextRunTime(c, entryId)
	})
	if err != nil {
		panic(err)
	}
	entryId = _entryId
	return entryId
}
