package form_request

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/junior-alves/go-test/internal/application"
)

type ProductFormRequest struct {
	service application.ProductService
}

func NewProductFormRequest(service application.ProductService) *ProductFormRequest {
	return &ProductFormRequest{service: service}
}

func (f *ProductFormRequest) CreateProductRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w = *HttpCommon(w)

	type ProductDTO struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,numeric"`
	}

	var data ProductDTO

	json.NewDecoder(r.Body).Decode(&data)

	product := f.service.CreateProduct(data.Name, int(data.Price))

	res, _ := json.Marshal(product)

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (f *ProductFormRequest) ListProductsRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w = *HttpCommon(w)

	id := p.ByName("id")

	if id != "list" {
		f.getProductById(w, r, id)
		return
	}

	products := f.service.ListProducts()

	res, _ := json.Marshal(products)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (f *ProductFormRequest) getProductById(w http.ResponseWriter, r *http.Request, id string) {

	product, err := f.service.GetProduct(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(product)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
