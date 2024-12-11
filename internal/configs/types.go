package configs

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	Mail struct {
		Sender struct {
			FromAlias string `yaml:"fromAlias"`
			Host      string `yaml:"host"`
			Port      int    `yaml:"port"`
			Username  string `yaml:"username"`
			Password  string `yaml:"password"`
		} `yaml:"sender"`
		Receivers []string `yaml:"receivers"`
	} `yaml:"mail"`
	Server struct {
		Cmc  string `yaml:"cmc"`
		Scan string `yaml:"scan"`
	} `yaml:"server"`
	PosAddress []common.Hash `yaml:"posAddress"`

	Service struct {
		Cron struct {
			Schedule string `yaml:"schedule"`
		} `yaml:"cron"`
	} `yaml:"service"`
}
