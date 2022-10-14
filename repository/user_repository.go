package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type UserProviderRepository struct {
	Database *gorm.DB
}

func (receiver *UserProviderRepository) CreateUser(reminder *entity.User) error {
	result := receiver.Database.Create(reminder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetUserById(id int, reminder *entity.User) *entity.User {
	receiver.Database.First(reminder, id)

	return reminder
}

func (receiver *UserProviderRepository) GetUserByUsername(username string, reminder *entity.User) error {
	result := receiver.Database.First(reminder, "username = ?", username)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetAllUsers(reminders []entity.User) []entity.User {

	receiver.Database.Find(&reminders)

	return reminders
}

func (receiver *UserProviderRepository) UpdateUser(r entity.User) {
	receiver.Database.Model(&r).Updates(r)
}

func (receiver *UserProviderRepository) DeleteUser(id int) {
	receiver.Database.Delete(&entity.User{}, id)
}
