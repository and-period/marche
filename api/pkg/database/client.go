package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// Client - DB操作用のクライアント構造体
type Client struct {
	DB *gorm.DB
}

type Params struct {
	Socket   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type options struct {
	logger     *zap.Logger
	now        func() time.Time
	timezone   string
	enabledTLS bool
}

type Option func(opts *options)

func WithTimeZone(timezone string) Option {
	return func(opts *options) {
		opts.timezone = timezone
	}
}

func WithTLS(enabled bool) Option {
	return func(opts *options) {
		opts.enabledTLS = enabled
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(opts *options) {
		opts.logger = logger
	}
}

func WithNow(now func() time.Time) Option {
	return func(opts *options) {
		opts.now = now
	}
}

// NewClient - DBクライアントの構造体
func NewClient(params *Params, opts ...Option) (*Client, error) {
	dopts := &options{
		logger:     zap.NewNop(),
		now:        time.Now,
		timezone:   "",
		enabledTLS: false,
	}
	for i := range opts {
		opts[i](dopts)
	}

	// プライマリレプリカの作成
	db, err := newDBClient(params, dopts)
	if err != nil {
		return nil, err
	}

	c := &Client{
		DB: db,
	}
	return c, nil
}

// Begin - トランザクションの開始処理
func (c *Client) Begin(ctx context.Context, opts ...*sql.TxOptions) (*gorm.DB, error) {
	tx := c.DB.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}
	return tx, nil
}

// Close - トランザクションの終了処理
func (c *Client) Close(tx *gorm.DB) func() {
	return func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}
}

// Transaction - トランザクション処理
func (c *Client) Transaction(
	ctx context.Context, f func(tx *gorm.DB) (interface{}, error),
) (data interface{}, err error) {
	tx, err := c.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit().Error
	}()
	data, err = f(tx)
	return
}

// Statement - セレクトクエリの生成
func (c *Client) Statement(ctx context.Context, tx *gorm.DB, table string, fields ...string) *gorm.DB {
	stmt := tx.WithContext(ctx).Table(table)
	if len(fields) == 0 {
		stmt = stmt.Select("*")
	} else {
		stmt = stmt.Select(fields)
	}
	return stmt
}

// Statement - カウントクエリの生成
func (c *Client) Count(ctx context.Context, tx *gorm.DB, table string) *gorm.DB {
	return tx.WithContext(ctx).Table(table).Select("COUNT(*)")
}

func newDBClient(params *Params, opts *options) (*gorm.DB, error) {
	conf := &gorm.Config{
		Logger:  zapgorm2.New(opts.logger),
		NowFunc: opts.now,
	}

	dsn := newDSN(params, opts)
	return gorm.Open(mysql.Open(dsn), conf)
}

func newDSN(params *Params, opts *options) string {
	switch params.Socket {
	case "tcp":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True%s%s",
			params.Username,
			params.Password,
			params.Host,
			params.Port,
			params.Database,
			withTLS(opts.enabledTLS),
			withTimeZone(opts.timezone),
		)
	case "unix":
		return fmt.Sprintf(
			"%s:%s@unix(%s)/%s?charset=utf8mb4&parseTime=true%s",
			params.Username,
			params.Password,
			params.Host,
			params.Database,
			withTLS(opts.enabledTLS),
		)
	default:
		return ""
	}
}

func withTLS(enabled bool) string {
	if !enabled {
		return ""
	}
	return "&tls=true"
}

func withTimeZone(tz string) string {
	if tz == "" {
		tz = "Asia%2FTokyo"
	} else {
		tz = strings.Replace(tz, "/", "%2F", -1)
	}
	return fmt.Sprintf("&loc=%s", tz)
}
