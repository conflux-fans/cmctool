package configs

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/conflux-fans/cmctool/internal/constants"
	"github.com/ethereum/go-ethereum/common"
)

type PosValidatorByContract struct {
	Name       string                    `yaml:"name"`
	PosAddress common.Hash               `yaml:"posAddress"`
	PowAddress cfxaddress.Address        `yaml:"powAddress"`
	QueryUser  cfxaddress.Address        `yaml:"queryUser"`
	Method     constants.PosRewardMethod `yaml:"method"`
}

type PosValidatorByScan struct {
	Name       string             `yaml:"name"`
	PosAddress common.Hash        `yaml:"posAddress"`
	PowAddress cfxaddress.Address `yaml:"powAddress"`
}

type Config struct {
	Mail struct {
		Sender struct {
			FromAlias string `yaml:"fromAlias"`
			Host      string `yaml:"host"`
			Port      int    `yaml:"port"`
			Username  string `yaml:"username"`
			Password  string `yaml:"password"`
		} `yaml:"sender"`
		Receivers struct {
			Volume    []string `yaml:"volume"`
			PosReward []string `yaml:"posReward"`
		} `yaml:"receivers"`
	} `yaml:"mail"`
	Server struct {
		Cmc           string `yaml:"cmc"`
		Scan          string `yaml:"scan"`
		CoreSpaceNode string `yaml:"coreSpaceNode"`
		EspaceNode    string `yaml:"espaceNode"`
	} `yaml:"server"`
	PosValidatorsByContract []*PosValidatorByContract `yaml:"posValidatorsByContract"`
	PosValidatorsByScan     []*PosValidatorByScan     `yaml:"posValidatorsByScan"`
	Service                 struct {
		Cron struct {
			Volume    string `yaml:"volume"`
			PosReward string `yaml:"posReward"`
		} `yaml:"cron"`
	} `yaml:"service"`
}
