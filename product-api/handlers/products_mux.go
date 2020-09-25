// Pacakge classification of product API
//
// Documentation for product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsMux struct {
	l *log.Logger
}

func NewProductsMux(l *log.Logger) *ProductsMux {
	return &ProductsMux{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from data store
func (p *ProductsMux) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//data, err := json.Marshal(lp) - using json Marshall method

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to Marshal products", http.StatusInternalServerError)
	}
	//rw.Write(data) - using json Marshall method
}

func (p *ProductsMux) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST Product")

	prdt := r.Context().Value(keyProduct{}).(data.Product)

	//p.l.Printf("Product: %#v", prdt)
	data.AddProduct(&prdt)
}

func (p *ProductsMux) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Product", id)

	prdt := r.Context().Value(keyProduct{}).(data.Product)

	p.l.Printf("Product: %#v", prdt)
	err = data.UpdateProduct(id, &prdt)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type keyProduct struct{}

func (p ProductsMux) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prdt := data.Product{}
		err := prdt.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Failed to unMarshaling json", http.StatusBadRequest)
			return
		}

		err = prdt.Validate()
		if err != nil {
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), keyProduct{}, prdt)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
