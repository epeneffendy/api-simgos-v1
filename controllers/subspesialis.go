package controllers

import (
	"api-simgos/entity"
	"api-simgos/helper"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Subspesialis struct {
	subspesialisService service.SubSpesialisService
}

func NewSubspesialis(subspesialisService service.SubSpesialisService) *Subspesialis {
	return &Subspesialis{subspesialisService}
}

type SubspesialisGetAllFormat struct {
	Id              int
	Kode            string
	Spesialis       string
	Subspesialis    string
	PoliHfis        string
	IdGroupLocation int
}

func FormatSubspesialisGetAllResponse(subs []entity.Subspesialis) []SubspesialisGetAllFormat {
	var response []SubspesialisGetAllFormat
	for _, sub := range subs {
		res := SubspesialisGetAllFormat{
			Id:              sub.Id,
			Kode:            sub.Kode,
			Spesialis:       sub.Spesialis,
			Subspesialis:    sub.Subspesialis,
			PoliHfis:        sub.PoliHfis,
			IdGroupLocation: sub.IdGroupLocation,
		}
		response = append(response, res)
	}
	return response
}

func (p *Subspesialis) GetAll(ctx *gin.Context) {
	subspesialis, err := p.subspesialisService.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success", http.StatusOK, "data", FormatSubspesialisGetAllResponse(subspesialis))
	ctx.JSON(http.StatusOK, response)
}
