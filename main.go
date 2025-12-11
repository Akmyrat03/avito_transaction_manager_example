package main

import (
	"context"
	"log"

	httpv0 "github.com/Akmyrat03/avito/controller/http/v0"
	"github.com/Akmyrat03/avito/repository/postgres"
	"github.com/Akmyrat03/avito/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, "postgres://user:pass@localhost:5432/db")
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}

	getter := trmpgx.DefaultCtxGetter

	userRepo := postgres.NewUserRepo(db, getter)

	trm := manager.Must(trmpgx.NewDefaultFactory(db))

	userUC := usecase.NewUserUsecase(userRepo, trm)

	app := fiber.New()

	handler := httpv0.NewUserHandler(userUC)

	app.Put("/users/username", handler.UpdateUsername)

	app.Listen(":8080")
}
