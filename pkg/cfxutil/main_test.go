package cfxutil

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/conflux-fans/cmctool/pkg/timeutils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSearchEpochOnTime(t *testing.T) {
	client, err := sdk.NewClient("https://cmain-rpc.nftrainbow.cn/gcnYNYGANW")
	assert.NoError(t, err)
	monthStart := timeutils.GetMonthStart(timeutils.GetChinaLocation())
	logrus.Info("month start: ", monthStart)
	epoch, time, err := SearchEpochOnTime(client, time.Now())
	assert.NoError(t, err)
	logrus.WithField("epoch", epoch).WithField("time", time).Info("get epoch on time")
	assert.NotNil(t, epoch)
	assert.NotNil(t, time)
	fmt.Println(epoch, time)
}
