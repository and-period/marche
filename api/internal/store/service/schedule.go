package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/and-period/furumaru/api/internal/exception"
	"github.com/and-period/furumaru/api/internal/store"
	"github.com/and-period/furumaru/api/internal/store/database"
	"github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/and-period/furumaru/api/internal/user"
	"golang.org/x/sync/errgroup"
)

func (s *service) ListSchedules(ctx context.Context, in *store.ListSchedulesInput) (entity.Schedules, int64, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, 0, exception.InternalError(err)
	}
	params := &database.ListSchedulesParams{
		Limit:  int(in.Limit),
		Offset: int(in.Offset),
	}
	var (
		schedules entity.Schedules
		total     int64
	)
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		schedules, err = s.db.Schedule.List(ectx, params)
		return
	})
	eg.Go(func() (err error) {
		total, err = s.db.Schedule.Count(ectx, params)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, exception.InternalError(err)
	}
	return schedules, total, nil
}

func (s *service) GetSchedule(ctx context.Context, in *store.GetScheduleInput) (*entity.Schedule, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, exception.InternalError(err)
	}
	schedule, err := s.db.Schedule.Get(ctx, in.ScheduleID)
	return schedule, exception.InternalError(err)
}

func (s *service) CreateSchedule(ctx context.Context, in *store.CreateScheduleInput) (*entity.Schedule, entity.Lives, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, nil, exception.InternalError(err)
	}
	params := &entity.NewScheduleParams{
		CoordinatorID:        in.CoordinatorID,
		ShippingID:           in.ShippingID,
		Title:                in.Title,
		Description:          in.Description,
		ThumbnailURL:         in.ThumbnailURL,
		OpeningVideoURL:      in.OpeningVideoURL,
		IntermissionVideoURL: in.IntermissionVideoURL,
		StartAt:              in.StartAt,
		EndAt:                in.EndAt,
	}
	schedule := entity.NewSchedule(params)
	lives := make(entity.Lives, len(in.Lives))
	products := make(entity.LiveProducts, 0, len(in.Lives))
	for i := range in.Lives {
		params := &entity.NewLiveParams{
			ScheduleID:  schedule.ID,
			ProducerID:  in.Lives[i].ProducerID,
			Title:       in.Lives[i].Title,
			Description: in.Lives[i].Description,
			StartAt:     in.Lives[i].StartAt,
			EndAt:       in.Lives[i].EndAt,
		}
		lives[i] = entity.NewLive(params)
		products = append(products, entity.NewLiveProducts(lives[i].ID, in.Lives[i].ProductIDs)...)
	}
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		in := &user.GetCoordinatorInput{
			CoordinatorID: in.CoordinatorID,
		}
		_, err := s.user.GetCoordinator(ectx, in)
		if errors.Is(err, exception.ErrNotFound) {
			return fmt.Errorf("service: not found coordinator: %w", exception.ErrInvalidArgument)
		}
		return err
	})
	eg.Go(func() error {
		producerIDs := lives.ProducerIDs()
		in := &user.MultiGetProducersInput{
			ProducerIDs: producerIDs,
		}
		ps, err := s.user.MultiGetProducers(ectx, in)
		if err != nil {
			return err
		}
		if len(ps) == len(producerIDs) {
			return nil
		}
		return fmt.Errorf("service: unmatch producers length: %w", exception.ErrInvalidArgument)
	})
	eg.Go(func() error {
		shippingID := schedule.ShippingID
		_, err := s.db.Shipping.Get(ectx, shippingID)
		if errors.Is(err, exception.ErrNotFound) {
			return fmt.Errorf("service: not found shipping: %w", exception.ErrNotFound)
		}
		return err
	})
	eg.Go(func() error {
		productIDs := products.ProductIDs()
		ps, err := s.db.Product.MultiGet(ectx, productIDs)
		if err != nil {
			return err
		}
		if len(ps) == len(productIDs) {
			return nil
		}
		return fmt.Errorf("service: unmatch products length: %w", exception.ErrInvalidArgument)
	})
	if err := eg.Wait(); err != nil {
		return nil, nil, exception.InternalError(err)
	}
	if err := s.db.Schedule.Create(ctx, schedule, lives, products); err != nil {
		return nil, nil, exception.InternalError(err)
	}
	return schedule, lives, nil
}
