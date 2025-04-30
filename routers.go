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
	caraBayar := r.Group("/cara-bayar")
	agama := r.Group("/agama")

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

	subspesialisRepository := repository.NewSubspesialisRepository(db2)
	subspesialisService := service.NewSubSpesialisService(subspesialisRepository)
	subspesialisHandler := controllers.NewSubspesialis(subspesialisService)
	subspesialis.GET("/", subspesialisHandler.GetAll)

	caraBayarRepository := repository.NewCaraBayarRepository(db)
	caraBayarService := service.NewCaraBayarService(caraBayarRepository)
	caraBayarHandler := controllers.NewCaraBayarHandler(caraBayarService)
	caraBayar.GET("/:subsistem", caraBayarHandler.GetAll)

	agamaRepository := repository.NewAgamaRepository(db)
	agamaService := service.NewAgamaService(agamaRepository)
	agamaHandler := controllers.NewAgamaHandler(agamaService)
	agama.GET("/", agamaHandler.GetAll)

	r.Run(root)
}
