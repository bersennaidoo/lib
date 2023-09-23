package config

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"

	"github.com/bersennaidoo/lib/pkg/infrastructure/rpc"
	"github.com/bersennaidoo/lib/pkg/infrastructure/swaggerweb"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	MYSQLConfig struct {
		Conn string `required:"true"`
	}

	HTTPServer struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		Server          HTTPServer
		PG              PGConfig
		MYSQL           MYSQLConfig
		Rpc             rpc.RpcConfig
		Web             swaggerweb.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
