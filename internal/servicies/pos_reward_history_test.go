package servicies

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
)

func TestPosRewardHistoryFetchAndWriteExcel(t *testing.T) {
	configs.Init()
	posRewardHistoryFetcher := NewPosRewardHistoryFetcher(configs.Get().PosValidatorsByContract)
	posRewardHistoryFetcher.FetchAndWriteExcel()
}
