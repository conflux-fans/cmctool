package lib

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
)

var (
	eip712DomainType = []apitypes.Type{
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
	}

	loginBodyType = []apitypes.Type{
		{
			Name: "nonce",
			Type: "string",
		},
	}
)

func GetEncodedForLoginSign(nonce string) ([]byte, error) {
	return eip712.EncodeForSigning(&apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": eip712DomainType,
			"Body":         loginBodyType,
		},
		PrimaryType: "Body",
		Domain: apitypes.TypedDataDomain{
			Name:    "BitUnion-Backend",
			Version: "1.0.0",
		},
		Message: apitypes.TypedDataMessage{
			"nonce": nonce,
		},
	})
}

func SignTypedData(data []byte, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	hash := crypto.Keccak256(data)
	sign, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] += 27 // 还没有研究为什么要加 27，只是eip712校验时要求 v值 是27 28
	return sign, nil
}
