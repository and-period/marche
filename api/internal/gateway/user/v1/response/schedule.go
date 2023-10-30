package response

// Schedule - マルシェ開催情報
type Schedule struct {
	ID              string   `json:"id"`              // スケジュールID
	CoordinatorID   string   `json:"coordinatorId"`   // コーディネータID
	Status          int32    `json:"status"`          // 開催状況
	Title           string   `json:"title"`           // タイトル
	Description     string   `json:"description"`     // 説明
	ThumbnailURL    string   `json:"thumbnailUrl"`    // サムネイルURL
	Thumbnails      []*Image `json:"thumbnails"`      // サムネイルURL(リサイズ済み)一覧
	DistributionURL string   `json:"distributionUrl"` // 映像配信URL
	StartAt         int64    `json:"startAt"`         // 配信開始日時
	EndAt           int64    `json:"endAt"`           // 配信終了日時
}

type ScheduleResponse struct {
	Schedule    *Schedule    `json:"schedule"`    // マルシェ開催情報
	Coordinator *Coordinator `json:"coordinator"` // コーディネータ情報
	Lives       []*Live      `json:"lives"`       // ライブ配信一覧
	Producers   []*Producer  `json:"producers"`   // 生産者一覧
	Products    []*Product   `json:"products"`    // 商品一覧
}
