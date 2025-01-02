package servicies

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/stretchr/testify/assert"
)

func TestPosReward(t *testing.T) {
	configs.Init()
	posRewardFetcher := NewPosRewardFetcher(configs.Get().PosValidatorsByScan, configs.Get().PosValidatorsByContract)
	err := posRewardFetcher.FetchAndMail()
	assert.NoError(t, err)
}
