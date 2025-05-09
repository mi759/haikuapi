package usecase

import (
	"context"

	"github.com/mi759/haikuapi/internal/domain"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, id int32) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) GetUserByID(ctx context.Context, id int32) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return uc.repo.Create(ctx, user)
}