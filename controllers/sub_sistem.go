package controllers

import (
	"api-simgos/entity"
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type subSistemHandler struct {
	subSistemService service.SubSistemService
}

type SubSistemGetAllFormat struct {
	Id            int
	NamaSubSistem string
}

func NewSubSistemHandler(subSistemService service.SubSistemService) *subSistemHandler {
	return &subSistemHandler{subSistemService}
}

func FormatSubSistemGetAllResponse(subs []entity.SubSistem) []SubSistemGetAllFormat {
	var response []SubSistemGetAllFormat
	for _, sub := range subs {
		res := SubSistemGetAllFormat{
			Id:            sub.Id,
			NamaSubSistem: sub.NamaSubSistem,
		}
		response = append(response, res)
	}
	return response
}

func (p *subSistemHandler) GetAllSubSistem(ctx *gin.Context) {
	subsistem, err := p.subSistemService.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success", http.StatusOK, "data", FormatSubSistemGetAllResponse(subsistem))
	ctx.JSON(http.StatusOK, response)
}
