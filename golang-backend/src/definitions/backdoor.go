package definitions

import "context"

type BackdoorUsecase interface {
	AddBulkGuest(ctx context.Context, csv [][]string) error
}
