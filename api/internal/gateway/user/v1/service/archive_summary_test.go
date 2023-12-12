package service

import (
	"testing"

	"github.com/and-period/furumaru/api/internal/common"
	"github.com/and-period/furumaru/api/internal/gateway/user/v1/response"
	"github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/stretchr/testify/assert"
)

func TestArchiveSummary(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name     string
		schedule *entity.Schedule
		expect   *ArchiveSummary
	}{
		{
			name: "success",
			schedule: &entity.Schedule{
				ID:            "schedule-id",
				CoordinatorID: "coordinator-id",
				Status:        entity.ScheduleStatusClosed,
				Title:         "スケジュールタイトル",
				Description:   "スケジュールの詳細です。",
				ThumbnailURL:  "https://example.com/thumbnail.png",
				Thumbnails: common.Images{
					{URL: "https://example.com/thumbnail_240.png", Size: common.ImageSizeSmall},
					{URL: "https://example.com/thumbnail_675.png", Size: common.ImageSizeMedium},
					{URL: "https://example.com/thumbnail_900.png", Size: common.ImageSizeLarge},
				},
				ImageURL:        "https://example.com/image.png",
				OpeningVideoURL: "https://example.com/opening-video.mp4",
				Public:          true,
				Approved:        true,
				ApprovedAdminID: "admin-id",
				StartAt:         now.AddDate(0, -1, 0),
				EndAt:           now.AddDate(0, 1, 0),
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			expect: &ArchiveSummary{
				ArchiveSummary: response.ArchiveSummary{
					ScheduleID:    "schedule-id",
					CoordinatorID: "coordinator-id",
					Title:         "スケジュールタイトル",
					ThumbnailURL:  "https://example.com/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewArchiveSummary(tt.schedule)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestArchiveSummary_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		archive *ArchiveSummary
		expect  *response.ArchiveSummary
	}{
		{
			name: "success",
			archive: &ArchiveSummary{
				ArchiveSummary: response.ArchiveSummary{
					ScheduleID:    "schedule-id",
					CoordinatorID: "coordinator-id",
					Title:         "スケジュールタイトル",
					ThumbnailURL:  "https://example.com/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
				},
			},
			expect: &response.ArchiveSummary{
				ScheduleID:    "schedule-id",
				CoordinatorID: "coordinator-id",
				Title:         "スケジュールタイトル",
				ThumbnailURL:  "https://example.com/thumbnail.png",
				Thumbnails: []*response.Image{
					{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
					{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
					{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.archive.Response())
		})
	}
}

func TestArchiveSummaries(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name      string
		schedules entity.Schedules
		expect    ArchiveSummaries
	}{
		{
			name: "success",
			schedules: entity.Schedules{
				{
					ID:            "schedule-id",
					CoordinatorID: "coordinator-id",
					Status:        entity.ScheduleStatusClosed,
					Title:         "スケジュールタイトル",
					Description:   "スケジュールの詳細です。",
					ThumbnailURL:  "https://example.com/thumbnail.png",
					Thumbnails: common.Images{
						{URL: "https://example.com/thumbnail_240.png", Size: common.ImageSizeSmall},
						{URL: "https://example.com/thumbnail_675.png", Size: common.ImageSizeMedium},
						{URL: "https://example.com/thumbnail_900.png", Size: common.ImageSizeLarge},
					},
					ImageURL:        "https://example.com/image.png",
					OpeningVideoURL: "https://example.com/opening-video.mp4",
					Public:          true,
					Approved:        true,
					ApprovedAdminID: "admin-id",
					StartAt:         now.AddDate(0, -1, 0),
					EndAt:           now.AddDate(0, 1, 0),
					CreatedAt:       now,
					UpdatedAt:       now,
				},
			},
			expect: ArchiveSummaries{
				{
					ArchiveSummary: response.ArchiveSummary{
						ScheduleID:    "schedule-id",
						CoordinatorID: "coordinator-id",
						Title:         "スケジュールタイトル",
						ThumbnailURL:  "https://example.com/thumbnail.png",
						Thumbnails: []*response.Image{
							{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
							{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
							{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewArchiveSummaries(tt.schedules)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestArchiveSummaries_Response(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		archives ArchiveSummaries
		expect   []*response.ArchiveSummary
	}{
		{
			name: "success",
			archives: ArchiveSummaries{
				{
					ArchiveSummary: response.ArchiveSummary{
						ScheduleID:    "schedule-id",
						CoordinatorID: "coordinator-id",
						Title:         "スケジュールタイトル",
						ThumbnailURL:  "https://example.com/thumbnail.png",
						Thumbnails: []*response.Image{
							{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
							{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
							{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
						},
					},
				},
			},
			expect: []*response.ArchiveSummary{
				{
					ScheduleID:    "schedule-id",
					CoordinatorID: "coordinator-id",
					Title:         "スケジュールタイトル",
					ThumbnailURL:  "https://example.com/thumbnail.png",
					Thumbnails: []*response.Image{
						{URL: "https://example.com/thumbnail_240.png", Size: int32(ImageSizeSmall)},
						{URL: "https://example.com/thumbnail_675.png", Size: int32(ImageSizeMedium)},
						{URL: "https://example.com/thumbnail_900.png", Size: int32(ImageSizeLarge)},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.archives.Response())
		})
	}
}
