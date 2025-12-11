package postgres

import (
	"context"

	"github.com/Akmyrat03/avito/domain"
	"github.com/Akmyrat03/avito/usecase"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ usecase.Repository = (*userRepo)(nil)

type userRepo struct {
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewUserRepo(db *pgxpool.Pool, getter *trmpgx.CtxGetter) *userRepo {
	return &userRepo{db: db, getter: getter}
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	u := &domain.User{}

	row := conn.QueryRow(ctx, `SELECT id, username FROM users WHERE id = $1`, id)

	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepo) Save(ctx context.Context, u *domain.User) error {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	if u.ID == 0 {
		return conn.QueryRow(ctx, `INSERT INTO users (username) VALUES ($1) RETURNING id`, u.Username).Scan(&u.ID)
	}

	_, err := conn.Exec(ctx, `UPDATE users SET username = $1 WHERE id=$2`, u.Username, u.ID)

	return err
}
