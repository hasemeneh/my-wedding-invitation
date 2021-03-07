package backdoor

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strings"

	"github.com/hasemeneh/my-wedding-invitation/golang-backend/helper/response"
)

func (i *Internal) HandleAddGuest(r *http.Request) response.HttpResponse {
	err := r.ParseMultipartForm(2 << 23)
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	file, header, err := r.FormFile("files")
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	defer file.Close()
	fileName := strings.Split(header.Filename, ".")
	if fileName[len(fileName)-1] != "csv" {
		err = errors.New("Invalid file format, file must be csv")
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return response.NewJsonResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = i.Service.Backdoor.AddBulkGuest(r.Context(), lines)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetMessage("Success")
}
