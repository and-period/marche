package handler

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/request"
	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/response"
	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/service"
	"github.com/and-period/furumaru/api/internal/store"
	sentity "github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/golang/mock/gomock"
)

func TestListPromotions(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	in := &store.ListPromotionsInput{
		Limit:  20,
		Offset: 0,
		Orders: []*store.ListPromotionsOrder{},
	}
	promotions := sentity.Promotions{
		{
			ID:           "promotion-id",
			Title:        "夏の採れたて野菜マルシェを開催!!",
			Description:  "採れたての夏野菜を紹介するマルシェを開催ます!!",
			Public:       true,
			PublishedAt:  now,
			DiscountType: sentity.DiscountTypeFreeShipping,
			DiscountRate: 0,
			Code:         "code0001",
			CodeType:     sentity.PromotionCodeTypeOnce,
			StartAt:      now,
			EndAt:        now.AddDate(0, 1, 0),
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	tests := []struct {
		name   string
		setup  func(t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		query  string
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				mocks.store.EXPECT().ListPromotions(gomock.Any(), in).Return(promotions, int64(1), nil)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.PromotionsResponse{
					Promotions: []*response.Promotion{
						{
							ID:           "promotion-id",
							Title:        "夏の採れたて野菜マルシェを開催!!",
							Description:  "採れたての夏野菜を紹介するマルシェを開催ます!!",
							Public:       true,
							PublishedAt:  1640962800,
							DiscountType: int32(service.DiscountTypeFreeShipping),
							DiscountRate: 0,
							Code:         "code0001",
							StartAt:      1640962800,
							EndAt:        1640966400,
						},
					},
					Total: 1,
				},
			},
		},
		{
			name:  "invalid limit",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?limit=a",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "invalid offset",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?offset=a",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "invalid orders",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?orders=public,other",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failedo to list promotions",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				mocks.store.EXPECT().ListPromotions(gomock.Any(), in).Return(nil, int64(0), errmock)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			const format = "/v1/promotions%s"
			path := fmt.Sprintf(format, tt.query)
			testGet(t, tt.setup, tt.expect, path)
		})
	}
}

func TestGetPromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setup       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		promotionID string
		expect      *testResponse
	}{
		{
			name:        "success",
			setup:       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			promotionID: "promotion-id",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.PromotionResponse{
					Promotion: &response.Promotion{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			const format = "/v1/promotions/%s"
			path := fmt.Sprintf(format, tt.promotionID)
			testGet(t, tt.setup, tt.expect, path)
		})
	}
}

func TestCreatePromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.CreatePromotionRequest
		expect *testResponse
	}{
		{
			name:  "success",
			setup: func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreatePromotionRequest{
				Title:        "プロモーションタイトル",
				Description:  "プロモーションの詳細です。",
				Public:       true,
				PublishedAt:  1640962800,
				DiscountType: 2,
				DiscountRate: 10,
				Code:         "excode01",
				StartAt:      1640962800,
				EndAt:        1640962800,
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.PromotionResponse{
					Promotion: &response.Promotion{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			const path = "/v1/promotions"
			testPost(t, tt.setup, tt.expect, path, tt.req)
		})
	}
}

func TestUpdatePromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setup       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		promotionID string
		req         *request.UpdatePromotionRequest
		expect      *testResponse
	}{
		{
			name:        "success",
			setup:       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			promotionID: "promotion-id",
			req: &request.UpdatePromotionRequest{
				Title:        "プロモーションタイトル",
				Description:  "プロモーションの詳細です。",
				Public:       true,
				PublishedAt:  1640962800,
				DiscountType: 2,
				DiscountRate: 10,
				Code:         "excode01",
				StartAt:      1640962800,
				EndAt:        1640962800,
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			const format = "/v1/promotions/%s"
			path := fmt.Sprintf(format, tt.promotionID)
			testPatch(t, tt.setup, tt.expect, path, tt.req)
		})
	}
}

func TestDeletePromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setup       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		promotionID string
		expect      *testResponse
	}{
		{
			name:        "success",
			setup:       func(t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			promotionID: "promotion-id",
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			const format = "/v1/promotions/%s"
			path := fmt.Sprintf(format, tt.promotionID)
			testDelete(t, tt.setup, tt.expect, path)
		})
	}
}
