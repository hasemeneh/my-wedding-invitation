package backdoor

import "github.com/hasemeneh/my-wedding-invitation/golang-backend/src/repositories"

type backdoorUsecase struct {
	Guest repositories.GuestDomain
}

func New(Guest repositories.GuestDomain) *backdoorUsecase {
	return &backdoorUsecase{
		Guest: Guest,
	}
}
