package backdoor

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/response"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/models"
)

var csvHeader = []string{`guest_name`, `guest_group`, `max_guest`, `sides`, `guess_key`}

func (b *backdoorUsecase) AddBulkGuest(ctx context.Context, csv [][]string) error {
	dbtx, err := b.Guest.StartTransaction(ctx)
	if err != nil {
		return err
	}
	defer dbtx.Rollback()
	for lines, row := range csv {
		if len(row) < len(csvHeader) {
			return response.NewResponseError(fmt.Sprintf("Invalid Row Number lines:%d", lines), 400)
		}
		if lines == 0 {
			continue
		}
		GuestName := row[0]
		GuestGroup := row[1]
		MaxGuest, err := strconv.ParseInt(row[2], 10, 64)
		if err != nil {
			return response.NewResponseError(fmt.Sprintf("Invalid MaxGuest value lines:%d", lines), 400)
		}
		Sides, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			return response.NewResponseError(fmt.Sprintf("Invalid Sides value lines:%d", lines), 400)
		}
		GuestKey := row[4]
		err = b.Guest.InsertGuest(ctx, dbtx, models.GuestModels{
			GuestName:  GuestName,
			GuestGroup: GuestGroup,
			MaxGuest:   MaxGuest,
			Sides:      Sides,
			GuestKey:   GuestKey,
		})
		if err != nil {
			return err
		}
	}
	return dbtx.Commit()

}
