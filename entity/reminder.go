package entity

import "time"

type Reminder struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	RemindMeAt  time.Time `json:"remindMeAt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
