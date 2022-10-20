package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type UserProviderRepository struct {
	Database *gorm.DB
}

func NewUserProviderRepository(database *gorm.DB) *UserProviderRepository {
	return &UserProviderRepository{Database: database}
}

func (receiver *UserProviderRepository) CreateUser(reminder *entity.User) error {
	result := receiver.Database.Create(reminder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetUserById(id int, reminder *entity.User) error {
	return receiver.Database.First(reminder, id).Error

}

func (receiver *UserProviderRepository) GetUserByUsername(username string, reminder *entity.User) error {
	result := receiver.Database.First(reminder, "username = ?", username)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetAllUsers(reminders []entity.User) error {
	return receiver.Database.Find(&reminders).Error
}

func (receiver *UserProviderRepository) UpdateUser(r entity.User) error {
	return receiver.Database.Model(&r).Updates(r).Error
}

func (receiver *UserProviderRepository) DeleteUser(id int) error {
	return receiver.Database.Delete(&entity.User{}, id).Error
}
