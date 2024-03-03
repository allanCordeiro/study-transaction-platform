package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type ChangePasswordUseCase struct {
	Gateway entity.UserInterface
}

type ChangePasswordUseCaseInput struct {
	UserId      string
	OldPassword string
	NewPassword string
}

type ChangePasswordUseCaseOutput struct {
	Success bool
}

func NewChangePasswordUseCase(gateway entity.UserInterface) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{
		Gateway: gateway,
	}
}

func (c *ChangePasswordUseCase) Execute(ctx context.Context, input ChangePasswordUseCaseInput) (ChangePasswordUseCaseOutput, error) {
	user, err := c.Gateway.FindByID(ctx, input.UserId)
	if err != nil {
		return ChangePasswordUseCaseOutput{Success: false}, err
	}

	if !user.IsPasswordValid(input.OldPassword) {
		return ChangePasswordUseCaseOutput{Success: false}, entity.ErrOldPasswordInvalid
	}
	user.Password = input.NewPassword
	err = user.Validate()
	if err != nil {
		return ChangePasswordUseCaseOutput{Success: false}, entity.ErrOldPasswordInvalid
	}

	user.NewPassword(user.Password)
	_, err = c.Gateway.Update(ctx, user)
	if err != nil {
		return ChangePasswordUseCaseOutput{}, nil
	}
	return ChangePasswordUseCaseOutput{Success: true}, nil
}
