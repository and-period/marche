package service

import (
	"errors"
	"net/url"
	"sync"

	"github.com/and-period/furumaru/api/internal/media"
	"github.com/and-period/furumaru/api/pkg/storage"
	"github.com/and-period/furumaru/api/pkg/validator"
	"go.uber.org/zap"
)

var (
	errParseURL   = errors.New("service: failed to parse url")
	errInvalidURL = errors.New("service: invalid url")
)

type Params struct {
	WaitGroup *sync.WaitGroup
	Tmp       storage.Bucket
	Storage   storage.Bucket
}

type service struct {
	logger     *zap.Logger
	waitGroup  *sync.WaitGroup
	validator  validator.Validator
	tmp        storage.Bucket
	storage    storage.Bucket
	tmpURL     func() *url.URL
	storageURL func() *url.URL
}

type options struct {
	logger *zap.Logger
}

type Option func(*options)

func WithLogger(logger *zap.Logger) Option {
	return func(opts *options) {
		opts.logger = logger
	}
}

func NewService(params *Params, opts ...Option) (media.Service, error) {
	dopts := &options{
		logger: zap.NewNop(),
	}
	for i := range opts {
		opts[i](dopts)
	}
	if params.Tmp == nil || params.Storage == nil {
		return nil, errors.New("tmp and storage is required")
	}
	turl, err := url.Parse(params.Tmp.GetFQDN())
	if err != nil {
		return nil, err
	}
	surl, err := url.Parse(params.Storage.GetFQDN())
	if err != nil {
		return nil, err
	}
	tmpURL := func() *url.URL {
		url := *turl // copy
		return &url
	}
	storageURL := func() *url.URL {
		url := *surl // copy
		return &url
	}
	return &service{
		logger:     dopts.logger,
		waitGroup:  params.WaitGroup,
		validator:  validator.NewValidator(),
		tmp:        params.Tmp,
		tmpURL:     tmpURL,
		storage:    params.Storage,
		storageURL: storageURL,
	}, nil
}
