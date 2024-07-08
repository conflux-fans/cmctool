package main

import (
	"log"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/internal/servicies/volume"
)

func main() {
	configs.Init()
	log.Println("config initialization done")
	volume.Loop()
	select {}
}
