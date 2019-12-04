package router

import (
	"github.com/gorilla/mux"
	"github.com/vm/handlers"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/instances", handlers.CreateInstance).Methods("POST")
	r.HandleFunc("/v1/instances", handlers.GetAllInstances).Methods("GET")
	r.HandleFunc("/v1/instances/{instanceid}", handlers.GetInstance).Methods("GET")
	r.HandleFunc("/v1/instances/{instanceid}", handlers.UpdateInstance).Methods("PUT")
	r.HandleFunc("/v1/instances/{instanceid}", handlers.DeleteInstance).Methods("DELETE")
	return r
}
