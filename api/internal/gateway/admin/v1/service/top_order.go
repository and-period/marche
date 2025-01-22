package service

import (
	"time"

	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/response"
	"github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/shopspring/decimal"
)

// TopOrderPeriodType - 期間種別
type TopOrderPeriodType string

const (
	TopOrderPeriodTypeDay   TopOrderPeriodType = "day"   // 日
	TopOrderPeriodTypeWeek  TopOrderPeriodType = "week"  // 週
	TopOrderPeriodTypeMonth TopOrderPeriodType = "month" // 月
)

func NewTopOrderPeriodType(periodType entity.AggregateOrderPeriodType) TopOrderPeriodType {
	switch periodType {
	case entity.AggregateOrderPeriodTypeDay:
		return TopOrderPeriodTypeDay
	case entity.AggregateOrderPeriodTypeWeek:
		return TopOrderPeriodTypeWeek
	case entity.AggregateOrderPeriodTypeMonth:
		return TopOrderPeriodTypeMonth
	default:
		return ""
	}
}

func NewTopOrderPeriodTypeFromRequest(periodType string) TopOrderPeriodType {
	return TopOrderPeriodType(periodType)
}

func (t TopOrderPeriodType) String(period time.Time) string {
	switch t {
	case TopOrderPeriodTypeDay:
		return period.Format(time.DateOnly)
	case TopOrderPeriodTypeWeek:
		days := int(period.Weekday())
		return period.AddDate(0, 0, -days).Format(time.DateOnly)
	case TopOrderPeriodTypeMonth:
		return period.Format("2006-01")
	default:
		return ""
	}
}

func (t TopOrderPeriodType) StoreEntity() entity.AggregateOrderPeriodType {
	switch t {
	case TopOrderPeriodTypeDay:
		return entity.AggregateOrderPeriodTypeDay
	case TopOrderPeriodTypeWeek:
		return entity.AggregateOrderPeriodTypeWeek
	case TopOrderPeriodTypeMonth:
		return entity.AggregateOrderPeriodTypeMonth
	default:
		return ""
	}
}

func (t TopOrderPeriodType) Response() string {
	switch t {
	case TopOrderPeriodTypeDay:
		return "day"
	case TopOrderPeriodTypeWeek:
		return "week"
	case TopOrderPeriodTypeMonth:
		return "month"
	default:
		return ""
	}
}

type TopOrderValue struct {
	response.TopOrderValue
}

func NewTopOrderValue(current, before int64) *TopOrderValue {
	if current == 0 && before == 0 {
		return newEmptyTopOrderValue()
	}

	var comparision float64
	if before == 0 {
		comparision = 100 // 0除算を避けるため
	} else {
		dcurrent := decimal.NewFromInt(current)
		dbefore := decimal.NewFromInt(before)
		diff := dcurrent.Sub(dbefore)
		dcomp := diff.Div(dbefore).Mul(decimal.NewFromInt(100))
		comparision, _ = dcomp.Float64()
	}

	return &TopOrderValue{
		TopOrderValue: response.TopOrderValue{
			Value:      current,
			Comparison: comparision,
		},
	}
}

func newEmptyTopOrderValue() *TopOrderValue {
	return &TopOrderValue{
		TopOrderValue: response.TopOrderValue{
			Value:      0,
			Comparison: 0,
		},
	}
}

func (v *TopOrderValue) Response() *response.TopOrderValue {
	return &v.TopOrderValue
}

type TopOrderSalesTrend struct {
	response.TopOrderSalesTrend
}

type TopOrderSalesTrends []*TopOrderSalesTrend

func NewTopOrderSalesTrend(aggregated *entity.AggregatedPeriodOrder, periodType TopOrderPeriodType) *TopOrderSalesTrend {
	return &TopOrderSalesTrend{
		TopOrderSalesTrend: response.TopOrderSalesTrend{
			Period:     periodType.String(aggregated.Period),
			SalesTotal: aggregated.SalesTotal,
		},
	}
}

func (t *TopOrderSalesTrend) Response() *response.TopOrderSalesTrend {
	return &t.TopOrderSalesTrend
}

func NewTopOrderSalesTrends(aggregated entity.AggregatedPeriodOrders, periodType TopOrderPeriodType) TopOrderSalesTrends {
	res := make(TopOrderSalesTrends, len(aggregated))
	for i := range aggregated {
		res[i] = NewTopOrderSalesTrend(aggregated[i], periodType)
	}
	return res
}

func (t TopOrderSalesTrends) Response() []*response.TopOrderSalesTrend {
	res := make([]*response.TopOrderSalesTrend, len(t))
	for i := range t {
		res[i] = t[i].Response()
	}
	return res
}
