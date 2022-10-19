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
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, r.ReminderRepository.Get(id))
}

func (r ReminderController) Create(c echo.Context) error {
	reminderEntity := &entity.Reminder{}
	if err := c.Bind(reminderEntity); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, r.ReminderRepository.Create(*reminderEntity))
}

func (r ReminderController) PUT(c echo.Context) error {
	reminderEntity := &entity.Reminder{}
	//get current reminder from dataBase
	//if empty return error not found
	if err := c.Bind(reminderEntity); err != nil {
		return err
	}
	//call the repo to update the existing reminder
	//save(entity)
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, r.ReminderRepository.PUT(*reminderEntity, id))
}

func (r ReminderController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	//call the repository to delete
	return c.JSON(http.StatusOK, r.ReminderRepository.Delete(id))
}
