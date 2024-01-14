package entity

import (
	"testing"
	"time"

	"github.com/and-period/furumaru/api/internal/codes"
	"github.com/and-period/furumaru/api/internal/user/entity"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/and-period/furumaru/api/pkg/set"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	t.Parallel()
	shikoku := []int32{
		codes.PrefectureValues["tokushima"],
		codes.PrefectureValues["kagawa"],
		codes.PrefectureValues["ehime"],
		codes.PrefectureValues["kochi"],
	}
	set := set.New(shikoku...)
	others := make([]int32, 0, 47-len(shikoku))
	for _, val := range codes.PrefectureValues {
		if set.Contains(val) {
			continue
		}
		others = append(others, val)
	}
	rates := ShippingRates{
		{Number: 1, Name: "四国", Price: 250, PrefectureCodes: shikoku},
		{Number: 2, Name: "その他", Price: 500, PrefectureCodes: others},
	}
	tests := []struct {
		name   string
		params *NewOrderParams
		expect *Order
		hasErr bool
	}{
		{
			name: "success",
			params: &NewOrderParams{
				OrderID:       "order-id",
				CoordinatorID: "coordinator-id",
				Customer: &entity.User{
					Member: entity.Member{
						UserID:       "user-id",
						CognitoID:    "cognito-id",
						AccountID:    "account-id",
						Username:     "username",
						ProviderType: entity.ProviderTypeEmail,
						Email:        "test@example.com",
						PhoneNumber:  "+819012345678",
						ThumbnailURL: "",
					},
					ID:         "user-id",
					Registered: true,
				},
				BillingAddress: &entity.Address{
					AddressRevision: entity.AddressRevision{
						ID:             1,
						AddressID:      "address-id",
						Lastname:       "&.",
						Firstname:      "購入者",
						PostalCode:     "1000014",
						PrefectureCode: 13,
						City:           "千代田区",
						AddressLine1:   "永田町1-7-1",
						AddressLine2:   "",
						PhoneNumber:    "+819012345678",
					},
					ID:     "address-id",
					UserID: "user-id",
				},
				ShippingAddress: &entity.Address{
					AddressRevision: entity.AddressRevision{
						ID:             1,
						AddressID:      "address-id",
						Lastname:       "&.",
						Firstname:      "購入者",
						PostalCode:     "1000014",
						PrefectureCode: 13,
						City:           "千代田区",
						AddressLine1:   "永田町1-7-1",
						AddressLine2:   "",
						PhoneNumber:    "+819012345678",
					},
					ID:     "address-id",
					UserID: "user-id",
				},
				Shipping: &Shipping{
					ID:            "coordinator-id",
					CoordinatorID: "coordinator-id",
					ShippingRevision: ShippingRevision{
						ShippingID:        "coordinator-id",
						Box60Rates:        rates,
						Box60Frozen:       800,
						Box80Rates:        rates,
						Box80Frozen:       800,
						Box100Rates:       rates,
						Box100Frozen:      800,
						HasFreeShipping:   true,
						FreeShippingRates: 3000,
					},
				},
				Baskets: CartBaskets{
					{
						BoxNumber: 1,
						BoxType:   ShippingTypeNormal,
						BoxSize:   ShippingSize60,
						BoxRate:   80,
						Items: []*CartItem{
							{
								ProductID: "product-id01",
								Quantity:  1,
							},
							{
								ProductID: "product-id02",
								Quantity:  2,
							},
						},
						CoordinatorID: "coordinator-id",
					},
				},
				Products: Products{
					{
						ID:   "product-id01",
						Name: "じゃがいも",
						ProductRevision: ProductRevision{
							ID:        1,
							ProductID: "product-id01",
							Price:     500,
						},
					},
					{
						ID:   "product-id02",
						Name: "人参",
						ProductRevision: ProductRevision{
							ID:        2,
							ProductID: "product-id02",
							Price:     1980,
						},
					},
				},
				PaymentMethodType: PaymentMethodTypeCreditCard,
				Promotion: &Promotion{
					Title:        "プロモーションタイトル",
					Description:  "プロモーションの詳細です。",
					Public:       true,
					PublishedAt:  jst.Date(2022, 8, 9, 18, 30, 0, 0),
					DiscountType: DiscountTypeRate,
					DiscountRate: 10,
					Code:         "excode01",
					CodeType:     PromotionCodeTypeAlways,
					StartAt:      jst.Date(2022, 8, 1, 0, 0, 0, 0),
					EndAt:        jst.Date(2022, 9, 1, 0, 0, 0, 0),
				},
			},
			expect: &Order{
				OrderPayment: OrderPayment{
					OrderID:           "order-id",
					AddressRevisionID: 1,
					Status:            PaymentStatusPending,
					TransactionID:     "",
					MethodType:        PaymentMethodTypeCreditCard,
					Subtotal:          4460,
					Discount:          446,
					ShippingFee:       0,
					Tax:               364,
					Total:             4014,
				},
				OrderFulfillments: OrderFulfillments{
					{
						OrderID:           "order-id",
						AddressRevisionID: 1,
						Status:            FulfillmentStatusUnfulfilled,
						TrackingNumber:    "",
						ShippingCarrier:   ShippingCarrierUnknown,
						ShippingType:      ShippingTypeNormal,
						BoxNumber:         1,
						BoxSize:           ShippingSize60,
						BoxRate:           80,
					},
				},
				OrderItems: OrderItems{
					{
						ProductRevisionID: 1,
						OrderID:           "order-id",
						Quantity:          1,
					},
					{
						ProductRevisionID: 2,
						OrderID:           "order-id",
						Quantity:          2,
					},
				},
				ID:              "order-id",
				UserID:          "user-id",
				CoordinatorID:   "coordinator-id",
				PromotionID:     "",
				ShippingMessage: "ご注文ありがとうございます！商品到着まで今しばらくお待ち下さい。",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewOrder(tt.params)
			if tt.hasErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			for _, f := range actual.OrderFulfillments {
				f.ID = "" // ignore
			}
			for _, i := range actual.OrderItems {
				i.FulfillmentID = "" // ignore
			}
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestOrder_Fill(t *testing.T) {
	t.Parallel()
	now := time.Now()
	tests := []struct {
		name         string
		order        *Order
		payment      *OrderPayment
		fulfillments OrderFulfillments
		items        OrderItems
		expect       *Order
	}{
		{
			name: "success",
			order: &Order{
				ID:            "order-id",
				UserID:        "user-id",
				CoordinatorID: "coordinator-id",
				PromotionID:   "promotion-id",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			payment:      &OrderPayment{OrderID: "order-id"},
			fulfillments: OrderFulfillments{{OrderID: "order-id"}},
			items:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
			expect: &Order{
				ID:                "order-id",
				UserID:            "user-id",
				CoordinatorID:     "coordinator-id",
				PromotionID:       "promotion-id",
				OrderPayment:      OrderPayment{OrderID: "order-id"},
				OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
				OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				CreatedAt:         now,
				UpdatedAt:         now,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.order.Fill(tt.payment, tt.fulfillments, tt.items)
			assert.Equal(t, tt.expect, tt.order)
		})
	}
}

func TestOrder_EnableAction(t *testing.T) {
	t.Parallel()
	type want struct {
		capturable  bool
		preservable bool
		completable bool
		cancelable  bool
		refundable  bool
	}
	tests := []struct {
		name  string
		order *Order
		want  want
	}{
		{
			name: "payment pending",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusPending},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusUnfulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  true,
				refundable:  false,
			},
		},
		{
			name: "payment authorized",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusAuthorized},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusUnfulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  true,
				preservable: false,
				completable: false,
				cancelable:  true,
				refundable:  false,
			},
		},
		{
			name: "payment captured and unfulfilled",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusCaptured},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusUnfulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: true,
				completable: false,
				cancelable:  false,
				refundable:  true,
			},
		},
		{
			name: "payment captured and fulfilled",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusCaptured},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusFulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: true,
				completable: true,
				cancelable:  false,
				refundable:  true,
			},
		},
		{
			name: "payment captured and already completed",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusCaptured},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusFulfilled}},
				CompletedAt:       time.Now(),
			},
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  false,
				refundable:  true,
			},
		},
		{
			name: "payment canceled",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusCanceled},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusUnfulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  false,
				refundable:  false,
			},
		},
		{
			name: "payment refunded",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusRefunded},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusFulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  false,
				refundable:  false,
			},
		},
		{
			name: "payment failed",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{Status: PaymentStatusFailed},
				OrderFulfillments: OrderFulfillments{{Status: FulfillmentStatusFulfilled}},
				CompletedAt:       time.Time{},
			},
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  false,
				refundable:  false,
			},
		},
		{
			name:  "empty",
			order: nil,
			want: want{
				capturable:  false,
				preservable: false,
				completable: false,
				cancelable:  false,
				refundable:  false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want.capturable, tt.order.Capturable())
			assert.Equal(t, tt.want.preservable, tt.order.Preservable())
			assert.Equal(t, tt.want.completable, tt.order.Completable())
			assert.Equal(t, tt.want.cancelable, tt.order.Cancelable())
			assert.Equal(t, tt.want.refundable, tt.order.Refundable())
		})
	}
}

