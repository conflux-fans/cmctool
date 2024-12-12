package servicies

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
)

func TestScheduleDataCollection(t *testing.T) {
	configs.Init()
	NewVolumeFetcher().FetchAndMail()

	cfg := configs.Get()
	NewPosRewardFetcher(cfg.PosValidatorsByScan, cfg.PosValidatorsByContract).FetchAndMail()
}
