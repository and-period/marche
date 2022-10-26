package database

import (
	"context"
	"time"

	"github.com/and-period/furumaru/api/internal/exception"
	"github.com/and-period/furumaru/api/internal/user/entity"
	"github.com/and-period/furumaru/api/pkg/database"
	"github.com/and-period/furumaru/api/pkg/jst"
	"gorm.io/gorm"
)

const coordinatorTable = "coordinators"

var coordinatorFields = []string{
	"admin_id", "phone_number", "company_name", "store_name", "thumbnail_url", "header_url",
	"twitter_account", "instagram_account", "facebook_account",
	"postal_code", "prefecture", "city", "address_line1", "address_line2",
	"created_at", "updated_at", "deleted_at",
}

type coordinator struct {
	db  *database.Client
	now func() time.Time
}

func NewCoordinator(db *database.Client) Coordinator {
	return &coordinator{
		db:  db,
		now: jst.Now,
	}
}

func (c *coordinator) List(
	ctx context.Context, params *ListCoordinatorsParams, fields ...string,
) (entity.Coordinators, error) {
	var coordinators entity.Coordinators
	if len(fields) == 0 {
		fields = coordinatorFields
	}

	stmt := c.db.DB.WithContext(ctx).Table(coordinatorTable).Select(fields)
	if params.Limit > 0 {
		stmt = stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt = stmt.Offset(params.Offset)
	}

	if err := stmt.Find(&coordinators).Error; err != nil {
		return nil, exception.InternalError(err)
	}
	if err := c.fill(ctx, c.db.DB, coordinators...); err != nil {
		return nil, exception.InternalError(err)
	}
	return coordinators, nil
}

func (c *coordinator) Count(ctx context.Context, params *ListCoordinatorsParams) (int64, error) {
	var total int64

	stmt := c.db.DB.WithContext(ctx).Table(coordinatorTable).Select("COUNT(*)")

	err := stmt.Count(&total).Error
	return total, exception.InternalError(err)
}

func (c *coordinator) MultiGet(
	ctx context.Context, coordinatorIDs []string, fields ...string,
) (entity.Coordinators, error) {
	var coordinators entity.Coordinators
	if len(fields) == 0 {
		fields = coordinatorFields
	}

	stmt := c.db.DB.WithContext(ctx).
		Table(coordinatorTable).Select(fields).
		Where("admin_id IN (?)", coordinatorIDs)

	if err := stmt.Find(&coordinators).Error; err != nil {
		return nil, exception.InternalError(err)
	}
	if err := c.fill(ctx, c.db.DB, coordinators...); err != nil {
		return nil, exception.InternalError(err)
	}
	return coordinators, nil
}

func (c *coordinator) Get(
	ctx context.Context, coordinatorID string, fields ...string,
) (*entity.Coordinator, error) {
	coordinator, err := c.get(ctx, c.db.DB, coordinatorID, fields...)
	if err != nil {
		return nil, exception.InternalError(err)
	}
	if err := c.fill(ctx, c.db.DB, coordinator); err != nil {
		return nil, exception.InternalError(err)
	}
	return coordinator, nil
}

func (c *coordinator) Create(
	ctx context.Context, coordinator *entity.Coordinator, auth func(ctx context.Context) error,
) error {
	_, err := c.db.Transaction(ctx, func(tx *gorm.DB) (interface{}, error) {
		now := c.now()
		coordinator.Admin.CreatedAt, coordinator.Admin.UpdatedAt = now, now
		if err := tx.WithContext(ctx).Table(adminTable).Create(&coordinator.Admin).Error; err != nil {
			return nil, err
		}
		coordinator.CreatedAt, coordinator.UpdatedAt = now, now
		if err := tx.WithContext(ctx).Table(coordinatorTable).Create(&coordinator).Error; err != nil {
			return nil, err
		}
		return nil, auth(ctx)
	})
	return exception.InternalError(err)
}

func (c *coordinator) Update(ctx context.Context, coordinatorID string, params *UpdateCoordinatorParams) error {
	_, err := c.db.Transaction(ctx, func(tx *gorm.DB) (interface{}, error) {
		if _, err := c.get(ctx, tx, coordinatorID); err != nil {
			return nil, err
		}

		now := c.now()
		adminParams := map[string]interface{}{
			"lastname":       params.Lastname,
			"firstname":      params.Firstname,
			"lastname_kana":  params.LastnameKana,
			"firstname_kana": params.FirstnameKana,
			"updated_at":     now,
		}
		coordinatorParams := map[string]interface{}{
			"company_name":      params.CompanyName,
			"store_name":        params.StoreName,
			"thumbnail_url":     params.ThumbnailURL,
			"header_url":        params.HeaderURL,
			"twitter_account":   params.TwitterAccount,
			"instagram_account": params.InstagramAccount,
			"facebook_account":  params.FacebookAccount,
			"phone_number":      params.PhoneNumber,
			"postal_code":       params.PostalCode,
			"city":              params.City,
			"address_line1":     params.AddressLine1,
			"address_line2":     params.AddressLine2,
			"updated_at":        now,
		}

		err := tx.WithContext(ctx).
			Table(adminTable).
			Where("id = ?", coordinatorID).
			Updates(adminParams).Error
		if err != nil {
			return nil, err
		}
		err = tx.WithContext(ctx).
			Table(coordinatorTable).
			Where("admin_id = ?", coordinatorID).
			Updates(coordinatorParams).Error
		return nil, err
	})
	return exception.InternalError(err)
}

func (c *coordinator) get(
	ctx context.Context, tx *gorm.DB, coordinatorID string, fields ...string,
) (*entity.Coordinator, error) {
	var coordinator *entity.Coordinator
	if len(fields) == 0 {
		fields = coordinatorFields
	}

	err := tx.WithContext(ctx).
		Table(coordinatorTable).Select(fields).
		Where("admin_id = ?", coordinatorID).
		First(&coordinator).Error
	return coordinator, err
}

func (c *coordinator) fill(ctx context.Context, tx *gorm.DB, coordinators ...*entity.Coordinator) error {
	var admins entity.Admins

	ids := entity.Coordinators(coordinators).IDs()
	if len(ids) == 0 {
		return nil
	}

	stmt := tx.WithContext(ctx).
		Table(adminTable).Select(adminFields).
		Where("id IN (?)", ids)
	if err := stmt.Find(&admins).Error; err != nil {
		return err
	}

	adminMap := admins.Map()

	for i, c := range coordinators {
		admin, ok := adminMap[c.AdminID]
		if !ok {
			admin = &entity.Admin{}
		}

		coordinators[i].Fill(admin)
	}
	return nil
}
