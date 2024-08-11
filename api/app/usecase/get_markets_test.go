package usecase

import (
	dbEntity "api/app/domain/entity/db"
	mock_db "api/app/domain/repository/db/mock"
	"api/app/infra/boiler"
	"context"
	"errors"
	"io"
	"log/slog"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_getMarketsUsecase_Run(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	type mocks struct {
		marketRepo *mock_db.MockMarketRepository
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		mockExpect func(args args, m mocks)
		want       []*dbEntity.Market
		wantErr    bool
	}{
		{
			name: "正常系_マーケットが存在しない",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				m.marketRepo.EXPECT().GetAll(args.ctx).Return([]*dbEntity.Market{}, nil)
			},
			want:    []*dbEntity.Market{},
			wantErr: false,
		},
		{
			name: "正常系_マーケットが複数存在する",
			args: args{ctx: context.Background()},
			mockExpect: func(args args, m mocks) {
				markets := []*dbEntity.Market{{
					Market: boiler.Market{
						ID:          1,
						Name:        "market name",
						Description: "market description",
					},
				}, {
					Market: boiler.Market{
						ID:          2,
						Name:        "market name2",
						Description: "market description2",
					},
				}}

				m.marketRepo.EXPECT().GetAll(args.ctx).Return(markets, nil)
			},
			want: []*dbEntity.Market{{
				Market: boiler.Market{
					ID:          1,
					Name:        "market name",
					Description: "market description",
				},
			}, {
				Market: boiler.Market{
					ID:          2,
					Name:        "market name2",
					Description: "market description2",
				},
			}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			marketRepo := mock_db.NewMockMarketRepository(ctrl)

			u := &getMarketsUsecase{
				marketRepo: marketRepo,
			}
			tt.mockExpect(tt.args, mocks{
				marketRepo: marketRepo,
			})
			got, err := u.Run(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMarketsUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMarketsUsecase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
