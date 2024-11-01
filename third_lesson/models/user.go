package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/usmonali4/jwt_auth/utils/token"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByID(uid uint) (User,error) {

	var u User

	if err := DB.First(&u,uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()
	
	return u,nil

}

func (u *User) PrepareGive(){
	u.Password = ""
}


type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {

	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
}