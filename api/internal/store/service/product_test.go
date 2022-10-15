package service

import (
	"context"
	"testing"

	"github.com/and-period/furumaru/api/internal/exception"
	"github.com/and-period/furumaru/api/internal/store"
	"github.com/and-period/furumaru/api/internal/store/database"
	"github.com/and-period/furumaru/api/internal/store/entity"
	"github.com/and-period/furumaru/api/internal/user"
	uentity "github.com/and-period/furumaru/api/internal/user/entity"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 6, 28, 18, 30, 0, 0)
	params := &database.ListProductsParams{
		Name:        "みかん",
		ProducerID:  "",
		ProducerIDs: []string{"producer-id"},
		Limit:       30,
		Offset:      0,
		Orders: []*database.ListProductsOrder{
			{Key: entity.ProductOrderByName, OrderByASC: true},
		},
	}
	products := entity.Products{
		{
			ID:              "product-id",
			TypeID:          "type-id",
			CategoryID:      "category-id",
			ProducerID:      "producer-id",
			Name:            "新鮮なじゃがいも",
			Description:     "新鮮なじゃがいもをお届けします。",
			Public:          true,
			Inventory:       100,
			Weight:          100,
			WeightUnit:      entity.WeightUnitGram,
			Item:            1,
			ItemUnit:        "袋",
			ItemDescription: "1袋あたり100gのじゃがいも",
			Media: entity.MultiProductMedia{
				{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
				{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
			},
			Price:            400,
			DeliveryType:     entity.DeliveryTypeNormal,
			Box60Rate:        50,
			Box80Rate:        40,
			Box100Rate:       30,
			OriginPrefecture: "滋賀県",
			OriginCity:       "彦根市",
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}

	tests := []struct {
		name        string
		setup       func(ctx context.Context, mocks *mocks)
		input       *store.ListProductsInput
		expect      entity.Products
		expectTotal int64
		expectErr   error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().List(gomock.Any(), params).Return(products, nil)
				mocks.db.Product.EXPECT().Count(gomock.Any(), params).Return(int64(1), nil)
			},
			input: &store.ListProductsInput{
				Name:        "みかん",
				ProducerID:  "",
				ProducerIDs: []string{"producer-id"},
				Limit:       30,
				Offset:      0,
				Orders: []*store.ListProductsOrder{
					{Key: entity.ProductOrderByName, OrderByASC: true},
				},
			},
			expect:      products,
			expectTotal: 1,
			expectErr:   nil,
		},
		{
			name:        "invalid argument",
			setup:       func(ctx context.Context, mocks *mocks) {},
			input:       &store.ListProductsInput{},
			expect:      nil,
			expectTotal: 0,
			expectErr:   exception.ErrInvalidArgument,
		},
		{
			name: "failed to list products",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.Product.EXPECT().Count(gomock.Any(), params).Return(int64(1), nil)
			},
			input: &store.ListProductsInput{
				Name:        "みかん",
				ProducerID:  "",
				ProducerIDs: []string{"producer-id"},
				Limit:       30,
				Offset:      0,
				Orders: []*store.ListProductsOrder{
					{Key: entity.ProductOrderByName, OrderByASC: true},
				},
			},
			expect:      nil,
			expectTotal: 0,
			expectErr:   exception.ErrUnknown,
		},
		{
			name: "failed to count products",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().List(gomock.Any(), params).Return(products, nil)
				mocks.db.Product.EXPECT().Count(gomock.Any(), params).Return(int64(0), errmock)
			},
			input: &store.ListProductsInput{
				Name:        "みかん",
				ProducerID:  "",
				ProducerIDs: []string{"producer-id"},
				Limit:       30,
				Offset:      0,
				Orders: []*store.ListProductsOrder{
					{Key: entity.ProductOrderByName, OrderByASC: true},
				},
			},
			expect:      nil,
			expectTotal: 0,
			expectErr:   exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			actual, total, err := service.ListProducts(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
			assert.ElementsMatch(t, tt.expect, actual)
			assert.Equal(t, tt.expectTotal, total)
		}))
	}
}

