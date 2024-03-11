package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type DeleteUserUseCase struct {
	Gateway entity.UserInterface
}

type DeleteUserInput struct {
	UserId string
}

func NewDeleteUserUseCase(gateway entity.UserInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{Gateway: gateway}
}

func (d *DeleteUserUseCase) Execute(ctx context.Context, input DeleteUserInput) error {
	user, err := d.Gateway.FindByID(ctx, input.UserId)
	if err != nil {
		return err
	}
	if user.IsActive {
		user.Deactivate()
	}

	_, err = d.Gateway.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
