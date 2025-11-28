package routes

import (
	"buildCourseAPI/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	handlers.SeedCourses()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to Course API</h1>"))
	}).Methods("GET")

	//routing(RESTFull)
	router.HandleFunc("/courses", handlers.GetAllCourses).Methods("GET")
	router.HandleFunc("/courses", handlers.CreateOneCourse).Methods("POST")

	router.HandleFunc("/courses/{id}", handlers.GetOneCourse).Methods("GET")
	router.HandleFunc("/courses/{id}", handlers.UpdateOneCourse).Methods("PUT")
	router.HandleFunc("/courses/{id}", handlers.DeleteOneCourse).Methods("DELETE")
}
