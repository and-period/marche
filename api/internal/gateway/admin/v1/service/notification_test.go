package service

import (
	"testing"

	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/response"
	"github.com/and-period/furumaru/api/internal/messenger/entity"
	uentity "github.com/and-period/furumaru/api/internal/user/entity"
	"github.com/and-period/furumaru/api/internal/user/service"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/stretchr/testify/assert"
)

func TestNotification(t *testing.T) {
	t.Parallel()
	var date int64 = 1640962800

	tests := []struct {
		name         string
		notification *entity.Notification
		expect       *Notification
	}{
		{
			name: "success",
			notification: &entity.Notification{
				ID:          "notification-id",
				CreatedBy:   "admin-id",
				CreatorName: "登録者",
				UpdatedBy:   "admin-id",
				Title:       "キャベツ祭り開催",
				Body:        "旬のキャベツを大安売り",
				Targets: []entity.TargetType{
					entity.PostTargetUsers,
					entity.PostTargetProducers,
				},
				Public:      true,
				PublishedAt: jst.ParseFromUnix(date),
				CreatedAt:   jst.ParseFromUnix(date),
				UpdatedAt:   jst.ParseFromUnix(date),
			},
			expect: &Notification{
				Notification: response.Notification{
					ID:          "notification-id",
					CreatedBy:   "admin-id",
					CreatorName: "登録者",
					UpdatedBy:   "admin-id",
					Title:       "キャベツ祭り開催",
					Body:        "旬のキャベツを大安売り",
					Targets: []response.TargetType{
						response.PostTargetUsers,
						response.PostTargetProducers,
					},
					PublishedAt: 1640962800,
					Public:      true,
					CreatedAt:   1640962800,
					UpdatedAt:   1640962800,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewNotification(tt.notification))
		})
	}
}

func TestNotification_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		notification *Notification
		admin        *service.Admin
		expect       *Notification
	}{
		{
			name: "success",
			notification: &Notification{
				Notification: response.Notification{
					ID:        "notification-id",
					CreatedBy: "admin-id",
					UpdatedBy: "admin-id",
					Title:     "キャベツ祭り開催",
					Body:      "旬のキャベツを大安売り",
					Targets: []response.TargetType{
						response.PostTargetUsers,
						response.PostTargetProducers,
					},
					PublishedAt: 1640962800,
					Public:      true,
					CreatedAt:   1640962800,
					UpdatedAt:   1640962800,
				},
			},
			admin: &service.Admin{
				Admin: response.Admin{
					ID:            "admin-id",
					Role:          uentity.AdminRoleAdministrator,
					Lastname:      "&.",
					Firstname:     "管理者",
					LastnameKana:  "あんどぴりおど",
					FirstnameKana: "かんりしゃ",
					Email:         "test-admin@and-period.jp",
					CreatedAt:     1640962800,
					UpdatedAt:     1640962800,
				},
			},
			expect: &Notification{
				Notification: response.Notification{
					ID:          "notification-id",
					CreatedBy:   "admin-id",
					CreatorName: "&. 管理者",
					UpdatedBy:   "admin-id",
					Title:       "キャベツ祭り開催",
					Body:        "旬のキャベツを大安売り",
					Targets: []response.TargetType{
						response.PostTargetUsers,
						response.PostTargetProducers,
					},
					PublishedAt: 1640962800,
					Public:      true,
					CreatedAt:   1640962800,
					UpdatedAt:   1640962800,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.notification.Fill(tt.admin)
			assert.Equal(t, tt.expect, tt.notification)
		})
	}
}

func TestNotification_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		notification *Notification
		expect       *response.Notification
	}{
		{
			name: "success",
			notification: &Notification{
				Notification: response.Notification{
					ID:          "notification-id",
					CreatedBy:   "admin-id",
					CreatorName: "&. 管理者",
					UpdatedBy:   "admin-id",
					Title:       "キャベツ祭り開催",
					Body:        "旬のキャベツを大安売り",
					Targets: []response.TargetType{
						response.PostTargetUsers,
						response.PostTargetProducers,
					},
					PublishedAt: 1640962800,
					Public:      true,
					CreatedAt:   1640962800,
					UpdatedAt:   1640962800,
				},
			},
			expect: &response.Notification{
				ID:          "notification-id",
				CreatedBy:   "admin-id",
				CreatorName: "&. 管理者",
				UpdatedBy:   "admin-id",
				Title:       "キャベツ祭り開催",
				Body:        "旬のキャベツを大安売り",
				Targets: []response.TargetType{
					response.PostTargetUsers,
					response.PostTargetProducers,
				},
				PublishedAt: 1640962800,
				Public:      true,
				CreatedAt:   1640962800,
				UpdatedAt:   1640962800,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.notification.Response())
		})
	}
}

