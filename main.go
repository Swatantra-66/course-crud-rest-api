package main

import (
	"buildCourseAPI/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Course API on port 4000...")

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":4000", router))
}
