package main

import (
	"buildCourseAPI/middleware"
	"buildCourseAPI/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Course API on port 4000...")

	router := mux.NewRouter()

	//logging middleware
	router.Use(middleware.LoggingMiddleware)
	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":4000", router))
}
