package sdk

import (
	"testing"

	"github.com/conflux-fans/cmctool/internal/sdk/types"
)

func TestMarketPairLatest(t *testing.T) {
	client := NewClient()
	resp, err := client.MarketPairLatest(types.MarketPairQueryParams{
		Slug:     "conflux",
		Category: "spot",
		Start:    1,
		Limit:    100,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.Data)
	t.Log(resp.Status)
}

func TestRemoveSheet(t *testing.T) {

}
