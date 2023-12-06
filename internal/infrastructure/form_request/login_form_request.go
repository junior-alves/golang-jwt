package form_request

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/junior-alves/go-test/internal/application"
)

type LoginFormRequest struct {
	service application.LoginService
}

func NewLoginFormRequest(service application.LoginService) *LoginFormRequest {
	return &LoginFormRequest{service: service}
}

func (f *LoginFormRequest) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w = *HttpCommon(w)

	type LoginDTO struct {
		Email string `json:"email" validate:"required"`
		Pass  string `json:"pass" validate:"required"`
	}

	var data LoginDTO

	json.NewDecoder(r.Body).Decode(&data)

	user := f.service.Login(data.Email, data.Pass)

	if user != nil {
		token, err := CreateToken(user)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
			return
		}

	}

	w.WriteHeader(http.StatusForbidden)
}
