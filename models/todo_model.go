package models

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
)

type Todo struct {
	Id     int64  `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status string `json:"status" db:"status"`
}

func GetTodo(sess *dbr.Session, id int64) (Todo, error) {
	var todo Todo
	_, err := sess.Select("*").From("todos").Where("id = ?", id).Load(&todo)

	return todo, err
}

func AllTodos(sess *dbr.Session) ([]Todo, error) {
	var todos []Todo
	_, err := sess.Select("*").From("todos").Load(&todos)

	return todos, err
}

func CreateTodo(sess *dbr.Session, todo Todo) (Todo, error) {
	var id int64
	err := sess.InsertInto("todos").Columns("title", "status").Values(todo.Title, todo.Status).Returning("id").Load(&id)
	return todo, err
}

func UpdateTodo(sess *dbr.Session, todo Todo, id int64) (Todo, error) {
	// Execute the update query
	_, err := sess.Update("todos").
		Set("title", todo.Title).
		Set("status", todo.Status).
		Where("id = ?", id).Exec() // Use Exec() here as Load() is unnecessary for update queries

	// Handle any errors during the update
	if err != nil {
		return Todo{}, err
	}

	// Update the ID field of the returned todo object
	todo.Id = id

	// Return the updated todo object
	return todo, nil
}

func DeleteTodo(sess *dbr.Session, id int64) (int64, error) {
	// Execute the delete query
	result, err := sess.DeleteFrom("todos").Where("id = ?", id).Exec()
	if err != nil {
		return 0, err // Return 0 rows and the error
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("no record found with ID %d", id)
	}

	return rowsAffected, nil // Return the number of affected rows and no error
}
