package main

import (
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/webservice"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/config"
	http_backdoor "github.com/hasemeneh/my-wedding-invitation/golang-backend/src/handler/http/backdoor"
	http_public "github.com/hasemeneh/my-wedding-invitation/golang-backend/src/handler/http/public"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/service"
)

func main() {
	cfg := config.New().Read()
	serviceObj := service.New(cfg)
	publicHttpHandler := http_public.NewHandler(serviceObj)
	internalHttpHandler := http_backdoor.NewHandler(serviceObj)
	ws := webservice.NewWebService(
		cfg.RunningPort,
		publicHttpHandler,
		internalHttpHandler,
	)
	ws.Run()
}
