package handlers

import (
	"crudApp/config"
	"crudApp/helpers"
	"crudApp/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gocraft/dbr/v2"
	"github.com/gorilla/mux"
)

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// Convert ID to int64
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helpers.HandleError(w, err, "Invalid ID format", http.StatusBadRequest)
		return
	}

	sess := config.DBConnect.NewSession(nil)
	defer sess.Close()

	data, err := models.GetTodo(sess, id)
	if err != nil {
		// Check for "not found" errors
		if errors.Is(err, dbr.ErrNotFound) { // Assuming dbr.ErrNotFound for dbr library
			helpers.HandleError(w, err, "Data not found", http.StatusNotFound)
			return
		}

		// Handle other errors
		helpers.HandleError(w, err, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	helpers.JsonResponse(w, http.StatusOK, true, "Todo retrieved successfully", data)
}

func AllTodoHandler(w http.ResponseWriter, r *http.Request) {
	sess := config.DBConnect.NewSession(nil)
	todos, err := models.AllTodos(sess)
	if err != nil {
		helpers.HandleError(w, err, "Failed to fetch todos", http.StatusBadRequest)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, true, "Todos retrieved successfully", todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		helpers.HandleError(w, err, "Invalid request body", http.StatusBadRequest)
		return
	}

	sess := config.DBConnect.NewSession(nil)

	createdTodo, err := models.CreateTodo(sess, todo)

	if err != nil {
		helpers.HandleError(w, err, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	defer sess.Close()

	// Respond with JSON
	helpers.JsonResponse(w, http.StatusCreated, true, "Todo Created successfully", createdTodo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the `id` parameter from the URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helpers.HandleError(w, err, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Parse the request body into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helpers.HandleError(w, err, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if todo.Title == "" || todo.Status == "" {
		helpers.HandleError(w, nil, "Title and Status are required", http.StatusBadRequest)
		return
	}

	// Open a new database session
	sess := config.DBConnect.NewSession(nil)
	defer sess.Close()

	// Call the model function to update the todo
	updatedTodo, err := models.UpdateTodo(sess, todo, id)
	if err != nil {
		helpers.HandleError(w, err, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	// Send a JSON response
	helpers.JsonResponse(w, http.StatusOK, true, "Todo updated successfully", updatedTodo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helpers.HandleError(w, err, "Invalid ID format", http.StatusBadRequest)
		return
	}

	sess := config.DBConnect.NewSession(nil)
	defer sess.Close()

	_, err = models.DeleteTodo(sess, id)
	if err != nil {
		helpers.HandleError(w, err, "Failed to delete todo", http.StatusInternalServerError)
		return
	} else {
		helpers.JsonResponse(w, http.StatusNoContent, true, "Todo deleted successfully", nil)
	}
}
