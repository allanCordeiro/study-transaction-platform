package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type CreateUserUseCase struct {
	Gateway entity.UserInterface
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	UserType string
}

type CreateUserOutput struct {
	UserId string
}

func NewCreateUserUseCase(gateway entity.UserInterface) *CreateUserUseCase {
	return &CreateUserUseCase{Gateway: gateway}
}

func (c *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.UserType, input.Password)
	if err != nil {
		return nil, err
	}

	user.Activate()

	err = c.Gateway.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutput{
		UserId: user.Id,
	}, nil
}
