package application

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/junior-alves/go-test/internal/domain/entity"
)

type LoginService struct {
	users []*entity.User
}

func NewLoginService() *LoginService {

	var hash = sha1.New()
	hash.Write([]byte("123"))
	pass := hex.EncodeToString(hash.Sum(nil))
	return &LoginService{
		users: []*entity.User{
			{Name: "Juneba", Email: "test@test.com", Pass: pass},
		},
	}
}

func (service LoginService) Login(email string, pass string) *entity.User {

	var hash = sha1.New()
	hash.Write([]byte(pass))

	for _, user := range service.users {
		if user.Email == email && user.Pass == hex.EncodeToString(hash.Sum(nil)) {
			return user
		}
	}

	return nil
}
