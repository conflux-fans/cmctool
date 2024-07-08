package main

import (
	"log"

	"github.com/conflux-fans/cmctool/internal/configs"
)

func main() {
	configs.Init()
	log.Println("config initialization done")
}
