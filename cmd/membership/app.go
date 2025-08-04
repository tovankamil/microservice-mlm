package main

import (
	cff "e-wallet-microservices/lib/config"
	"os"

	"github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2/log"
	zlog "github.com/rs/zerolog"
)

func main() {

	zlogT := zlog.New(os.Stdout).With().CallerWithSkipFrameCount(4).Timestamp().Logger()
	logger := zerolog.NewLogger(&zlogT)
	log.SetLogger(logger)
	log.Info("Starting payment service...")

	cfg, err := cff.InitConfig("config.membership")
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}
	// Initialize the service
	startService(cfg)
}
