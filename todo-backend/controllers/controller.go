// Package controllers contains the todo backend controllers
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"todo-rest-backend/models"
	"todo-rest-backend/models/persistence"

	"github.com/julienschmidt/httprouter"
)

const BackendHostUrl string = ":8080"

// Run does the running of the web server
func Run() {
	models.SetTodoPersistence(persistence.CsvFileTodoPersistence{})
	fmt.Println("Backend running at:", BackendHostUrl)
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodosGet)
	router.GET("/todos/:id", TodoGetById)
	router.POST("/todos", TodoPost)
	router.PUT("/todos/:id", TodoPutById)

	err := http.ListenAndServe(BackendHostUrl, router)
	log.Fatal(err)
}

// Index Handler for the index action
// GET /
func Index(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(writer, "Welcome to the Todo REST API!\n")
	if err != nil {
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
}

// TodosGet Handler for the todos get action
// GET /todos
func TodosGet(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	todos, err := models.ReadTodos()
	if err != nil {
		handleError(writer, http.StatusNotFound, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	sortedTodos := models.SortTodosAfterIdAscending(todos)
	response := models.JsonDataResponse{Data: sortedTodos}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

func handleError(writer http.ResponseWriter, statusCode int, text string) {
	writer.WriteHeader(statusCode)
	response := models.JsonErrorResponse{Error: models.ApiError{Status: statusCode, Title: text}}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// TodoGetById Handler for a todo get by id action
func TodoGetById(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get todo id from url parameters
	id := params.ByName("id")
	todo, err := models.ReadTodoById(id)
	if err != nil {
		handleError(writer, http.StatusNotFound, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := models.JsonExtendedResponse{Data: todo}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// TodoPost Handler for the todos post action
func TodoPost(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var todo models.Todo
	err := decodeTodo(request, &todo)
	if err != nil {
		handleError(writer, http.StatusBadRequest, "invalid Body")
		return
	}

	todoAdded, err := models.CreateTodo(todo)
	if err != nil {
		handleError(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusCreated)
	response := models.JsonExtendedResponse{Data: todoAdded}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// decodeTodo decodes the json request body into a Todo
func decodeTodo(request *http.Request, todo *models.Todo) error {
	if request.Body == nil {
		return errors.New("invalid body")
	}
	err := json.NewDecoder(request.Body).Decode(todo)
	if err != nil {
		return err
	}
	return nil
}

// TodoPutById Handler for updating a todo by ID
func TodoPutById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//create a Todo
	var newToDoValues models.Todo
	//save new Values to the newToDoValues Variable
	err := decodeTodo(request, &newToDoValues)
	if err != nil {
		handleError(writer, http.StatusBadRequest, "invalid Body")
		return
	}

	//store the id from url parameters
	id := params.ByName("id")
	//Check for existence of id
	_, err = models.ReadTodoById(id)
	if err != nil {
		handleError(writer, http.StatusNotFound, "Id doesn't exist")
	}

	todoUpdated, err := models.UpdateTodoById(newToDoValues)
	if err != nil {
		handleError(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := models.JsonExtendedResponse{Data: todoUpdated}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}
