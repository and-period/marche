// 管理者を登録します
// usage: go run ./main.go \
//          -db-host=127.0.0.1 -db-port=3316 -db-password=12345678 \
//          -aws-access-key=xxx -aws-secret-key=xxx \
//          -cognito-client-id=xxx -cognito-pool-id=xxx \
//          -email=test-admin@and-period.jp
package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/and-period/furumaru/api/internal/messenger"
	messengersrv "github.com/and-period/furumaru/api/internal/messenger/service"
	"github.com/and-period/furumaru/api/internal/user"
	"github.com/and-period/furumaru/api/internal/user/database"
	usersrv "github.com/and-period/furumaru/api/internal/user/service"
	"github.com/and-period/furumaru/api/pkg/cognito"
	db "github.com/and-period/furumaru/api/pkg/database"
	"github.com/and-period/furumaru/api/pkg/log"
	"github.com/and-period/furumaru/api/pkg/sqs"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awscredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"go.uber.org/zap"
)

const (
	dbName    = "users"
	awsRegion = "ap-northeast-1"

	fromName    = "&. コマンド実行"
	fromAddress = "info@and-period.jp"

	defaultTemplatePath = "./../../config/messenger/mailer/dev.yaml"
)

type app struct {
	db        *db.Client
	config    aws.Config
	auth      cognito.Client
	user      user.Service
	messenger messenger.Service
	waitGroup *sync.WaitGroup
	logger    *zap.Logger
}

func main() {
	start := time.Now()
	fmt.Println("Start..")
	if err := run(); err != nil {
		panic(err)
	}
	fmt.Printf("Done: %s\n", time.Since(start))
}

func run() error {
	var (
		dbHost, dbPort             string
		dbUsername, dbPassword     string
		dbEnabledTLS               bool
		awsAccessKey, awsSecretKey string
		authClientID, authPoolID   string
		email                      string
		err                        error
	)

	app := app{waitGroup: &sync.WaitGroup{}}
	flag.StringVar(&dbHost, "db-host", "mysql", "target mysql host")
	flag.StringVar(&dbPort, "db-port", "3306", "target mysql port")
	flag.StringVar(&dbUsername, "db-username", "root", "target mysql username")
	flag.StringVar(&dbPassword, "db-password", "12345678", "target mysql password")
	flag.BoolVar(&dbEnabledTLS, "db-enabled-tls", false, "target mysql enabled tls")
	flag.StringVar(&awsAccessKey, "aws-access-key", "", "aws access key for cognito")
	flag.StringVar(&awsSecretKey, "aws-secret-key", "", "aws secret key for cognito")
	flag.StringVar(&authClientID, "cognito-client-id", "", "target cognito client id")
	flag.StringVar(&authPoolID, "cognito-pool-id", "", "target cognito user pool id")
	flag.StringVar(&email, "email", "", "target email for created admin")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.logger, err = log.NewLogger(log.WithLogLevel("debug"))
	if err != nil {
		return err
	}
	app.db, err = app.setupDB(dbHost, dbPort, dbUsername, dbPassword, dbEnabledTLS)
	if err != nil {
		return err
	}
	app.config, err = app.setupAWSConfig(ctx, awsAccessKey, awsSecretKey)
	if err != nil {
		return err
	}
	app.auth, err = app.setupAuth(ctx, authClientID, authPoolID)
	if err != nil {
		return err
	}

	app.messenger = app.newMessengerService(ctx)
	app.user = app.newUserService()

	in := &user.CreateAdministratorInput{
		Lastname:      "&.",
		Firstname:     "管理者",
		LastnameKana:  "あんどどっと",
		FirstnameKana: "かんりしゃ",
		Email:         email,
		PhoneNumber:   "+819012345678",
	}
	_, err = app.user.CreateAdministrator(ctx, in)
	app.waitGroup.Wait()
	return err
}

func (a *app) newUserService() user.Service {
	params := &usersrv.Params{
		Database:  database.NewDatabase(&database.Params{Database: a.db}),
		AdminAuth: a.auth,
		Messenger: a.messenger,
		WaitGroup: a.waitGroup,
	}
	return usersrv.NewService(params, usersrv.WithLogger(a.logger))
}

func (a *app) newMessengerService(ctx context.Context) messenger.Service {
	params := &messengersrv.Params{
		WaitGroup: a.waitGroup,
		Producer:  sqs.NewProducer(a.config, &sqs.Params{}, sqs.WithDryRun(true)),
	}
	return messengersrv.NewService(params, messengersrv.WithLogger(a.logger))
}

func (a *app) setupDB(host, port, username, password string, tls bool) (*db.Client, error) {
	params := &db.Params{
		Socket:   "tcp",
		Host:     host,
		Port:     port,
		Database: dbName,
		Username: username,
		Password: password,
	}
	return db.NewClient(params, db.WithTLS(tls), db.WithLogger(a.logger))
}

func (a *app) setupAWSConfig(ctx context.Context, accessKey, secretKey string) (aws.Config, error) {
	awscreds := aws.NewCredentialsCache(
		awscredentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
	)
	return awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(awsRegion),
		awsconfig.WithCredentialsProvider(awscreds),
	)
}

func (a *app) setupAuth(ctx context.Context, clientID, poolID string) (cognito.Client, error) {
	params := &cognito.Params{
		UserPoolID:  poolID,
		AppClientID: clientID,
	}
	return cognito.NewClient(a.config, params, cognito.WithLogger(a.logger)), nil
}
