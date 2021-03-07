package guest

import (
	"context"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/database"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/models"
	"github.com/jmoiron/sqlx"
)

func (g *guestDomain) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return g.DB.BeginTxx(ctx, nil)
}

const insertGuestQuery = "INSERT INTO `guest_list` (`guest_name`, `guest_group`, `max_guest`, `sides`, `guess_key`) VALUES (?, ?, ?, ?, ?);"

func (g *guestDomain) InsertGuest(ctx context.Context, dbtx *sqlx.Tx, obj models.GuestModels) error {
	var executor database.ExecFunc
	if dbtx != nil {
		executor = dbtx.ExecContext
	} else {
		executor = g.DB.ExecContext
	}
	_, err := executor(ctx, insertGuestQuery, obj.GuestName, obj.GuestGroup, obj.MaxGuest, obj.Sides, obj.GuestKey)
	return err
}

const selectGuestByGuessKeyQuery = "SELECT id, guest_name, guest_group, max_guest, sides, guest_key  FROM `guest_list` WHERE `guess_key` = ?"

func (g *guestDomain) GetGuestByGuessKey(ctx context.Context, guessKey string) (*models.GuestModels, error) {
	resp := models.GuestModels{}
	err := g.DB.GetContext(ctx, &resp, selectGuestByGuessKeyQuery, guessKey)
	return &resp, err
}

const selectAllGuestQuery = "SELECT id, guest_name, guest_group, max_guest, sides, guest_key  FROM `guest_list`"

func (g *guestDomain) GeAllGuest(ctx context.Context) ([]models.GuestModels, error) {
	resp := make([]models.GuestModels, 0)
	err := g.DB.GetContext(ctx, &resp, selectAllGuestQuery)
	return resp, err
}
