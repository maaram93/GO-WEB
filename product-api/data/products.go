// Pacakge classification of product API using standard HTTP package.
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure of an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU is of format abc-defg-hijk

	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

func (p *Product) FromJson(r io.Reader) error {
	dcdr := json.NewDecoder(r)
	return dcdr.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	encdr := json.NewEncoder(w)
	return encdr.Encode(p)
}

// GetProducts returns list of products
func GetProducts() Products {
	return listOfProducts
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	listOfProducts = append(listOfProducts, p)
}

func getNextID() int {
	lastProduct := listOfProducts[len(listOfProducts)-1]
	return lastProduct.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	listOfProducts[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, prdt := range listOfProducts {
		if prdt.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

var listOfProducts = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Contains milk",
		Price:       3.50,
		SKU:         "phl789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Does not contains milk",
		Price:       2.50,
		SKU:         "jkf789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
