package controllers

import (
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgamaHandler struct {
	agamaService service.AgamaService
}

func NewAgamaHandler(agamaService service.AgamaService) *AgamaHandler {
	return &AgamaHandler{agamaService}
}

func (a *AgamaHandler) GetAll(ctx *gin.Context) {
	agama, err := a.agamaService.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Get poliklinik by subsistem success", http.StatusOK, "success", agama)
	ctx.JSON(http.StatusOK, response)
}
