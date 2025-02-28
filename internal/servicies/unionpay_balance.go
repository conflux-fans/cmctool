package servicies

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/lib"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types/admintypes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/sirupsen/logrus"
)

type UnionPayBalanceService struct {
	client          *unionpaysdk.AdminClient
	adminPrivateKey *ecdsa.PrivateKey
	adminAddress    common.Address
}

func NewUnionPayBalanceService(adminClientUrl string, adminPrivateKey string) *UnionPayBalanceService {
	key, err := privatekeyhelper.NewFromKeyString(adminPrivateKey)
	if err != nil {
		panic(err)
	}
	return &UnionPayBalanceService{
		client:          unionpaysdk.NewAdminClient(adminClientUrl),
		adminPrivateKey: key,
		adminAddress:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func (u *UnionPayBalanceService) Login() (string, error) {
	nonceResp, err := u.client.GetSignNonce(u.adminAddress)
	if err != nil {
		return "", err
	}
	eip712EncodedForLogin, err := lib.GetEncodedForLoginSign(nonceResp.Nonce)
	if err != nil {
		return "", err
	}
	sign, err := lib.SignTypedData(eip712EncodedForLogin, u.adminPrivateKey)
	if err != nil {
		return "", err
	}

	loginReq := types.UserLoginRequest{
		RegisterAddress: u.adminAddress.String(),
		Signature:       hex.EncodeToString(sign),
	}
	loginResp, err := u.client.Login(loginReq)
	if err != nil {
		return "", err
	}

	u.client.SetAuth(loginResp.Token)
	return loginResp.Token, nil
}

// 1. login
// 2. request api
/*
curl 'https://admin-api.bitunion.io/v1/admin/finance/balance' \
  -H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3Mjk2ODYsImlkIjoxMzUsIm9yaWdfaWF0IjoxNzQwNzI2MDg2LCJ1c2VyX3R5cGUiOjF9.gmNf7mleTgHe2xmRZUaQlq8PbAyoe_BbYgm1moqeCkU' \
*/
func (u *UnionPayBalanceService) GetMerchantCardBalance() (*admintypes.FinanceBalanceResponse, error) {
	_, err := u.Login()
	if err != nil {
		return nil, err
	}

	balanceResp, err := u.client.GetFinanceBalance()
	if err != nil {
		return nil, err
	}

	return balanceResp, nil
}

func (u *UnionPayBalanceService) FetchAndMail() error {
	balanceResp, err := u.GetMerchantCardBalance()
	if err != nil {
		return err
	}
	reporter := NewReporter(nil, nil, nil, balanceResp)
	excelPath, err := reporter.WriteToExcel(&WriteEnables{
		IncludeUnionPay: true,
	})
	if err != nil {
		logrus.Errorf("[UnionPayBalanceService] Write to excel failed: %v", err)
		return err
	}
	logrus.Infof("[UnionPayBalanceService] Write to excel success: %s", excelPath)

	err = reporter.SendMail(excelPath, configs.Get().Mail.Receivers.UnionPay)
	if err != nil {
		logrus.Errorf("[UnionPayBalanceService] Send mail failed: %v", err)
		return err
	}

	logrus.Info("[UnionPayBalanceService] === Send mail Done ===")
	return nil
}
