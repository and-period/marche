//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/messenger/$GOPACKAGE/$GOFILE
package database

import (
	"context"

	"github.com/and-period/furumaru/api/internal/messenger/entity"
	"github.com/and-period/furumaru/api/pkg/database"
)

type Params struct {
	Database *database.Client
}

type Database struct {
	Contact        Contact
	Notification   Notification
	ReceivedQueue  ReceivedQueue
	ReportTemplate ReportTemplate
}

func NewDatabase(params *Params) *Database {
	return &Database{
		Contact:        NewContact(params.Database),
		Notification:   NewNotification(params.Database),
		ReceivedQueue:  NewReceivedQueue(params.Database),
		ReportTemplate: NewReportTemplate(params.Database),
	}
}

/**
 * interface
 */
type Contact interface {
	List(ctx context.Context, params *ListContactsParams, fields ...string) (entity.Contacts, error)
	Count(ctx context.Context, params *ListContactsParams) (int64, error)
	Get(ctx context.Context, contactID string, fields ...string) (*entity.Contact, error)
	Create(ctx context.Context, contact *entity.Contact) error
	Update(ctx context.Context, contactID string, params *UpdateContactParams) error
	Delete(ctx context.Context, contactID string) error
}

type Notification interface {
	Create(ctx context.Context, notification *entity.Notification) error
}

type ReceivedQueue interface {
	Get(ctx context.Context, queueID string, fields ...string) (*entity.ReceivedQueue, error)
	Create(ctx context.Context, queue *entity.ReceivedQueue) error
	UpdateDone(ctx context.Context, queueID string, done bool) error
}

type ReportTemplate interface {
	Get(ctx context.Context, reportID string, fields ...string) (*entity.ReportTemplate, error)
}

/**
 * params
 */
type ListContactsParams struct {
	Limit  int
	Offset int
}

type UpdateContactParams struct {
	Status   entity.ContactStatus
	Priority entity.ContactPriority
	Note     string
}
