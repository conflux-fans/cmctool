package unionpaysdk

import (
	"encoding/json"

	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types/admintypes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
)

type AdminClient struct {
	inner *resty.Client
	auth  string
}

func NewAdminClient(baseUrl string) *AdminClient {
	client := resty.New()
	client.BaseURL = baseUrl
	return &AdminClient{
		inner: client,
	}
}

func (c *AdminClient) R() *resty.Request {
	return c.inner.R()
}

func (c *AdminClient) SetAuth(auth string) {
	// return c.R().SetHeader("Authorization", "Bearer "+auth)
	c.auth = auth
}

func (c *AdminClient) Login(req types.UserLoginRequest) (val *types.UserLoginResponse, err error) {
	resp, err := c.R().EnableTrace().SetBody(req).Post("/v1/admin/login")
	if err != nil {
		return nil, err
	}

	return parseResponse[types.UserLoginResponse](resp)
}

func (c *AdminClient) GetSignNonce(addr common.Address) (val *types.EIP712NonceResp, err error) {
	resp, err := c.R().EnableTrace().SetQueryParam("address", addr.String()).Get("/v1/admin/sign/nonce")
	if err != nil {
		return nil, err
	}

	return parseResponse[types.EIP712NonceResp](resp)
}

func (c *AdminClient) GetFinanceBalance() (val *admintypes.FinanceBalanceResponse, err error) {
	resp, err := c.R().EnableTrace().SetHeader("Authorization", "Bearer "+c.auth).Get("/v1/admin/finance/balance")
	if err != nil {
		return nil, err
	}

	// logrus.WithField("is error", resp.IsError()).WithField("resp", resp).Info("get finance balance")

	return parseResponse[admintypes.FinanceBalanceResponse](resp)
}

func parseResponse[T any](resp *resty.Response) (val *T, err error) {
	if resp.IsError() {
		var apiErr *types.ApiError
		if err = json.Unmarshal(resp.Body(), &apiErr); err != nil {
			return nil, err
		}
		return nil, apiErr
	}

	if err = json.Unmarshal(resp.Body(), &val); err != nil {
		return nil, err
	}

	return val, nil
}
