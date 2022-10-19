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

func (r ReminderRepository) Get(id int) entity.Reminder {
	/*this methode will take a specific id of a reminder and return all data of this id  */
	var reminder = entity.Reminder{Id: id}
	_ = r.DB.First(&reminder)
	return reminder
}

func (r ReminderRepository) Create(reminder entity.Reminder) entity.Reminder {
	var data = entity.Reminder{Name: reminder.Name, RemindMeAt: reminder.RemindMeAt, Description: reminder.Description}
	r.DB.Create(&data)
	return data
}

func (r ReminderRepository) PUT(reminder entity.Reminder, id int) bool {
	//get the old object based on the 'id' if exist modifies content
	var remNew = entity.Reminder{Id: id}
	result := r.DB.First(&remNew)
	if result.RowsAffected > 0 {
		remNew.Name = reminder.Name
		remNew.RemindMeAt = reminder.RemindMeAt
		remNew.Description = reminder.Description
		r.DB.Save(&remNew)
		return true
	}
	return false
}

func (r ReminderRepository) Delete(id int) bool {
	//new query to know if the object exist then delete it
	var remNew = entity.Reminder{Id: id}
	result := r.DB.First(&remNew)

	if result.RowsAffected > 0 {
		var reminder = entity.Reminder{Id: id}
		r.DB.Delete(&reminder)
		return true
	}
	return false
}