func TestMultiGetProducts(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 6, 28, 18, 30, 0, 0)
	products := entity.Products{
		{
			ID:              "product-id",
			TypeID:          "type-id",
			CategoryID:      "category-id",
			ProducerID:      "producer-id",
			Name:            "新鮮なじゃがいも",
			Description:     "新鮮なじゃがいもをお届けします。",
			Public:          true,
			Inventory:       100,
			Weight:          100,
			WeightUnit:      entity.WeightUnitGram,
			Item:            1,
			ItemUnit:        "袋",
			ItemDescription: "1袋あたり100gのじゃがいも",
			Media: entity.MultiProductMedia{
				{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
				{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
			},
			Price:            400,
			DeliveryType:     entity.DeliveryTypeNormal,
			Box60Rate:        50,
			Box80Rate:        40,
			Box100Rate:       30,
			OriginPrefecture: "滋賀県",
			OriginCity:       "彦根市",
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, mocks *mocks)
		input     *store.MultiGetProductsInput
		expect    entity.Products
		expectErr error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().MultiGet(ctx, []string{"product-id"}).Return(products, nil)
			},
			input: &store.MultiGetProductsInput{
				ProductIDs: []string{"product-id"},
			},
			expect:    products,
			expectErr: nil,
		},
		{
			name:  "invalid argument",
			setup: func(ctx context.Context, mocks *mocks) {},
			input: &store.MultiGetProductsInput{
				ProductIDs: []string{""},
			},
			expect:    nil,
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name: "failed to get products",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().MultiGet(ctx, []string{"product-id"}).Return(nil, errmock)
			},
			input: &store.MultiGetProductsInput{
				ProductIDs: []string{"product-id"},
			},
			expect:    nil,
			expectErr: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			actual, err := service.MultiGetProducts(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
			assert.ElementsMatch(t, tt.expect, actual)
		}))
	}
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 6, 28, 18, 30, 0, 0)
	product := &entity.Product{
		ID:              "product-id",
		TypeID:          "type-id",
		CategoryID:      "category-id",
		ProducerID:      "producer-id",
		Name:            "新鮮なじゃがいも",
		Description:     "新鮮なじゃがいもをお届けします。",
		Public:          true,
		Inventory:       100,
		Weight:          100,
		WeightUnit:      entity.WeightUnitGram,
		Item:            1,
		ItemUnit:        "袋",
		ItemDescription: "1袋あたり100gのじゃがいも",
		Media: entity.MultiProductMedia{
			{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
			{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
		},
		Price:            400,
		DeliveryType:     entity.DeliveryTypeNormal,
		Box60Rate:        50,
		Box80Rate:        40,
		Box100Rate:       30,
		OriginPrefecture: "滋賀県",
		OriginCity:       "彦根市",
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, mocks *mocks)
		input     *store.GetProductInput
		expect    *entity.Product
		expectErr error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().Get(ctx, "product-id").Return(product, nil)
			},
			input: &store.GetProductInput{
				ProductID: "product-id",
			},
			expect:    product,
			expectErr: nil,
		},
		{
			name:      "invalid argument",
			setup:     func(ctx context.Context, mocks *mocks) {},
			input:     &store.GetProductInput{},
			expect:    nil,
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name: "failed to get product",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().Get(ctx, "product-id").Return(nil, errmock)
			},
			input: &store.GetProductInput{
				ProductID: "product-id",
			},
			expect:    nil,
			expectErr: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			actual, err := service.GetProduct(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
			assert.Equal(t, tt.expect, actual)
		}))
	}
}

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	producerIn := &user.GetProducerInput{
		ProducerID: "producer-id",
	}
	producer := &uentity.Producer{
		AdminID: "producer-id",
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, mocks *mocks)
		input     *store.CreateProductInput
		expectErr error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(producer, nil)
				mocks.db.Product.EXPECT().
					Create(ctx, gomock.Any()).
					DoAndReturn(func(ctx context.Context, product *entity.Product) error {
						expect := &entity.Product{
							ID:              product.ID, // ignore
							TypeID:          "product-type-id",
							CategoryID:      "category-id",
							ProducerID:      "producer-id",
							Name:            "新鮮なじゃがいも",
							Description:     "新鮮なじゃがいもをお届けします。",
							Public:          true,
							Inventory:       100,
							Weight:          100,
							WeightUnit:      entity.WeightUnitGram,
							Item:            1,
							ItemUnit:        "袋",
							ItemDescription: "1袋あたり100gのじゃがいも",
							Media: entity.MultiProductMedia{
								{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
								{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
							},
							Price:            400,
							DeliveryType:     entity.DeliveryTypeNormal,
							Box60Rate:        50,
							Box80Rate:        40,
							Box100Rate:       30,
							OriginPrefecture: "滋賀県",
							OriginCity:       "彦根市",
						}
						assert.Equal(t, expect, product)
						return nil
					})
			},
			input: &store.CreateProductInput{
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.CreateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: nil,
		},
		{
			name:      "invalid argument",
			setup:     func(ctx context.Context, mocks *mocks) {},
			input:     &store.CreateProductInput{},
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name:  "invalid media format",
			setup: func(ctx context.Context, mocks *mocks) {},
			input: &store.CreateProductInput{
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.CreateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: true},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name: "failed to get producer",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(nil, errmock)
			},
			input: &store.CreateProductInput{
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.CreateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrUnknown,
		},
		{
			name: "failed to create product",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(producer, nil)
				mocks.db.Product.EXPECT().Create(ctx, gomock.Any()).Return(errmock)
			},
			input: &store.CreateProductInput{
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.CreateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			_, err := service.CreateProduct(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
		}))
	}
}

