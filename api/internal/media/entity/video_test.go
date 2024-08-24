package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVideo_SetStatus(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		video  *Video
		expect VideoStatus
	}{
		{
			name: "private",
			video: &Video{
				Public:      false,
				Limited:     false,
				PublishedAt: now.AddDate(0, 0, -1),
			},
			expect: VideoStatusPrivate,
		},
		{
			name: "waiting",
			video: &Video{
				Public:      true,
				Limited:     false,
				PublishedAt: now.AddDate(0, 0, 1),
			},
			expect: VideoStatusWaiting,
		},
		{
			name: "limited",
			video: &Video{
				Public:      true,
				Limited:     true,
				PublishedAt: now.AddDate(0, 0, -1),
			},
			expect: VideoStatusLimited,
		},
		{
			name: "public",
			video: &Video{
				Public:      true,
				Limited:     false,
				PublishedAt: now.AddDate(0, 0, -1),
			},
			expect: VideoStatusPublished,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.video.SetStatus(now)
			assert.Equal(t, tt.expect, tt.video.Status)
		})
	}
}

func TestVideos_IDs(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		videos Videos
		expect []string
	}{
		{
			name: "success",
			videos: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					ProductIDs:    []string{"product-id"},
					ExperienceIDs: []string{"experience-id"},
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusPublished,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					VideoProducts: []*VideoProduct{{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					}},
					VideoExperiences: []*VideoExperience{{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					}},
					PublishedAt: now.AddDate(0, 0, -1),
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []string{"video-id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.videos.IDs()
			assert.ElementsMatch(t, tt.expect, actual)
		})
	}
}

func TestVideos_CoordinatorIDs(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		videos Videos
		expect []string
	}{
		{
			name: "success",
			videos: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					ProductIDs:    []string{"product-id"},
					ExperienceIDs: []string{"experience-id"},
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusPublished,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					VideoProducts: []*VideoProduct{{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					}},
					VideoExperiences: []*VideoExperience{{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					}},
					PublishedAt: now.AddDate(0, 0, -1),
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []string{"coordinator-id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.videos.CoordinatorIDs()
			assert.ElementsMatch(t, tt.expect, actual)
		})
	}
}

func TestVideos_ProductIDs(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		videos Videos
		expect []string
	}{
		{
			name: "success",
			videos: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					ProductIDs:    []string{"product-id"},
					ExperienceIDs: []string{"experience-id"},
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusPublished,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					VideoProducts: []*VideoProduct{{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					}},
					VideoExperiences: []*VideoExperience{{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					}},
					PublishedAt: now.AddDate(0, 0, -1),
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []string{"product-id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.videos.ProductIDs()
			assert.ElementsMatch(t, tt.expect, actual)
		})
	}
}

func TestVideos_ExperienceIDs(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		videos Videos
		expect []string
	}{
		{
			name: "success",
			videos: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					ProductIDs:    []string{"product-id"},
					ExperienceIDs: []string{"experience-id"},
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusPublished,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					VideoProducts: []*VideoProduct{{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					}},
					VideoExperiences: []*VideoExperience{{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					}},
					PublishedAt: now.AddDate(0, 0, -1),
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []string{"experience-id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.videos.ExperienceIDs()
			assert.ElementsMatch(t, tt.expect, actual)
		})
	}
}

func TestVideos_Fill(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name        string
		videos      Videos
		products    map[string]VideoProducts
		experiences map[string]VideoExperiences
		expect      Videos
	}{
		{
			name: "success",
			videos: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusUnknown,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					PublishedAt:   now.AddDate(0, 0, -1),
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
			products: map[string]VideoProducts{
				"video-id": {
					{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			experiences: map[string]VideoExperiences{
				"video-id": {
					{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					},
				},
			},
			expect: Videos{
				{
					ID:            "video-id",
					CoordinatorID: "coordinator-id",
					ProductIDs:    []string{"product-id"},
					ExperienceIDs: []string{"experience-id"},
					Title:         "じゃがいもの育て方",
					Description:   "じゃがいもの育て方の動画です。",
					Status:        VideoStatusPublished,
					ThumbnailURL:  "https://example.com/thumbnail.jpg",
					VideoURL:      "https://example.com/video.mp4",
					Public:        true,
					Limited:       false,
					VideoProducts: []*VideoProduct{{
						VideoID:   "video-id",
						ProductID: "product-id",
						Priority:  1,
						CreatedAt: now,
						UpdatedAt: now,
					}},
					VideoExperiences: []*VideoExperience{{
						VideoID:      "video-id",
						ExperienceID: "experience-id",
						Priority:     1,
						CreatedAt:    now,
						UpdatedAt:    now,
					}},
					PublishedAt: now.AddDate(0, 0, -1),
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.videos.Fill(tt.products, tt.experiences, now)
			assert.Equal(t, tt.expect, tt.videos)
		})
	}
}
