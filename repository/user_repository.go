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
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.Create(user).Error
		c <- err
	}(c)

	err := <-c

	if err != nil {

		return err
	}
	return nil
}

func (receiver *UserProviderRepository) GetUserById(id int, user *entity.User) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.First(user, id).Error
		c <- err
	}(c)

	err := <-c

	return err
}

func (receiver *UserProviderRepository) GetUserByUsername(username string, user *entity.User) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.First(user, "username = ?", username).Error
		c <- err
	}(c)

	err := <-c

	return err
}

func (receiver *UserProviderRepository) GetAllUsers(users []entity.User) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.Find(&users).Error
		c <- err
	}(c)
	err := <-c

	return err
}

func (receiver *UserProviderRepository) UpdateUser(user entity.User) int64 {
	c := make(chan int64)

	go func(c chan int64) {
		count := receiver.Database.Model(&user).Updates(user).RowsAffected
		c <- count
	}(c)
	count := <-c

	return count
}

func (receiver *UserProviderRepository) DeleteUser(id int) int64 {
	c := make(chan int64)

	go func(c chan int64) {
		count := receiver.Database.Delete(&entity.User{}, id).RowsAffected
		c <- count
	}(c)
	count := <-c

	return count
}
