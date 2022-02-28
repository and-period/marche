package cmd

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port             int64  `envconfig:"PORT" default:"8080"`
	MetricsPort      int64  `envconfig:"METRICS_PORT" default:"9090"`
	ShutdownDelaySec int64  `envconfig:"SHUTDOWN_DELAY_SEC" default:"20"`
	LogPath          string `envconfig:"LOG_PATH" default:""`
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	GRPCInsecure     bool   `envconfig:"GRPC_INSECURE" default:"true"`
	UserServiceURL   string `envconfig:"USER_SERVICE_URL" default:""`
}

func newConfig() (*config, error) {
	conf := &config{}
	if err := envconfig.Process("", conf); err != nil {
		return conf, fmt.Errorf("config: failed to new config: %w", err)
	}
	return conf, nil
}
