package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/junior-alves/go-test/internal/application"
	"github.com/junior-alves/go-test/internal/infrastructure/form_request"
	"github.com/junior-alves/go-test/internal/infrastructure/repository"
)

func main() {
	println("Start a server!")

	router := httprouter.New()

	login_form := form_request.NewLoginFormRequest(*application.NewLoginService())
	router.POST("/auth/login", login_form.Login)

	repository := repository.NewMemoryProductRepository()
	product_form := form_request.NewProductFormRequest(*application.NewProductService(repository))

	router.POST("/product/create", form_request.ValidateToken(product_form.CreateProductRequest))
	router.GET("/product/:id", form_request.ValidateToken(product_form.ListProductsRequest))

	http.ListenAndServe(":8080", router)
}
