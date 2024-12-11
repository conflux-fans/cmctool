package servicies

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/stretchr/testify/assert"
)

func TestScheduleDataCollection(t *testing.T) {
	configs.Init()
	err := scheduleDataCollection()
	assert.NoError(t, err)
}
