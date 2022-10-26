package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type RemindersProviderRepository struct {
	Database *gorm.DB
}

func NewRemindersProviderRepository(database *gorm.DB) *RemindersProviderRepository {
	return &RemindersProviderRepository{Database: database}
}

func (receiver *RemindersProviderRepository) CreateReminder(reminder *entity.Reminder) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.Create(reminder).Error
		c <- err
	}(c)

	err := <-c

	return err
}

func (receiver *RemindersProviderRepository) GetReminder(id int, reminder *entity.Reminder) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.First(reminder, id).Error
		c <- err
	}(c)
	err := <-c

	return err
}

func (receiver *RemindersProviderRepository) GetAllReminders(reminders *[]entity.Reminder) error {
	c := make(chan error)

	go func(c chan error) {
		err := receiver.Database.Find(reminders).Error
		c <- err
	}(c)
	err := <-c

	return err
}

func (receiver *RemindersProviderRepository) UpdateReminder(r entity.Reminder) int64 {
	c := make(chan int64)

	go func(c chan int64) {
		count := receiver.Database.Model(&r).Updates(r).RowsAffected
		c <- count
	}(c)
	count := <-c

	return count

}

func (receiver *RemindersProviderRepository) DeleteReminder(id int) int64 {
	c := make(chan int64)

	go func(c chan int64) {
		count := receiver.Database.Delete(&entity.Reminder{}, id).RowsAffected
		c <- count
	}(c)
	count := <-c

	return count
}
