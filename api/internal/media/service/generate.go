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

func (s *service) GenerateCoordinatorPromotionVideo(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.CoordinatorPromotionVideoRegulation)
}

func (s *service) GenerateCoordinatorBonusVideo(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.CoordinatorBonusVideoRegulation)
}

func (s *service) GenerateProducerThumbnail(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProducerThumbnailRegulation)
}

func (s *service) GenerateProducerHeader(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProducerHeaderRegulation)
}

func (s *service) GenerateProductMediaImage(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductMediaImageRegulation)
}

func (s *service) GenerateProductMediaVideo(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductMediaVideoRegulation)
}

func (s *service) GenerateProductTypeIcon(ctx context.Context, in *media.GenerateFileInput) (string, error) {
	return s.generateFile(ctx, in, entity.ProductTypeIconRegulation)
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
	url, err := s.tmp.Upload(ctx, path, &buf)
	return url, exception.InternalError(err)
}
