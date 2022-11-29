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

func (receiver *UserProviderRepository) CreateUser(user *entity.User) error {
	return receiver.Database.Create(user).Error
}

func (receiver *UserProviderRepository) GetUserById(id int, user *entity.User) error {
	return receiver.Database.First(user, id).Error
}

func (receiver *UserProviderRepository) GetUserByUsername(username string, user *entity.User) error {
	return receiver.Database.First(user, "username = ?", username).Error
}

func (receiver *UserProviderRepository) GetAllUsers(users []entity.User) error {
	return receiver.Database.Find(&users).Error
}

func (receiver *UserProviderRepository) UpdateUser(user entity.User) int64 {
	return receiver.Database.Model(&user).Updates(user).RowsAffected
}

func (receiver *UserProviderRepository) DeleteUser(id int) int64 {
	return receiver.Database.Delete(&entity.User{}, id).RowsAffected
}
