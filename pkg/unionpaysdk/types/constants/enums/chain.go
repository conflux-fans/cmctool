package enums

import (
	"github.com/nft-rainbow/rainbow-goutils/utils/enumutils"
)

type ChainType int8

const (
	CHAIN_ESPACE_MAINNET ChainType = iota + 1
	CHAIN_ESPACE_TESTNET
)

// ChainInfo return block chain name and chain id
func (c ChainType) ChainInfo() (string, int64) {
	switch c {
	case CHAIN_ESPACE_MAINNET:
		return "CONFLUX-ESPACE", 1030
	case CHAIN_ESPACE_TESTNET:
		return "CONFLUX-ESPACE", 71
	default:
		return "", -1
	}
}

var chainEb enumutils.EnumBase[ChainType]

func init() {
	chainEb = enumutils.NewEnumBase("CHAIN", map[ChainType]string{
		CHAIN_ESPACE_MAINNET: "espace-mainnet",
		CHAIN_ESPACE_TESTNET: "espace-testnet",
	})
}

func (t ChainType) MarshalText() ([]byte, error) {
	return chainEb.MarshalText(t)
}

func (t *ChainType) UnmarshalText(data []byte) error {
	val, err := chainEb.UnmarshalText(data)
	if err != nil {
		return err
	}
	*t = val
	return nil
}

func (t ChainType) String() string {
	return chainEb.String(t)
}

func ParseChainType(s string) (ChainType, error) {
	return chainEb.Parse(s)
}

func MustParseChainType(s string) ChainType {
	c, err := chainEb.Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}
