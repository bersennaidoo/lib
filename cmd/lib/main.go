package main

import (
	"fmt"

	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Errorf("%w", err)
	}

	fmt.Printf("%s %s %s %s %s %s\n", cfg.Rpc.Host, cfg.Rpc.Port, cfg.Environment, cfg.PG.Conn,
		cfg.Web.Port, cfg.Web.Host)
}
