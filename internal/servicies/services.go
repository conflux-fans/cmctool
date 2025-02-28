package servicies

import (
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"

	"github.com/conflux-fans/cmctool/pkg/common"
	"github.com/conflux-fans/cmctool/pkg/timeutils"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func Start() {
	c := cron.New(cron.WithLocation(timeutils.GetChinaLocation()))

	volumeEntryId := runVolumeTask(c)
	posRewardEntryId := runPosRewardTask(c)
	unionPayEntryId := runUnionPayTask(c)
	defer func() {
		logNextRunTime(c, "volume", volumeEntryId)
		logNextRunTime(c, "posReward", posRewardEntryId)
		logNextRunTime(c, "unionPay", unionPayEntryId)
	}()

	c.Start()
}

func logNextRunTime(c *cron.Cron, name string, entryId cron.EntryID) {
	nextTime := c.Entry(cron.EntryID(entryId)).Next
	logrus.Infof("[Services] Next %s task execution time: %v\n", name, nextTime)
}

func runVolumeTask(c *cron.Cron) cron.EntryID {
	var entryId cron.EntryID
	_entryId, err := c.AddFunc(configs.Get().Service.Cron.Volume, func() {
		err := common.Retry(3, time.Minute, NewVolumeFetcher().FetchAndMail)
		if err != nil {
			logrus.Errorf("[Services] Run volume task failed: %v", err)
		}
		logNextRunTime(c, "volume", entryId)
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
		logNextRunTime(c, "posReward", entryId)
	})
	if err != nil {
		panic(err)
	}
	entryId = _entryId
	return entryId
}

func runUnionPayTask(c *cron.Cron) cron.EntryID {
	var entryId cron.EntryID
	cfg := configs.Get()
	_entryId, err := c.AddFunc(cfg.Service.Cron.UnionPay, func() {
		err := common.Retry(3, time.Minute, NewUnionPayBalanceService(cfg.UnionPay.AdminClientUrl, cfg.UnionPay.AdminPrivateKey).FetchAndMail)
		if err != nil {
			logrus.Errorf("[Services] Run union pay task failed: %v", err)
		}
		logNextRunTime(c, "unionPay", entryId)
	})
	if err != nil {
		panic(err)
	}
	entryId = _entryId
	return entryId
}
