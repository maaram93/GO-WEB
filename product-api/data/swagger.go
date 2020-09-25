package data

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the data store
	// in: body
	Body []Product
}
