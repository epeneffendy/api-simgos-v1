package main

import (
	"api-simgos/controllers"
	"api-simgos/repository"
	"api-simgos/service"
	"github.com/gin-gonic/gin"
)

func StartServer(root string) {

	r := gin.Default()
	pasien := r.Group("/pasien")
	subSistem := r.Group("/sub-sistem")
	poliklinik := r.Group("/poliklinik")
	subspesialis := r.Group("/sub-spesialis")

	//Pasien Endpoint Handler
	pasienRepository := repository.NewPasienRepository(db)
	pasienService := service.NewPasienService(pasienRepository)
	pasienHandler := controllers.NewPasienHandler(pasienService)

	pasien.GET("/", pasienHandler.GetAllPasien)

	subSistemRepository := repository.NewSubSistemRepository(db)
	subSistemService := service.NewSubSistemService(subSistemRepository)
	subSistemHandler := controllers.NewSubSistemHandler(subSistemService)

	subSistem.GET("/", subSistemHandler.GetAllSubSistem)

	poliklinikRepository := repository.NewPoliklinikRepository(db)
	poliklinikService := service.NewPoliklinikService(poliklinikRepository)
	poliklinikHandler := controllers.NewPoliklinikHandler(poliklinikService)

	poliklinik.GET("/:subsistem", poliklinikHandler.GetBySubSistem)

	subspesialisRepository := repository.NewSubspesialisRepository(db)
	subspesialisService := service.NewSubSpesialisService(subspesialisRepository)
	subspesialisHandler := controllers.NewSubspesialis(subspesialisService)

	subspesialis.GET("/", subspesialisHandler.GetAll)

	r.Run(root)
}
