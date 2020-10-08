package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from data store
func (p *ProductsMux) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all products")

	pdts := data.GetProducts()
	//data, err := json.Marshal(lp) - using json Marshall method

	err := pdts.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to Marshal products", http.StatusInternalServerError)
	}
}
