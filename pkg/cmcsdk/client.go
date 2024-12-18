package cmcsdk

import (
	"encoding/json"
	"fmt"

	"github.com/conflux-fans/cmctool/pkg/cmcsdk/types"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	inner *resty.Client
}

func NewClient(baseUrl string) *Client {
	client := resty.New()
	client.BaseURL = baseUrl
	return &Client{
		inner: client,
	}
}

func (c *Client) R() *resty.Request {
	return c.inner.R()
}

func (c *Client) MarketPairLatest(args types.MarketPairQueryParams) (val *types.MarketPairsResp, err error) {
	resp, err := c.R().EnableTrace().SetQueryParams(args.Build()).Get("/cryptocurrency/market-pairs/latest")
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp.Body(), &val); err != nil {
		return nil, err
	}

	if val.Status.ErrorCode != "0" {
		return nil, fmt.Errorf("error_code: %s, error_message: %s", val.Status.ErrorCode, val.Status.ErrorMessage)
	}

	return
}
