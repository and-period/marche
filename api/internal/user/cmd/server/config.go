package cmd

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port                int64  `envconfig:"PORT" default:"8080"`
	MetricsPort         int64  `envconfig:"METRICS_PORT" default:"9090"`
	ShutdownDelaySec    int64  `envconfig:"SHUTDOWN_DELAY_SEC" default:"20"`
	LogPath             string `envconfig:"LOG_PATH" default:""`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"info"`
	DBSocket            string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost              string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort              string `envconfig:"DB_PORT" default:"3306"`
	DBUsername          string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword          string `envconfig:"DB_PASSWORD" default:""`
	DBDatabase          string `envconfig:"DB_DATABASE" default:"users"`
	DBTimeZone          string `envconfig:"DB_TIMEZONE" default:""`
	DBEnabledTLS        bool   `envconfig:"DB_ENABLED_TLS" default:"false"`
	AWSRegion           string `envconfig:"AWS_REGION" default:"ap-northeast-1"`
	AWSAccessKey        string `envconfig:"AWS_ACCESS_KEY" default:""`
	AWSSecretKey        string `envconfig:"AWS_SECRET_KEY" default:""`
	CognitoUserPoolID   string `envconfig:"COGNITO_USER_POOL_ID" default:""`
	CognitoClientID     string `envconfig:"COGNITO_CLIENT_ID" default:""`
	CognitoClientSecret string `envconfig:"COGNITO_CLIENT_SECRET" default:""`
}

func newConfig() (*config, error) {
	conf := &config{}
	if err := envconfig.Process("", conf); err != nil {
		return conf, fmt.Errorf("config: failed to new config: %w", err)
	}
	return conf, nil
}
