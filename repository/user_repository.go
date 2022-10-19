package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) AddNewUser(user *entity.User) *entity.User {

	if dbc := r.DB.Save(&user); dbc.Error != nil {
		// Create failed, do something e.g. return, panic etc.
		return nil
	}

	return user
}

func (r UserRepository) GetUserByEmail(email string) *entity.User {
	var user entity.User

	if dbc := r.DB.Where("email = ?", email).First(&user); dbc.Error != nil {
		return nil
	}
	return &user
}
