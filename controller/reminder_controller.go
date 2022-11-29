package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
)

type ReminderController struct {
	ReminderService *repository.RemindersProviderRepository
}

func NewReminderController(reminderService *repository.RemindersProviderRepository) *ReminderController {
	return &ReminderController{ReminderService: reminderService}
}

func (receiver *ReminderController) PostCreateReminder(c echo.Context) error {
	reminder := &entity.Reminder{}

	err := c.Bind(reminder)
	if err != nil {
		return err
	}

	err = receiver.ReminderService.CreateReminder(reminder)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, reminder)
}

func (receiver *ReminderController) GetReminder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var reminder entity.Reminder

	err = receiver.ReminderService.GetReminder(id, &reminder)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, reminder)
}

func (receiver *ReminderController) GetAllReminders(c echo.Context) error {
	var reminders []entity.Reminder

	err := receiver.ReminderService.GetAllReminders(&reminders)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, reminders)
}

func (receiver *ReminderController) PutUpdateReminder(c echo.Context) error {
	r := new(entity.Reminder)

	if err := c.Bind(r); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	r.Id = id
	count := receiver.ReminderService.UpdateReminder(*r)
	if count == 0 {
		return errors.New("record not found")
	}

	return c.JSON(http.StatusOK, r)
}

func (receiver *ReminderController) DeleteReminder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	count := receiver.ReminderService.DeleteReminder(id)
	if count == 0 {
		return errors.New("record not found")
	}
	return c.NoContent(http.StatusNoContent)
}
