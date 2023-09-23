package main

import (
	"fmt"

	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
	"github.com/bersennaidoo/lib/pkg/infrastructure/logger"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Errorf("%w", err)
	}

	trace := logger.TRACE

	logcfg := logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    trace,
	}

	logsrv := logger.NewZeroLoggerSrv(logcfg)

	logsrv.Info().Msg("Hello from zeroLog")

	fmt.Printf("%s %s %s %s %s %s\n", cfg.Rpc.Host, cfg.Rpc.Port, cfg.Environment, cfg.PG.Conn,
		cfg.Web.Port, cfg.Web.Host)
}
