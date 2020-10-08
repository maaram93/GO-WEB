package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

// Products handler for getting and updating products.
type Products struct {
	l *log.Logger
}

// NewProducts returns a  new products handler with the given logger.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL)

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		p.l.Println("id", id)
		p.updateProduct(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//data, err := json.Marshal(lp) - using json Marshall method

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to Marshal products", http.StatusInternalServerError)
	}
	//rw.Write(data) - using json Marshall method
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST Product")

	prdt := &data.Product{}
	err := prdt.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Failed to unMarshaling json", http.StatusBadRequest)
	}

	//p.l.Printf("Product: %#v", prdt)
	data.AddProduct(prdt)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prdt := &data.Product{}
	err := prdt.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Failed to unMarshaling json", http.StatusBadRequest)
	}

	p.l.Printf("Product: %#v", prdt)
	err = data.UpdateProduct(id, prdt)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
