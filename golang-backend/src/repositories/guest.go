package repositories

import (
	"context"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/models"
	"github.com/jmoiron/sqlx"
)

type GuestDomain interface {
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
	InsertGuest(ctx context.Context, dbtx *sqlx.Tx, obj models.GuestModels) error
	GetGuestByGuessKey(ctx context.Context, guessKey string) (*models.GuestModels, error)
	GeAllGuest(ctx context.Context) ([]models.GuestModels, error)
}
