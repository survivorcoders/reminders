package entity

type Date string

type Reminder struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	RemindMeAt  Date   `json:"remindMeAt"`
	Description string `json:"description"`
}
