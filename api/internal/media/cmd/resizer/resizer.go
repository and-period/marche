package resizer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/and-period/furumaru/api/internal/media/resizer"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type app struct {
	*cobra.Command
	logger           *zap.Logger
	waitGroup        *sync.WaitGroup
	resizer          resizer.Resizer
	AppName          string `envconfig:"APP_NAME" default:"media-resizer"`
	Environment      string `envconfig:"ENV" default:"none"`
	RunMethod        string `envconfig:"RUN_METHOD" default:"lambda"`
	LogPath          string `envconfig:"LOG_PATH" default:""`
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	DBSocket         string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost           string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort           string `envconfig:"DB_PORT" default:"3306"`
	DBUsername       string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword       string `envconfig:"DB_PASSWORD" default:""`
	DBTimeZone       string `envconfig:"DB_TIMEZONE" default:"Asia/Tokyo"`
	DBEnabledTLS     bool   `envconfig:"DB_ENABLED_TLS" default:"false"`
	DBSecretName     string `envconfig:"DB_SECRET_NAME" default:""`
	SentryDsn        string `envconfig:"SENTRY_DSN" default:""`
	SentrySecretName string `envconfig:"SENTRY_SECRET_NAME" default:""`
	AWSRegion        string `envconfig:"AWS_REGION" default:"ap-northeast-1"`
	S3Bucket         string `envconfig:"S3_BUCKET" default:""`
}

//nolint:revive
func NewApp() *app {
	cmd := &cobra.Command{
		Use:   "resizer",
		Short: "media resizer",
	}
	app := &app{Command: cmd}
	app.RunE = func(c *cobra.Command, args []string) error {
		return app.run()
	}
	return app
}

func (a *app) run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 環境変数の読み込み
	if err := envconfig.Process("", a); err != nil {
		return fmt.Errorf("resizer: failed to load environment: %w", err)
	}

	// 依存関係の解決
	if err := a.inject(ctx); err != nil {
		return fmt.Errorf("resizer: failed to new registry: %w", err)
	}
	defer a.logger.Sync() //nolint:errcheck

	// Workerの起動
	a.logger.Info("Started")
	switch a.RunMethod {
	case "lambda":
		lambda.StartWithOptions(a.resizer.Lambda, lambda.WithContext(ctx))
	default:
		return errors.New("not implemented")
	}

	// Workerの停止
	a.logger.Info("Shutdown...")
	a.waitGroup.Wait()
	return nil
}
