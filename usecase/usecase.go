package usecase

import (
	"context"

	"github.com/Akmyrat03/avito/domain"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	Save(ctx context.Context, u *domain.User) error
}

type UserUsecase struct {
	repo      Repository
	trManager *manager.Manager
}

func NewUserUsecase(repo Repository, tr *manager.Manager) *UserUsecase {
	return &UserUsecase{
		repo:      repo,
		trManager: tr,
	}
}

func (u *UserUsecase) UpdateUsername(ctx context.Context, id int64, newName string) error {
	return u.trManager.Do(ctx, func(ctx context.Context) error {
		user, err := u.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		user.Username = newName

		if err := u.repo.Save(ctx, user); err != nil {
			return err
		}

		return nil
	})
}
