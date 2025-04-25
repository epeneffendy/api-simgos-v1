package controllers

import (
	"api-simgos/entity"
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pasienHandler struct {
	pasienService service.PasienService
}

type PasienGetAllFormat struct {
	Norm   string
	Nama   string
	Alamat string
}

func NewPasienHandler(pasienService service.PasienService) *pasienHandler {
	return &pasienHandler{pasienService}
}

func FormatPasienGetAllResponse(cats []entity.Pasien) []PasienGetAllFormat {
	var response []PasienGetAllFormat
	for _, cat := range cats {
		res := PasienGetAllFormat{
			Norm:   cat.Norm,
			Nama:   cat.Nama,
			Alamat: cat.Alamat,
		}
		response = append(response, res)
	}
	return response
}

func (p *pasienHandler) GetAllPasien(ctx *gin.Context) {
	pasiens, err := p.pasienService.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := FormatPasienGetAllResponse(pasiens)
	response := helper.APIResponse("Success get data", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}
