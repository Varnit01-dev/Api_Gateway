package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.AuthenticationMiddleware)
	r.Use(middleware.LoggingMiddleware)

	r.Handlefunc("/api/v1/endpoint", handlers.HandlerFunc).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// 	router := mux.NewRouter()
	// 	router.Use(middleware.JwtAuthentication)

	//	r.HandleFunc("/users", userHandlers).Methods(http.MethodGet)
	//	r.HandleFunc("/products",productsHandlers).Methods(http.MethodGet)

	log.Printf("Starting API Gateway server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
