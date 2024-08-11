package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepoMock "api/app/domain/repository/db/mock"
	db "api/app/infra/repository"
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_createUserUsecase_Run(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	type args struct {
		ctx  context.Context
		name string
		sub  string
	}
	type mocks struct {
		userRepo    *dbRepoMock.MockUserRepository
		subjectRepo *dbRepoMock.MockAuthSubjectRepository
	}
	tests := []struct {
		name       string
		args       args
		mockExpect func(args args, m mocks)
		wantErr    bool
	}{
		{
			name: "正常系",

			args: args{
				ctx:  context.Background(),
				name: "user name",
				sub:  "subject",
			},
			mockExpect: func(args args, m mocks) {
				user := &dbEntity.User{}
				user.Name = args.name
				subject := &dbEntity.AuthSubject{}
				subject.Subject = args.sub

				m.userRepo.EXPECT().Create(args.ctx, user).Return(nil)
				m.subjectRepo.EXPECT().Create(args.ctx, subject).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ユーザー作成エラー",
			args: args{
				ctx:  context.Background(),
				name: "user name",
				sub:  "subject",
			},
			mockExpect: func(args args, m mocks) {
				user := &dbEntity.User{}
				user.Name = args.name

				m.userRepo.EXPECT().Create(args.ctx, user).Return(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "サブジェクト作成エラー",
			args: args{
				ctx:  context.Background(),
				name: "user name",
				sub:  "subject",
			},
			mockExpect: func(args args, m mocks) {
				user := &dbEntity.User{}
				user.Name = args.name
				subject := &dbEntity.AuthSubject{}
				subject.Subject = args.sub

				m.userRepo.EXPECT().Create(args.ctx, user).Return(nil)
				m.subjectRepo.EXPECT().Create(args.ctx, subject).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			txRepo := db.NewMockTransactionRepository()
			userRepo := dbRepoMock.NewMockUserRepository(ctrl)
			subjectRepo := dbRepoMock.NewMockAuthSubjectRepository(ctrl)

			u := &createUserUsecase{
				txRepo:      txRepo,
				userRepo:    userRepo,
				subjectRepo: subjectRepo,
			}
			tt.mockExpect(tt.args, mocks{
				userRepo:    userRepo,
				subjectRepo: subjectRepo,
			})
			if err := u.Run(tt.args.ctx, tt.args.name, tt.args.sub); (err != nil) != tt.wantErr {
				t.Errorf("createUserUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
