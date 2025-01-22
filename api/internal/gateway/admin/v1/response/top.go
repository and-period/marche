package response

type TopOrderValue struct {
	Value      int64   `json:"value"`      // 値
	Comparison float64 `json:"comparison"` // 比較値（％：前日比など）
}

type TopOrderSalesTrend struct {
	Period     string `json:"period"`     // 期間
	SalesTotal int64  `json:"salesTotal"` // 売上合計
}

type TopOrdersResponse struct {
	StartAt     int64                 `json:"startAt"`     // 開始日時
	EndAt       int64                 `json:"endAt"`       // 終了日時
	PeriodType  string                `json:"periodType"`  // 期間種別
	Orders      *TopOrderValue        `json:"orders"`      // 注文数
	Users       *TopOrderValue        `json:"users"`       // ユーザー数
	Sales       *TopOrderValue        `json:"sales"`       // 売上合計
	SalesTrends []*TopOrderSalesTrend `json:"salesTrends"` // 売上推移（グラフ描画用）
}
