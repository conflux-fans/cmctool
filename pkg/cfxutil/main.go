package cfxutil

import (
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// func SearchEpochOnTime(client *sdk.Client, target time.Time) (*big.Int, time.Time, error) {
// 	// 按epoch 时间算，每1 个快 0.5 往前推，直到查找到 |epoch_time - time|< 60s的 epoch
// 	epochNumber, epochTime, err := GetEpochTime(client, types.EpochLatestState)
// 	if err != nil {
// 		return nil, time.Time{}, errors.WithMessage(err, "failed to get epoch time")
// 	}
// 	timeDiff := epochTime.Sub(target).Seconds()
// 	targetEpochNumber := new(big.Int).Sub(epochNumber, big.NewInt(int64(timeDiff/0.5)))

// 	for i := 0; i < 100; i++ {
// 		if timeDiff < 60 && timeDiff > -60 {
// 			return targetEpochNumber, epochTime, nil
// 		}

// 		if timeDiff > 60 {
// 			targetEpochNumber = new(big.Int).Sub(epochNumber, (big.NewInt(1)))
// 		} else {
// 			targetEpochNumber = new(big.Int).Add(epochNumber, (big.NewInt(1)))
// 		}

// 		_, epochTime, err = GetEpochTime(client, types.NewEpochNumberBig(targetEpochNumber))
// 		if err != nil {
// 			return nil, time.Time{}, errors.WithMessage(err, "failed to get epoch time")
// 		}
// 		timeDiff = epochTime.Sub(target).Seconds()
// 	}
// 	return targetEpochNumber, epochTime, nil
// }

func SearchEpochOnTime(client *sdk.Client, target time.Time) (*big.Int, time.Time, error) {
	// 按epoch 时间算，每秒 1 个 epoch 往前推，直到查找到 |epoch_time - time|< 60s的 epoch
	maxLoopCnt := 20
	epochNumber := types.EpochLatestState
	for i := 0; i < maxLoopCnt; i++ {
		currentEpoch, epochTime, err := GetEpochTime(client, epochNumber)
		if err != nil {
			return nil, time.Time{}, errors.WithMessage(err, "failed to get epoch time")
		}
		timeDiff := epochTime.Sub(target).Seconds()
		if timeDiff/60 == 0 || i == maxLoopCnt-1 {
			return currentEpoch, epochTime, nil
		}

		targetEpoch := new(big.Int).Sub(currentEpoch, big.NewInt(int64(timeDiff)))
		logrus.WithField("current epochTime", epochTime).WithField("timeDiff", timeDiff).WithField("current epochNumber", currentEpoch).WithField("target epochNumber", targetEpoch).Info("search epoch on time")
		epochNumber = types.NewEpochNumberBig(targetEpoch)
	}
	return nil, time.Time{}, errors.New("failed to find epoch on time")

	// for i := 0; i < 100; i++ {
	// 	if timeDiff < 60 && timeDiff > -60 {
	// 		return targetEpochNumber, epochTime, nil
	// 	}

	// 	if timeDiff > 60 {
	// 		targetEpochNumber = new(big.Int).Sub(epochNumber, (big.NewInt(1)))
	// 	} else {
	// 		targetEpochNumber = new(big.Int).Add(epochNumber, (big.NewInt(1)))
	// 	}

	// 	_, epochTime, err = GetEpochTime(client, types.NewEpochNumberBig(targetEpochNumber))
	// 	if err != nil {
	// 		return nil, time.Time{}, errors.WithMessage(err, "failed to get epoch time")
	// 	}
	// 	timeDiff = epochTime.Sub(target).Seconds()
	// }
	// return targetEpochNumber, epochTime, nil
}

func GetEpochTime(client *sdk.Client, epochNumber *types.Epoch) (*big.Int, time.Time, error) {
	block, err := client.GetBlockByEpoch(epochNumber)
	if err != nil {
		return nil, time.Time{}, errors.WithMessage(err, "failed to get block by epoch")
	}
	return block.EpochNumber.ToInt(), time.Unix(block.Timestamp.ToInt().Int64(), 0), nil
}
