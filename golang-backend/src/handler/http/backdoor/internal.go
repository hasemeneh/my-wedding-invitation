package backdoor

import (
	"net/http"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/httphandler"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/response"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/service"
	"github.com/julienschmidt/httprouter"
)

type Internal struct {
	Service *service.Service
}

func NewHandler(Service *service.Service) *Internal {
	return &Internal{
		Service: Service,
	}
}
func (i *Internal) Register(r *httprouter.Router) {
	apiHttpHandlers := httphandler.New("/internal", r)
	apiHttpHandlers.GET("/ping", i.PING)
	apiHttpHandlers.POST("/guest", i.HandleAddGuest)

}
func (i *Internal) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
