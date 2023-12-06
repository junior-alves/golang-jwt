package form_request

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/junior-alves/go-test/internal/domain/entity"
)

func HttpCommon(w http.ResponseWriter) *http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	return &w
}

var SECRET = "45F131D8-5630-4497-BBD9-8B5881005799"

func CreateToken(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["sub"] = user.Email

	tokenStr, err := token.SignedString([]byte(SECRET))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token != "" {
			token = strings.TrimSpace(strings.Split(token, "Bearer ")[1])

			token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("not authorized")
				}

				return []byte(SECRET), nil
			})

			if err == nil && token.Valid {
				handle(w, r, ps)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	}

}
