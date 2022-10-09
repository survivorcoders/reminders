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

func (r *ReminderRepository) Create(reminder **entity.Reminder) {
	r.DB.Create(reminder)
}

//result := db.First(&user)

func (r *ReminderRepository) Delete(id int) {
	r.DB.Delete(&entity.Reminder{}, id)
}

func (r *ReminderRepository) Exists(id int) bool {
	reminderEntity := &entity.Reminder{}
	result := r.DB.First(&reminderEntity, id)
	if result.Error != nil {
		return false
	}
	return true
}
