package usecases

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
)

type FindUserUseCase struct {
	Gateway entity.UserInterface
}

type FindUserInput struct {
	Id    string
	Email string
}

type FindUserOutput struct {
	UserId   string
	Name     string
	Email    string
	Password string
	UserType string
}

func NewFindUserUseCase(gateway entity.UserInterface) *FindUserUseCase {
	return &FindUserUseCase{Gateway: gateway}
}

func (c *FindUserUseCase) Execute(ctx context.Context, input FindUserInput) (*FindUserOutput, error) {
	if input.Id != "" {
		user, err := c.Gateway.FindByID(ctx, input.Id)
		if err != nil {
			return nil, err
		}
		return &FindUserOutput{
			UserId:   user.Id,
			Name:     user.Name,
			Email:    user.Email.GetEmail(),
			Password: user.Password,
			UserType: user.UserType.String(),
		}, nil
	}

	user, err := c.Gateway.FindByMail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	return &FindUserOutput{
		UserId:   user.Id,
		Name:     user.Name,
		Email:    user.Email.GetEmail(),
		Password: user.Password,
		UserType: user.UserType.String(),
	}, nil
}
