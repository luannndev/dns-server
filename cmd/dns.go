package main

import (
	"dns-server/internal/config"
	"dns-server/internal/dnsserver"
	"dns-server/internal/logger"
	"flag"
)

func main() {
	initServices()
	dnsServer := dnsserver.New(config.Cfg)

	if err := dnsServer.Listen(); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to start DNS server")
	}

}

func initServices() {
	configPathFlag := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	logger.Init()
	config.Init(*configPathFlag)
	dnsserver.Init()
}
