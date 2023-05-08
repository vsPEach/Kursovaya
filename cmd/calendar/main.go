package main

import (
	"flag"
	"log"

	"github.com/vsPEach/Kursovaya/config"
	"github.com/vsPEach/Kursovaya/internal/app"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