func TestOrder_IsCanceled(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		order  *Order
		expect bool
	}{
		{
			name: "canceled",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{OrderID: "order-id", Status: PaymentStatusRefunded},
				OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
				OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
			},
			expect: true,
		},
		{
			name: "not canceled",
			order: &Order{
				ID:                "order-id",
				OrderPayment:      OrderPayment{OrderID: "order-id", Status: PaymentStatusPending},
				OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
				OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
			},
			expect: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.order.IsCanceled())
		})
	}
}

func TestOrders_IDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []string
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id",
					OrderPayment:      OrderPayment{OrderID: "order-id"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []string{"order-id"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.orders.IDs())
		})
	}
}

func TestOrders_UserIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []string
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id",
					UserID:            "user-id",
					CoordinatorID:     "coordinator-id",
					OrderPayment:      OrderPayment{OrderID: "order-id"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []string{"user-id"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.orders.UserIDs())
		})
	}
}

func TestOrders_CoordinatorID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []string
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id",
					UserID:            "user-id",
					CoordinatorID:     "coordinator-id",
					OrderPayment:      OrderPayment{OrderID: "order-id"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []string{"coordinator-id"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.orders.CoordinatorIDs())
		})
	}
}

func TestOrders_PromotionIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []string
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id01",
					UserID:            "user-id",
					CoordinatorID:     "coordinator-id",
					PromotionID:       "promotion-id",
					OrderPayment:      OrderPayment{OrderID: "order-id"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
				{
					ID:                "order-id02",
					UserID:            "user-id",
					CoordinatorID:     "coordinator-id",
					OrderPayment:      OrderPayment{OrderID: "order-id"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id"}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []string{"promotion-id"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.orders.PromotionIDs())
		})
	}
}

