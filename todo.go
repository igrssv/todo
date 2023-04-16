package todo

import "errors"

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Titel       string `json:"title" db:"title" binding:"required"`
	Descriprion string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Titel       string `json:"Titel"`
	Descriprion string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemID int
}

type UpdateListInput struct {
	Titel       *string `json:"title"`
	Descriprion *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Titel == nil && i.Descriprion == nil {
		return errors.New("update strucrure has no values")
	}

	return nil
}
