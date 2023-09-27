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

	fmt.Printf("%q %q %q %q\n", cfg.PG.Conn, cfg.Server.Conn, cfg.MYSQL.Conn, cfg.Environment)

	fmt.Printf("%v %v %q\n", cfg.MaxOpenConns, cfg.MaxIdleConns, cfg.MaxIdleTime)
}
