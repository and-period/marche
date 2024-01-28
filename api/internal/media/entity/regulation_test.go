package entity

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/and-period/furumaru/api/pkg/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegulation_Validate(t *testing.T) {
	t.Parallel()
	reg := &Regulation{
		MaxSize: 10 << 20, // 10MB
		Formats: set.New("image/png", "image/jpeg"),
		dir:     "image/png",
	}
	type args struct {
		fileType string
		fileSize int64
	}
	tests := []struct {
		name string
		args
		expect error
	}{
		{
			name: "success",
			args: args{
				fileType: "image/png",
				fileSize: 2 << 20, // 2MB
			},
			expect: nil,
		},
		{
			name: "invalid size",
			args: args{
				fileType: "image/png",
				fileSize: 20 << 20, // 20MB
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name: "invalid format",
			args: args{
				fileType: "video/mp4",
				fileSize: 2 << 20, // 2MB
			},
			expect: ErrInvalidFileFormat,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := reg.Validate(tt.fileType, tt.fileSize)
			assert.ErrorIs(t, err, tt.expect)
		})
	}
}

func TestRegulation_ValidateV1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		regulation *Regulation
		input      func(t *testing.T) (io.Reader, *multipart.FileHeader)
		expect     error
	}{
		// BroadcastArchive
		{
			name:       "success broadcast archive",
			regulation: BroadcastArchiveRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for broadcast archive",
			regulation: BroadcastArchiveRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for broadcast archive",
			regulation: BroadcastArchiveRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for broadcast archive",
			regulation: BroadcastArchiveRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// BroadcastLiveMP4
		{
			name:       "success broadcast live mp4",
			regulation: BroadcastLiveMP4Regulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for broadcast live mp4",
			regulation: BroadcastLiveMP4Regulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for broadcast live mp4",
			regulation: BroadcastLiveMP4Regulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for broadcast live mp4",
			regulation: BroadcastLiveMP4Regulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// CoordinatorThumbnail
		{
			name:       "success coordinator thumbnail",
			regulation: CoordinatorThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for coordinator thumbnail",
			regulation: CoordinatorThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for coordinator thumbnail",
			regulation: CoordinatorThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for coordinator thumbnail",
			regulation: CoordinatorThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// CoordinatorHeader
		{
			name:       "success coordinator header",
			regulation: CoordinatorHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for coordinator header",
			regulation: CoordinatorHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for coordinator header",
			regulation: CoordinatorHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for coordinator header",
			regulation: CoordinatorHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// CoordinatorPromotionVideo
		{
			name:       "success coordinator promotion video",
			regulation: CoordinatorPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for coordinator promotion video",
			regulation: CoordinatorPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for coordinator promotion video",
			regulation: CoordinatorPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for coordinator promotion video",
			regulation: CoordinatorPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// CoordinatorBonusVideo
		{
			name:       "success coordinator bonus video",
			regulation: CoordinatorBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for coordinator bonus video",
			regulation: CoordinatorBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for coordinator bonus video",
			regulation: CoordinatorBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for coordinator bonus video",
			regulation: CoordinatorBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProducerThumbnail
		{
			name:       "success producer thumbnail",
			regulation: ProducerThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for producer thumbnail",
			regulation: ProducerThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for producer thumbnail",
			regulation: ProducerThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for producer thumbnail",
			regulation: ProducerThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProducerHeader
		{
			name:       "success producer header",
			regulation: ProducerHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for producer header",
			regulation: ProducerHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for producer header",
			regulation: ProducerHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for producer header",
			regulation: ProducerHeaderRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProducerPromotionVideo
		{
			name:       "success producer promotion video",
			regulation: ProducerPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for producer promotion video",
			regulation: ProducerPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for producer promotion video",
			regulation: ProducerPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for producer promotion video",
			regulation: ProducerPromotionVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProducerBonusVideo
		{
			name:       "success producer bonus video",
			regulation: ProducerBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for producer bonus video",
			regulation: ProducerBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for producer bonus video",
			regulation: ProducerBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for producer bonus video",
			regulation: ProducerBonusVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// UserThumbnail
		{
			name:       "success user thumbnail",
			regulation: UserThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for user thumbnail",
			regulation: UserThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for user thumbnail",
			regulation: UserThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		// ProductMediaImage
		{
			name:       "success product media image",
			regulation: ProductMediaImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for product media image",
			regulation: ProductMediaImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for product media image",
			regulation: ProductMediaImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for product media image",
			regulation: ProductMediaImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProductMediaVideo
		{
			name:       "success product media video",
			regulation: ProductMediaVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for product media video",
			regulation: ProductMediaVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for product media video",
			regulation: ProductMediaVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for product media video",
			regulation: ProductMediaVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ProductTypeIcon
		{
			name:       "success product type icon",
			regulation: ProductTypeIconRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for product type icon",
			regulation: ProductTypeIconRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for product type icon",
			regulation: ProductTypeIconRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for product type icon",
			regulation: ProductTypeIconRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ScheduleThumbnail
		{
			name:       "success schedule thumbnail",
			regulation: ScheduleThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for schedule thumbnail",
			regulation: ScheduleThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for schedule thumbnail",
			regulation: ScheduleThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for schedule thumbnail",
			regulation: ScheduleThumbnailRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ScheduleImage
		{
			name:       "success schedule image",
			regulation: ScheduleImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for schedule image",
			regulation: ScheduleImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testImageFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for schedule image",
			regulation: ScheduleImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testImageFile(t)
				header.Size = 10<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for schedule image",
			regulation: ScheduleImageRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
		// ScheduleOpeningVideo
		{
			name:       "success schedule opening video",
			regulation: ScheduleOpeningVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testVideoFile(t)
			},
			expect: nil,
		},
		{
			name:       "required for schedule opening video",
			regulation: ScheduleOpeningVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				_, header := testVideoFile(t)
				return nil, header
			},
			expect: ErrInvalidFileFormat,
		},
		{
			name:       "invalid size for schedule opening video",
			regulation: ScheduleOpeningVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				file, header := testVideoFile(t)
				header.Size = 200<<20 + 1
				return file, header
			},
			expect: ErrTooLargeFileSize,
		},
		{
			name:       "invalid format for schedule opening video",
			regulation: ScheduleOpeningVideoRegulation,
			input: func(t *testing.T) (io.Reader, *multipart.FileHeader) {
				return testImageFile(t)
			},
			expect: ErrInvalidFileFormat,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(t, tt.regulation.ValidateV1(tt.input(t)), tt.expect)
		})
	}
}

func TestRegulation_GenerateFilePath(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		regulation *Regulation
		header     *multipart.FileHeader
		args       []interface{}
		expect     string
	}{
		{
			name:       "success",
			regulation: CoordinatorThumbnailRegulation,
			header:     &multipart.FileHeader{Filename: "and-period.png"},
			expect:     "coordinators/thumbnail/[a-zA-Z0-9]+.png",
		},
		{
			name:       "success with params",
			regulation: BroadcastArchiveRegulation,
			header:     &multipart.FileHeader{Filename: "and-period.mp4"},
			args:       []interface{}{"broadcast-id"},
			expect:     "schedules/archives/broadcast-id/mp4/[a-zA-Z0-9]+.mp4",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Regexp(t, regexp.MustCompile(tt.expect), tt.regulation.GenerateFilePath(tt.header, tt.args...))
		})
	}
}

func TestRegulation_GetObjectKey(t *testing.T) {
	t.Parallel()
	reg := &Regulation{
		MaxSize: 0,
		Formats: set.New[string]("image/jpeg", "image/png", "image/webp", "video/mp4"),
		dir:     "test",
	}
	tests := []struct {
		name        string
		regulation  *Regulation
		contentType string
		args        []interface{}
		expect      string
		expectErr   error
	}{
		{
			name:        "success image/jpeg",
			regulation:  reg,
			contentType: "image/jpeg",
			args:        []interface{}{},
			expect:      "test/[a-zA-Z0-9]+.jpg",
			expectErr:   nil,
		},
		{
			name:        "success image/png",
			regulation:  reg,
			contentType: "image/png",
			expect:      "test/[a-zA-Z0-9]+.png",
			expectErr:   nil,
		},
		{
			name:        "success with params",
			regulation:  BroadcastArchiveRegulation,
			contentType: "video/mp4",
			args:        []interface{}{"broadcast-id"},
			expect:      "schedules/archives/broadcast-id/mp4/[a-zA-Z0-9]+",
			expectErr:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.regulation.GetObjectKey(tt.contentType, tt.args...)
			assert.ErrorIs(t, err, tt.expectErr)
			assert.Regexp(t, regexp.MustCompile(tt.expect), actual, actual)
		})
	}
}

func testImageFile(t *testing.T) (io.Reader, *multipart.FileHeader) {
	const filename, format = "and-period.png", "image"

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	filepath := testFilePath(t, filename)
	file, err := os.Open(filepath)
	require.NoError(t, err, err)

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, format, filename))
	header.Set("Content-Type", "multipart/form-data")
	part := &multipart.FileHeader{
		Filename: filepath,
		Header:   header,
		Size:     3 << 20, // 3MB
	}

	return file, part
}

func testVideoFile(t *testing.T) (io.Reader, *multipart.FileHeader) {
	const filename, format = "and-period.mp4", "video"

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	filepath := testFilePath(t, filename)
	file, err := os.Open(filepath)
	require.NoError(t, err, err)

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, format, filename))
	header.Set("Content-Type", "multipart/form-data")
	part := &multipart.FileHeader{
		Filename: filepath,
		Header:   header,
		Size:     10 << 20, // 10MB
	}

	return file, part
}

func testFilePath(t *testing.T, filename string) string {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	strs := strings.Split(dir, "api/internal")
	if len(strs) == 0 {
		t.Fatal("test: invalid file path")
	}
	return filepath.Join(strs[0], "/api/tmp", filename)
}
