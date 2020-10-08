package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products products updateProduct
// Updates product details in data store
// responses:
// 	200: noContentResponse

// UpdateProduct updates product in data store based on ID.
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
