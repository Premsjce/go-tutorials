package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/premsjce/go-tuts/3_book_store_management/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", r))

	fmt.Printf("Server is shutdown\n")
}
