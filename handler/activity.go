package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"vandyahmad/skyshi/activity"
	"vandyahmad/skyshi/helper"

	"github.com/gin-gonic/gin"
)

type activityHandler struct {
	activityService activity.Service
}

func NewActivityHandler(activityService activity.Service) *activityHandler {
	return &activityHandler{
		activityService: activityService,
	}
}

func (h *activityHandler) ListActivity(c *gin.Context) {

	result, err := h.activityService.GetAll()
	if err != nil {
		response := helper.ApiResponse("Not Found", "Activity not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.ApiResponse("Success", "Success", result)
	c.JSON(http.StatusOK, response)
	return

}

func (h *activityHandler) DetailActivity(c *gin.Context) {
	id := c.Param("activityId")
	idInt, _ := strconv.Atoi(id)
	result, err := h.activityService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse("Success", "Success", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *activityHandler) CreateActivity(c *gin.Context) {
	var input activity.InputActivity
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidation(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newCashier, err := h.activityService.CreateActivity(&input)
	if err != nil {
		response := helper.ApiResponse("Bad Request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success", "Success", newCashier)
	c.JSON(http.StatusOK, response)
	return

}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	id := c.Param("activityId")
	idInt, _ := strconv.Atoi(id)
	_, err := h.activityService.GetById(idInt)
	if err != nil {
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	err = h.activityService.Delete(idInt)
	response := helper.ApiWithOutData("Success", "Success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *activityHandler) UpdateActivity(c *gin.Context) {
	var input activity.InputActivity
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.FormatErrorValidation(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println("input ", input)

	id := c.Param("activityId")
	idInt, _ := strconv.Atoi(id)
	_, err = h.activityService.GetById(idInt)
	if err != nil {
		fmt.Println("Error ", err)
		response := helper.ApiResponse("Not Found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	// maka update
	res, err := h.activityService.UpdateCashier(idInt, &input)
	if err != nil {
		response := helper.ApiResponse("failed", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success", "Success", res)
	c.JSON(http.StatusOK, response)
	return
}
