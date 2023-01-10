package router

import (
	"github.com/dunzoit/diy/lalit/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", controller.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/add", controller.AddProduct).Methods("POST")
	// router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/product/{id}", controller.DeleteAProduct).Methods("DELETE")
	router.HandleFunc("/api/deleteallproduct", controller.DeleteAllProducts).Methods("DELETE")

	return router

}