func TestUpdateProduct(t *testing.T) {
	t.Parallel()

	producerIn := &user.GetProducerInput{
		ProducerID: "producer-id",
	}
	producer := &uentity.Producer{
		AdminID: "producer-id",
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, mocks *mocks)
		input     *store.UpdateProductInput
		expectErr error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(producer, nil)
				mocks.db.Product.EXPECT().
					Update(ctx, "product-id", gomock.Any()).
					DoAndReturn(func(ctx context.Context, productID string, params *database.UpdateProductParams) error {
						expect := &database.UpdateProductParams{
							ProducerID:      "producer-id",
							CategoryID:      "category-id",
							TypeID:          "product-type-id",
							Name:            "新鮮なじゃがいも",
							Description:     "新鮮なじゃがいもをお届けします。",
							Public:          true,
							Inventory:       100,
							Weight:          100,
							WeightUnit:      entity.WeightUnitGram,
							Item:            1,
							ItemUnit:        "袋",
							ItemDescription: "1袋あたり100gのじゃがいも",
							Media: entity.MultiProductMedia{
								{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
								{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
							},
							Price:            400,
							DeliveryType:     entity.DeliveryTypeNormal,
							Box60Rate:        50,
							Box80Rate:        40,
							Box100Rate:       30,
							OriginPrefecture: "滋賀県",
							OriginCity:       "彦根市",
						}
						assert.Equal(t, expect, params)
						return nil
					})
			},
			input: &store.UpdateProductInput{
				ProductID:       "product-id",
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.UpdateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: nil,
		},
		{
			name:      "invalid argument",
			setup:     func(ctx context.Context, mocks *mocks) {},
			input:     &store.UpdateProductInput{},
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name:  "invalid media format",
			setup: func(ctx context.Context, mocks *mocks) {},
			input: &store.UpdateProductInput{
				ProductID:       "product-id",
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.UpdateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: true},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name: "failed to get producer",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(nil, errmock)
			},
			input: &store.UpdateProductInput{
				ProductID:       "product-id",
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.UpdateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrUnknown,
		},
		{
			name: "failed to update product",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.user.EXPECT().GetProducer(gomock.Any(), producerIn).Return(producer, nil)
				mocks.db.Product.EXPECT().Update(ctx, "product-id", gomock.Any()).Return(errmock)
			},
			input: &store.UpdateProductInput{
				ProductID:       "product-id",
				ProducerID:      "producer-id",
				CategoryID:      "category-id",
				TypeID:          "product-type-id",
				Name:            "新鮮なじゃがいも",
				Description:     "新鮮なじゃがいもをお届けします。",
				Public:          true,
				Inventory:       100,
				Weight:          100,
				WeightUnit:      entity.WeightUnitGram,
				Item:            1,
				ItemUnit:        "袋",
				ItemDescription: "1袋あたり100gのじゃがいも",
				Media: []*store.UpdateProductMedia{
					{URL: "https://and-period.jp/thumbnail01.png", IsThumbnail: true},
					{URL: "https://and-period.jp/thumbnail02.png", IsThumbnail: false},
				},
				Price:            400,
				DeliveryType:     entity.DeliveryTypeNormal,
				Box60Rate:        50,
				Box80Rate:        40,
				Box100Rate:       30,
				OriginPrefecture: "滋賀県",
				OriginCity:       "彦根市",
			},
			expectErr: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			err := service.UpdateProduct(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
		}))
	}
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		setup     func(ctx context.Context, mocks *mocks)
		input     *store.DeleteProductInput
		expectErr error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().Delete(ctx, "product-id").Return(nil)
			},
			input: &store.DeleteProductInput{
				ProductID: "product-id",
			},
			expectErr: nil,
		},
		{
			name:      "invalid argument",
			setup:     func(ctx context.Context, mocks *mocks) {},
			input:     &store.DeleteProductInput{},
			expectErr: exception.ErrInvalidArgument,
		},
		{
			name: "failed to delete product",
			setup: func(ctx context.Context, mocks *mocks) {
				mocks.db.Product.EXPECT().Delete(ctx, "product-id").Return(errmock)
			},
			input: &store.DeleteProductInput{
				ProductID: "product-id",
			},
			expectErr: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testService(tt.setup, func(ctx context.Context, t *testing.T, service *service) {
			err := service.DeleteProduct(ctx, tt.input)
			assert.ErrorIs(t, err, tt.expectErr)
		}))
	}
}
