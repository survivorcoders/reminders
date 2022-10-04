package repository

import (
	"gorm.io/gorm"
	"survivorcoders.com/reminders/entity"
)

type ReminderRepository struct {
	DB *gorm.DB
}

func (r ReminderRepository) GetAll() []entity.Reminder {
	var reminders []entity.Reminder
	_ = r.DB.Find(&reminders)
	return reminders
}
