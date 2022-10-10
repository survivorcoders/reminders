package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
)

type ReminderController struct {
	ReminderRepository repository.ReminderRepository
}

func (r ReminderController) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, r.ReminderRepository.GetAll())
}

func (r ReminderController) Get(c echo.Context) error {
	id := c.Param("id")
	reminderEntity := r.ReminderRepository.Get(id)
	return c.JSON(http.StatusOK, reminderEntity)
}

func (r ReminderController) Create(c echo.Context) error {
	reminderEntity := &entity.Reminder{}
	if err := c.Bind(reminderEntity); err != nil {
		return err
	}
	//save into database
	r.ReminderRepository.Create(&reminderEntity)
	return c.JSON(http.StatusCreated, reminderEntity)
}

func (r ReminderController) PUT(c echo.Context) error {
	reminderEntity := &entity.Reminder{Name: "Call my mom"}
	//get current reminder from dataBase
	//if empty return error not found
	if err := c.Bind(reminderEntity); err != nil {
		return err
	}

	//call the repo to update the existing reminder
	//save(entity)
	return c.JSON(http.StatusOK, reminderEntity)
}

func (r ReminderController) Delete(c echo.Context) error {

	id := c.Param("id")
	//call repository to validate the existence
	i, _ := strconv.Atoi(id)
	exists := r.ReminderRepository.Exists(i)
	if exists {
		//call the repository to delete
		r.ReminderRepository.Delete(i)
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusNotFound, nil)

}
