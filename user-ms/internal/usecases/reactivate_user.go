package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type ReactivateUserUseCase struct {
	Gateway entity.UserInterface
}

type ReactivateUserInput struct {
	UserId string
}

func NewReactivateUserUseCase(gateway entity.UserInterface) *ReactivateUserUseCase {
	return &ReactivateUserUseCase{Gateway: gateway}
}

func (d *ReactivateUserUseCase) Execute(ctx context.Context, input ReactivateUserInput) error {
	user, err := d.Gateway.FindByID(ctx, input.UserId)
	if err != nil {
		return err
	}
	if user.IsActive {
		return entity.ErrUserAlreadyActivated
	}
	user.Activate()

	_, err = d.Gateway.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