func TestNotifications(t *testing.T) {
	t.Parallel()

	var date int64 = 1640962800
	tests := []struct {
		name          string
		notifications entity.Notifications
		expect        Notifications
	}{
		{
			name: "success",
			notifications: entity.Notifications{
				{
					ID:        "notification-id",
					CreatedBy: "admin-id",
					UpdatedBy: "admin-id",
					Title:     "キャベツ祭り開催",
					Body:      "旬のキャベツを大安売り",
					Targets: []entity.TargetType{
						entity.PostTargetUsers,
						entity.PostTargetProducers,
					},
					Public:      true,
					PublishedAt: jst.ParseFromUnix(date),
					CreatedAt:   jst.ParseFromUnix(date),
					UpdatedAt:   jst.ParseFromUnix(date),
				},
			},
			expect: Notifications{
				{
					Notification: response.Notification{
						ID:          "notification-id",
						CreatedBy:   "admin-id",
						CreatorName: "",
						UpdatedBy:   "admin-id",
						Title:       "キャベツ祭り開催",
						Body:        "旬のキャベツを大安売り",
						Targets: []response.TargetType{
							response.PostTargetUsers,
							response.PostTargetProducers,
						},
						PublishedAt: 1640962800,
						Public:      true,
						CreatedAt:   1640962800,
						UpdatedAt:   1640962800,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewNotifications(tt.notifications))
		})
	}
}

func TestNotifications_AdminIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		notifications Notifications
		expect        []string
	}{
		{
			name: "success",
			notifications: Notifications{
				{
					Notification: response.Notification{
						ID:          "notification-id",
						CreatedBy:   "admin-id",
						CreatorName: "&. 管理者",
						UpdatedBy:   "admin-id",
						Title:       "キャベツ祭り開催",
						Body:        "旬のキャベツを大安売り",
						Targets: []response.TargetType{
							response.PostTargetUsers,
							response.PostTargetProducers,
						},
						PublishedAt: 1640962800,
						Public:      true,
						CreatedAt:   1640962800,
						UpdatedAt:   1640962800,
					},
				},
			},
			expect: []string{"admin-id"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, tt.expect, tt.notifications.AdminIDs())
		})
	}
}

func TestNotifications_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		notifications Notifications
		admins        map[string]*service.Admin
		expect        Notifications
	}{
		{
			name: "success",
			notifications: Notifications{
				{
					Notification: response.Notification{
						ID:        "notification-id",
						CreatedBy: "admin-id",
						UpdatedBy: "admin-id",
						Title:     "キャベツ祭り開催",
						Body:      "旬のキャベツを大安売り",
						Targets: []response.TargetType{
							response.PostTargetUsers,
							response.PostTargetProducers,
						},
						PublishedAt: 1640962800,
						Public:      true,
						CreatedAt:   1640962800,
						UpdatedAt:   1640962800,
					},
				},
			},
			admins: map[string]*service.Admin{
				"admin-id": {
					Admin: response.Admin{
						ID:            "admin-id",
						Role:          uentity.AdminRoleAdministrator,
						Lastname:      "&.",
						Firstname:     "管理者",
						LastnameKana:  "あんどぴりおど",
						FirstnameKana: "かんりしゃ",
						Email:         "test-admin@and-period.jp",
						CreatedAt:     1640962800,
						UpdatedAt:     1640962800,
					},
				},
			},
			expect: Notifications{
				{
					Notification: response.Notification{
						ID:          "notification-id",
						CreatedBy:   "admin-id",
						CreatorName: "&. 管理者",
						UpdatedBy:   "admin-id",
						Title:       "キャベツ祭り開催",
						Body:        "旬のキャベツを大安売り",
						Targets: []response.TargetType{
							response.PostTargetUsers,
							response.PostTargetProducers,
						},
						PublishedAt: 1640962800,
						Public:      true,
						CreatedAt:   1640962800,
						UpdatedAt:   1640962800,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.notifications.Fill(tt.admins)
			assert.Equal(t, tt.expect, tt.notifications)
		})
	}
}

func TestNotifications_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		notifications Notifications
		expect        []*response.Notification
	}{
		{
			name: "success",
			notifications: Notifications{
				{
					Notification: response.Notification{
						ID:          "notification-id",
						CreatedBy:   "admin-id",
						CreatorName: "&. 管理者",
						UpdatedBy:   "admin-id",
						Title:       "キャベツ祭り開催",
						Body:        "旬のキャベツを大安売り",
						Targets: []response.TargetType{
							response.PostTargetUsers,
							response.PostTargetProducers,
						},
						PublishedAt: 1640962800,
						Public:      true,
						CreatedAt:   1640962800,
						UpdatedAt:   1640962800,
					},
				},
			},
			expect: []*response.Notification{
				{
					ID:          "notification-id",
					CreatedBy:   "admin-id",
					CreatorName: "&. 管理者",
					UpdatedBy:   "admin-id",
					Title:       "キャベツ祭り開催",
					Body:        "旬のキャベツを大安売り",
					Targets: []response.TargetType{
						response.PostTargetUsers,
						response.PostTargetProducers,
					},
					PublishedAt: 1640962800,
					Public:      true,
					CreatedAt:   1640962800,
					UpdatedAt:   1640962800,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.notifications.Response())
		})
	}
}
