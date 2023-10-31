package service

import (
	"testing"

	"github.com/and-period/furumaru/api/internal/common"
	"github.com/and-period/furumaru/api/internal/gateway/user/v1/response"
	mentity "github.com/and-period/furumaru/api/internal/media/entity"
	sentity "github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/stretchr/testify/assert"
)

func TestScheduleStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		status   sentity.ScheduleStatus
		archived bool
		expect   ScheduleStatus
	}{
		{
			name:     "private",
			status:   sentity.ScheduleStatusPrivate,
			archived: false,
			expect:   ScheduleStatusUnknown,
		},
		{
			name:     "in progress",
			status:   sentity.ScheduleStatusInProgress,
			archived: false,
			expect:   ScheduleStatusUnknown,
		},
		{
			name:     "waiting",
			status:   sentity.ScheduleStatusWaiting,
			archived: false,
			expect:   ScheduleStatusWaiting,
		},
		{
			name:     "live",
			status:   sentity.ScheduleStatusLive,
			archived: false,
			expect:   ScheduleStatusLive,
		},
		{
			name:     "closed",
			status:   sentity.ScheduleStatusClosed,
			archived: false,
			expect:   ScheduleStatusClosed,
		},
		{
			name:     "archived",
			status:   sentity.ScheduleStatusClosed,
			archived: true,
			expect:   ScheduleStatusArchived,
		},
		{
			name:     "unknown",
			status:   sentity.ScheduleStatusUnknown,
			archived: false,
			expect:   ScheduleStatusUnknown,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewScheduleStatus(tt.status, tt.archived))
		})
	}
}

func TestSchedule(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name      string
		schedule  *sentity.Schedule
		broadcast *mentity.Broadcast
		expect    *Schedule
	}{
		{
			name: "success 開催前",
			schedule: &sentity.Schedule{
				ID:            "schedule-id",
				CoordinatorID: "coordinator-id",
				Status:        sentity.ScheduleStatusWaiting,
				Title:         "スケジュールタイトル",
				Description:   "スケジュールの詳細です。",
				ThumbnailURL:  "https://and-period.jp/thumbnail.png",
				Thumbnails: common.Images{
					{URL: "https://and-period.jp/thumbnail_240.png", Size: common.ImageSizeSmall},
					{URL: "https://and-period.jp/thumbnail_675.png", Size: common.ImageSizeMedium},
					{URL: "https://and-period.jp/thumbnail_900.png", Size: common.ImageSizeLarge},
				},
				ImageURL:        "https://and-period.jp/image.png",
				OpeningVideoURL: "https://and-period.jp/opening-video.mp4",
				Public:          true,
				Approved:        true,
				ApprovedAdminID: "admin-id",
				StartAt:         now.AddDate(0, -1, 0),
				EndAt:           now.AddDate(0, 1, 0),
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			broadcast: nil,
			expect: &Schedule{
				Schedule: response.Schedule{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        ScheduleStatusWaiting.Response(),
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://and-period.jp/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://and-period.jp/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://and-period.jp/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
					DistributionURL: "",
					StartAt:         1638284400,
					EndAt:           1643641200,
				},
			},
		},
		{
			name: "success 開催中",
			schedule: &sentity.Schedule{
				ID:            "schedule-id",
				CoordinatorID: "coordinator-id",
				Status:        sentity.ScheduleStatusLive,
				Title:         "スケジュールタイトル",
				Description:   "スケジュールの詳細です。",
				ThumbnailURL:  "https://and-period.jp/thumbnail.png",
				Thumbnails: common.Images{
					{URL: "https://and-period.jp/thumbnail_240.png", Size: common.ImageSizeSmall},
					{URL: "https://and-period.jp/thumbnail_675.png", Size: common.ImageSizeMedium},
					{URL: "https://and-period.jp/thumbnail_900.png", Size: common.ImageSizeLarge},
				},
				ImageURL:        "https://and-period.jp/image.png",
				OpeningVideoURL: "https://and-period.jp/opening-video.mp4",
				Public:          true,
				Approved:        true,
				ApprovedAdminID: "admin-id",
				StartAt:         now.AddDate(0, -1, 0),
				EndAt:           now.AddDate(0, 1, 0),
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			broadcast: &mentity.Broadcast{
				ID:         "broadcast-id",
				ScheduleID: "schedule-id",
				Status:     mentity.BroadcastStatusActive,
				InputURL:   "rtmp://127.0.0.1:1935/app/instance",
				OutputURL:  "http://example.com/index.m3u8",
				CreatedAt:  jst.Date(2022, 1, 1, 0, 0, 0, 0),
				UpdatedAt:  jst.Date(2022, 1, 1, 0, 0, 0, 0),
			},
			expect: &Schedule{
				Schedule: response.Schedule{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        ScheduleStatusLive.Response(),
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://and-period.jp/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://and-period.jp/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://and-period.jp/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
					DistributionURL: "http://example.com/index.m3u8",
					StartAt:         1638284400,
					EndAt:           1643641200,
				},
			},
		},
		{
			name: "success 開催後",
			schedule: &sentity.Schedule{
				ID:            "schedule-id",
				CoordinatorID: "coordinator-id",
				Status:        sentity.ScheduleStatusClosed,
				Title:         "スケジュールタイトル",
				Description:   "スケジュールの詳細です。",
				ThumbnailURL:  "https://and-period.jp/thumbnail.png",
				Thumbnails: common.Images{
					{URL: "https://and-period.jp/thumbnail_240.png", Size: common.ImageSizeSmall},
					{URL: "https://and-period.jp/thumbnail_675.png", Size: common.ImageSizeMedium},
					{URL: "https://and-period.jp/thumbnail_900.png", Size: common.ImageSizeLarge},
				},
				ImageURL:        "https://and-period.jp/image.png",
				OpeningVideoURL: "https://and-period.jp/opening-video.mp4",
				Public:          true,
				Approved:        true,
				ApprovedAdminID: "admin-id",
				StartAt:         now.AddDate(0, -1, 0),
				EndAt:           now.AddDate(0, 1, 0),
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			broadcast: &mentity.Broadcast{
				ID:         "broadcast-id",
				ScheduleID: "schedule-id",
				Status:     mentity.BroadcastStatusDisabled,
				InputURL:   "rtmp://127.0.0.1:1935/app/instance",
				OutputURL:  "http://example.com/index.m3u8",
				ArchiveURL: "http://example.com/movie.mp4",
				CreatedAt:  jst.Date(2022, 1, 1, 0, 0, 0, 0),
				UpdatedAt:  jst.Date(2022, 1, 1, 0, 0, 0, 0),
			},
			expect: &Schedule{
				Schedule: response.Schedule{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        ScheduleStatusArchived.Response(),
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://and-period.jp/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://and-period.jp/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://and-period.jp/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
					DistributionURL: "http://example.com/movie.mp4",
					StartAt:         1638284400,
					EndAt:           1643641200,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchedule(tt.schedule, tt.broadcast))
		})
	}
}

func TestSchedule_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *Schedule
		expect   *response.Schedule
	}{
		{
			name: "success",
			schedule: &Schedule{
				Schedule: response.Schedule{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        ScheduleStatusLive.Response(),
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					StartAt:       1638284400,
					EndAt:         1643641200,
				},
			},
			expect: &response.Schedule{
				ID:            "schedule-id",
				CoordinatorID: "coordinator-id",
				Status:        ScheduleStatusLive.Response(),
				Title:         "スケジュールタイトル",
				Description:   "スケジュールの詳細です。",
				ThumbnailURL:  "https://and-period.jp/thumbnail.png",
				StartAt:       1638284400,
				EndAt:         1643641200,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedule.Response())
		})
	}
}

