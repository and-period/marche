//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../mock/$GOPACKAGE/$GOFILE
package media

import (
	"context"

	"github.com/and-period/furumaru/api/internal/media/entity"
)

type Service interface {
	GetUploadEvent(ctx context.Context, in *GetUploadEventInput) (*entity.UploadEvent, error) // ファイルアップロード結果取得
	// ライブ配信
	ListBroadcasts(ctx context.Context, in *ListBroadcastsInput) (entity.Broadcasts, int64, error)                                // 一覧取得
	GetBroadcastByScheduleID(ctx context.Context, in *GetBroadcastByScheduleIDInput) (*entity.Broadcast, error)                   // 一覧取得(マルシェ開催スケジュールID指定)
	CreateBroadcast(ctx context.Context, in *CreateBroadcastInput) (*entity.Broadcast, error)                                     // 登録
	GetBroadcastArchiveMP4UploadURL(ctx context.Context, in *GenerateBroadcastArchiveMP4UploadInput) (*entity.UploadEvent, error) // アーカイブ動画アップロード用URLの生成
	UpdateBroadcastArchive(ctx context.Context, in *UpdateBroadcastArchiveInput) error                                            // アーカイブ動画の更新
	GetBroadcastLiveMP4UploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)                    // ライブ配信アップロード用URLの生成
	PauseBroadcast(ctx context.Context, in *PauseBroadcastInput) error                                                            // ライブ配信の一時停止
	UnpauseBroadcast(ctx context.Context, in *UnpauseBroadcastInput) error                                                        // ライブ配信の一時停止を解除
	ActivateBroadcastRTMP(ctx context.Context, in *ActivateBroadcastRTMPInput) error                                              // ライブ配信の入力をRTMPに切り替え
	ActivateBroadcastMP4(ctx context.Context, in *ActivateBroadcastMP4Input) error                                                // ライブ配信の入力をMP4に切り替え
	ActivateBroadcastStaticImage(ctx context.Context, in *ActivateBroadcastStaticImageInput) error                                // ライブ配信のふた絵を有効化
	DeactivateBroadcastStaticImage(ctx context.Context, in *DeactivateBroadcastStaticImageInput) error                            // ライブ配信のふた絵を無効化
	// ライブ視聴履歴
	CreateBroadcastViewerLog(ctx context.Context, in *CreateBroadcastViewerLogInput) error // ライブ配信視聴履歴登録
	// ライブコメント
	ListBroadcastComments(ctx context.Context, in *ListBroadcastCommentsInput) (entity.BroadcastComments, string, error)     // ライブコメント一覧取得
	CreateBroadcastComment(ctx context.Context, in *CreateBroadcastCommentInput) (*entity.BroadcastComment, error)           // ライブコメント登録
	CreateBroadcastGuestComment(ctx context.Context, in *CreateBroadcastGuestCommentInput) (*entity.BroadcastComment, error) // ライブゲストコメント登録
	// コーディネータ
	GetCoordinatorThumbnailUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)      // サムネイル画像アップロード用URLの生成
	ResizeCoordinatorThumbnail(ctx context.Context, in *ResizeFileInput) error                                          // サムネイル画像リサイズ
	GetCoordinatorHeaderUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)         // ヘッダー画像アップロード用URLの生成
	ResizeCoordinatorHeader(ctx context.Context, in *ResizeFileInput) error                                             // ヘッダー画像リサイズ
	GetCoordinatorPromotionVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // 紹介映像アップロード用URLの生成
	GetCoordinatorBonusVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)     // 購入特典映像アップロード用URLの生成
	// 生産者
	GetProducerThumbnailUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)      // サムネイル画像アップロード用URLの生成
	ResizeProducerThumbnail(ctx context.Context, in *ResizeFileInput) error                                          // サムネイル画像リサイズ
	GetProducerHeaderUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)         // ヘッダー画像アップロード用URLの生成
	ResizeProducerHeader(ctx context.Context, in *ResizeFileInput) error                                             // ヘッダー画像リサイズ
	GetProducerPromotionVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // 紹介映像アップロード用URLの生成
	GetProducerBonusVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)     // 購入特典映像アップロード用URLの生成
	// 購入者
	GetUserThumbnailUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // サムネイル画像アップロード用URLの生成
	ResizeUserThumbnail(ctx context.Context, in *ResizeFileInput) error                                     // サムネイル画像リサイズ
	// 商品
	GetProductMediaImageUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // メディア(画像)アップロード用URLの生成
	GetProductMediaVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // メディア(映像)アップロード用URLの生成
	ResizeProductMedia(ctx context.Context, in *ResizeFileInput) error                                          // メディアリサイズ
	// 品目
	GetProductTypeIconUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // アイコン画像アップロード用URLの生成
	ResizeProductTypeIcon(ctx context.Context, in *ResizeFileInput) error                                     // アイコン画像リサイズ
	// 開催スケジュール
	GetScheduleThumbnailUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)    // アイコン画像アップロード用URLの生成
	ResizeScheduleThumbnail(ctx context.Context, in *ResizeFileInput) error                                        // サムネイル画像リサイズ
	GetScheduleImageUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error)        // 蓋絵画像アップロード用URLの生成
	GetScheduleOpeningVideoUploadURL(ctx context.Context, in *GenerateUploadURLInput) (*entity.UploadEvent, error) // オープニング動画アップロード用URLの生成
}
