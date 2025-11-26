package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course //fake DB

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - MenScriptCourses.com")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Go Backend Developer Course",
		CoursePrice: 799,
		Author:      &Author{FullName: "Swatantra Yadav", Website: "go.dev"}})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "ReactJS",
		CoursePrice: 599,
		Author:      &Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})
	courses = append(courses, Course{
		CourseId:    "6",
		CourseName:  "Solana Blockchain Full Course",
		CoursePrice: 999,
		Author:      &Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})
	courses = append(courses, Course{
		CourseId:    "10",
		CourseName:  "Rust Course",
		CoursePrice: 1499,
		Author:      &Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})

	//routing(RESTFull)
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/create/{id}", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// handlers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by MenScriptOnline.com</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		json.NewEncoder(w).Encode("Invalid JSON")
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate (almost) unique id
	id := strconv.Itoa(rand.Intn(100))
	for _, c := range courses {
		if c.CourseId == id {
			id = strconv.Itoa(rand.Intn(100))
		}
	}
	course.CourseId = id

	courses = append(courses, course) //save in "DB"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, c := range courses {
		if c.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)

			var updated Course
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Invalid JSON payload")
				return
			}

			if updated.IsEmpty() {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("No data inside JSON")
				return
			}

			updated.CourseId = id
			courses = append(courses, updated)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //grad the id to be deleted

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course with given id is deleted")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with given id")
}
