package service

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/and-period/furumaru/api/internal/exception"
	"github.com/and-period/furumaru/api/internal/media"
	"github.com/and-period/furumaru/api/internal/media/entity"
)

func (s *service) GenerateCoordinatorThumbnail(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.CoordinatorThumbnailRegulation)
}

func (s *service) GenerateCoordinatorHeader(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.CoordinatorHeaderRegulation)
}

func (s *service) GenerateProducerThumbnail(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProducerThumbnailRegulation)
}

func (s *service) GenerateProducerHeader(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProducerHeaderRegulation)
}

func (s *service) GenerateProductImage(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductImageRegulation)
}

func (s *service) GenerateProductVideo(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductVideoRegulation)
}

func (s *service) GenerateProductTypeIcon(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductTypeIconPathRegulation)
}

func (s *service) generateFile(
	ctx context.Context, in *media.GenerateFileInput, reg *entity.Regulation,
) (string, error) {
	if err := s.validator.Struct(in); err != nil {
		return "", exception.InternalError(err)
	}
	var buf bytes.Buffer
	teeReader := io.TeeReader(in.File, &buf)
	if err := reg.Validate(teeReader, in.Header); err != nil {
		return "", fmt.Errorf("%w: %s", exception.ErrInvalidArgument, err.Error())
	}
	path := reg.GenerateFilePath(in.Header)
	url, err := s.temp.Upload(ctx, path, in.File)
	return url, exception.InternalError(err)
}
