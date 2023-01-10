package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dunzoit/diy/lalit/db_config"
	"github.com/dunzoit/diy/lalit/models"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allProducts := db_config.GetAllProducts()
	json.NewEncoder(w).Encode(allProducts)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// if r.Body == nil {
	// 	json.NewEncoder(w).Encode("Please send some data")
	// }

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	db_config.AddProduct(product)
	json.NewEncoder(w).Encode(product)

}

func DeleteAProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	db_config.DeleteOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := db_config.DeleteAllProducts()
	json.NewEncoder(w).Encode(count)
}
