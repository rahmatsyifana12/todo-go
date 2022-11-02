package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id			string	`json:"id"`
	Item 		string	`json:"item"`
	Completed 	bool	`json:"completed"`
}

var todos = []Todo{
	{ Id: "1", Item: "Clean room", Completed: false},
	{ Id: "2", Item: "Clean bathroom", Completed: false},
	{ Id: "3", Item: "Clean my car", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.JSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)	
}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if (t.Id == id) {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.Run("localhost:5000")
}