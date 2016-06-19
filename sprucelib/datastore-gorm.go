package sprucelib

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type SqlDataStore struct {
	DB *gorm.DB
}

func (ds SqlDataStore) GetUserByUsernameAndPassword(username string, password string) (user User, err error) {
	err = ds.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrUnknownUser
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		err = ErrIncorrectPassword
	}
	return
}

func (db SqlDataStore) NewTokenForUser(user User) (token string, err error) {
	return "NOTAREALTOKEN", nil
}

func (db SqlDataStore) GetUserByToken(token string) (user User, err error) {
	err = db.DB.Where("EXISTS (SELECT user_id FROM user_tokens ut WHERE ut.user_id = users.id AND ut.token = ?)", token).First(&user).Error
	return
}
