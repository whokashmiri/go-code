package router

import (
	"github.com/gorilla/mux"
	"controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/insertMovie", controllers.insertOneMovie).Methods("GET")

	return router

}
