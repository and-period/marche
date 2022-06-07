package service

import (
	"context"

	"github.com/and-period/furumaru/api/internal/exception"
	"github.com/and-period/furumaru/api/internal/user"
	"github.com/and-period/furumaru/api/internal/user/database"
	"github.com/and-period/furumaru/api/internal/user/entity"
	"github.com/and-period/furumaru/api/pkg/random"
	"github.com/and-period/furumaru/api/pkg/uuid"
	"go.uber.org/zap"
)

func (s *userService) ListProducers(ctx context.Context, in *user.ListProducersInput) (entity.Producers, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, exception.InternalError(err)
	}
	params := &database.ListProducersParams{
		Limit:  int(in.Limit),
		Offset: int(in.Offset),
	}
	producers, err := s.db.Producer.List(ctx, params)
	return producers, exception.InternalError(err)
}

func (s *userService) GetProducer(ctx context.Context, in *user.GetProducerInput) (*entity.Producer, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, exception.InternalError(err)
	}
	producer, err := s.db.Producer.Get(ctx, in.ProducerID)
	return producer, exception.InternalError(err)
}

func (s *userService) CreateProducer(ctx context.Context, in *user.CreateProducerInput) (*entity.Producer, error) {
	const size = 8
	if err := s.validator.Struct(in); err != nil {
		return nil, exception.InternalError(err)
	}
	cognitoID := uuid.Base58Encode(uuid.New())
	password := random.NewStrings(size)
	if err := s.createCognitoAdmin(ctx, cognitoID, in.Email, password); err != nil {
		return nil, exception.InternalError(err)
	}
	producerID := uuid.Base58Encode(uuid.New())
	params := &entity.NewProducerParams{
		ID:            producerID,
		Lastname:      in.Lastname,
		Firstname:     in.Firstname,
		LastnameKana:  in.Lastname,
		FirstnameKana: in.Firstname,
		StoreName:     in.StoreName,
		ThumbnailURL:  in.ThumbnailURL,
		HeaderURL:     in.HeaderURL,
		Email:         in.Email,
		PhoneNumber:   in.PhoneNumber,
		PostalCode:    in.PostalCode,
		Prefecture:    in.Prefecture,
		City:          in.City,
		AddressLine1:  in.AddressLine1,
		AddressLine2:  in.AddressLine2,
	}
	auth := entity.NewAdminAuth(producerID, cognitoID, entity.AdminRoleProducer)
	producer := entity.NewProducer(params)
	if err := s.db.Producer.Create(ctx, auth, producer); err != nil {
		return nil, exception.InternalError(err)
	}
	s.logger.Debug("Create producer", zap.String("producerId", producerID), zap.String("password", password))
	s.waitGroup.Add(1)
	go func() {
		defer s.waitGroup.Done()
		err := s.notifyRegisterAdmin(context.Background(), producer.Name(), producer.Email, password)
		if err != nil {
			s.logger.Warn("Failed to notify register admin", zap.String("producerId", producerID), zap.Error(err))
		}
	}()
	return producer, nil
}
