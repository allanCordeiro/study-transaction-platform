package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type UpdateUserUseCase struct {
	Gateway entity.UserInterface
}

type UpdateUserUseCaseInput struct {
	UserId   string
	Name     string
	Password string
	UserType string
}

func NewUpdateUserUseCase(gateway entity.UserInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{Gateway: gateway}
}

func (u *UpdateUserUseCase) Execute(ctx context.Context, input UpdateUserUseCaseInput) error {
	user, err := u.Gateway.FindByID(ctx, input.UserId)
	if err != nil {
		return err
	}
	user.Name = input.Name
	user.NewPassword(input.Password)
	user.SetProfile(input.UserType)
	_, err = u.Gateway.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
