package controllers

import (
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PoliklinikHandler struct {
	PoliklinikService service.PoliklinikService
}

func NewPoliklinikHandler(poliklinikService service.PoliklinikService) *PoliklinikHandler {
	return &PoliklinikHandler{poliklinikService}
}

func (p *PoliklinikHandler) GetBySubSistem(ctx *gin.Context) {
	subSistemID, err := strconv.Atoi(ctx.Param("subsistem"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get reviews by Sub Sistem ID failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	poliklinik, err := p.PoliklinikService.GetBySubSistem(subSistemID)
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

	response := helper.APIResponse("Get poliklinik by subsistem success", http.StatusOK, "success", poliklinik)
	ctx.JSON(http.StatusOK, response)
}
