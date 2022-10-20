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
	return receiver.Database.Create(reminder).Error
}

func (receiver *RemindersProviderRepository) GetReminder(id int, reminder *entity.Reminder) error {
	return receiver.Database.First(reminder, id).Error
}

func (receiver *RemindersProviderRepository) GetAllReminders(reminders *[]entity.Reminder) error {
	return receiver.Database.Find(reminders).Error
}

func (receiver *RemindersProviderRepository) UpdateReminder(r entity.Reminder) error {
	return receiver.Database.Model(&r).Updates(r).Error
}

func (receiver *RemindersProviderRepository) DeleteReminder(id int) error {
	return receiver.Database.Delete(&entity.Reminder{}, id).Error
}
