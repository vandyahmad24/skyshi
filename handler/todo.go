package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"vandyahmad/skyshi/helper"
	"vandyahmad/skyshi/todo"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func (h *todoHandler) ListTodo(c *gin.Context) {
	result, err := h.todoService.GetAll(c)
	if err != nil {
		response := helper.ApiResponse("Not Found", "Todo not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.ApiResponse("Success", "Success", result)
	c.JSON(http.StatusOK, response)
	return

}

func (h *todoHandler) DetailTodo(c *gin.Context) {
	id := c.Param("todoId")
	idInt, _ := strconv.Atoi(id)
	result, err := h.todoService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse("Success", "Success", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.InputTodo
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidation(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(input)
	newCashier, err := h.todoService.CreateActivity(&input)
	if err != nil {
		response := helper.ApiResponse("Bad Request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success", "Success", newCashier)
	c.JSON(http.StatusCreated, response)
	return

}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("todoId")
	idInt, _ := strconv.Atoi(id)
	_, err := h.todoService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.todoService.Delete(idInt)
	response := helper.ApiWithOutData("Success", "Success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var input todo.InputTodo
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidation(err)
		c.JSON(http.StatusNotFound, response)
		return
	}
	fmt.Println("input ", input)

	id := c.Param("todoId")
	idInt, _ := strconv.Atoi(id)
	_, err = h.todoService.GetById(idInt)
	if err != nil {
		fmt.Println("Error ", err)
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	// maka update
	res, err := h.todoService.UpdateTodo(idInt, &input)
	if err != nil {
		response := helper.ApiResponse("failed", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse("Success", "Success", res)
	c.JSON(http.StatusOK, response)
	return
}
