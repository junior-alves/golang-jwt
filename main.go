package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/junior-alves/go-test/internal/application"
	form "github.com/junior-alves/go-test/internal/infrastructure/http"
	"github.com/junior-alves/go-test/internal/infrastructure/repository"
)

func main() {
	println("Start a server!")

	router := httprouter.New()

	repository := repository.NewMemoryProductRepository()
	service := application.NewProductService(repository)
	form := form.NewFormRequest(*service)

	router.POST("/product/create", form.CreateProductRequest)
	router.GET("/product/:id", form.ListProductsRequest)

	http.ListenAndServe(":8080", router)
}
