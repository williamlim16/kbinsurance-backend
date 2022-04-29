package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        uint   `json:"id" form:"id"`
	FirstName string `json:firstname; form:"firstname"`
	LastName  string `json:lastname;form:"lastname"`
	Email     string `json:email;form:"email"`
	Password  []byte `json:"-";form:"password"`
	Phone     string `json:"phone" form:"phone"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