func TestOrders_AddressRevisionIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []int64
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id",
					UserID:            "user-id",
					OrderPayment:      OrderPayment{OrderID: "order-id", AddressRevisionID: 1},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id", AddressRevisionID: 2}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []int64{1, 2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, tt.expect, tt.orders.AddressRevisionIDs())
		})
	}
}

func TestOrders_ProductRevisionIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders Orders
		expect []int64
	}{
		{
			name: "success",
			orders: Orders{
				{
					ID:                "order-id",
					UserID:            "user-id",
					OrderPayment:      OrderPayment{OrderID: "order-id", AddressRevisionID: 1},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id", AddressRevisionID: 2}},
					OrderItems:        OrderItems{{OrderID: "order-id", ProductRevisionID: 1}},
				},
			},
			expect: []int64{1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, tt.expect, tt.orders.ProductRevisionIDs())
		})
	}
}

func TestOrders_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		orders       Orders
		payments     map[string]*OrderPayment
		fulfillments map[string]OrderFulfillments
		items        map[string]OrderItems
		expect       Orders
	}{
		{
			name: "success",
			orders: Orders{
				{ID: "order-id01"},
				{ID: "order-id02"},
				{ID: "order-id03"},
			},
			payments: map[string]*OrderPayment{
				"order-id01": {OrderID: "order-id01"},
				"order-id02": {OrderID: "order-id02"},
			},
			fulfillments: map[string]OrderFulfillments{
				"order-id01": {{OrderID: "order-id01"}},
				"order-id02": {{OrderID: "order-id02"}},
			},
			items: map[string]OrderItems{
				"order-id01": {{OrderID: "order-id01", ProductRevisionID: 1}},
				"order-id02": {{OrderID: "order-id02", ProductRevisionID: 1}},
			},
			expect: Orders{
				{
					ID:                "order-id01",
					OrderPayment:      OrderPayment{OrderID: "order-id01"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id01"}},
					OrderItems:        OrderItems{{OrderID: "order-id01", ProductRevisionID: 1}},
				},
				{
					ID:                "order-id02",
					OrderPayment:      OrderPayment{OrderID: "order-id02"},
					OrderFulfillments: OrderFulfillments{{OrderID: "order-id02"}},
					OrderItems:        OrderItems{{OrderID: "order-id02", ProductRevisionID: 1}},
				},
				{
					ID:           "order-id03",
					OrderPayment: OrderPayment{},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.orders.Fill(tt.payments, tt.fulfillments, tt.items)
			assert.Equal(t, tt.expect, tt.orders)
		})
	}
}

func TestAggregatedOrders_Map(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		orders AggregatedOrders
		expect map[string]*AggregatedOrder
	}{
		{
			name: "success",
			orders: AggregatedOrders{
				{
					UserID:     "user-id",
					OrderCount: 2,
					Subtotal:   3000,
					Discount:   0,
				},
			},
			expect: map[string]*AggregatedOrder{
				"user-id": {
					UserID:     "user-id",
					OrderCount: 2,
					Subtotal:   3000,
					Discount:   0,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.orders.Map())
		})
	}
}
