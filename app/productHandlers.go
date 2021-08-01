package app

import (
	"net/http"

	"github.com/ashishjuyal/banking/service"
)

type ProductHandlers struct {
	service service.ProductService
}

func (ch *ProductHandlers) getAllProducts(w http.ResponseWriter, r *http.Request) {

	discontinued := r.URL.Query().Get("discontinued")

	products, err := ch.service.GetAllProduct(discontinued)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, products)
	}
}

/*
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}
*/
/*
func writeProductResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
*/
