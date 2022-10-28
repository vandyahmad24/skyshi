package router

import (
	"vandyahmad/skyshi/config"
	"vandyahmad/skyshi/handler"
	"vandyahmad/skyshi/todo"

	"github.com/gin-gonic/gin"
)

func TodoRouter(app *gin.Engine) {

	todoRepository := todo.NewRepository(config.DB)
	todoService := todo.NewService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	app.GET("/todo-items", todoHandler.ListTodo)
	app.GET("/todo-items/:todoId", todoHandler.DetailTodo)
	app.POST("/todo-items", todoHandler.CreateTodo)
	app.PATCH("/todo-items/:todoId", todoHandler.UpdateTodo)
	app.DELETE("/todo-items/:todoId", todoHandler.DeleteTodo)

}
