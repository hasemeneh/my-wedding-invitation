package public

import (
	"net/http"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/httphandler"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/response"
	"github.com/hasemeneh/my-wedding-invitation/golang-backend/src/service"
	"github.com/julienschmidt/httprouter"
)

type Public struct {
	Service *service.Service
}

func NewHandler(Service *service.Service) *Public {
	return &Public{
		Service: Service,
	}
}
func (p *Public) Register(r *httprouter.Router) {
	apiHttpHandlers := httphandler.New("/api", r)
	apiHttpHandlers.GET("/ping", p.PING)

}
func (p *Public) PING(r *http.Request) response.HttpResponse {
	return response.NewJsonResponse().SetMessage("Pong").SetData("Pung")
}
