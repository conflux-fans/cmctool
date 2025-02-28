package unionpaysdk

import (
	"encoding/hex"
	"testing"

	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/lib"
	"github.com/conflux-fans/cmctool/pkg/unionpaysdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/sirupsen/logrus"
)

func TestGetFinanceBalance(t *testing.T) {
	client := NewAdminClient("https://admin-api.bitunion.io")
	nonceResp, err := client.GetSignNonce(common.HexToAddress("0x70355A9006ACD9359836CbA51EC8ace67186a177"))
	if err != nil {
		t.Fatal(err)
	}

	logrus.WithField("nonce", nonceResp.Nonce).Info("get nonce")

	data, err := lib.GetEncodedForLoginSign(nonceResp.Nonce)
	if err != nil {
		t.Fatal(err)
	}

	logrus.WithField("data", hex.EncodeToString(data)).Info("get encoded for login sign")
	hash := crypto.Keccak256(data)

	logrus.WithField("hash", hex.EncodeToString(hash)).Info("get hash")

	key, err := privatekeyhelper.NewFromKeyString("0x27f89200bda88058e0234ea877fb2eb1b6e929e2e3f8dd7bfd7de9ed28ae960c")
	if err != nil {
		t.Fatal(err)
	}

	sign, err := crypto.Sign(hash, key)
	if err != nil {
		t.Fatal(err)
	}
	sign[64] += 27 // 还没有研究为什么要加 27，只是eip712校验时要求 v值 是27 28

	logrus.WithField("sign", hex.EncodeToString(sign)).Info("get sign")
	resp, err := client.Login(types.UserLoginRequest{
		RegisterAddress: "0x70355A9006ACD9359836CbA51EC8ace67186a177",
		IsBimWallet:     false,
		Signature:       hex.EncodeToString(sign),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)

	client.SetAuth(resp.Token)

	balance, err := client.GetFinanceBalance()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(balance)
}
