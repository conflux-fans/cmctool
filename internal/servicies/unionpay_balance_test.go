package servicies

import "testing"

func TestGetMerchantCardBalance(t *testing.T) {
	_, err := NewUnionPayBalanceService("https://admin-api.bitunion.io", "0x27f89200bda88058e0234ea877fb2eb1b6e929e2e3f8dd7bfd7de9ed28ae960c").GetMerchantCardBalance()
	if err != nil {
		t.Error(err)
	}
}