func TestSchedules(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name       string
		schedules  sentity.Schedules
		broadcasts map[string]*mentity.Broadcast
		expect     Schedules
	}{
		{
			name: "success",
			schedules: sentity.Schedules{
				{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        sentity.ScheduleStatusLive,
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					Thumbnails: common.Images{
						{URL: "https://and-period.jp/thumbnail_240.png", Size: common.ImageSizeSmall},
						{URL: "https://and-period.jp/thumbnail_675.png", Size: common.ImageSizeMedium},
						{URL: "https://and-period.jp/thumbnail_900.png", Size: common.ImageSizeLarge},
					},
					ImageURL:        "https://and-period.jp/image.png",
					OpeningVideoURL: "https://and-period.jp/opening-video.mp4",
					Public:          true,
					Approved:        true,
					ApprovedAdminID: "admin-id",
					StartAt:         now.AddDate(0, -1, 0),
					EndAt:           now.AddDate(0, 1, 0),
					CreatedAt:       now,
					UpdatedAt:       now,
				},
			},
			expect: Schedules{
				{
					Schedule: response.Schedule{
						ID:            "schedule-id",
						CoordinatorID: "coordinator-id",
						Status:        ScheduleStatusLive.Response(),
						Title:         "スケジュールタイトル",
						Description:   "スケジュールの詳細です。",
						ThumbnailURL:  "https://and-period.jp/thumbnail.png",
						Thumbnails: []*response.Image{
							{URL: "https://and-period.jp/thumbnail_240.png", Size: int32(ImageSizeSmall)},
							{URL: "https://and-period.jp/thumbnail_675.png", Size: int32(ImageSizeMedium)},
							{URL: "https://and-period.jp/thumbnail_900.png", Size: int32(ImageSizeLarge)},
						},
						StartAt: 1638284400,
						EndAt:   1643641200,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchedules(tt.schedules, tt.broadcasts))
		})
	}
}

func TestSchedules_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules Schedules
		expect    []*response.Schedule
	}{
		{
			name: "success",
			schedules: Schedules{
				{
					Schedule: response.Schedule{
						ID:            "schedule-id",
						CoordinatorID: "coordinator-id",
						Status:        ScheduleStatusLive.Response(),
						Title:         "スケジュールタイトル",
						Description:   "スケジュールの詳細です。",
						ThumbnailURL:  "https://and-period.jp/thumbnail.png",
						StartAt:       1638284400,
						EndAt:         1643641200,
					},
				},
			},
			expect: []*response.Schedule{
				{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        ScheduleStatusLive.Response(),
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://and-period.jp/thumbnail.png",
					StartAt:       1638284400,
					EndAt:         1643641200,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedules.Response())
		})
	}
}
