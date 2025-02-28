package types

type UserLoginRequest struct {
	RegisterAddress string `form:"registerAddress" json:"registerAddress" binding:"ethAddress,required"`
	IsBimWallet     bool   `form:"isBimWallet" json:"isBimWallet"`
	Signature       string `form:"signature" json:"signature" binding:"required"`
}

type UserLoginResponse struct {
	Expire   string `json:"expire"`
	Token    string `json:"token"`
	UserType int    `json:"userType"`
}

type EIP712NonceResp struct {
	Nonce string `json:"nonce"`
}
