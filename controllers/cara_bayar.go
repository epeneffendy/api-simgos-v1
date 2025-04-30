package controllers

import (
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CaraBayarHandler struct {
	CaraBayarService service.CaraBayarService
}

func NewCaraBayarHandler(caraBayarService service.CaraBayarService) *CaraBayarHandler {
	return &CaraBayarHandler{caraBayarService}
}

func (c *CaraBayarHandler) GetAll(ctx *gin.Context) {
	subSistemID, err := strconv.Atoi(ctx.Param("subsistem"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Get review Cara Bayar Failed",
			http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	caraBayar, err := c.CaraBayarService.GetAll(subSistemID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get Poliklinik By Sub Sistem failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get carabayar by subsistem success", http.StatusOK, "success", caraBayar)
	ctx.JSON(http.StatusOK, response)
}
