package database

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func NewDabtaseMock() *DatabaseMock {
	return &DatabaseMock{}
}

func (m *DatabaseMock) Save(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *DatabaseMock) FindByMail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entity.User), args.Error(0)
}

func (m *DatabaseMock) FindByID(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.User), args.Error(0)
}
func (m *DatabaseMock) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*entity.User), args.Error(0)
}
