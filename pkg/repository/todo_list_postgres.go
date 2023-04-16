package repository

import (
	"fmt"
	"todo"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	//create transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int // list id
	// create list for db todo list table
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Titel, list.Descriprion)
	if err := row.Scan(&id); err != nil {
		tx.Rollback() // exit transaction
		return 0, err
	}

	// add list for user
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	if err != nil {
		return nil, err
	}

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
