package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type RemindersProviderRepository struct {
	Database *gorm.DB
}

func (receiver *RemindersProviderRepository) CreateReminder(reminder *entity.Reminder) {
	receiver.Database.Create(reminder)
}

func (receiver *RemindersProviderRepository) GetReminder(id int, reminder *entity.Reminder) *entity.Reminder {
	receiver.Database.First(reminder, id)

	return reminder
}

func (receiver *RemindersProviderRepository) GetAllReminders(reminders *[]entity.Reminder) {
	receiver.Database.Find(reminders)
}

func (receiver *RemindersProviderRepository) UpdateReminder(r entity.Reminder) {
	receiver.Database.Model(&r).Updates(r)
}

func (receiver *RemindersProviderRepository) DeleteReminder(id int) {
	receiver.Database.Delete(&entity.Reminder{}, id)
}
