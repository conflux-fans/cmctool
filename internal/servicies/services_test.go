package servicies

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/stretchr/testify/assert"
)

func TestScheduleDataCollection(t *testing.T) {
	configs.Init()
	NewVolumeFetcher().FetchAndMail()

	cfg := configs.Get()
	NewPosRewardFetcher(cfg.PosValidatorsByScan, cfg.PosValidatorsByContract).FetchAndMail()
}

func TestUnionPayBalanceService(t *testing.T) {
	configs.Init()
	cfg := configs.Get()
	err := NewUnionPayBalanceService(cfg.UnionPay.AdminClientUrl, cfg.UnionPay.AdminPrivateKey).FetchAndMail()
	assert.NoError(t, err)
}
