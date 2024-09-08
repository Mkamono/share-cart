package usecase

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"math/rand"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	dbEntity "api/app/domain/entity/db"
	mock_db "api/app/domain/repository/db/mock"
	"api/app/oas"
	"api/app/shared"
)

func Test_getMarketAllUsecase_Run(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	// テストデータ
	type fakeMarket struct {
		ID          string
		Name        string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	var markets []*dbEntity.Market
	for i := 0; i < 3+rand.Intn(10); i++ {
		var fm fakeMarket
		gofakeit.Struct(&fm)
		market := shared.Projection[dbEntity.Market](fm)
		markets = append(markets, &market)
	}

	marketIDs := lo.Map(markets, func(m *dbEntity.Market, _ int) string {
		return m.ID.String()
	})

	type fakeMarketImage struct {
		ID        string
		MarketID  string
		CreatedAt time.Time
		UpdatedAt time.Time
		ImageID   string
	}

	var marketImages []*dbEntity.MarketImage
	for i := 0; i < 3+rand.Intn(10); i++ {
		var fmi fakeMarketImage
		gofakeit.Struct(&fmi)
		marketImage := shared.Projection[dbEntity.MarketImage](fmi)
		marketImages = append(marketImages, &marketImage)
	}

	type mocks struct {
		marketRepo      *mock_db.MockMarketRepository
		marketImageRepo *mock_db.MockMarketImageRepository
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name       string
		args       args
		mockExpect func(args args, m mocks)
		want       []*oas.Market
		wantErr    bool
	}{
		{
			name: "正常系_マーケットが存在しない",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				m.marketRepo.EXPECT().GetAll(args.ctx).Return([]*dbEntity.Market{}, nil)
			},
			want:    []*oas.Market{},
			wantErr: false,
		},
		{
			name: "正常系_画像が設定されていないマーケットが存在する",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				m.marketRepo.EXPECT().GetAll(args.ctx).Return(markets, nil)
				m.marketImageRepo.EXPECT().GetAllByMarketIDs(args.ctx, marketIDs).Return([]*dbEntity.MarketImage{}, nil)
			},
			want: lo.Map(markets, func(m *dbEntity.Market, _ int) *oas.Market {
				return &oas.Market{
					ID:          m.ID.String(),
					Name:        m.Name,
					Description: m.Description,
					Images:      []string{},
				}
			}),
		},
		{
			name: "正常系_画像が設定されているマーケットが存在する",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				copyMarketImages := make([]*dbEntity.MarketImage, len(marketImages))
				copy(copyMarketImages, marketImages)

				copyMarketImages[0].MarketID = marketIDs[0]
				copyMarketImages[1].MarketID = marketIDs[0]
				copyMarketImages[2].MarketID = marketIDs[1]

				m.marketRepo.EXPECT().GetAll(args.ctx).Return([]*dbEntity.Market{markets[0], markets[1], markets[2]}, nil)
				m.marketImageRepo.EXPECT().GetAllByMarketIDs(args.ctx, []string{marketIDs[0], marketIDs[1], marketIDs[2]}).Return(copyMarketImages, nil)
			},
			want: []*oas.Market{
				{
					ID:          markets[0].ID.String(),
					Name:        markets[0].Name,
					Description: markets[0].Description,
					Images:      []string{marketImages[0].ID.String(), marketImages[1].ID.String()},
				},
				{
					ID:          markets[1].ID.String(),
					Name:        markets[1].Name,
					Description: markets[1].Description,
					Images:      []string{marketImages[2].ID.String()},
				},
				{
					ID:          markets[2].ID.String(),
					Name:        markets[2].Name,
					Description: markets[2].Description,
					Images:      []string{},
				},
			},
			wantErr: false,
		},
		{
			name: "異常系_マーケット取得エラー",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				m.marketRepo.EXPECT().GetAll(args.ctx).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系_マーケット画像取得エラー",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				m.marketRepo.EXPECT().GetAll(args.ctx).Return(markets, nil)
				m.marketImageRepo.EXPECT().GetAllByMarketIDs(args.ctx, marketIDs).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			marketRepo := mock_db.NewMockMarketRepository(ctrl)
			marketImageRepo := mock_db.NewMockMarketImageRepository(ctrl)

			u := &getMarketAllUsecase{
				marketRepo:      marketRepo,
				marketImageRepo: marketImageRepo,
			}
			tt.mockExpect(tt.args, mocks{
				marketRepo:      marketRepo,
				marketImageRepo: marketImageRepo,
			})
			got, err := u.Run(tt.args.ctx)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
