package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Descriprion *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Descriprion *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Descriprion == nil {
		return errors.New("update strucrure has no values")
	}

	return nil
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Descriprion == nil && i.Done == nil {
		return errors.New("update strucrure has no values")
	}

	return nil
}
