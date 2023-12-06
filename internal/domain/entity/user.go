package entity

type User struct {
	Name  string
	Email string
	Pass  string
}

func (u User) NewUser(name string, email string, pass string) *User {
	return &User{
		Name:  name,
		Email: email,
		Pass:  pass,
	}
}
