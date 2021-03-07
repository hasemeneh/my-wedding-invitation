package service

import (
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/database"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/definitions"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/domain/guest"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/models"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/usecase/backdoor"
)

type Service struct {
	Backdoor definitions.BackdoorUsecase
	cfg      *models.MainConfig
}

func New(cfg *models.MainConfig) *Service {
	db := database.New().Connect(cfg.DBConnectionString)
	guestDomain := guest.New(db)
	backdoorUsecase := backdoor.New(guestDomain)
	serviceObj := Service{
		Backdoor: backdoorUsecase,
	}
	return &serviceObj
}
