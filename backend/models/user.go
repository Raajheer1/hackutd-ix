package models

import (
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Email     string    `gorm:"size:255; not null;" json:"email"`
	Password  string    `gorm:"size:255; not null;" json:"password"`
	Income    uint      `json:"income"`
	Age       uint      `json:"age"`
}

func (u *User) CreateUser() (*User, error) {
	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateUser() (*User, error) {
	var user User
	if err := DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return &User{}, err
	}

	if err := DB.Model(&user).Updates(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) DeleteUser() (*User, error) {
	var user User
	if err := DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return &User{}, err
	}

	if err := DB.Delete(&user).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

// RemovePassword - This cleans the data, so we are not returning the password to the user.
func (u *User) RemovePassword() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}

func Login(email string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	userToken, err := token.Generate(u.ID)

	if err != nil {
		return "", err
	}

	return userToken, nil

}
