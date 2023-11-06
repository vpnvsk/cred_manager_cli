package models

import "github.com/google/uuid"

type PSList struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" `
}
type AllList struct {
	Data []PSList `json:"data"`
}
type SingleManager struct {
	Userlogin string `json:"userlogin"`
	Password  string `json:"password"`
}
type UpdateManager struct {
	Title       string `json:"title,omitempty"`
	Userlogin   string `json:"userlogin,omitempty"`
	Password    string `json:"password,omitempty" `
	Description string `json:"description,omitempty""`
}
