package main

import (
	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/internal/servicies/volume"
	"github.com/conflux-fans/cmctool/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.Init()
	logger.Init()
	logrus.Println("config initialization done")
	volume.Loop()
	select {}
}
