// Pacakge classification of product API using GorillaMux package.
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
)

type ProductsMux struct {
	l *log.Logger
}

func NewProductsMux(l *log.Logger) *ProductsMux {
	return &ProductsMux{l}
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
