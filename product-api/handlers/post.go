package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route POST /products products createProduct
// Adds a new product to data store
// responses:
// 	200: productResponse

// AddProduct adds a product to data store.
func (p *ProductsMux) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Handling POST Product")

	prdt := r.Context().Value(keyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting Product: %#v\n", prdt)
	data.AddProduct(&prdt)
}
