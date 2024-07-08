package volume

import (
	"fmt"
	"testing"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestCollectVolumeAndMail(t *testing.T) {
	configs.Init()
	err := collectVolumeAndMail()
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveSheet(t *testing.T) {
	f := excelize.NewFile()
	f.NewSheet("MySheet")
	id, err := f.GetSheetIndex("Sheet1")
	assert.NoError(t, err)
	fmt.Printf("Sheet1 index: %v\n", id)

	err = f.DeleteSheet("Sheet1")
	assert.NoError(t, err)

	err = f.SaveAs("/Users/dayong/tmp/test.xlsx")
	assert.NoError(t, err)
}
